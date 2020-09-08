// +build m3o

package signup

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/micro/micro/v3/test"
	"github.com/stripe/stripe-go/v71"
	stripe_client "github.com/stripe/stripe-go/v71/client"
)

const (
	retryCount          = 1
	signupSuccessString = "Signup complete"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// generates test emails
func testEmail(nth int) string {
	uid := randStringRunes(8)
	if nth == 0 {
		return fmt.Sprintf("platform+citest+%v@m3o.com", uid)
	}
	return fmt.Sprintf("platform+cites+%v+%v@m3o.com", nth, uid)
}

func TestSignupFlow(t *testing.T) {
	test.TrySuite(t, testSignupFlow, retryCount)
}

func setupM3Tests(serv test.Server, t *test.T) {
	envToConfigKey := map[string][]string{
		"MICRO_STRIPE_API_KEY":                      {"micro.payments.stripe.api_key"},
		"MICRO_SENDGRID_API_KEY":                    {"micro.signup.sendgrid.api_key", "micro.invite.sendgrid.api_key"},
		"MICRO_SENDGRID_TEMPLATE_ID":                {"micro.signup.sendgrid.template_id"},
		"MICRO_SENDGRID_INVITE_TEMPLATE_ID":         {"micro.invite.sendgrid.invite_template_id"},
		"MICRO_STRIPE_PLAN_ID":                      {"micro.subscriptions.plan_id"},
		"MICRO_STRIPE_ADDITIONAL_USERS_PRICE_ID":    {"micro.subscriptions.additional_users_price_id"},
		"MICRO_EMAIL_FROM":                          {"micro.signup.email_from"},
		"MICRO_TEST_ENV":                            {"micro.signup.test_env", "micro.invite.test_env"},
		"MICRO_STRIPE_ADDITIONAL_SERVICES_PRICE_ID": {"micro.subscriptions.additional_services_price_id"},
	}

	for envKey, configKeys := range envToConfigKey {
		val := os.Getenv(envKey)
		if len(val) == 0 {
			t.Fatalf("'%v' flag is missing", envKey)
		}
		for _, configKey := range configKeys {
			outp, err := serv.Command().Exec("config", "set", configKey, val)
			if err != nil {
				t.Fatal(string(outp))
			}
		}
	}
	serv.Command().Exec("config", "set", "micro.billing.max_included_services", "3")

	services := []struct {
		envVar string
		deflt  string
	}{
		{envVar: "M3O_INVITE_SVC", deflt: "../../../invite"},
		{envVar: "M3O_SIGNUP_SVC", deflt: "../../../signup"},
		{envVar: "M3O_STRIPE_SVC", deflt: "../../../payments/provider/stripe"},
		{envVar: "M3O_CUSTOMERS_SVC", deflt: "../../../customers"},
		{envVar: "M3O_NAMESPACES_SVC", deflt: "../../../namespaces"},
		{envVar: "M3O_SUBSCRIPTIONS_SVC", deflt: "../../../subscriptions"},
		{envVar: "M3O_PLATFORM_SVC", deflt: "../../../platform"},
	}

	for _, v := range services {
		outp, err := serv.Command().Exec("run", getSrcString(v.envVar, v.deflt))
		if err != nil {
			t.Fatal(string(outp))
			return
		}
	}

	if err := test.Try("Find signup, invite and stripe in list", t, func() ([]byte, error) {
		outp, err := serv.Command().Exec("services")
		if err != nil {
			return outp, err
		}
		if !strings.Contains(string(outp), "stripe") ||
			!strings.Contains(string(outp), "signup") ||
			!strings.Contains(string(outp), "invite") ||
			!strings.Contains(string(outp), "customers") {
			return outp, errors.New("Can't find signup or stripe or invite in list")
		}
		return outp, err
	}, 180*time.Second); err != nil {
		return
	}

	// setup rules

	// Adjust rules before we signup into a non admin account
	outp, err := serv.Command().Exec("auth", "create", "rule", "--access=granted", "--scope=''", "--resource=\"service:invite:*\"", "invite")
	if err != nil {
		t.Fatalf("Error setting up rules: %v", outp)
		return
	}

	// Adjust rules before we signup into a non admin account
	outp, err = serv.Command().Exec("auth", "create", "rule", "--access=granted", "--scope=''", "--resource=\"service:signup:*\"", "signup")
	if err != nil {
		t.Fatalf("Error setting up rules: %v", outp)
		return
	}

	// Adjust rules before we signup into a non admin account
	outp, err = serv.Command().Exec("auth", "create", "rule", "--access=granted", "--scope=''", "--resource=\"service:auth:*\"", "auth")
	if err != nil {
		t.Fatalf("Error setting up rules: %v", outp)
		return
	}

	// copy the config with the admin logged in so we can use it for reading logs
	// we dont want to have an open access rule for logs as it's not how it works in live
	confPath := serv.Command().Config
	outp, err = exec.Command("cp", "-rf", confPath, confPath+".admin").CombinedOutput()
	if err != nil {
		t.Fatalf("Error copying config: %v", outp)
		return
	}
}

func logout(serv test.Server, t *test.T) {
	// Log out and switch namespace back to micro
	outp, err := serv.Command().Exec("user", "config", "set", "micro.auth."+serv.Env())
	if err != nil {
		t.Fatal(string(outp))
		return
	}
	outp, err = serv.Command().Exec("user", "config", "set", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatal(string(outp))
		return
	}
}

func testSignupFlow(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)

	// flags
	envFlag := "-e=" + serv.Env()
	confFlag := "-c=" + serv.Command().Config

	email := testEmail(0)

	time.Sleep(5 * time.Second)

	cmd := exec.Command("micro", envFlag, confFlag, "signup")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		outp, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("Expected an error for login but got none")
		} else if !strings.Contains(string(outp), "signup.notallowed") {
			t.Fatal(string(outp))
		}
		wg.Done()
	}()
	go func() {
		time.Sleep(20 * time.Second)
		cmd.Process.Kill()
	}()
	_, err = io.WriteString(stdin, email+"\n")
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	if t.Failed() {
		return
	}

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	// Log out of the admin account to start testing signups
	logout(serv, t)

	password := "PassWord1@"
	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})
	if t.Failed() {
		return
	}
	outp, err := serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	ns := strings.TrimSpace(string(outp))

	if strings.Count(ns, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	t.T().Logf("Namespace set is %v", ns)

	test.Try("Find account", t, func() ([]byte, error) {
		outp, err = serv.Command().Exec("auth", "list", "accounts")
		if err != nil {
			return outp, err
		}
		if !strings.Contains(string(outp), email) {
			return outp, errors.New("Account not found")
		}
		if strings.Contains(string(outp), "default") {
			return outp, errors.New("Default account should not be present in the namespace")
		}
		return outp, nil
	}, 5*time.Second)

	newEmail := testEmail(1)
	newEmail2 := testEmail(2)

	test.Login(serv, t, email, password)

	if err := test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+newEmail, "--namespace="+ns)
	}, 7*time.Second); err != nil {
		t.Fatal(err)
		return
	}
	if err := test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+newEmail2, "--namespace="+ns)
	}, 7*time.Second); err != nil {
		t.Fatal(err)
		return
	}

	logout(serv, t)

	signup(serv, t, newEmail, password, signupOptions{inviterEmail: email, xthInvitee: 1, isInvitedToNamespace: true, shouldJoin: true})
	if t.Failed() {
		return
	}
	outp, err = serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	newNs := strings.TrimSpace(string(outp))
	if newNs != ns {
		t.Fatalf("Namespaces should match, old: %v, new: %v", ns, newNs)
		return
	}

	t.T().Logf("Namespace joined: %v", string(outp))

	logout(serv, t)

	signup(serv, t, newEmail2, password, signupOptions{inviterEmail: email, xthInvitee: 2, isInvitedToNamespace: true, shouldJoin: true})
	if t.Failed() {
		return
	}
	outp, err = serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	newNs = strings.TrimSpace(string(outp))
	if newNs != ns {
		t.Fatalf("Namespaces should match, old: %v, new: %v", ns, newNs)
		return
	}

	t.T().Logf("Namespace joined: %v", string(outp))
}

