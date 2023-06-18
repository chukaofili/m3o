// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/carbon.proto

package carbon

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

// Api Endpoints for Carbon service

func NewCarbonEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Carbon service

type CarbonService interface {
	Offset(ctx context.Context, in *OffsetRequest, opts ...client.CallOption) (*OffsetResponse, error)
}

type carbonService struct {
	c    client.Client
	name string
}

func NewCarbonService(name string, c client.Client) CarbonService {
	return &carbonService{
		c:    c,
		name: name,
	}
}

func (c *carbonService) Offset(ctx context.Context, in *OffsetRequest, opts ...client.CallOption) (*OffsetResponse, error) {
	req := c.c.NewRequest(c.name, "Carbon.Offset", in)
	out := new(OffsetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Carbon service

type CarbonHandler interface {
	Offset(context.Context, *OffsetRequest, *OffsetResponse) error
}

func RegisterCarbonHandler(s server.Server, hdlr CarbonHandler, opts ...server.HandlerOption) error {
	type carbon interface {
		Offset(ctx context.Context, in *OffsetRequest, out *OffsetResponse) error
	}
	type Carbon struct {
		carbon
	}
	h := &carbonHandler{hdlr}
	return s.Handle(s.NewHandler(&Carbon{h}, opts...))
}

type carbonHandler struct {
	CarbonHandler
}

func (h *carbonHandler) Offset(ctx context.Context, in *OffsetRequest, out *OffsetResponse) error {
	return h.CarbonHandler.Offset(ctx, in, out)
}
