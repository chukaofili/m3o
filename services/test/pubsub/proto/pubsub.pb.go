// Code generated by protoc-gen-go. DO NOT EDIT.
// source: m3o.dev/services/test/pubsub/proto/pubsub.proto

package pubsub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fce66331057326, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "pubsub.Message")
}

func init() {
	proto.RegisterFile("m3o.dev/services/test/pubsub/proto/pubsub.proto", fileDescriptor_59fce66331057326)
}

var fileDescriptor_59fce66331057326 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2f, 0x28, 0x4d, 0x2a, 0x2e,
	0x4d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x72, 0xf4, 0xc0, 0x1c, 0x21, 0x36, 0x08, 0x4f,
	0x49, 0x97, 0x8b, 0xdd, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x29,
	0x3f, 0xa5, 0x52, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x4e, 0x62, 0x03, 0xeb, 0x36,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x26, 0xef, 0xb4, 0x87, 0x79, 0x00, 0x00, 0x00,
}