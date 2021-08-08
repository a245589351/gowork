// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	proto/hello.proto

It has these top-level messages:
	SayHelloRequest
	SayHelloResponse
*/
package hello

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type SayHelloRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SayHelloRequest) Reset()                    { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string            { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()               {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SayHelloRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SayHelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SayHelloResponse) Reset()                    { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string            { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()               {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SayHelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "SayHelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := grpc.Invoke(ctx, "/Hello/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x95, 0xb4, 0xb9, 0xf8, 0x83, 0x13,
	0x2b, 0x3d, 0x40, 0x22, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92,
	0x0e, 0x97, 0x00, 0x42, 0x71, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x6e, 0xd5, 0x46, 0x56, 0x5c,
	0xac, 0x60, 0xa5, 0x42, 0x86, 0x5c, 0x1c, 0x30, 0x6d, 0x42, 0x02, 0x7a, 0x68, 0xd6, 0x49, 0x09,
	0xea, 0xa1, 0x9b, 0xa9, 0xc4, 0xe0, 0xc4, 0x1f, 0xc5, 0xab, 0xa7, 0x8f, 0xe4, 0xda, 0x24, 0x36,
	0x30, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x64, 0xb6, 0x2f, 0x7a, 0xc3, 0x00, 0x00, 0x00,
}
