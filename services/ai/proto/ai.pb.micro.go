// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ai.proto

package ai

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

// Api Endpoints for Ai service

func NewAiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Ai service

type AiService interface {
	Complete(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*CompleteResponse, error)
	Edit(ctx context.Context, in *EditRequest, opts ...client.CallOption) (*EditResponse, error)
	Moderate(ctx context.Context, in *ModerateRequest, opts ...client.CallOption) (*ModerateResponse, error)
	Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error)
}

type aiService struct {
	c    client.Client
	name string
}

func NewAiService(name string, c client.Client) AiService {
	return &aiService{
		c:    c,
		name: name,
	}
}

func (c *aiService) Complete(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*CompleteResponse, error) {
	req := c.c.NewRequest(c.name, "Ai.Complete", in)
	out := new(CompleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiService) Edit(ctx context.Context, in *EditRequest, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "Ai.Edit", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiService) Moderate(ctx context.Context, in *ModerateRequest, opts ...client.CallOption) (*ModerateResponse, error) {
	req := c.c.NewRequest(c.name, "Ai.Moderate", in)
	out := new(ModerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiService) Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error) {
	req := c.c.NewRequest(c.name, "Ai.Generate", in)
	out := new(GenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ai service

type AiHandler interface {
	Complete(context.Context, *CompleteRequest, *CompleteResponse) error
	Edit(context.Context, *EditRequest, *EditResponse) error
	Moderate(context.Context, *ModerateRequest, *ModerateResponse) error
	Generate(context.Context, *GenerateRequest, *GenerateResponse) error
}

func RegisterAiHandler(s server.Server, hdlr AiHandler, opts ...server.HandlerOption) error {
	type ai interface {
		Complete(ctx context.Context, in *CompleteRequest, out *CompleteResponse) error
		Edit(ctx context.Context, in *EditRequest, out *EditResponse) error
		Moderate(ctx context.Context, in *ModerateRequest, out *ModerateResponse) error
		Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error
	}
	type Ai struct {
		ai
	}
	h := &aiHandler{hdlr}
	return s.Handle(s.NewHandler(&Ai{h}, opts...))
}

type aiHandler struct {
	AiHandler
}

func (h *aiHandler) Complete(ctx context.Context, in *CompleteRequest, out *CompleteResponse) error {
	return h.AiHandler.Complete(ctx, in, out)
}

func (h *aiHandler) Edit(ctx context.Context, in *EditRequest, out *EditResponse) error {
	return h.AiHandler.Edit(ctx, in, out)
}

func (h *aiHandler) Moderate(ctx context.Context, in *ModerateRequest, out *ModerateResponse) error {
	return h.AiHandler.Moderate(ctx, in, out)
}

func (h *aiHandler) Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error {
	return h.AiHandler.Generate(ctx, in, out)
}