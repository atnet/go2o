// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code_service.proto

package proto // import "./"

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

type SaveQrTemplateRequest struct {
	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	// * 模板标题
	Title string `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"Title"`
	// * 背景图片
	BgImage string `protobuf:"bytes,3,opt,name=BgImage,json=bgImage,proto3" json:"BgImage"`
	// * 垂直偏离量
	OffsetX int32 `protobuf:"varint,4,opt,name=OffsetX,json=offsetX,proto3" json:"OffsetX"`
	// * 垂直偏移量
	OffsetY int32 `protobuf:"varint,5,opt,name=OffsetY,json=offsetY,proto3" json:"OffsetY"`
	// * 二维码模板文本
	Comment string `protobuf:"bytes,6,opt,name=Comment,json=comment,proto3" json:"Comment"`
	// * 回调地址
	CallbackUrl string `protobuf:"bytes,7,opt,name=CallbackUrl,json=callbackUrl,proto3" json:"CallbackUrl"`
	// * 是否启用
	Enabled              int32    `protobuf:"varint,8,opt,name=Enabled,json=enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveQrTemplateRequest) Reset()         { *m = SaveQrTemplateRequest{} }
func (m *SaveQrTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*SaveQrTemplateRequest) ProtoMessage()    {}
func (*SaveQrTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_326a43dc9ca342fb, []int{0}
}
func (m *SaveQrTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveQrTemplateRequest.Unmarshal(m, b)
}
func (m *SaveQrTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveQrTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *SaveQrTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveQrTemplateRequest.Merge(dst, src)
}
func (m *SaveQrTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_SaveQrTemplateRequest.Size(m)
}
func (m *SaveQrTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveQrTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveQrTemplateRequest proto.InternalMessageInfo

func (m *SaveQrTemplateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaveQrTemplateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetBgImage() string {
	if m != nil {
		return m.BgImage
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetOffsetX() int32 {
	if m != nil {
		return m.OffsetX
	}
	return 0
}

func (m *SaveQrTemplateRequest) GetOffsetY() int32 {
	if m != nil {
		return m.OffsetY
	}
	return 0
}

func (m *SaveQrTemplateRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetCallbackUrl() string {
	if m != nil {
		return m.CallbackUrl
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

type SaveQrTemplateResponse struct {
	ErrCode              int64    `protobuf:"varint,1,opt,name=ErrCode,json=errCode,proto3" json:"ErrCode"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,json=errMsg,proto3" json:"ErrMsg"`
	Id                   int64    `protobuf:"varint,3,opt,name=Id,json=id,proto3" json:"Id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveQrTemplateResponse) Reset()         { *m = SaveQrTemplateResponse{} }
func (m *SaveQrTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*SaveQrTemplateResponse) ProtoMessage()    {}
func (*SaveQrTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_326a43dc9ca342fb, []int{1}
}
func (m *SaveQrTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveQrTemplateResponse.Unmarshal(m, b)
}
func (m *SaveQrTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveQrTemplateResponse.Marshal(b, m, deterministic)
}
func (dst *SaveQrTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveQrTemplateResponse.Merge(dst, src)
}
func (m *SaveQrTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_SaveQrTemplateResponse.Size(m)
}
func (m *SaveQrTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveQrTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveQrTemplateResponse proto.InternalMessageInfo

func (m *SaveQrTemplateResponse) GetErrCode() int64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SaveQrTemplateResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SaveQrTemplateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CommQrTemplateId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommQrTemplateId) Reset()         { *m = CommQrTemplateId{} }
func (m *CommQrTemplateId) String() string { return proto.CompactTextString(m) }
func (*CommQrTemplateId) ProtoMessage()    {}
func (*CommQrTemplateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_326a43dc9ca342fb, []int{2}
}
func (m *CommQrTemplateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommQrTemplateId.Unmarshal(m, b)
}
func (m *CommQrTemplateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommQrTemplateId.Marshal(b, m, deterministic)
}
func (dst *CommQrTemplateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommQrTemplateId.Merge(dst, src)
}
func (m *CommQrTemplateId) XXX_Size() int {
	return xxx_messageInfo_CommQrTemplateId.Size(m)
}
func (m *CommQrTemplateId) XXX_DiscardUnknown() {
	xxx_messageInfo_CommQrTemplateId.DiscardUnknown(m)
}

var xxx_messageInfo_CommQrTemplateId proto.InternalMessageInfo

func (m *CommQrTemplateId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SQrTemplate struct {
	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	// * 模板标题
	Title string `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"Title"`
	// * 背景图片
	BgImage string `protobuf:"bytes,3,opt,name=BgImage,json=bgImage,proto3" json:"BgImage"`
	// * 垂直偏离量
	OffsetX int32 `protobuf:"varint,4,opt,name=OffsetX,json=offsetX,proto3" json:"OffsetX"`
	// * 垂直偏移量
	OffsetY int32 `protobuf:"varint,5,opt,name=OffsetY,json=offsetY,proto3" json:"OffsetY"`
	// * 二维码模板文本
	Comment string `protobuf:"bytes,6,opt,name=Comment,json=comment,proto3" json:"Comment"`
	// * 回调地址
	CallbackUrl string `protobuf:"bytes,7,opt,name=CallbackUrl,json=callbackUrl,proto3" json:"CallbackUrl"`
	// * 是否启用
	Enabled              int32    `protobuf:"varint,8,opt,name=Enabled,json=enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SQrTemplate) Reset()         { *m = SQrTemplate{} }
func (m *SQrTemplate) String() string { return proto.CompactTextString(m) }
func (*SQrTemplate) ProtoMessage()    {}
func (*SQrTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_326a43dc9ca342fb, []int{3}
}
func (m *SQrTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SQrTemplate.Unmarshal(m, b)
}
func (m *SQrTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SQrTemplate.Marshal(b, m, deterministic)
}
func (dst *SQrTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SQrTemplate.Merge(dst, src)
}
func (m *SQrTemplate) XXX_Size() int {
	return xxx_messageInfo_SQrTemplate.Size(m)
}
func (m *SQrTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SQrTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SQrTemplate proto.InternalMessageInfo

func (m *SQrTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SQrTemplate) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SQrTemplate) GetBgImage() string {
	if m != nil {
		return m.BgImage
	}
	return ""
}

func (m *SQrTemplate) GetOffsetX() int32 {
	if m != nil {
		return m.OffsetX
	}
	return 0
}

func (m *SQrTemplate) GetOffsetY() int32 {
	if m != nil {
		return m.OffsetY
	}
	return 0
}

func (m *SQrTemplate) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *SQrTemplate) GetCallbackUrl() string {
	if m != nil {
		return m.CallbackUrl
	}
	return ""
}

