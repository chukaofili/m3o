// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/event.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
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

// Api Endpoints for Event service

func NewEventEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Event service

type EventService interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...client.CallOption) (Event_ConsumeService, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
}

type eventService struct {
	c    client.Client
	name string
}

func NewEventService(name string, c client.Client) EventService {
	return &eventService{
		c:    c,
		name: name,
	}
}

func (c *eventService) Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Publish", in)
	out := new(PublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) Consume(ctx context.Context, in *ConsumeRequest, opts ...client.CallOption) (Event_ConsumeService, error) {
	req := c.c.NewRequest(c.name, "Event.Consume", &ConsumeRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &eventServiceConsume{stream}, nil
}

type Event_ConsumeService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ConsumeResponse, error)
}

type eventServiceConsume struct {
	stream client.Stream
}

func (x *eventServiceConsume) Close() error {
	return x.stream.Close()
}

func (x *eventServiceConsume) Context() context.Context {
	return x.stream.Context()
}

func (x *eventServiceConsume) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventServiceConsume) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventServiceConsume) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Event service

type EventHandler interface {
	Publish(context.Context, *PublishRequest, *PublishResponse) error
	Consume(context.Context, *ConsumeRequest, Event_ConsumeStream) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
}

func RegisterEventHandler(s server.Server, hdlr EventHandler, opts ...server.HandlerOption) error {
	type event interface {
		Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error
		Consume(ctx context.Context, stream server.Stream) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
	}
	type Event struct {
		event
	}
	h := &eventHandler{hdlr}
	return s.Handle(s.NewHandler(&Event{h}, opts...))
}

type eventHandler struct {
	EventHandler
}

func (h *eventHandler) Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error {
	return h.EventHandler.Publish(ctx, in, out)
}

func (h *eventHandler) Consume(ctx context.Context, stream server.Stream) error {
	m := new(ConsumeRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.EventHandler.Consume(ctx, m, &eventConsumeStream{stream})
}

type Event_ConsumeStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ConsumeResponse) error
}

type eventConsumeStream struct {
	stream server.Stream
}

func (x *eventConsumeStream) Close() error {
	return x.stream.Close()
}

func (x *eventConsumeStream) Context() context.Context {
	return x.stream.Context()
}

func (x *eventConsumeStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventConsumeStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventConsumeStream) Send(m *ConsumeResponse) error {
	return x.stream.Send(m)
}

func (h *eventHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.EventHandler.Read(ctx, in, out)
}
