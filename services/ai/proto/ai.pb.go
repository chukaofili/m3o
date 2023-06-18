// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/ai.proto

package ai

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Make a request to the AI
type CompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// input to pass in
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CompleteRequest) Reset() {
	*x = CompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRequest) ProtoMessage() {}

func (x *CompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRequest.ProtoReflect.Descriptor instead.
func (*CompleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Response from the AI
type CompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text returned
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CompleteResponse) Reset() {
	*x = CompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteResponse) ProtoMessage() {}

func (x *CompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteResponse.ProtoReflect.Descriptor instead.
func (*CompleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Edit or edit prompt/code
type EditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text/code to check
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// instruction hint e.g check the grammar
	Instruction string `protobuf:"bytes,2,opt,name=instruction,proto3" json:"instruction,omitempty"`
}

func (x *EditRequest) Reset() {
	*x = EditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRequest) ProtoMessage() {}

func (x *EditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRequest.ProtoReflect.Descriptor instead.
func (*EditRequest) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{2}
}

func (x *EditRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EditRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

type EditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response output
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *EditResponse) Reset() {
	*x = EditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditResponse) ProtoMessage() {}

func (x *EditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditResponse.ProtoReflect.Descriptor instead.
func (*EditResponse) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{3}
}

func (x *EditResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Moderate hate speech
type ModerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text to check
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ModerateRequest) Reset() {
	*x = ModerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModerateRequest) ProtoMessage() {}

func (x *ModerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModerateRequest.ProtoReflect.Descriptor instead.
func (*ModerateRequest) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{4}
}

func (x *ModerateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ModerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// whether it was flagged or not
	Flagged bool `protobuf:"varint,1,opt,name=flagged,proto3" json:"flagged,omitempty"`
	// categories tested and identified
	Categories map[string]bool `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// related scores
	Scores map[string]float64 `protobuf:"bytes,3,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *ModerateResponse) Reset() {
	*x = ModerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModerateResponse) ProtoMessage() {}

func (x *ModerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModerateResponse.ProtoReflect.Descriptor instead.
func (*ModerateResponse) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{5}
}

func (x *ModerateResponse) GetFlagged() bool {
	if x != nil {
		return x.Flagged
	}
	return false
}

func (x *ModerateResponse) GetCategories() map[string]bool {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ModerateResponse) GetScores() map[string]float64 {
	if x != nil {
		return x.Scores
	}
	return nil
}

// Generage an image from prompt
type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text description of image
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// number of images to generate (max 10)
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// size of image 256x256, 512x512, 1024x1024
	Size string `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GenerateRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GenerateRequest) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// image url
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// base64 encoded
	Base64 string `protobuf:"bytes,2,opt,name=base64,proto3" json:"base64,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{7}
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetBase64() string {
	if x != nil {
		return x.Base64
	}
	return ""
}

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// image urls
	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ai_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{8}
}

func (x *GenerateResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

var File_proto_ai_proto protoreflect.FileDescriptor

var file_proto_ai_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x61, 0x69, 0x22, 0x25, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x43, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x25, 0x0a, 0x0f,
	0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0xa6, 0x02, 0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6c, 0x61, 0x67,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x67,
	0x65, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x69, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0f,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x31, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65,
	0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34,
	0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x32, 0xdc, 0x01, 0x0a, 0x02, 0x41, 0x69, 0x12, 0x37,
	0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x69, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12,
	0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x61, 0x69, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x61, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x69, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x61, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ai_proto_rawDescOnce sync.Once
	file_proto_ai_proto_rawDescData = file_proto_ai_proto_rawDesc
)

func file_proto_ai_proto_rawDescGZIP() []byte {
	file_proto_ai_proto_rawDescOnce.Do(func() {
		file_proto_ai_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ai_proto_rawDescData)
	})
	return file_proto_ai_proto_rawDescData
}

var file_proto_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_ai_proto_goTypes = []interface{}{
	(*CompleteRequest)(nil),  // 0: ai.CompleteRequest
	(*CompleteResponse)(nil), // 1: ai.CompleteResponse
	(*EditRequest)(nil),      // 2: ai.EditRequest
	(*EditResponse)(nil),     // 3: ai.EditResponse
	(*ModerateRequest)(nil),  // 4: ai.ModerateRequest
	(*ModerateResponse)(nil), // 5: ai.ModerateResponse
	(*GenerateRequest)(nil),  // 6: ai.GenerateRequest
	(*Image)(nil),            // 7: ai.Image
	(*GenerateResponse)(nil), // 8: ai.GenerateResponse
	nil,                      // 9: ai.ModerateResponse.CategoriesEntry
	nil,                      // 10: ai.ModerateResponse.ScoresEntry
}
var file_proto_ai_proto_depIdxs = []int32{
	9,  // 0: ai.ModerateResponse.categories:type_name -> ai.ModerateResponse.CategoriesEntry
	10, // 1: ai.ModerateResponse.scores:type_name -> ai.ModerateResponse.ScoresEntry
	7,  // 2: ai.GenerateResponse.images:type_name -> ai.Image
	0,  // 3: ai.Ai.Complete:input_type -> ai.CompleteRequest
	2,  // 4: ai.Ai.Edit:input_type -> ai.EditRequest
	4,  // 5: ai.Ai.Moderate:input_type -> ai.ModerateRequest
	6,  // 6: ai.Ai.Generate:input_type -> ai.GenerateRequest
	1,  // 7: ai.Ai.Complete:output_type -> ai.CompleteResponse
	3,  // 8: ai.Ai.Edit:output_type -> ai.EditResponse
	5,  // 9: ai.Ai.Moderate:output_type -> ai.ModerateResponse
	8,  // 10: ai.Ai.Generate:output_type -> ai.GenerateResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_ai_proto_init() }
func file_proto_ai_proto_init() {
	if File_proto_ai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModerateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModerateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ai_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ai_proto_goTypes,
		DependencyIndexes: file_proto_ai_proto_depIdxs,
		MessageInfos:      file_proto_ai_proto_msgTypes,
	}.Build()
	File_proto_ai_proto = out.File
	file_proto_ai_proto_rawDesc = nil
	file_proto_ai_proto_goTypes = nil
	file_proto_ai_proto_depIdxs = nil
}