func (m *SQrTemplate) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func init() {
	proto.RegisterType((*SaveQrTemplateRequest)(nil), "SaveQrTemplateRequest")
	proto.RegisterType((*SaveQrTemplateResponse)(nil), "SaveQrTemplateResponse")
	proto.RegisterType((*CommQrTemplateId)(nil), "CommQrTemplateId")
	proto.RegisterType((*SQrTemplate)(nil), "SQrTemplate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CodeServiceClient is the client API for CodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodeServiceClient interface {
	// 保存CommQrTemplate
	SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error)
}

type codeServiceClient struct {
	cc *grpc.ClientConn
}

func NewCodeServiceClient(cc *grpc.ClientConn) CodeServiceClient {
	return &codeServiceClient{cc}
}

func (c *codeServiceClient) SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error) {
	out := new(SaveQrTemplateResponse)
	err := c.cc.Invoke(ctx, "/CodeService/SaveQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error) {
	out := new(SQrTemplate)
	err := c.cc.Invoke(ctx, "/CodeService/GetQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CodeService/DeleteQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServiceServer is the server API for CodeService service.
type CodeServiceServer interface {
	// 保存CommQrTemplate
	SaveQrTemplate(context.Context, *SaveQrTemplateRequest) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(context.Context, *CommQrTemplateId) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(context.Context, *CommQrTemplateId) (*Result, error)
}

func RegisterCodeServiceServer(s *grpc.Server, srv CodeServiceServer) {
	s.RegisterService(&_CodeService_serviceDesc, srv)
}

func _CodeService_SaveQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQrTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/SaveQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, req.(*SaveQrTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_GetQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/GetQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_DeleteQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/DeleteQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CodeService",
	HandlerType: (*CodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveQrTemplate",
			Handler:    _CodeService_SaveQrTemplate_Handler,
		},
		{
			MethodName: "GetQrTemplate",
			Handler:    _CodeService_GetQrTemplate_Handler,
		},
		{
			MethodName: "DeleteQrTemplate",
			Handler:    _CodeService_DeleteQrTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_service.proto",
}

func init() { proto.RegisterFile("code_service.proto", fileDescriptor_code_service_326a43dc9ca342fb) }

var fileDescriptor_code_service_326a43dc9ca342fb = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xc1, 0x6e, 0x9b, 0x30,
	0x18, 0xc7, 0x43, 0x32, 0x20, 0x33, 0x59, 0x94, 0x79, 0x5b, 0x66, 0xe5, 0x32, 0xc4, 0x89, 0x93,
	0x37, 0x65, 0xc7, 0xdd, 0xc2, 0xa2, 0x29, 0x87, 0x69, 0x1a, 0xc9, 0xa6, 0x26, 0x97, 0xca, 0xc0,
	0x17, 0x8a, 0x6a, 0x62, 0x6a, 0x4c, 0x9e, 0xad, 0x6f, 0xd2, 0x73, 0xdf, 0xa4, 0x02, 0x27, 0x22,
	0x8d, 0xa2, 0x3e, 0x40, 0x4f, 0xe8, 0xe7, 0x3f, 0xfe, 0xcb, 0xdf, 0x0f, 0x83, 0x70, 0x2c, 0x12,
	0xb8, 0x2e, 0x41, 0xee, 0xb3, 0x18, 0x68, 0x21, 0x85, 0x12, 0x93, 0x41, 0xca, 0x45, 0xc4, 0xb8,
	0x26, 0xef, 0xd1, 0x40, 0x9f, 0x96, 0x6c, 0x0f, 0x7f, 0xe5, 0x0a, 0xf2, 0x82, 0x33, 0x05, 0x21,
	0xdc, 0x55, 0x50, 0x2a, 0x3c, 0x44, 0xdd, 0x45, 0x42, 0x0c, 0xd7, 0xf0, 0x7b, 0x61, 0x37, 0x4b,
	0xf0, 0x47, 0x64, 0xae, 0x32, 0xc5, 0x81, 0x74, 0x5d, 0xc3, 0x7f, 0x1b, 0x9a, 0xaa, 0x06, 0x4c,
	0x90, 0x3d, 0x4b, 0x17, 0x39, 0x4b, 0x81, 0xf4, 0x9a, 0x75, 0x3b, 0xd2, 0x58, 0x27, 0x7f, 0xb6,
	0xdb, 0x12, 0xd4, 0x15, 0x79, 0xe3, 0x1a, 0xbe, 0x19, 0xda, 0x42, 0x63, 0x9b, 0xac, 0x89, 0x79,
	0x9a, 0xac, 0xeb, 0x24, 0x10, 0x79, 0x0e, 0x3b, 0x45, 0x2c, 0xdd, 0x16, 0x6b, 0xc4, 0x2e, 0x72,
	0x02, 0xc6, 0x79, 0xc4, 0xe2, 0xdb, 0x7f, 0x92, 0x13, 0xbb, 0x49, 0x9d, 0xb8, 0x5d, 0xaa, 0xf7,
	0xce, 0x77, 0x2c, 0xe2, 0x90, 0x90, 0xbe, 0x6e, 0x05, 0x8d, 0xde, 0x06, 0x8d, 0xcf, 0x47, 0x2c,
	0x0b, 0xb1, 0x2b, 0x9b, 0x33, 0xce, 0xa5, 0x0c, 0x44, 0x02, 0x87, 0x41, 0x6d, 0xd0, 0x88, 0xc7,
	0xc8, 0x9a, 0x4b, 0xf9, 0xbb, 0x4c, 0x0f, 0xe3, 0x5a, 0xd0, 0xd0, 0xc1, 0x4a, 0xef, 0x68, 0xc5,
	0xf3, 0xd1, 0xa8, 0x3e, 0x71, 0xdb, 0xbd, 0x68, 0x4c, 0xfd, 0x67, 0xbc, 0x3a, 0x76, 0x9a, 0xfb,
	0x1a, 0xbc, 0x07, 0x03, 0x39, 0xcb, 0xf6, 0xbd, 0xd7, 0xe3, 0x77, 0x7a, 0x6f, 0x20, 0xa7, 0x96,
	0xb6, 0xd4, 0xf7, 0x0c, 0x07, 0x68, 0xf8, 0xdc, 0x37, 0x1e, 0xd3, 0x8b, 0x77, 0x6c, 0xf2, 0x99,
	0x5e, 0xfe, 0x30, 0x5e, 0x07, 0x4f, 0xd1, 0xbb, 0x5f, 0xa0, 0x4e, 0x3a, 0xde, 0xd3, 0x73, 0xd1,
	0x93, 0x01, 0x3d, 0x11, 0xea, 0x75, 0xf0, 0x37, 0x34, 0xfa, 0x09, 0x1c, 0x14, 0xbc, 0xbc, 0xcd,
	0xa6, 0x21, 0x94, 0x15, 0x57, 0x5e, 0x67, 0xf6, 0x05, 0x7d, 0x88, 0x45, 0x4e, 0xd3, 0x4c, 0xdd,
	0x54, 0x11, 0x4d, 0xc5, 0x54, 0x50, 0x59, 0xc4, 0x9b, 0x3e, 0xfd, 0xfa, 0xa3, 0xf9, 0x3f, 0x22,
	0xab, 0x79, 0x7c, 0x7f, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x03, 0xb3, 0x6c, 0x96, 0x4a, 0x03, 0x00,
	0x00,
}
