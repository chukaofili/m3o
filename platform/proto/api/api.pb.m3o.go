// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "m3o.dev/platform/service/api"
	client "m3o.dev/platform/service/client"
	server "m3o.dev/platform/service/server"
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

// Api Endpoints for Api service

func NewApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Api service

type ApiService interface {
	AddToBlockList(ctx context.Context, in *AddToBlockListRequest, opts ...client.CallOption) (*AddToBlockListResponse, error)
	RemoveFromBlockList(ctx context.Context, in *RemoveFromBlockListRequest, opts ...client.CallOption) (*RemoveFromBlockListResponse, error)
	ReadBlockList(ctx context.Context, in *ReadBlockListRequest, opts ...client.CallOption) (*ReadBlockListResponse, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) AddToBlockList(ctx context.Context, in *AddToBlockListRequest, opts ...client.CallOption) (*AddToBlockListResponse, error) {
	req := c.c.NewRequest(c.name, "Api.AddToBlockList", in)
	out := new(AddToBlockListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) RemoveFromBlockList(ctx context.Context, in *RemoveFromBlockListRequest, opts ...client.CallOption) (*RemoveFromBlockListResponse, error) {
	req := c.c.NewRequest(c.name, "Api.RemoveFromBlockList", in)
	out := new(RemoveFromBlockListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) ReadBlockList(ctx context.Context, in *ReadBlockListRequest, opts ...client.CallOption) (*ReadBlockListResponse, error) {
	req := c.c.NewRequest(c.name, "Api.ReadBlockList", in)
	out := new(ReadBlockListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	AddToBlockList(context.Context, *AddToBlockListRequest, *AddToBlockListResponse) error
	RemoveFromBlockList(context.Context, *RemoveFromBlockListRequest, *RemoveFromBlockListResponse) error
	ReadBlockList(context.Context, *ReadBlockListRequest, *ReadBlockListResponse) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) error {
	type api interface {
		AddToBlockList(ctx context.Context, in *AddToBlockListRequest, out *AddToBlockListResponse) error
		RemoveFromBlockList(ctx context.Context, in *RemoveFromBlockListRequest, out *RemoveFromBlockListResponse) error
		ReadBlockList(ctx context.Context, in *ReadBlockListRequest, out *ReadBlockListResponse) error
	}
	type Api struct {
		api
	}
	h := &apiHandler{hdlr}
	return s.Handle(s.NewHandler(&Api{h}, opts...))
}

type apiHandler struct {
	ApiHandler
}

func (h *apiHandler) AddToBlockList(ctx context.Context, in *AddToBlockListRequest, out *AddToBlockListResponse) error {
	return h.ApiHandler.AddToBlockList(ctx, in, out)
}

func (h *apiHandler) RemoveFromBlockList(ctx context.Context, in *RemoveFromBlockListRequest, out *RemoveFromBlockListResponse) error {
	return h.ApiHandler.RemoveFromBlockList(ctx, in, out)
}

func (h *apiHandler) ReadBlockList(ctx context.Context, in *ReadBlockListRequest, out *ReadBlockListResponse) error {
	return h.ApiHandler.ReadBlockList(ctx, in, out)
}