func TestAdminInvites(t *testing.T) {
	test.TrySuite(t, testAdminInvites, retryCount)
}

func testAdminInvites(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	time.Sleep(2 * time.Second)

	logout(serv, t)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	outp, err := serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	ns := strings.TrimSpace(string(outp))
	if ns == "micro" {
		t.Fatal("SECURITY FLAW: invited user ended up in micro namespace")
	}
	if strings.Count(ns, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	t.T().Logf("Namespace joined: %v", string(outp))
}

func TestAdminInviteNoLimit(t *testing.T) {
	test.TrySuite(t, testAdminInviteNoLimit, retryCount)
}

func testAdminInviteNoLimit(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)

	// Make sure test mod is on otherwise this will spam
	for i := 0; i < 10; i++ {
		test.Try("Send invite", t, func() ([]byte, error) {
			return serv.Command().Exec("invite", "user", "--email="+testEmail(i))
		}, 5*time.Second)
	}
}

func TestUserInviteLimit(t *testing.T) {
	test.TrySuite(t, testUserInviteLimit, retryCount)
}

func testUserInviteLimit(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	logout(serv, t)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	// Make sure test mod is on otherwise this will spam
	for i := 0; i < 5; i++ {
		test.Try("Send invite", t, func() ([]byte, error) {
			return serv.Command().Exec("invite", "user", "--email="+testEmail(i+1))
		}, 5*time.Second)
	}

	outp, err := serv.Command().Exec("invite", "user", "--email="+testEmail(7))
	if err == nil {
		t.Fatalf("Sending 6th invite should fail: %v", outp)
	}
}

