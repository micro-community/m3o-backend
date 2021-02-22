// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/quota.proto

package quota

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Quota service

func NewQuotaEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Quota service

type QuotaService interface {
	// Creates a quota
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	// Registers a user against a quota
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...client.CallOption) (*RegisterUserResponse, error)
	// Lists current quota usage
	ListUsage(ctx context.Context, in *ListUsageRequest, opts ...client.CallOption) (*ListUsageResponse, error)
	// Debug
	ResetQuotas(ctx context.Context, in *ResetRequest, opts ...client.CallOption) (*ResetResponse, error)
	// Lists quotas
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type quotaService struct {
	c    client.Client
	name string
}

func NewQuotaService(name string, c client.Client) QuotaService {
	return &quotaService{
		c:    c,
		name: name,
	}
}

func (c *quotaService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Quota.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaService) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...client.CallOption) (*RegisterUserResponse, error) {
	req := c.c.NewRequest(c.name, "Quota.RegisterUser", in)
	out := new(RegisterUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaService) ListUsage(ctx context.Context, in *ListUsageRequest, opts ...client.CallOption) (*ListUsageResponse, error) {
	req := c.c.NewRequest(c.name, "Quota.ListUsage", in)
	out := new(ListUsageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaService) ResetQuotas(ctx context.Context, in *ResetRequest, opts ...client.CallOption) (*ResetResponse, error) {
	req := c.c.NewRequest(c.name, "Quota.ResetQuotas", in)
	out := new(ResetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Quota.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Quota service

type QuotaHandler interface {
	// Creates a quota
	Create(context.Context, *CreateRequest, *CreateResponse) error
	// Registers a user against a quota
	RegisterUser(context.Context, *RegisterUserRequest, *RegisterUserResponse) error
	// Lists current quota usage
	ListUsage(context.Context, *ListUsageRequest, *ListUsageResponse) error
	// Debug
	ResetQuotas(context.Context, *ResetRequest, *ResetResponse) error
	// Lists quotas
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterQuotaHandler(s server.Server, hdlr QuotaHandler, opts ...server.HandlerOption) error {
	type quota interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		RegisterUser(ctx context.Context, in *RegisterUserRequest, out *RegisterUserResponse) error
		ListUsage(ctx context.Context, in *ListUsageRequest, out *ListUsageResponse) error
		ResetQuotas(ctx context.Context, in *ResetRequest, out *ResetResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Quota struct {
		quota
	}
	h := &quotaHandler{hdlr}
	return s.Handle(s.NewHandler(&Quota{h}, opts...))
}

type quotaHandler struct {
	QuotaHandler
}

func (h *quotaHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.QuotaHandler.Create(ctx, in, out)
}

func (h *quotaHandler) RegisterUser(ctx context.Context, in *RegisterUserRequest, out *RegisterUserResponse) error {
	return h.QuotaHandler.RegisterUser(ctx, in, out)
}

func (h *quotaHandler) ListUsage(ctx context.Context, in *ListUsageRequest, out *ListUsageResponse) error {
	return h.QuotaHandler.ListUsage(ctx, in, out)
}

func (h *quotaHandler) ResetQuotas(ctx context.Context, in *ResetRequest, out *ResetResponse) error {
	return h.QuotaHandler.ResetQuotas(ctx, in, out)
}

func (h *quotaHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.QuotaHandler.List(ctx, in, out)
}
