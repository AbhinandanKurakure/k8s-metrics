// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protofiles/k8services.proto

package clientAPI

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

type InputValues struct {
	NameSpace            string   `protobuf:"bytes,1,opt,name=nameSpace,proto3" json:"nameSpace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputValues) Reset()         { *m = InputValues{} }
func (m *InputValues) String() string { return proto.CompactTextString(m) }
func (*InputValues) ProtoMessage()    {}
func (*InputValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a704f5ed8d36550, []int{0}
}

func (m *InputValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputValues.Unmarshal(m, b)
}
func (m *InputValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputValues.Marshal(b, m, deterministic)
}
func (m *InputValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputValues.Merge(m, src)
}
func (m *InputValues) XXX_Size() int {
	return xxx_messageInfo_InputValues.Size(m)
}
func (m *InputValues) XXX_DiscardUnknown() {
	xxx_messageInfo_InputValues.DiscardUnknown(m)
}

var xxx_messageInfo_InputValues proto.InternalMessageInfo

func (m *InputValues) GetNameSpace() string {
	if m != nil {
		return m.NameSpace
	}
	return ""
}

type OutputValues struct {
	Errors               string   `protobuf:"bytes,1,opt,name=errors,proto3" json:"errors,omitempty"`
	Pods                 []string `protobuf:"bytes,2,rep,name=pods,proto3" json:"pods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputValues) Reset()         { *m = OutputValues{} }
func (m *OutputValues) String() string { return proto.CompactTextString(m) }
func (*OutputValues) ProtoMessage()    {}
func (*OutputValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a704f5ed8d36550, []int{1}
}

func (m *OutputValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputValues.Unmarshal(m, b)
}
func (m *OutputValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputValues.Marshal(b, m, deterministic)
}
func (m *OutputValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputValues.Merge(m, src)
}
func (m *OutputValues) XXX_Size() int {
	return xxx_messageInfo_OutputValues.Size(m)
}
func (m *OutputValues) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputValues.DiscardUnknown(m)
}

var xxx_messageInfo_OutputValues proto.InternalMessageInfo

func (m *OutputValues) GetErrors() string {
	if m != nil {
		return m.Errors
	}
	return ""
}

func (m *OutputValues) GetPods() []string {
	if m != nil {
		return m.Pods
	}
	return nil
}

type ListNameSpaceInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNameSpaceInput) Reset()         { *m = ListNameSpaceInput{} }
func (m *ListNameSpaceInput) String() string { return proto.CompactTextString(m) }
func (*ListNameSpaceInput) ProtoMessage()    {}
func (*ListNameSpaceInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a704f5ed8d36550, []int{2}
}

func (m *ListNameSpaceInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNameSpaceInput.Unmarshal(m, b)
}
func (m *ListNameSpaceInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNameSpaceInput.Marshal(b, m, deterministic)
}
func (m *ListNameSpaceInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNameSpaceInput.Merge(m, src)
}
func (m *ListNameSpaceInput) XXX_Size() int {
	return xxx_messageInfo_ListNameSpaceInput.Size(m)
}
func (m *ListNameSpaceInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNameSpaceInput.DiscardUnknown(m)
}

var xxx_messageInfo_ListNameSpaceInput proto.InternalMessageInfo