func TestUserInviteNoJoin(t *testing.T) {
	test.TrySuite(t, testUserInviteNoJoin, retryCount)
}

func testUserInviteNoJoin(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	logout(serv, t)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	outp, err := serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	ns := strings.TrimSpace(string(outp))
	if strings.Count(ns, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	newEmail := testEmail(1)

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+newEmail)
	}, 5*time.Second)

	logout(serv, t)

	signup(serv, t, newEmail, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	outp, err = serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	newNs := strings.TrimSpace(string(outp))
	if strings.Count(newNs, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	if ns == newNs {
		t.Fatal("User should not have joined invitees namespace")
	}
}

func TestUserInviteJoinDecline(t *testing.T) {
	test.TrySuite(t, testUserInviteJoinDecline, retryCount)
}

func testUserInviteJoinDecline(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	logout(serv, t)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	outp, err := serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	ns := strings.TrimSpace(string(outp))
	if strings.Count(ns, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	newEmail := testEmail(1)

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+newEmail, "--namespace="+ns)
	}, 5*time.Second)

	logout(serv, t)

	signup(serv, t, newEmail, password, signupOptions{isInvitedToNamespace: true, shouldJoin: false})

	outp, err = serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	newNs := strings.TrimSpace(string(outp))
	if strings.Count(newNs, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	if ns == newNs {
		t.Fatal("User should not have joined invitees namespace")
	}
}

func TestUserInviteToNotOwnedNamespace(t *testing.T) {
	test.TrySuite(t, testUserInviteToNotOwnedNamespace, retryCount)
}

func testUserInviteToNotOwnedNamespace(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	logout(serv, t)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	outp, err := serv.Command().Exec("user", "config", "get", "namespaces."+serv.Env()+".current")
	if err != nil {
		t.Fatalf("Error getting namespace: %v", err)
		return
	}
	ns := strings.TrimSpace(string(outp))
	if strings.Count(ns, "-") != 2 {
		t.Fatalf("Expected 2 dashes in namespace but namespace is: %v", ns)
		return
	}

	newEmail := testEmail(1)

	outp, err = serv.Command().Exec("invite", "user", "--email="+newEmail, "--namespace=not-my-namespace")
	if err == nil {
		t.Fatalf("Should not be able to invite to an unowned namespace, output: %v", string(outp))
	}

	// Testing for micro namespace just to be sure as it's the worst case
	outp, err = serv.Command().Exec("invite", "user", "--email="+newEmail, "--namespace=micro")
	if err == nil {
		t.Fatalf("Should not be able to invite to an unowned namespace, output: %v", string(outp))
	}
}

func TestServicesSubscription(t *testing.T) {
	test.TrySuite(t, testServicesSubscription, retryCount)
}

