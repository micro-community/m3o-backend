// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/event.proto

package go_micro_api_events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
	microClient "github.com/micro/micro/v3/service/client"
	microServer "github.com/micro/micro/v3/service/server"
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
var _ = microServer.Handle
var _ = microClient.Call

// Api Endpoints for Events service

func NewEventsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Events service

type EventsService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
}

type eventsService struct {
	name string
}

func NewEventsService(name string) EventsService {
	return &eventsService{name: name}
}

func (c *eventsService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := microClient.NewRequest(c.name, "Events.Create", in)
	out := new(CreateResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Events service

type EventsHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
}

func RegisterEventsHandler(hdlr EventsHandler, opts ...server.HandlerOption) error {
	type events interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
	}
	type Events struct {
		events
	}
	h := &eventsHandler{hdlr}
	return microServer.Handle(microServer.NewHandler(&Events{h}, opts...))
}

type eventsHandler struct {
	EventsHandler
}

func (h *eventsHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.EventsHandler.Create(ctx, in, out)
}