type NameSpaceOutput struct {
	NameSpaceList        []string `protobuf:"bytes,1,rep,name=nameSpaceList,proto3" json:"nameSpaceList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameSpaceOutput) Reset()         { *m = NameSpaceOutput{} }
func (m *NameSpaceOutput) String() string { return proto.CompactTextString(m) }
func (*NameSpaceOutput) ProtoMessage()    {}
func (*NameSpaceOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a704f5ed8d36550, []int{3}
}

func (m *NameSpaceOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameSpaceOutput.Unmarshal(m, b)
}
func (m *NameSpaceOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameSpaceOutput.Marshal(b, m, deterministic)
}
func (m *NameSpaceOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameSpaceOutput.Merge(m, src)
}
func (m *NameSpaceOutput) XXX_Size() int {
	return xxx_messageInfo_NameSpaceOutput.Size(m)
}
func (m *NameSpaceOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_NameSpaceOutput.DiscardUnknown(m)
}

var xxx_messageInfo_NameSpaceOutput proto.InternalMessageInfo

func (m *NameSpaceOutput) GetNameSpaceList() []string {
	if m != nil {
		return m.NameSpaceList
	}
	return nil
}

type CheckPodInput struct {
	NameSpace            string   `protobuf:"bytes,1,opt,name=nameSpace,proto3" json:"nameSpace,omitempty"`
	PodName              string   `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPodInput) Reset()         { *m = CheckPodInput{} }
func (m *CheckPodInput) String() string { return proto.CompactTextString(m) }
func (*CheckPodInput) ProtoMessage()    {}
func (*CheckPodInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a704f5ed8d36550, []int{4}
}