func testServicesSubscription(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	serv.Command().Exec("run", "github.com/micro/services/helloworld")
	serv.Command().Exec("run", "github.com/micro/services/blog/posts")
	serv.Command().Exec("run", "github.com/micro/services/blog/tags")
	serv.Command().Exec("run", "github.com/micro/services/pubsub")

	test.Try("Get changes", t, func() ([]byte, error) {
		outp, err := serv.Command().Exec("status")
		if !strings.Contains(string(outp), "helloworld") || !strings.Contains(string(outp), "posts") || !strings.Contains(string(outp), "posts") ||
			!strings.Contains(string(outp), "tags") || !strings.Contains(string(outp), "pubsub") {
			return outp, errors.New("Can't find services")
		}
		return outp, err
	}, 30*time.Second)

	adminConfFlag := "-c=" + serv.Command().Config + ".admin"
	envFlag := "-e=" + serv.Env()
	exec.Command("micro", envFlag, adminConfFlag, "kill", "billing").CombinedOutput()
	time.Sleep(2 * time.Second)
	exec.Command("micro", envFlag, adminConfFlag, "run", "../../../usage").CombinedOutput()
	time.Sleep(4 * time.Second)
	exec.Command("micro", envFlag, adminConfFlag, "run", "../../../billing").CombinedOutput()

	changeId := ""
	test.Try("Get changes", t, func() ([]byte, error) {
		outp, err := exec.Command("micro", envFlag, adminConfFlag, "billing", "updates").CombinedOutput()
		if err != nil {
			return outp, err
		}
		updatesRsp := map[string]interface{}{}
		err = json.Unmarshal(outp, &updatesRsp)
		if err != nil {
			return outp, err
		}
		updates, ok := updatesRsp["updates"].([]interface{})
		if !ok {
			return outp, errors.New("Unexpected output")
		}
		if len(updates) == 0 {
			return outp, errors.New("No updates found")
		}
		if updates[0].(map[string]interface{})["quantityTo"].(string) != "1" {
			return outp, errors.New("Quantity should be 1")
		}
		changeId = updates[0].(map[string]interface{})["id"].(string)
		if !strings.Contains(string(outp), "Additional services") {
			return outp, errors.New("unexpected output")
		}
		if strings.Contains(string(outp), "Additional users") {
			return outp, errors.New("unexpected output")
		}
		return outp, err
	}, 40*time.Second)

	test.Try("Apply change", t, func() ([]byte, error) {
		return exec.Command("micro", envFlag, adminConfFlag, "billing", "apply", "--id="+changeId).CombinedOutput()
	}, 5*time.Second)

	time.Sleep(4 * time.Second)
	subs := getSubscriptions(t, email)
	priceID := os.Getenv("MICRO_STRIPE_ADDITIONAL_SERVICES_PRICE_ID")
	sub, ok := subs[priceID]
	if !ok {
		t.Fatalf("Sub with id %v not found", priceID)
		return
	}
	if sub.Quantity != 1 {
		t.Fatalf("Quantity should be 1, but it's %v", sub.Quantity)
	}
}

func TestUsersSubscription(t *testing.T) {
	test.TrySuite(t, testUsersSubscription, retryCount)
}

func testUsersSubscription(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)
	password := "PassWord1@"

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)

	signup(serv, t, email, password, signupOptions{isInvitedToNamespace: false, shouldJoin: false})

	serv.Command().Exec("auth", "create", "account", "create", "hi@there.com")

	adminConfFlag := "-c=" + serv.Command().Config + ".admin"
	envFlag := "-e=" + serv.Env()
	exec.Command("micro", envFlag, adminConfFlag, "kill", "billing").CombinedOutput()
	time.Sleep(2 * time.Second)
	exec.Command("micro", envFlag, adminConfFlag, "run", "../../../usage").CombinedOutput()
	time.Sleep(4 * time.Second)
	exec.Command("micro", envFlag, adminConfFlag, "run", "../../../billing").CombinedOutput()

	changeId := ""
	test.Try("Get changes", t, func() ([]byte, error) {
		outp, err := exec.Command("micro", envFlag, adminConfFlag, "billing", "updates").CombinedOutput()
		if err != nil {
			return outp, err
		}
		updatesRsp := map[string]interface{}{}
		err = json.Unmarshal(outp, &updatesRsp)
		if err != nil {
			return outp, err
		}
		updates, ok := updatesRsp["updates"].([]interface{})
		if !ok {
			return outp, errors.New("Unexpected output")
		}
		if len(updates) == 0 {
			return outp, errors.New("No updates found")
		}
		if updates[0].(map[string]interface{})["quantityTo"].(string) != "1" {
			return outp, errors.New("Quantity should be 1")
		}
		changeId = updates[0].(map[string]interface{})["id"].(string)
		if !strings.Contains(string(outp), "Additional users") {
			return outp, errors.New("unexpected output")
		}
		if strings.Contains(string(outp), "Additional services") {
			return outp, errors.New("unexpected output")
		}
		return outp, err
	}, 40*time.Second)

	test.Try("Apply change", t, func() ([]byte, error) {
		return exec.Command("micro", envFlag, adminConfFlag, "billing", "apply", "--id="+changeId).CombinedOutput()
	}, 5*time.Second)

	time.Sleep(4 * time.Second)
	subs := getSubscriptions(t, email)
	priceID := os.Getenv("MICRO_STRIPE_ADDITIONAL_USERS_PRICE_ID")
	sub, ok := subs[priceID]
	if !ok {
		t.Fatalf("Sub with id %v not found", priceID)
		return
	}
	if sub.Quantity != 1 {
		t.Fatalf("Quantity should be 1, but it's %v", sub.Quantity)
	}
}

