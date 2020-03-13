// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/login/login.proto

package go_micro_srv_login

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Login service

type LoginService interface {
	CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...client.CallOption) (*CreateLoginResponse, error)
	VerifyLogin(ctx context.Context, in *VerifyLoginRequest, opts ...client.CallOption) (*VerifyLoginResponse, error)
	// TODO: Remove the RPC and replace with consuming users update events
	// once we have update diff implemented.
	UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...client.CallOption) (*UpdateEmailResponse, error)
}

type loginService struct {
	c    client.Client
	name string
}

func NewLoginService(name string, c client.Client) LoginService {
	return &loginService{
		c:    c,
		name: name,
	}
}

func (c *loginService) CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...client.CallOption) (*CreateLoginResponse, error) {
	req := c.c.NewRequest(c.name, "Login.CreateLogin", in)
	out := new(CreateLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginService) VerifyLogin(ctx context.Context, in *VerifyLoginRequest, opts ...client.CallOption) (*VerifyLoginResponse, error) {
	req := c.c.NewRequest(c.name, "Login.VerifyLogin", in)
	out := new(VerifyLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginService) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...client.CallOption) (*UpdateEmailResponse, error) {
	req := c.c.NewRequest(c.name, "Login.UpdateEmail", in)
	out := new(UpdateEmailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginHandler interface {
	CreateLogin(context.Context, *CreateLoginRequest, *CreateLoginResponse) error
	VerifyLogin(context.Context, *VerifyLoginRequest, *VerifyLoginResponse) error
	// TODO: Remove the RPC and replace with consuming users update events
	// once we have update diff implemented.
	UpdateEmail(context.Context, *UpdateEmailRequest, *UpdateEmailResponse) error
}

func RegisterLoginHandler(s server.Server, hdlr LoginHandler, opts ...server.HandlerOption) error {
	type login interface {
		CreateLogin(ctx context.Context, in *CreateLoginRequest, out *CreateLoginResponse) error
		VerifyLogin(ctx context.Context, in *VerifyLoginRequest, out *VerifyLoginResponse) error
		UpdateEmail(ctx context.Context, in *UpdateEmailRequest, out *UpdateEmailResponse) error
	}
	type Login struct {
		login
	}
	h := &loginHandler{hdlr}
	return s.Handle(s.NewHandler(&Login{h}, opts...))
}

type loginHandler struct {
	LoginHandler
}

func (h *loginHandler) CreateLogin(ctx context.Context, in *CreateLoginRequest, out *CreateLoginResponse) error {
	return h.LoginHandler.CreateLogin(ctx, in, out)
}

func (h *loginHandler) VerifyLogin(ctx context.Context, in *VerifyLoginRequest, out *VerifyLoginResponse) error {
	return h.LoginHandler.VerifyLogin(ctx, in, out)
}

func (h *loginHandler) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, out *UpdateEmailResponse) error {
	return h.LoginHandler.UpdateEmail(ctx, in, out)
}