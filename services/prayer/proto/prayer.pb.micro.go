// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/prayer.proto

package prayer

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

// Api Endpoints for Prayer service

func NewPrayerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Prayer service

type PrayerService interface {
	Times(ctx context.Context, in *TimesRequest, opts ...client.CallOption) (*TimesResponse, error)
}

type prayerService struct {
	c    client.Client
	name string
}

func NewPrayerService(name string, c client.Client) PrayerService {
	return &prayerService{
		c:    c,
		name: name,
	}
}

func (c *prayerService) Times(ctx context.Context, in *TimesRequest, opts ...client.CallOption) (*TimesResponse, error) {
	req := c.c.NewRequest(c.name, "Prayer.Times", in)
	out := new(TimesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Prayer service

type PrayerHandler interface {
	Times(context.Context, *TimesRequest, *TimesResponse) error
}

func RegisterPrayerHandler(s server.Server, hdlr PrayerHandler, opts ...server.HandlerOption) error {
	type prayer interface {
		Times(ctx context.Context, in *TimesRequest, out *TimesResponse) error
	}
	type Prayer struct {
		prayer
	}
	h := &prayerHandler{hdlr}
	return s.Handle(s.NewHandler(&Prayer{h}, opts...))
}

type prayerHandler struct {
	PrayerHandler
}

func (h *prayerHandler) Times(ctx context.Context, in *TimesRequest, out *TimesResponse) error {
	return h.PrayerHandler.Times(ctx, in, out)
}