// returns map witj plan (price) id -> subscriptions
func getSubscriptions(t *test.T, email string) map[string]*stripe.Subscription {
	sc := stripe_client.New(os.Getenv("MICRO_STRIPE_API_KEY"), nil)
	subListParams := &stripe.SubscriptionListParams{}
	subListParams.Limit = stripe.Int64(20)
	subListParams.AddExpand("data.customer")
	iter := sc.Subscriptions.List(subListParams)
	count := 0
	// email -> plan/price id -> subscription
	plans := map[string]*stripe.Subscription{}
	for iter.Next() {
		if count > 20 {
			break
		}
		count++

		c := iter.Subscription()
		if c.Customer.Email == email {
			if c.Plan != nil {
				plans[c.Plan.ID] = c
			}
		}
	}
	return plans
}

type signupOptions struct {
	isInvitedToNamespace bool
	shouldJoin           bool
	inviterEmail         string
	xthInvitee           int
}

func signup(serv test.Server, t *test.T, email, password string, opts signupOptions) {
	envFlag := "-e=" + serv.Env()
	confFlag := "-c=" + serv.Command().Config
	adminConfFlag := "-c=" + serv.Command().Config + ".admin"

	cmd := exec.Command("micro", envFlag, confFlag, "signup", "--password", password)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		outp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(string(outp), err)
			return
		}
		if !strings.Contains(string(outp), signupSuccessString) {
			t.Fatal(string(outp))
			return
		}
	}()
	go func() {
		time.Sleep(40 * time.Second)
		cmd.Process.Kill()
	}()

	time.Sleep(1 * time.Second)
	_, err = io.WriteString(stdin, email+"\n")
	if err != nil {
		t.Fatal(err)
		return
	}

	code := ""
	// careful: there might be multiple codes in the logs
	codes := []string{}
	time.Sleep(2 * time.Second)

	t.Log("looking for code now", email)
	if err := test.Try("Find latest verification token in logs", t, func() ([]byte, error) {
		outp, err := exec.Command("micro", envFlag, adminConfFlag, "logs", "-n", "300", "signup").CombinedOutput()
		if err != nil {
			return outp, err
		}
		if !strings.Contains(string(outp), email) {
			return outp, errors.New("Output does not contain email")
		}
		if !strings.Contains(string(outp), "Sending verification token") {
			return outp, errors.New("Output does not contain expected")
		}
		for _, line := range strings.Split(string(outp), "\n") {
			if strings.Contains(line, "Sending verification token") {
				codes = append(codes, strings.Split(line, "'")[1])
			}
		}
		return outp, nil
	}, 15*time.Second); err != nil {
		return
	}

	if len(codes) == 0 {
		t.Fatal("No code found")
		return
	}
	code = codes[len(codes)-1]

	t.Log("Code is ", code, " for email ", email)
	if code == "" {
		t.Fatal("Code not found")
		return
	}
	_, err = io.WriteString(stdin, code+"\n")
	if err != nil {
		t.Fatal(err)
		return
	}

	if opts.isInvitedToNamespace {
		time.Sleep(3 * time.Second)
		answer := "own"
		if opts.shouldJoin {
			t.Log("Joining a namespace now")
			answer = "join"
		}
		_, err = io.WriteString(stdin, answer+"\n")
		if err != nil {
			t.Fatal(err)
			return
		}
	}

	if !opts.shouldJoin {
		time.Sleep(5 * time.Second)
		sc := stripe_client.New(os.Getenv("MICRO_STRIPE_API_KEY"), nil)
		pm, err := sc.PaymentMethods.New(
			&stripe.PaymentMethodParams{
				Card: &stripe.PaymentMethodCardParams{
					Number:   stripe.String("4242424242424242"),
					ExpMonth: stripe.String("7"),
					ExpYear:  stripe.String("2021"),
					CVC:      stripe.String("314"),
				},
				Type: stripe.String("card"),
			})
		if err != nil {
			t.Fatal(err)
			return
		}

		_, err = io.WriteString(stdin, pm.ID+"\n")
		if err != nil {
			t.Fatal(err)
		}
	}

	// Instead of a sleep shoud probably use Try
	// Some gotchas for this: while the stripe api documentation online
	// shows prices and plans being separate entitires, even v71 version of the
	// go library only has plans. However, it seems like the prices are under plans too.
	time.Sleep(5 * time.Second)

	// Testing if stripe subscriptions exist

	sc := stripe_client.New(os.Getenv("MICRO_STRIPE_API_KEY"), nil)
	subListParams := &stripe.SubscriptionListParams{}
	subListParams.Limit = stripe.Int64(20)
	subListParams.AddExpand("data.customer")
	iter := sc.Subscriptions.List(subListParams)
	count := 0
	// email -> plan/price id -> subscription
	userPlans := map[string]*stripe.Subscription{}
	inviterPlans := map[string]*stripe.Subscription{}
	for iter.Next() {
		if count > 20 {
			break
		}
		count++

		c := iter.Subscription()
		if len(opts.inviterEmail) > 0 && c.Customer.Email == opts.inviterEmail {
			if c.Plan != nil {
				inviterPlans[c.Plan.ID] = c
			}
		}
		if c.Customer.Email == email {
			if c.Plan != nil {
				userPlans[c.Plan.ID] = c
			}
		}
	}

	if opts.shouldJoin {
		priceID := os.Getenv("MICRO_STRIPE_ADDITIONAL_USERS_PRICE_ID")
		sub, found := inviterPlans[priceID]
		if !found {
			t.Fatalf("Subscription with price ID %v not found", priceID)
		}
		if sub.Quantity != int64(opts.xthInvitee) {
			t.Fatalf("Subscription quantity '%v' should match invitee number '%v", sub.Quantity, opts.xthInvitee)
		}
	} else {
		planID := os.Getenv("MICRO_STRIPE_PLAN_ID")
		sub, found := userPlans[planID]
		if !found {
			t.Fatalf("Subscription with plan ID %v not found", planID)
		}
		if sub.Quantity != 1 {
			t.Fatalf("Subscription quantity should be 1 but is %v", sub.Quantity)
		}
	}

	test.Try("Check customer marked active", t, func() ([]byte, error) {
		outp, err := exec.Command("micro", envFlag, adminConfFlag, "customers", "read", "--id="+email).CombinedOutput()
		if err != nil {
			return outp, err
		}
		if !strings.Contains(string(outp), `"status": "active"`) {
			return outp, fmt.Errorf("Customer status is not active")
		}
		return nil, nil
	}, 60*time.Second)

	// Don't wait if a test is already failed, this is a quirk of the
	// test framework @todo fix this quirk
	if t.Failed() {
		return
	}
	wg.Wait()
}

