package handler

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	"github.com/micro/go-micro/v3/errors"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro/v3/store"
	"github.com/micro/micro/v3/service"
	mconfig "github.com/micro/micro/v3/service/config"
	mstore "github.com/micro/micro/v3/service/store"

	pb "github.com/m3o/services/secrets/proto"
)

// New returns an initialised handler
func New(srv *service.Service) *Secrets {
	// todo: debug why explicity setting service name is required
	secret := mconfig.Get("go", "micro", "service", "secrets", "secret").String("")
	if len(secret) == 0 {
		logger.Fatal("Missing required config: secret")
	}

	return &Secrets{
		secret: secret,
		name:   srv.Name(),
	}
}

// Secrets implements the secrets service interface
type Secrets struct {
	name   string
	secret string
}

// pathJoiner is the character used to join the path when writing to the store
const pathJoiner = "/"

// Create a secret
func (h *Secrets) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	// validate the request
	if req.Path == nil || len(req.Path) == 0 {
		return errors.BadRequest(h.name, "Missing path")
	}
	if len(req.Value) == 0 {
		return errors.BadRequest(h.name, "Missing value")
	}

	// encode the secret
	secret, err := h.encrypt(req.Value)
	if err != nil {
		return errors.InternalServerError(h.name, "Error encrypting secret: %v", err)
	}

	// write to the store
	key := strings.Join(req.Path, pathJoiner)
	if err := mstore.Write(&store.Record{Key: key, Value: secret}); err != nil {
		return errors.InternalServerError(h.name, "Error writing to the store: %v", err)
	}

	return nil
}

// Read a secret
func (h *Secrets) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	// validate the request
	if req.Path == nil || len(req.Path) == 0 {
		return errors.BadRequest(h.name, "Missing path")
	}

	// read from the store
	recs, err := mstore.Read(strings.Join(req.Path, pathJoiner))
	if err == store.ErrNotFound {
		return errors.NotFound(h.name, "Secret not found")
	} else if err != nil {
		return errors.InternalServerError(h.name, "Error reading from the store: %v", err)
	}

	// decrypt the secret
	secret, err := h.decrypt(recs[0].Value)
	if err != nil {
		return errors.InternalServerError(h.name, "Error decrypting secret: %v", err)
	}

	rsp.Value = secret
	return nil
}

// Update a secret
func (h *Secrets) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UpdateResponse) error {
	// validate the request
	if req.Path == nil || len(req.Path) == 0 {
		return errors.BadRequest(h.name, "Missing path")
	}
	if len(req.Value) == 0 {
		return errors.BadRequest(h.name, "Missing value")
	}

	// encode the secret
	secret, err := h.encrypt(req.Value)
	if err != nil {
		return errors.InternalServerError(h.name, "Error encrypting secret: %v", err)
	}

	// write to the store
	key := strings.Join(req.Path, pathJoiner)
	if err := mstore.Write(&store.Record{Key: key, Value: secret}); err != nil {
		return errors.InternalServerError(h.name, "Error writing to the store: %v", err)
	}

	return nil
}

// Delete a secret
func (h *Secrets) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	// validate the request
	if req.Path == nil || len(req.Path) == 0 {
		return errors.BadRequest(h.name, "Missing path")
	}

	// deletre from the store
	key := strings.Join(req.Path, pathJoiner)
	err := mstore.Delete(key)
	if err == store.ErrNotFound {
		return errors.NotFound(h.name, "Secret not found")
	} else if err != nil {
		return errors.InternalServerError(h.name, "Error reading from the store: %v", err)
	}

	return nil
}

func (h *Secrets) encrypt(text string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(h.secret))
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func (h *Secrets) decrypt(text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(h.secret))
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.InternalServerError(h.name, "Ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	return string(data), nil
}
