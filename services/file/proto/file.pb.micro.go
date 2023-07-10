// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/file.proto

package file

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

// Api Endpoints for File service

func NewFileEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for File service

type FileService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "File.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "File.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error) {
	req := c.c.NewRequest(c.name, "File.Save", in)
	out := new(SaveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "File.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for File service

type FileHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Save(context.Context, *SaveRequest, *SaveResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterFileHandler(s server.Server, hdlr FileHandler, opts ...server.HandlerOption) error {
	type file interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type File struct {
		file
	}
	h := &fileHandler{hdlr}
	return s.Handle(s.NewHandler(&File{h}, opts...))
}

type fileHandler struct {
	FileHandler
}

func (h *fileHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.FileHandler.Read(ctx, in, out)
}

func (h *fileHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.FileHandler.List(ctx, in, out)
}

func (h *fileHandler) Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error {
	return h.FileHandler.Save(ctx, in, out)
}

func (h *fileHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.FileHandler.Delete(ctx, in, out)
}
