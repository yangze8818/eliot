// Code generated by protoc-gen-go.
// source: core/metadata.proto
// DO NOT EDIT!

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core/metadata.proto

It has these top-level messages:
	ResourceMetadata
*/
package core

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

// ResourceMetadata is metadata that all resources must have
type ResourceMetadata struct {
	// Name must be unique within a namespace.
	// Cannot be updated.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Namespace defines space within each name must be unique.
	// An empty namespace is equivalent to the default namespace.
	// Cannot be updated.
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ResourceMetadata) Reset()                    { *m = ResourceMetadata{} }
func (m *ResourceMetadata) String() string            { return proto.CompactTextString(m) }
func (*ResourceMetadata) ProtoMessage()               {}
func (*ResourceMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResourceMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceMetadata)(nil), "cand.core.ResourceMetadata")
}

func init() { proto.RegisterFile("core/metadata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0x2f, 0x4a,
	0xd5, 0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4c, 0x4e, 0xcc, 0x4b, 0xd1, 0x03, 0xc9, 0x28, 0xb9, 0x70, 0x09, 0x04, 0xa5, 0x16, 0xe7,
	0x97, 0x16, 0x25, 0xa7, 0xfa, 0x42, 0x15, 0x09, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x32, 0x5c, 0x9c, 0x20, 0xba, 0xb8, 0x20,
	0x31, 0x39, 0x55, 0x82, 0x09, 0x2c, 0x81, 0x10, 0x70, 0xd2, 0x8e, 0xd2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2d, 0xca, 0xcb, 0x4f, 0x4c, 0x2c, 0x48, 0xd4,
	0x4f, 0x4e, 0xcc, 0xd3, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x07, 0xd9, 0x66, 0x0d,
	0x22, 0x92, 0xd8, 0xc0, 0x8e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x60, 0x2e, 0x75, 0xb6,
	0x9b, 0x00, 0x00, 0x00,
}
