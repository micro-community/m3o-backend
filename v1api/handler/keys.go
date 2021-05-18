package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	m3oauth "github.com/m3o/services/pkg/auth"
	v1api "github.com/m3o/services/v1api/proto"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/events"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
)

type keyStatus string

const (
	keyStatusActive  = "active"
	keyStatusBlocked = "blocked" // blocked - probably because they've run out of money
)

type apiKeyRecord struct {
	ID            string    `json:"id"`            // id of the key
	ApiKey        string    `json:"apiKey"`        // hashed api key
	Scopes        []string  `json:"scopes"`        // the scopes this key has granted
	UserID        string    `json:"userID"`        // the ID of the key's owner
	AccID         string    `json:"accID"`         // the ID of the service account
	Description   string    `json:"description"`   // optional description of the API key as given by user
	Namespace     string    `json:"namespace"`     // the namespace that this user belongs to (only because technically user IDs aren't globally unique)
	Token         string    `json:"token"`         // the short lived JWT token
	Created       int64     `json:"created"`       // creation time
	Status        keyStatus `json:"status"`        // status of the key
	StatusMessage string    `json:"statusMessage"` // message to go with the status updated, presented to users
}

// GenerateKey generates an API key
func (v1 *V1) GenerateKey(ctx context.Context, req *v1api.GenerateKeyRequest, rsp *v1api.GenerateKeyResponse) error {
	if len(req.Scopes) == 0 {
		return errors.BadRequest("v1api.generate", "Missing scopes field")
	}
	if len(req.Description) == 0 {
		return errors.BadRequest("v1api.generate", "Missing description field")
	}

	acc, err := m3oauth.VerifyMicroCustomer(ctx, "v1api.generate")
	if err != nil {
		return err
	}
	// are they allowed to generate with the requested scopes?
	if !v1.checkRequestedScopes(ctx, req.Scopes) {
		return errors.Forbidden("v1api.generate", "Not allowed to generate a key with requested scopes")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return errors.InternalServerError("v1api.generate", "Failed to generate api key")
	}

	apiKey := base64.StdEncoding.EncodeToString([]byte(id.String()))
	hashedKey, err := hashSecret(apiKey)
	if err != nil {
		log.Errorf("Error hashing api key %s", err)
		return errors.InternalServerError("v1api.generate", "Failed to generate api key")
	}

	// api key is the secret for a new account
	// generate the new account + short lived access token for it
	authAcc, err := auth.Generate(
		uuid.New().String(),
		auth.WithSecret(apiKey),
		auth.WithIssuer(acc.Issuer),
		auth.WithType("apikey"),
		auth.WithScopes(req.Scopes...),
		auth.WithMetadata(map[string]string{"apikey_owner": acc.ID}),
	)
	if err != nil {
		log.Errorf("Error generating auth account %s", err)
		return errors.InternalServerError("v1api.generate", "Failed to generate api key")
	}
	tok, err := auth.Token(
		auth.WithCredentials(authAcc.ID, apiKey),
		auth.WithTokenIssuer(acc.Issuer),
		auth.WithExpiry(1*time.Hour))
	if err != nil {
		log.Errorf("Error generating token %s", err)
		return errors.InternalServerError("v1api.generate", "Failed to generate api key")
	}
	// hash API key and store with scopes
	rec := apiKeyRecord{
		ID:          uuid.New().String(),
		ApiKey:      hashedKey,
		Scopes:      req.Scopes,
		UserID:      acc.ID,
		Namespace:   acc.Issuer,
		Description: req.Description,
		AccID:       authAcc.ID,
		Token:       tok.AccessToken,
		Created:     time.Now().Unix(),
		Status:      keyStatusBlocked,
	}
	if err := v1.writeAPIRecord(&rec); err != nil {
		log.Errorf("Failed to write api record %s", err)
		return errors.InternalServerError("v1api.generate", "Failed to generate api key")
	}

	if err := events.Publish("v1api", v1api.Event{Type: "APIKeyCreate",
		ApiKeyCreate: &v1api.APIKeyCreateEvent{
			UserId:    rec.UserID,
			Namespace: rec.Namespace,
			ApiKeyId:  rec.ID,
			Scopes:    rec.Scopes,
		}}); err != nil {
		log.Errorf("Error publishing event %s", err)
	}
	// return the unhashed key
	rsp.ApiKey = apiKey
	return nil
}

