// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HelloRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Eat                  string   `protobuf:"bytes,2,opt,name=eat,proto3" json:"eat,omitempty"`
	Work                 string   `protobuf:"bytes,3,opt,name=work,proto3" json:"work,omitempty"`
	Food                 []*Food  `protobuf:"bytes,4,rep,name=food,proto3" json:"food,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloRequest) GetEat() string {
	if m != nil {
		return m.Eat
	}
	return ""
}

func (m *HelloRequest) GetWork() string {
	if m != nil {
		return m.Work
	}
	return ""
}

func (m *HelloRequest) GetFood() []*Food {
	if m != nil {
		return m.Food
	}
	return nil
}

type HelloReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Food struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  string   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Food) Reset()         { *m = Food{} }
func (m *Food) String() string { return proto.CompactTextString(m) }
func (*Food) ProtoMessage()    {}
func (*Food) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *Food) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Food.Unmarshal(m, b)
}
func (m *Food) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Food.Marshal(b, m, deterministic)
}
func (m *Food) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Food.Merge(m, src)
}
func (m *Food) XXX_Size() int {
	return xxx_messageInfo_Food.Size(m)
}
func (m *Food) XXX_DiscardUnknown() {
	xxx_messageInfo_Food.DiscardUnknown(m)
}

var xxx_messageInfo_Food proto.InternalMessageInfo

func (m *Food) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Food) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "HelloReply")
	proto.RegisterType((*Food)(nil), "Food")
}

func init() {
	proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce)
}

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0x8d, 0xc9, 0xee, 0xe2, 0x54, 0x41, 0xe6, 0x20, 0xd1, 0xd3, 0x12, 0x3d, 0x14, 0x85,
	0x20, 0xeb, 0x7f, 0x10, 0xcf, 0xed, 0xc9, 0x63, 0xa4, 0xe3, 0x07, 0x8d, 0x4c, 0x6d, 0xd2, 0x4a,
	0xff, 0xbd, 0x24, 0xa4, 0xe8, 0xcd, 0x9c, 0xde, 0xf7, 0xc9, 0xc7, 0x93, 0x04, 0xaa, 0x77, 0xf2,
	0x9e, 0xed, 0x30, 0x72, 0x64, 0xf3, 0x01, 0xa7, 0x4f, 0xa9, 0x36, 0xf4, 0x35, 0x51, 0x88, 0xa8,
	0x61, 0xf7, 0x49, 0x21, 0xb8, 0x37, 0xd2, 0x62, 0x2f, 0xea, 0x93, 0x66, 0xad, 0x78, 0x0e, 0x92,
	0x5c, 0xd4, 0xc7, 0x99, 0xa6, 0x88, 0x08, 0xea, 0x9b, 0xc7, 0x5e, 0xcb, 0x8c, 0x72, 0xc6, 0x4b,
	0x50, 0xaf, 0xcc, 0x9d, 0x56, 0x7b, 0x59, 0x57, 0x87, 0x8d, 0x7d, 0x64, 0xee, 0x9a, 0x8c, 0xcc,
	0x0d, 0x40, 0x51, 0x0d, 0x7e, 0xc1, 0x0b, 0xd8, 0x8e, 0x14, 0x26, 0x1f, 0x8b, 0xa7, 0x34, 0x73,
	0x0b, 0x2a, 0xed, 0x49, 0xba, 0x9e, 0x96, 0x32, 0x99, 0x62, 0x22, 0xb3, 0xf3, 0xeb, 0x05, 0x66,
	0xe7, 0x0f, 0xcf, 0xb0, 0xc9, 0x27, 0xe2, 0x35, 0xc8, 0xd6, 0x2d, 0x78, 0x66, 0xff, 0xbe, 0xe5,
	0xaa, 0xb2, 0xbf, 0x3e, 0x73, 0x84, 0x77, 0xb0, 0x6b, 0x5d, 0x1a, 0xff, 0x2c, 0xac, 0xc5, 0xbd,
	0x78, 0xd9, 0xe6, 0xef, 0x79, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xf8, 0x83, 0xda, 0x2d,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Say(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Saaaaay(ctx context.Context, opts ...grpc.CallOption) (Hello_SaaaaayClient, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Hello/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Saaaaay(ctx context.Context, opts ...grpc.CallOption) (Hello_SaaaaayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[0], "/Hello/Saaaaay", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloSaaaaayClient{stream}
	return x, nil
}

type Hello_SaaaaayClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type helloSaaaaayClient struct {
	grpc.ClientStream
}

func (x *helloSaaaaayClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloSaaaaayClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Say(context.Context, *HelloRequest) (*HelloReply, error)
	Saaaaay(Hello_SaaaaayServer) error
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) Say(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}
func (*UnimplementedHelloServer) Saaaaay(srv Hello_SaaaaayServer) error {
	return status.Errorf(codes.Unimplemented, "method Saaaaay not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Say(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Saaaaay_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).Saaaaay(&helloSaaaaayServer{stream})
}

type Hello_SaaaaayServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloSaaaaayServer struct {
	grpc.ServerStream
}

func (x *helloSaaaaayServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloSaaaaayServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Saaaaay",
			Handler:       _Hello_Saaaaay_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
