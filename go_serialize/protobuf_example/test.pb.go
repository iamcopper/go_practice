// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package protobuf_example

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
	Tag                  *uint64  `protobuf:"varint,1,req,name=tag" json:"tag,omitempty"`
	Length               *uint64  `protobuf:"varint,2,req,name=length" json:"length,omitempty"`
	Value                *string  `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
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

func (m *Message) GetTag() uint64 {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return 0
}

func (m *Message) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Message) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "protobuf_example.Message")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00, 0x53, 0x49, 0xa5, 0x69, 0xf1, 0xa9, 0x15,
	0x89, 0xb9, 0x05, 0x39, 0xa9, 0x4a, 0x9e, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9,
	0x42, 0x02, 0x5c, 0xcc, 0x25, 0x89, 0xe9, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x2c, 0x41, 0x20, 0xa6,
	0x90, 0x18, 0x17, 0x5b, 0x4e, 0x6a, 0x5e, 0x7a, 0x49, 0x86, 0x04, 0x13, 0x58, 0x10, 0xca, 0x13,
	0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x56, 0x60, 0xd2, 0xe0, 0x0c, 0x82,
	0x70, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0xdc, 0xb2, 0xe0, 0x69, 0x00, 0x00, 0x00,
}
