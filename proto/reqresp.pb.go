// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reqresp.proto

package proto

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

type Request struct {
	RequestID            int32             `protobuf:"varint,1,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
	Mapper               map[string]string `protobuf:"bytes,2,rep,name=mapper,proto3" json:"mapper,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e6a846fc0301fb, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetRequestID() int32 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *Request) GetMapper() map[string]string {
	if m != nil {
		return m.Mapper
	}
	return nil
}

type Response struct {
	ResponseID           int32             `protobuf:"varint,1,opt,name=ResponseID,proto3" json:"ResponseID,omitempty"`
	Mapper               map[string]string `protobuf:"bytes,2,rep,name=mapper,proto3" json:"mapper,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e6a846fc0301fb, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseID() int32 {
	if m != nil {
		return m.ResponseID
	}
	return 0
}

func (m *Response) GetMapper() map[string]string {
	if m != nil {
		return m.Mapper
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterMapType((map[string]string)(nil), "proto.Request.MapperEntry")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterMapType((map[string]string)(nil), "proto.Response.MapperEntry")
}

func init() { proto.RegisterFile("reqresp.proto", fileDescriptor_c6e6a846fc0301fb) }

var fileDescriptor_c6e6a846fc0301fb = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4a, 0x2d, 0x2e, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd3, 0x18,
	0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0xa1, 0x4c, 0x4f,
	0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x84, 0x80, 0x90, 0x11, 0x17, 0x5b, 0x6e, 0x62,
	0x41, 0x41, 0x6a, 0x91, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x14, 0xc4, 0x20, 0x3d, 0xa8,
	0x0a, 0x3d, 0x5f, 0xb0, 0xa4, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x54, 0xa5, 0x94, 0x25, 0x17,
	0x37, 0x92, 0xb0, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0x25, 0xd8, 0x68, 0xce, 0x20, 0x10, 0x53,
	0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x58, 0x31,
	0x59, 0x30, 0x2a, 0xcd, 0x62, 0xe4, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x92, 0xe3, 0xe2, 0x82, 0xb1, 0xe1, 0x4e, 0x43, 0x12, 0x11, 0x32, 0x46, 0x73, 0x9b, 0x34, 0xdc,
	0x6d, 0x10, 0x25, 0x54, 0x76, 0x9c, 0x91, 0x33, 0x17, 0x7f, 0x68, 0x71, 0x6a, 0x91, 0x4b, 0x62,
	0x49, 0x62, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x01, 0x17, 0x77, 0x70, 0x6a, 0x5e,
	0x0a, 0x2c, 0x2c, 0xf9, 0x50, 0x43, 0x47, 0x8a, 0x1f, 0xcd, 0x45, 0x4a, 0x0c, 0x49, 0x6c, 0x60,
	0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x1c, 0x8d, 0x1c, 0x99, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserDataServiceClient interface {
	SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type userDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserDataServiceClient(cc *grpc.ClientConn) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.UserDataService/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServiceServer is the server API for UserDataService service.
type UserDataServiceServer interface {
	SendRequest(context.Context, *Request) (*Response, error)
}

// UnimplementedUserDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (*UnimplementedUserDataServiceServer) SendRequest(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}

func RegisterUserDataServiceServer(s *grpc.Server, srv UserDataServiceServer) {
	s.RegisterService(&_UserDataService_serviceDesc, srv)
}

func _UserDataService_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserDataService/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).SendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _UserDataService_SendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reqresp.proto",
}