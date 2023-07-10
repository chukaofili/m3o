// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/search.proto

package search

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

// Api Endpoints for Search service

func NewSearchEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Search service

type SearchService interface {
	Index(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*IndexResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...client.CallOption) (*CreateIndexResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...client.CallOption) (*DeleteIndexResponse, error)
}

type searchService struct {
	c    client.Client
	name string
}

func NewSearchService(name string, c client.Client) SearchService {
	return &searchService{
		c:    c,
		name: name,
	}
}

func (c *searchService) Index(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*IndexResponse, error) {
	req := c.c.NewRequest(c.name, "Search.Index", in)
	out := new(IndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Search.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Search.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchService) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...client.CallOption) (*CreateIndexResponse, error) {
	req := c.c.NewRequest(c.name, "Search.CreateIndex", in)
	out := new(CreateIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchService) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...client.CallOption) (*DeleteIndexResponse, error) {
	req := c.c.NewRequest(c.name, "Search.DeleteIndex", in)
	out := new(DeleteIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Search service

type SearchHandler interface {
	Index(context.Context, *IndexRequest, *IndexResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	CreateIndex(context.Context, *CreateIndexRequest, *CreateIndexResponse) error
	DeleteIndex(context.Context, *DeleteIndexRequest, *DeleteIndexResponse) error
}

func RegisterSearchHandler(s server.Server, hdlr SearchHandler, opts ...server.HandlerOption) error {
	type search interface {
		Index(ctx context.Context, in *IndexRequest, out *IndexResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
		CreateIndex(ctx context.Context, in *CreateIndexRequest, out *CreateIndexResponse) error
		DeleteIndex(ctx context.Context, in *DeleteIndexRequest, out *DeleteIndexResponse) error
	}
	type Search struct {
		search
	}
	h := &searchHandler{hdlr}
	return s.Handle(s.NewHandler(&Search{h}, opts...))
}

type searchHandler struct {
	SearchHandler
}

func (h *searchHandler) Index(ctx context.Context, in *IndexRequest, out *IndexResponse) error {
	return h.SearchHandler.Index(ctx, in, out)
}

func (h *searchHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.SearchHandler.Delete(ctx, in, out)
}

func (h *searchHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.SearchHandler.Search(ctx, in, out)
}

func (h *searchHandler) CreateIndex(ctx context.Context, in *CreateIndexRequest, out *CreateIndexResponse) error {
	return h.SearchHandler.CreateIndex(ctx, in, out)
}

func (h *searchHandler) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, out *DeleteIndexResponse) error {
	return h.SearchHandler.DeleteIndex(ctx, in, out)
}
