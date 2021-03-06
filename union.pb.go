// Code generated by protoc-gen-go.
// source: union.proto
// DO NOT EDIT!

/*
Package protounion is a generated protocol buffer package.

It is generated from these files:
	union.proto

It has these top-level messages:
	TypeHeader
*/
package protounion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TypeHeader struct {
	Type uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
}

func (m *TypeHeader) Reset()                    { *m = TypeHeader{} }
func (m *TypeHeader) String() string            { return proto.CompactTextString(m) }
func (*TypeHeader) ProtoMessage()               {}
func (*TypeHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TypeHeader) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*TypeHeader)(nil), "protounion.TypeHeader")
}

func init() { proto.RegisterFile("union.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xcd, 0xcb, 0xcc,
	0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x53, 0x60, 0x11, 0x25, 0x05, 0x2e,
	0xae, 0x90, 0xca, 0x82, 0x54, 0x8f, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0x21, 0x21, 0x2e, 0x96, 0x92,
	0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x30, 0x3b, 0x89, 0x0d, 0xac, 0xda,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x7d, 0x52, 0xe8, 0x43, 0x00, 0x00, 0x00,
}
