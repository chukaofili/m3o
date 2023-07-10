// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/admin.proto

package vehicle

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

// Api Endpoints for VehicleAdmin service

func NewVehicleAdminEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VehicleAdmin service

type VehicleAdminService interface {
	SetLogo(ctx context.Context, in *SetLogoRequest, opts ...client.CallOption) (*SetLogoResponse, error)
}

type vehicleAdminService struct {
	c    client.Client
	name string
}

func NewVehicleAdminService(name string, c client.Client) VehicleAdminService {
	return &vehicleAdminService{
		c:    c,
		name: name,
	}
}

func (c *vehicleAdminService) SetLogo(ctx context.Context, in *SetLogoRequest, opts ...client.CallOption) (*SetLogoResponse, error) {
	req := c.c.NewRequest(c.name, "VehicleAdmin.SetLogo", in)
	out := new(SetLogoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VehicleAdmin service

type VehicleAdminHandler interface {
	SetLogo(context.Context, *SetLogoRequest, *SetLogoResponse) error
}

func RegisterVehicleAdminHandler(s server.Server, hdlr VehicleAdminHandler, opts ...server.HandlerOption) error {
	type vehicleAdmin interface {
		SetLogo(ctx context.Context, in *SetLogoRequest, out *SetLogoResponse) error
	}
	type VehicleAdmin struct {
		vehicleAdmin
	}
	h := &vehicleAdminHandler{hdlr}
	return s.Handle(s.NewHandler(&VehicleAdmin{h}, opts...))
}

type vehicleAdminHandler struct {
	VehicleAdminHandler
}

func (h *vehicleAdminHandler) SetLogo(ctx context.Context, in *SetLogoRequest, out *SetLogoResponse) error {
	return h.VehicleAdminHandler.SetLogo(ctx, in, out)
}