func (m *CheckPodInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPodInput.Unmarshal(m, b)
}
func (m *CheckPodInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPodInput.Marshal(b, m, deterministic)
}
func (m *CheckPodInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPodInput.Merge(m, src)
}
func (m *CheckPodInput) XXX_Size() int {
	return xxx_messageInfo_CheckPodInput.Size(m)
}
func (m *CheckPodInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPodInput.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPodInput proto.InternalMessageInfo

func (m *CheckPodInput) GetNameSpace() string {
	if m != nil {
		return m.NameSpace
	}
	return ""
}

func (m *CheckPodInput) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

type CheckPodOutput struct {
	CheckPodOutput       string   `protobuf:"bytes,1,opt,name=checkPodOutput,proto3" json:"checkPodOutput,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPodOutput) Reset()         { *m = CheckPodOutput{} }
func (m *CheckPodOutput) String() string { return proto.CompactTextString(m) }
func (*CheckPodOutput) ProtoMessage()    {}
func (*CheckPodOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a704f5ed8d36550, []int{5}
}

func (m *CheckPodOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPodOutput.Unmarshal(m, b)
}
func (m *CheckPodOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPodOutput.Marshal(b, m, deterministic)
}
func (m *CheckPodOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPodOutput.Merge(m, src)
}
func (m *CheckPodOutput) XXX_Size() int {
	return xxx_messageInfo_CheckPodOutput.Size(m)
}
func (m *CheckPodOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPodOutput.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPodOutput proto.InternalMessageInfo

func (m *CheckPodOutput) GetCheckPodOutput() string {
	if m != nil {
		return m.CheckPodOutput
	}
	return ""
}

func init() {
	proto.RegisterType((*InputValues)(nil), "InputValues")
	proto.RegisterType((*OutputValues)(nil), "OutputValues")
	proto.RegisterType((*ListNameSpaceInput)(nil), "ListNameSpaceInput")
	proto.RegisterType((*NameSpaceOutput)(nil), "NameSpaceOutput")
	proto.RegisterType((*CheckPodInput)(nil), "CheckPodInput")
	proto.RegisterType((*CheckPodOutput)(nil), "CheckPodOutput")
}

func init() { proto.RegisterFile("protofiles/k8services.proto", fileDescriptor_7a704f5ed8d36550) }

var fileDescriptor_7a704f5ed8d36550 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xef, 0x4a, 0xc2, 0x50,
	0x18, 0xc6, 0xd9, 0x0a, 0x73, 0xaf, 0x4e, 0xe5, 0x14, 0x31, 0x96, 0x90, 0x1c, 0x2c, 0xc4, 0x60,
	0xa7, 0x7f, 0x90, 0xf8, 0xad, 0xfa, 0x10, 0x42, 0x94, 0x28, 0x04, 0xf5, 0x6d, 0x6e, 0x27, 0x1d,
	0xce, 0x9d, 0xb1, 0x33, 0x25, 0x90, 0xbe, 0x74, 0x0b, 0xdd, 0x4d, 0xb7, 0xd1, 0x2d, 0x74, 0x21,
	0xb1, 0xe3, 0x36, 0x9d, 0x46, 0xdf, 0x76, 0x9e, 0xf3, 0x9e, 0xe7, 0x79, 0xde, 0x1f, 0x83, 0x03,
	0x3f, 0x60, 0x21, 0x7b, 0x75, 0x5c, 0xca, 0xc9, 0xb8, 0xc5, 0x69, 0x30, 0x73, 0x2c, 0xca, 0x0d,
	0xa1, 0xea, 0xd5, 0x21, 0x63, 0x43, 0x97, 0x12, 0xd3, 0x77, 0x88, 0xe9, 0x79, 0x2c, 0x34, 0x43,
	0x87, 0x79, 0xf1, 0x2d, 0x3e, 0x81, 0x42, 0xc7, 0xf3, 0xa7, 0xe1, 0x93, 0xe9, 0x4e, 0x29, 0x47,
	0x55, 0x50, 0x3c, 0x73, 0x42, 0xfb, 0xbe, 0x69, 0x51, 0x4d, 0xaa, 0x49, 0x0d, 0xa5, 0xb7, 0x14,
	0x70, 0x1b, 0x8a, 0x8f, 0xd3, 0x70, 0x39, 0xbd, 0x0f, 0x39, 0x1a, 0x04, 0x2c, 0xe0, 0xf1, 0x68,
	0x7c, 0x42, 0x08, 0xb6, 0x7d, 0x66, 0x73, 0x4d, 0xae, 0x6d, 0x35, 0x94, 0x9e, 0xf8, 0xc6, 0x7b,
	0x80, 0xee, 0x1d, 0x1e, 0x3e, 0x24, 0x66, 0x22, 0x15, 0x5f, 0x41, 0x39, 0x55, 0x16, 0xd6, 0xa8,
	0x0e, 0x6a, 0x9a, 0x18, 0xbd, 0xd0, 0x24, 0xe1, 0x92, 0x15, 0xf1, 0x1d, 0xa8, 0xb7, 0x23, 0x6a,
	0x8d, 0xbb, 0xcc, 0x16, 0x4e, 0xff, 0x37, 0x47, 0x1a, 0xec, 0xf8, 0xcc, 0x8e, 0xa2, 0x34, 0x59,
	0xdc, 0x25, 0x47, 0xdc, 0x82, 0x52, 0x62, 0x14, 0x17, 0x38, 0x86, 0x92, 0x95, 0x51, 0x62, 0xbb,
	0x35, 0xf5, 0xfc, 0x4b, 0x86, 0x8a, 0xe5, 0x3a, 0xd4, 0x0b, 0xaf, 0xbb, 0x9d, 0xfe, 0x02, 0x3a,
	0x7a, 0x86, 0x7c, 0xd4, 0xaf, 0xcb, 0x6c, 0x8e, 0x8a, 0xc6, 0x0a, 0x5a, 0x5d, 0x35, 0x56, 0xd9,
	0xe1, 0xd3, 0x8f, 0xef, 0x9f, 0x4f, 0xb9, 0x89, 0x8f, 0xc8, 0xec, 0x8c, 0xd0, 0x37, 0x73, 0xe2,
	0xbb, 0x94, 0x50, 0x6b, 0xc4, 0x88, 0x1b, 0xbf, 0x27, 0xf3, 0xb4, 0xff, 0x7b, 0x5b, 0x6a, 0xa2,
	0x01, 0xa8, 0x19, 0x82, 0x68, 0xd7, 0xd8, 0x24, 0xaa, 0x57, 0x8c, 0x35, 0xa0, 0xb8, 0x29, 0x92,
	0xea, 0xf8, 0xf0, 0xcf, 0xa4, 0x74, 0x9a, 0x47, 0x19, 0x43, 0xc8, 0x27, 0x34, 0x50, 0xc9, 0xc8,
	0x10, 0xd6, 0xcb, 0x46, 0x16, 0x14, 0x6e, 0x0b, 0xe3, 0x4b, 0x4c, 0x36, 0x8c, 0x13, 0x52, 0xab,
	0x2b, 0x90, 0x79, 0x8c, 0x3c, 0x5a, 0xe6, 0xa6, 0xf0, 0xa2, 0xa4, 0xec, 0x06, 0x39, 0xf1, 0x2f,
	0x5e, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x64, 0x03, 0x1c, 0x19, 0xc8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientAPIServiceClient is the client API for ClientAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientAPIServiceClient interface {
	ListPods(ctx context.Context, in *InputValues, opts ...grpc.CallOption) (*OutputValues, error)
	ListNameSpace(ctx context.Context, in *ListNameSpaceInput, opts ...grpc.CallOption) (*NameSpaceOutput, error)
	CheckPod(ctx context.Context, in *CheckPodInput, opts ...grpc.CallOption) (*CheckPodOutput, error)
}

type clientAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientAPIServiceClient(cc grpc.ClientConnInterface) ClientAPIServiceClient {
	return &clientAPIServiceClient{cc}
}

func (c *clientAPIServiceClient) ListPods(ctx context.Context, in *InputValues, opts ...grpc.CallOption) (*OutputValues, error) {
	out := new(OutputValues)
	err := c.cc.Invoke(ctx, "/clientAPIService/ListPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAPIServiceClient) ListNameSpace(ctx context.Context, in *ListNameSpaceInput, opts ...grpc.CallOption) (*NameSpaceOutput, error) {
	out := new(NameSpaceOutput)
	err := c.cc.Invoke(ctx, "/clientAPIService/ListNameSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAPIServiceClient) CheckPod(ctx context.Context, in *CheckPodInput, opts ...grpc.CallOption) (*CheckPodOutput, error) {
	out := new(CheckPodOutput)
	err := c.cc.Invoke(ctx, "/clientAPIService/CheckPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAPIServiceServer is the server API for ClientAPIService service.
type ClientAPIServiceServer interface {
	ListPods(context.Context, *InputValues) (*OutputValues, error)
	ListNameSpace(context.Context, *ListNameSpaceInput) (*NameSpaceOutput, error)
	CheckPod(context.Context, *CheckPodInput) (*CheckPodOutput, error)
}

// UnimplementedClientAPIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientAPIServiceServer struct {
}

func (*UnimplementedClientAPIServiceServer) ListPods(ctx context.Context, req *InputValues) (*OutputValues, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPods not implemented")
}
func (*UnimplementedClientAPIServiceServer) ListNameSpace(ctx context.Context, req *ListNameSpaceInput) (*NameSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNameSpace not implemented")
}
func (*UnimplementedClientAPIServiceServer) CheckPod(ctx context.Context, req *CheckPodInput) (*CheckPodOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPod not implemented")
}

func RegisterClientAPIServiceServer(s *grpc.Server, srv ClientAPIServiceServer) {
	s.RegisterService(&_ClientAPIService_serviceDesc, srv)
}

func _ClientAPIService_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputValues)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServiceServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientAPIService/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServiceServer).ListPods(ctx, req.(*InputValues))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAPIService_ListNameSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNameSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServiceServer).ListNameSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientAPIService/ListNameSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServiceServer).ListNameSpace(ctx, req.(*ListNameSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAPIService_CheckPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPodInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServiceServer).CheckPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientAPIService/CheckPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServiceServer).CheckPod(ctx, req.(*CheckPodInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientAPIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clientAPIService",
	HandlerType: (*ClientAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPods",
			Handler:    _ClientAPIService_ListPods_Handler,
		},
		{
			MethodName: "ListNameSpace",
			Handler:    _ClientAPIService_ListNameSpace_Handler,
		},
		{
			MethodName: "CheckPod",
			Handler:    _ClientAPIService_CheckPod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/k8services.proto",
}
