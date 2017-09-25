// Code generated by protoc-gen-go.
// source: pods.proto
// DO NOT EDIT!

/*
Package pods is a generated protocol buffer package.

It is generated from these files:
	pods.proto

It has these top-level messages:
	ListPodsRequest
	ListPodsResponse
	GetLogsRequest
	GetLogsResponse
	Pod
	PodSpec
	Container
	PodStatus
	ContainerStatus
*/
package pods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetLogsResponse_Type int32

const (
	GetLogsResponse_STDOUT GetLogsResponse_Type = 0
	GetLogsResponse_STDERR GetLogsResponse_Type = 1
)

var GetLogsResponse_Type_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
}
var GetLogsResponse_Type_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
}

func (x GetLogsResponse_Type) String() string {
	return proto.EnumName(GetLogsResponse_Type_name, int32(x))
}
func (GetLogsResponse_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type ListPodsRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ListPodsRequest) Reset()                    { *m = ListPodsRequest{} }
func (m *ListPodsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPodsRequest) ProtoMessage()               {}
func (*ListPodsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListPodsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListPodsResponse struct {
	Pods []*Pod `protobuf:"bytes,1,rep,name=pods" json:"pods,omitempty"`
}

func (m *ListPodsResponse) Reset()                    { *m = ListPodsResponse{} }
func (m *ListPodsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPodsResponse) ProtoMessage()               {}
func (*ListPodsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListPodsResponse) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type GetLogsRequest struct {
	Namespace   string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	ContainerID string `protobuf:"bytes,2,opt,name=containerID" json:"containerID,omitempty"`
}

func (m *GetLogsRequest) Reset()                    { *m = GetLogsRequest{} }
func (m *GetLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogsRequest) ProtoMessage()               {}
func (*GetLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetLogsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetLogsRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

type GetLogsResponse struct {
	Type GetLogsResponse_Type `protobuf:"varint,1,opt,name=type,enum=cand.services.pods.v1.GetLogsResponse_Type" json:"type,omitempty"`
	Line []byte               `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
}

func (m *GetLogsResponse) Reset()                    { *m = GetLogsResponse{} }
func (m *GetLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLogsResponse) ProtoMessage()               {}
func (*GetLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetLogsResponse) GetType() GetLogsResponse_Type {
	if m != nil {
		return m.Type
	}
	return GetLogsResponse_STDOUT
}

func (m *GetLogsResponse) GetLine() []byte {
	if m != nil {
		return m.Line
	}
	return nil
}

type Pod struct {
	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec     *PodSpec          `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *PodStatus        `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pod) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Pod) GetSpec() *PodSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Pod) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PodSpec struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
}

func (m *PodSpec) Reset()                    { *m = PodSpec{} }
func (m *PodSpec) String() string            { return proto.CompactTextString(m) }
func (*PodSpec) ProtoMessage()               {}
func (*PodSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PodSpec) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type Container struct {
	ID    string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Container) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type PodStatus struct {
	ContainerStatuses []*ContainerStatus `protobuf:"bytes,1,rep,name=containerStatuses" json:"containerStatuses,omitempty"`
}

func (m *PodStatus) Reset()                    { *m = PodStatus{} }
func (m *PodStatus) String() string            { return proto.CompactTextString(m) }
func (*PodStatus) ProtoMessage()               {}
func (*PodStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PodStatus) GetContainerStatuses() []*ContainerStatus {
	if m != nil {
		return m.ContainerStatuses
	}
	return nil
}

type ContainerStatus struct {
	ContainerID string `protobuf:"bytes,1,opt,name=containerID" json:"containerID,omitempty"`
	Image       string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *ContainerStatus) Reset()                    { *m = ContainerStatus{} }
func (m *ContainerStatus) String() string            { return proto.CompactTextString(m) }
func (*ContainerStatus) ProtoMessage()               {}
func (*ContainerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ContainerStatus) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *ContainerStatus) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*ListPodsRequest)(nil), "cand.services.pods.v1.ListPodsRequest")
	proto.RegisterType((*ListPodsResponse)(nil), "cand.services.pods.v1.ListPodsResponse")
	proto.RegisterType((*GetLogsRequest)(nil), "cand.services.pods.v1.GetLogsRequest")
	proto.RegisterType((*GetLogsResponse)(nil), "cand.services.pods.v1.GetLogsResponse")
	proto.RegisterType((*Pod)(nil), "cand.services.pods.v1.Pod")
	proto.RegisterType((*PodSpec)(nil), "cand.services.pods.v1.PodSpec")
	proto.RegisterType((*Container)(nil), "cand.services.pods.v1.Container")
	proto.RegisterType((*PodStatus)(nil), "cand.services.pods.v1.PodStatus")
	proto.RegisterType((*ContainerStatus)(nil), "cand.services.pods.v1.ContainerStatus")
	proto.RegisterEnum("cand.services.pods.v1.GetLogsResponse_Type", GetLogsResponse_Type_name, GetLogsResponse_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pods service

type PodsClient interface {
	List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	Logs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (Pods_LogsClient, error)
}

type podsClient struct {
	cc *grpc.ClientConn
}

func NewPodsClient(cc *grpc.ClientConn) PodsClient {
	return &podsClient{cc}
}

func (c *podsClient) List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) Logs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (Pods_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Pods_serviceDesc.Streams[0], c.cc, "/cand.services.pods.v1.Pods/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &podsLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pods_LogsClient interface {
	Recv() (*GetLogsResponse, error)
	grpc.ClientStream
}

type podsLogsClient struct {
	grpc.ClientStream
}

func (x *podsLogsClient) Recv() (*GetLogsResponse, error) {
	m := new(GetLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Pods service

type PodsServer interface {
	List(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	Logs(*GetLogsRequest, Pods_LogsServer) error
}

func RegisterPodsServer(s *grpc.Server, srv PodsServer) {
	s.RegisterService(&_Pods_serviceDesc, srv)
}

func _Pods_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).List(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PodsServer).Logs(m, &podsLogsServer{stream})
}

type Pods_LogsServer interface {
	Send(*GetLogsResponse) error
	grpc.ServerStream
}

type podsLogsServer struct {
	grpc.ServerStream
}

func (x *podsLogsServer) Send(m *GetLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Pods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cand.services.pods.v1.Pods",
	HandlerType: (*PodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Pods_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _Pods_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pods.proto",
}

func init() { proto.RegisterFile("pods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x51, 0x8b, 0x13, 0x31,
	0x10, 0xc7, 0xcd, 0x76, 0x3d, 0xdd, 0xa9, 0xb6, 0x35, 0x28, 0x94, 0xa2, 0x67, 0x09, 0x78, 0x16,
	0x84, 0x5d, 0x6f, 0x7d, 0x39, 0xbc, 0x07, 0xe5, 0xdc, 0x22, 0xc5, 0x13, 0x4b, 0xae, 0x22, 0xf8,
	0x22, 0x71, 0x37, 0x94, 0xe5, 0xae, 0x9b, 0xd8, 0xa4, 0x85, 0x7e, 0x01, 0x3f, 0x93, 0x9f, 0xcc,
	0x67, 0x49, 0x36, 0xdd, 0xd6, 0x9e, 0x6b, 0xef, 0x69, 0x93, 0xc9, 0x7f, 0x66, 0x7e, 0xc9, 0xcc,
	0x2c, 0x80, 0x14, 0x99, 0x0a, 0xe5, 0x5c, 0x68, 0x81, 0x1f, 0xa5, 0xac, 0xc8, 0x42, 0xc5, 0xe7,
	0xcb, 0x3c, 0xe5, 0x2a, 0xb4, 0x27, 0xcb, 0x63, 0x12, 0x41, 0xfb, 0x3c, 0x57, 0x7a, 0x2c, 0x32,
	0x45, 0xf9, 0x8f, 0x05, 0x57, 0x1a, 0x3f, 0x86, 0xa0, 0x60, 0x33, 0xae, 0x24, 0x4b, 0x79, 0x17,
	0xf5, 0xd1, 0x20, 0xa0, 0x1b, 0x03, 0x39, 0x83, 0xce, 0xc6, 0x41, 0x49, 0x51, 0x28, 0x8e, 0x43,
	0xf0, 0x4d, 0xbc, 0x2e, 0xea, 0x37, 0x06, 0xcd, 0xb8, 0x17, 0xfe, 0x33, 0x55, 0x38, 0x16, 0x19,
	0xb5, 0x3a, 0x32, 0x86, 0xd6, 0x7b, 0xae, 0xcf, 0xc5, 0xf4, 0x66, 0x39, 0x71, 0x1f, 0x9a, 0xa9,
	0x28, 0x34, 0xcb, 0x0b, 0x3e, 0x1f, 0x25, 0x5d, 0xcf, 0x9e, 0x6f, 0x9b, 0xc8, 0x4f, 0x04, 0xed,
	0x2a, 0xa4, 0xa3, 0x7a, 0x03, 0xbe, 0x5e, 0xc9, 0x32, 0x5c, 0x2b, 0x7e, 0x51, 0x43, 0xb5, 0xe3,
	0x15, 0x4e, 0x56, 0x92, 0x53, 0xeb, 0x88, 0x31, 0xf8, 0x57, 0x79, 0xc1, 0x6d, 0xbe, 0x7b, 0xd4,
	0xae, 0xc9, 0x21, 0xf8, 0x46, 0x81, 0x01, 0x0e, 0x2e, 0x26, 0xc9, 0xa7, 0xcf, 0x93, 0xce, 0x2d,
	0xb7, 0x1e, 0x52, 0xda, 0x41, 0xe4, 0x37, 0x82, 0xc6, 0x58, 0x64, 0x38, 0x81, 0xbb, 0x33, 0xae,
	0x59, 0xc6, 0x34, 0x73, 0xcf, 0x32, 0xa8, 0x7f, 0x96, 0xf0, 0xa3, 0x93, 0x0e, 0x0b, 0x3d, 0x5f,
	0xd1, 0xca, 0x13, 0xc7, 0xe0, 0x2b, 0xc9, 0x53, 0x4b, 0xd0, 0x8c, 0x0f, 0xeb, 0x23, 0x5c, 0x48,
	0x9e, 0x52, 0xab, 0xc5, 0x27, 0x70, 0xa0, 0x34, 0xd3, 0x0b, 0xd5, 0x6d, 0x58, 0xaf, 0xfe, 0x7f,
	0xbc, 0xac, 0x8e, 0x3a, 0x7d, 0xef, 0x14, 0xee, 0xff, 0x05, 0x82, 0x3b, 0xd0, 0xb8, 0xe4, 0x2b,
	0x57, 0x0f, 0xb3, 0xc4, 0x0f, 0xe1, 0xf6, 0x92, 0x5d, 0x2d, 0xb8, 0xab, 0x41, 0xb9, 0x79, 0xed,
	0x9d, 0x20, 0xf2, 0x01, 0xee, 0x38, 0x0e, 0xfc, 0x16, 0xa0, 0xaa, 0xcd, 0xba, 0x29, 0xea, 0x28,
	0xde, 0xad, 0x85, 0x74, 0xcb, 0x87, 0x0c, 0x21, 0xa8, 0x0e, 0x70, 0x0b, 0xbc, 0x51, 0xe2, 0x20,
	0xbc, 0x51, 0x62, 0xca, 0x62, 0x5a, 0xc3, 0x21, 0xd8, 0xb5, 0xe1, 0xca, 0x67, 0x6c, 0xca, 0xed,
	0x9d, 0x03, 0x5a, 0x6e, 0x08, 0x83, 0xa0, 0xba, 0x25, 0x9e, 0xc0, 0x83, 0x2a, 0x43, 0x69, 0xe2,
	0x6b, 0xb8, 0xa3, 0x7d, 0x70, 0xee, 0xa1, 0xae, 0x07, 0x20, 0xdf, 0xa0, 0xbd, 0xa3, 0xda, 0xed,
	0x56, 0x74, 0xad, 0x5b, 0x37, 0xb4, 0xde, 0x16, 0xad, 0xb1, 0x9a, 0x42, 0x54, 0x77, 0xb0, 0x9b,
	0xf8, 0x17, 0x02, 0xdf, 0x0c, 0x1b, 0xfe, 0x02, 0xbe, 0x19, 0x3c, 0x5c, 0x07, 0xbb, 0x33, 0xc6,
	0xbd, 0xe7, 0x7b, 0x75, 0x6e, 0x4e, 0x4c, 0x60, 0x31, 0x55, 0xf8, 0xd9, 0xbe, 0x09, 0x29, 0xe3,
	0x1e, 0xdd, 0x6c, 0x90, 0x5e, 0xa2, 0xb3, 0xa7, 0x5f, 0x9f, 0xc8, 0xcb, 0x69, 0xc4, 0x64, 0x1e,
	0xad, 0xd5, 0x91, 0x51, 0x47, 0xcb, 0xe3, 0x53, 0xf3, 0xfd, 0x7e, 0x60, 0x7f, 0x4d, 0xaf, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x69, 0xdd, 0x6f, 0xa8, 0x04, 0x00, 0x00,
}