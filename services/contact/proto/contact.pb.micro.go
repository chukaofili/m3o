// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/contact.proto

package contact

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

// Api Endpoints for Contact service

func NewContactEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Contact service

type ContactService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type contactService struct {
	c    client.Client
	name string
}

func NewContactService(name string, c client.Client) ContactService {
	return &contactService{
		c:    c,
		name: name,
	}
}

func (c *contactService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Contact.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Contact.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Contact.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Contact.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Contact.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contact service

type ContactHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterContactHandler(s server.Server, hdlr ContactHandler, opts ...server.HandlerOption) error {
	type contact interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Contact struct {
		contact
	}
	h := &contactHandler{hdlr}
	return s.Handle(s.NewHandler(&Contact{h}, opts...))
}

type contactHandler struct {
	ContactHandler
}

func (h *contactHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ContactHandler.Create(ctx, in, out)
}

func (h *contactHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ContactHandler.Read(ctx, in, out)
}

func (h *contactHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ContactHandler.Update(ctx, in, out)
}

func (h *contactHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ContactHandler.Delete(ctx, in, out)
}

func (h *contactHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ContactHandler.List(ctx, in, out)
}