// ListKeys lists all keys for a user
func (v1 *V1) ListKeys(ctx context.Context, req *v1api.ListRequest, rsp *v1api.ListResponse) error {
	// Check account
	acc, err := m3oauth.VerifyMicroCustomer(ctx, "v1api.listkeys")
	if err != nil {
		return err
	}
	recs, err := listKeysForUser(acc.Issuer, acc.ID)
	if err != nil {
		log.Errorf("Error listing keys %s", err)
		return errors.InternalServerError("v1aapi.listkeys", "Error listing keys")
	}
	rsp.ApiKeys = make([]*v1api.APIKey, len(recs))
	for i, apiRec := range recs {
		rsp.ApiKeys[i] = &v1api.APIKey{
			Id:          apiRec.ID,
			Description: apiRec.Description,
			CreatedTime: apiRec.Created,
			Scopes:      apiRec.Scopes,
		}
	}
	return nil
}

func listKeysForUser(ns, userID string) ([]*apiKeyRecord, error) {
	recs, err := store.Read(fmt.Sprintf("%s:%s:%s:", storePrefixUserID, ns, userID), store.ReadPrefix())
	if err != nil {
		if err == store.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	ret := make([]*apiKeyRecord, len(recs))
	for i, rec := range recs {
		apiRec := &apiKeyRecord{}
		if err := json.Unmarshal(rec.Value, apiRec); err != nil {
			return nil, err
		}
		ret[i] = apiRec
	}
	return ret, nil
}

func (v1 *V1) RevokeKey(ctx context.Context, request *v1api.RevokeRequest, response *v1api.RevokeResponse) error {
	if len(request.Id) == 0 {
		return errors.BadRequest("v1api.Revoke", "Missing ID field")
	}
	acc, err := m3oauth.VerifyMicroCustomer(ctx, "v1api.Revoke")
	if err != nil {
		return err
	}

	rec, err := readAPIRecordByKeyID(acc.Issuer, acc.ID, request.Id)
	if err != nil {
		if err == store.ErrNotFound {
			return errors.NotFound("v1api.Revoke", "Not found")
		}
		log.Errorf("Error reading API key record %s", err)
		return errors.InternalServerError("v1pi.Revoke", "Error revoking key")
	}
	if err := v1.deleteAPIRecord(rec); err != nil {
		log.Errorf("Error deleting API key record %s", err)
		return errors.InternalServerError("v1pi.Revoke", "Error revoking key")
	}

	if err := events.Publish("v1api", v1api.Event{Type: "APIKeyRevoke",
		ApiKeyRevoke: &v1api.APIKeyRevokeEvent{
			UserId:    acc.ID,
			Namespace: acc.Issuer,
			ApiKeyId:  request.Id,
		}}); err != nil {
		log.Errorf("Error publishing event %s", err)
	}

	return nil
}

func (v1 *V1) BlockKey(ctx context.Context, request *v1api.BlockKeyRequest, response *v1api.BlockKeyResponse) error {
	return v1.updateKeyStatus(ctx, "v1api.BlockKey", request.Namespace, request.UserId, request.KeyId, keyStatusBlocked, request.Message)
}

func (v1 *V1) UnblockKey(ctx context.Context, request *v1api.UnblockKeyRequest, response *v1api.UnblockKeyResponse) error {
	return v1.updateKeyStatus(ctx, "v1api.UnblockKey", request.Namespace, request.UserId, request.KeyId, keyStatusActive, "")
}

func (v1 *V1) updateKeyStatus(ctx context.Context, methodName, ns, userID, keyID string, status keyStatus, statusMessage string) error {

	if _, err := m3oauth.VerifyMicroAdmin(ctx, methodName); err != nil {
		return err
	}

	var keys []*apiKeyRecord
	if len(keyID) > 0 {
		rec, err := readAPIRecordByKeyID(ns, userID, keyID)
		if err != nil {
			log.Errorf("Error reading key %s", err)
			return errors.InternalServerError(methodName, "Error updating key")
		}
		keys = []*apiKeyRecord{rec}
	} else {
		recs, err := listKeysForUser(ns, userID)
		if err != nil {
			log.Errorf("Error listing keys %s", err)
			return errors.InternalServerError(methodName, "Error updating key")
		}
		keys = recs
	}
	for _, k := range keys {
		k.Status = status
		k.StatusMessage = statusMessage
		if err := v1.writeAPIRecord(k); err != nil {
			log.Errorf("Error updating api key record %s", err)
			return errors.InternalServerError(methodName, "Error updating key")
		}
	}

	return nil
}
