package provider

import (
	"context"
	"testing"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/client"

	pb "github.com/m3o/services/payments/provider/proto"
)

type testprovider struct{}

func (t testprovider) DeleteCustomer(ctx context.Context, request *pb.DeleteCustomerRequest, response *pb.DeleteCustomerResponse) error {
	panic("implement me")
}

func (t testprovider) UpdateSubscription(ctx context.Context, request *pb.UpdateSubscriptionRequest, response *pb.UpdateSubscriptionResponse) error {
	panic("implement me")
}

func (t testprovider) VerifyPaymentMethod(ctx context.Context, request *pb.VerifyPaymentMethodRequest, response *pb.VerifyPaymentMethodResponse) error {
	panic("implement me")
}

func (t testprovider) CancelSubscription(ctx context.Context, request *pb.CancelSubscriptionRequest, response *pb.CancelSubscriptionResponse) error {
	panic("implement me")
}

func (t testprovider) CreateProduct(ctx context.Context, req *pb.CreateProductRequest, rsp *pb.CreateProductResponse) error {
	return nil
}
func (t testprovider) CreatePlan(ctx context.Context, req *pb.CreatePlanRequest, rsp *pb.CreatePlanResponse) error {
	return nil
}
func (t testprovider) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest, rsp *pb.CreateCustomerResponse) error {
	return nil
}
func (t testprovider) CreateSubscription(ctx context.Context, req *pb.CreateSubscriptionRequest, rsp *pb.CreateSubscriptionResponse) error {
	return nil
}
func (t testprovider) CreatePaymentMethod(ctx context.Context, req *pb.CreatePaymentMethodRequest, rsp *pb.CreatePaymentMethodResponse) error {
	return nil
}
func (t testprovider) DeletePaymentMethod(ctx context.Context, req *pb.DeletePaymentMethodRequest, rsp *pb.DeletePaymentMethodResponse) error {
	return nil
}
func (t testprovider) ListPaymentMethods(ctx context.Context, req *pb.ListPaymentMethodsRequest, rsp *pb.ListPaymentMethodsResponse) error {
	return nil
}
func (t testprovider) ListPlans(ctx context.Context, req *pb.ListPlansRequest, rsp *pb.ListPlansResponse) error {
	return nil
}
func (t testprovider) ListSubscriptions(ctx context.Context, req *pb.ListSubscriptionsRequest, rsp *pb.ListSubscriptionsResponse) error {
	return nil
}
func (t testprovider) SetDefaultPaymentMethod(ctx context.Context, req *pb.SetDefaultPaymentMethodRequest, rsp *pb.SetDefaultPaymentMethodResponse) error {
	return nil
}

func TestNewProvider(t *testing.T) {
	// test the provider returns ErrNotFound when not registered
	t.Run("no provider set", func(t *testing.T) {
		_, err := NewProvider("test", client.DefaultClient)
		if err != ErrNotFound {
			t.Errorf("Expected ErrNotFound, got %v", err)
		}
	})

	// test the provider returns a provider when one is registered
	t.Run("provider set", func(t *testing.T) {
		testSrv := service.New(service.Name(ServicePrefix + "test"))
		if err := pb.RegisterProviderHandler(testSrv.Server(), new(testprovider)); err != nil {
			t.Fatalf("Error registering test handler: %v", err)
		}
		go testSrv.Run()

		// TODO: Find way of improving this test so the delay is not hardcoded
		// and the testSrv is stopped at the end of the function
		time.Sleep(200 * time.Millisecond)

		_, err := NewProvider("test", testSrv.Client())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