func getSrcString(envvar, dflt string) string {
	if env := os.Getenv(envvar); env != "" {
		return env
	}
	return dflt
}

func TestDuplicateInvites(t *testing.T) {
	test.TrySuite(t, testDuplicateInvites, retryCount)
}

func testDuplicateInvites(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)
	email := testEmail(0)

	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)
	test.Try("Send invite again", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email)
	}, 5*time.Second)
	outp, err := serv.Command().Exec("logs", "invite")
	if err != nil {
		t.Fatalf("Unexpected error retrieving logs %s", err)
	}
	if !strings.Contains(string(outp), "Invite already sent for user "+email) {
		t.Fatalf("Invite was sent multiple times")
	}

	// test a force resend
	email2 := testEmail(1)
	test.Try("Send invite", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email2)
	}, 5*time.Second)
	test.Try("Send invite again", t, func() ([]byte, error) {
		return serv.Command().Exec("invite", "user", "--email="+email2, "--resend=true")
	}, 5*time.Second)
	outp, err = serv.Command().Exec("logs", "invite")
	if err != nil {
		t.Fatalf("Unexpected error retrieving logs %s", err)
	}
	if strings.Contains(string(outp), "Invite already sent for user "+email2) {
		t.Fatalf("Invite should have been sent multiple times but was blocked")
	}

}

func TestInviteEmailValidation(t *testing.T) {
	test.TrySuite(t, testInviteEmailValidation, retryCount)
}

func testInviteEmailValidation(t *test.T) {
	t.Parallel()

	serv := test.NewServer(t, test.WithLogin())
	defer serv.Close()
	if err := serv.Run(); err != nil {
		return
	}

	setupM3Tests(serv, t)

	outp, _ := serv.Command().Exec("invite", "user", "--email=notanemail.com")
	if !strings.Contains(string(outp), "400") {
		t.Fatalf("Expected a 400 bad request error %s", string(outp))
	}

}
