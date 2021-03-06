// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zas/tmp/ping.proto

package pkg

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PingRequest struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7857c3c85066e9d3, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PingResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7857c3c85066e9d3, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "sms.core.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "sms.core.PingResponse")
}

func init() { proto.RegisterFile("zas/tmp/ping.proto", fileDescriptor_7857c3c85066e9d3) }

var fileDescriptor_7857c3c85066e9d3 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x43, 0x31,
	0x14, 0x85, 0x79, 0x55, 0xac, 0x4d, 0x5d, 0x8c, 0x28, 0xb5, 0x74, 0x28, 0x6f, 0x2a, 0x08, 0x09,
	0xe8, 0xea, 0x62, 0x05, 0x5d, 0x1f, 0xcf, 0xcd, 0x45, 0x62, 0x7a, 0x49, 0x03, 0x4d, 0xee, 0x35,
	0x37, 0x3a, 0x28, 0x2e, 0xfe, 0x05, 0x7f, 0x9a, 0x7f, 0xc1, 0x1f, 0x22, 0x7d, 0xa1, 0x2a, 0x0e,
	0x8e, 0xf7, 0x70, 0x38, 0xe7, 0xbb, 0x47, 0xc8, 0x67, 0xc3, 0x3a, 0x07, 0xd2, 0xe4, 0xa3, 0x53,
	0x94, 0x30, 0xa3, 0xdc, 0xe5, 0xc0, 0xca, 0x62, 0x82, 0xf1, 0xc4, 0x21, 0xba, 0x15, 0x68, 0x43,
	0x5e, 0x9b, 0x18, 0x31, 0x9b, 0xec, 0x31, 0x72, 0xf1, 0xd5, 0x27, 0x62, 0xd8, 0xf8, 0xe8, 0x5a,
	0x78, 0x78, 0x04, 0xce, 0x72, 0x22, 0x06, 0xd9, 0x07, 0xe0, 0x6c, 0x02, 0x8d, 0xaa, 0x69, 0x35,
	0xdb, 0x6a, 0x7f, 0x84, 0xfa, 0x4a, 0xec, 0x15, 0x33, 0x13, 0x46, 0x86, 0xff, 0xdd, 0x72, 0x24,
	0xfa, 0x01, 0x98, 0x8d, 0x83, 0x51, 0x6f, 0x5a, 0xcd, 0x06, 0xed, 0xe6, 0x3c, 0xbd, 0x13, 0xc3,
	0x4b, 0x4c, 0x70, 0x03, 0xe9, 0xc9, 0x5b, 0x90, 0x8d, 0xd8, 0x5e, 0xc7, 0xca, 0x43, 0xb5, 0x81,
	0x56, 0xbf, 0x98, 0xc6, 0x47, 0x7f, 0xe5, 0xd2, 0x5e, 0x1f, 0xbf, 0x7d, 0x7c, 0xbe, 0xf7, 0x0e,
	0xe4, 0x7e, 0xf7, 0xb7, 0x7e, 0xf9, 0x6e, 0x7e, 0x9d, 0x9f, 0x8b, 0xda, 0x62, 0x50, 0x79, 0xe9,
	0xe3, 0x72, 0x61, 0x14, 0x53, 0x5a, 0x2f, 0xe3, 0x12, 0xd9, 0x92, 0x13, 0x70, 0x01, 0xab, 0x79,
	0x07, 0x71, 0x9d, 0xc8, 0x5e, 0x90, 0x6f, 0xaa, 0xdb, 0x3e, 0x17, 0x9e, 0xfb, 0x9d, 0x6e, 0x9a,
	0xb3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xef, 0xd7, 0xc6, 0x58, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/sms.core.CoreService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedCoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (*UnimplementedCoreServiceServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.core.CoreService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CoreService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zas/tmp/ping.proto",
}
