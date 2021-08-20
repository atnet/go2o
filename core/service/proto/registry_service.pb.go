// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry_service.proto

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

type RegistriesResponse struct {
	Value                []*SRegistry `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegistriesResponse) Reset()         { *m = RegistriesResponse{} }
func (m *RegistriesResponse) String() string { return proto.CompactTextString(m) }
func (*RegistriesResponse) ProtoMessage()    {}
func (*RegistriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_service_d11359167f93674f, []int{0}
}
func (m *RegistriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistriesResponse.Unmarshal(m, b)
}
func (m *RegistriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistriesResponse.Marshal(b, m, deterministic)
}
func (dst *RegistriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistriesResponse.Merge(dst, src)
}
func (m *RegistriesResponse) XXX_Size() int {
	return xxx_messageInfo_RegistriesResponse.Size(m)
}
func (m *RegistriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistriesResponse proto.InternalMessageInfo

func (m *RegistriesResponse) GetValue() []*SRegistry {
	if m != nil {
		return m.Value
	}
	return nil
}

type RegistryPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryPair) Reset()         { *m = RegistryPair{} }
func (m *RegistryPair) String() string { return proto.CompactTextString(m) }
func (*RegistryPair) ProtoMessage()    {}
func (*RegistryPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_service_d11359167f93674f, []int{1}
}
func (m *RegistryPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryPair.Unmarshal(m, b)
}
func (m *RegistryPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryPair.Marshal(b, m, deterministic)
}
func (dst *RegistryPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryPair.Merge(dst, src)
}
func (m *RegistryPair) XXX_Size() int {
	return xxx_messageInfo_RegistryPair.Size(m)
}
func (m *RegistryPair) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryPair.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryPair proto.InternalMessageInfo

func (m *RegistryPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RegistryPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RegistryValueResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryValueResponse) Reset()         { *m = RegistryValueResponse{} }
func (m *RegistryValueResponse) String() string { return proto.CompactTextString(m) }
func (*RegistryValueResponse) ProtoMessage()    {}
func (*RegistryValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_service_d11359167f93674f, []int{2}
}
func (m *RegistryValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryValueResponse.Unmarshal(m, b)
}
func (m *RegistryValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryValueResponse.Marshal(b, m, deterministic)
}
func (dst *RegistryValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryValueResponse.Merge(dst, src)
}
func (m *RegistryValueResponse) XXX_Size() int {
	return xxx_messageInfo_RegistryValueResponse.Size(m)
}
func (m *RegistryValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryValueResponse proto.InternalMessageInfo

func (m *RegistryValueResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RegistryValueResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type RegistryCreateRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryCreateRequest) Reset()         { *m = RegistryCreateRequest{} }
func (m *RegistryCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RegistryCreateRequest) ProtoMessage()    {}
func (*RegistryCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_service_d11359167f93674f, []int{3}
}
func (m *RegistryCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryCreateRequest.Unmarshal(m, b)
}
func (m *RegistryCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryCreateRequest.Marshal(b, m, deterministic)
}
func (dst *RegistryCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryCreateRequest.Merge(dst, src)
}
func (m *RegistryCreateRequest) XXX_Size() int {
	return xxx_messageInfo_RegistryCreateRequest.Size(m)
}
func (m *RegistryCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryCreateRequest proto.InternalMessageInfo

func (m *RegistryCreateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RegistryCreateRequest) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func (m *RegistryCreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type RegistrySearchRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrySearchRequest) Reset()         { *m = RegistrySearchRequest{} }
func (m *RegistrySearchRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrySearchRequest) ProtoMessage()    {}
func (*RegistrySearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_service_d11359167f93674f, []int{4}
}
func (m *RegistrySearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrySearchRequest.Unmarshal(m, b)
}
func (m *RegistrySearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrySearchRequest.Marshal(b, m, deterministic)
}
func (dst *RegistrySearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrySearchRequest.Merge(dst, src)
}
func (m *RegistrySearchRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrySearchRequest.Size(m)
}
func (m *RegistrySearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrySearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrySearchRequest proto.InternalMessageInfo

func (m *RegistrySearchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// * 注册表
type SRegistry struct {
	// * 键
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	// * 值
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	// * 默认值
	DefaultValue string `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value"`
	// * 可选值
	Options string `protobuf:"bytes,4,opt,name=options,proto3" json:"options"`
	// * 标志
	Flag int32 `protobuf:"zigzag32,5,opt,name=flag,proto3" json:"flag"`
	// * 描述
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SRegistry) Reset()         { *m = SRegistry{} }
func (m *SRegistry) String() string { return proto.CompactTextString(m) }
func (*SRegistry) ProtoMessage()    {}
func (*SRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_service_d11359167f93674f, []int{5}
}
func (m *SRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SRegistry.Unmarshal(m, b)
}
func (m *SRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SRegistry.Marshal(b, m, deterministic)
}
func (dst *SRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SRegistry.Merge(dst, src)
}
func (m *SRegistry) XXX_Size() int {
	return xxx_messageInfo_SRegistry.Size(m)
}
func (m *SRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_SRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_SRegistry proto.InternalMessageInfo

func (m *SRegistry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SRegistry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SRegistry) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func (m *SRegistry) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *SRegistry) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SRegistry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistriesResponse)(nil), "RegistriesResponse")
	proto.RegisterType((*RegistryPair)(nil), "RegistryPair")
	proto.RegisterType((*RegistryValueResponse)(nil), "RegistryValueResponse")
	proto.RegisterType((*RegistryCreateRequest)(nil), "RegistryCreateRequest")
	proto.RegisterType((*RegistrySearchRequest)(nil), "RegistrySearchRequest")
	proto.RegisterType((*SRegistry)(nil), "SRegistry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryServiceClient is the client API for RegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryServiceClient interface {
	// * 获取注册表键值,key
	GetValue(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistryValueResponse, error)
	// * 更新注册表值
	UpdateValue(ctx context.Context, in *RegistryPair, opts ...grpc.CallOption) (*Result, error)
	// * 获取键值存储数据字典,keys
	GetValues(ctx context.Context, in *StringArray, opts ...grpc.CallOption) (*StringMap, error)
	// * 搜索键值
	Search(ctx context.Context, in *RegistrySearchRequest, opts ...grpc.CallOption) (*StringMap, error)
	// * 按键前缀获取键数据,prefix
	FindRegistries(ctx context.Context, in *String, opts ...grpc.CallOption) (*StringMap, error)
	// * 更新注册表键值
	UpdateValues(ctx context.Context, in *StringMap, opts ...grpc.CallOption) (*Result, error)
	// * 搜索注册表,keyword
	SearchRegistry(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistriesResponse, error)
	// * 创建自定义注册表项,@defaultValue 默认值,如需更改,使用UpdateRegistry方法
	CreateRegistry(ctx context.Context, in *RegistryCreateRequest, opts ...grpc.CallOption) (*Result, error)
}

type registryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistryServiceClient(cc *grpc.ClientConn) RegistryServiceClient {
	return &registryServiceClient{cc}
}

func (c *registryServiceClient) GetValue(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistryValueResponse, error) {
	out := new(RegistryValueResponse)
	err := c.cc.Invoke(ctx, "/RegistryService/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) UpdateValue(ctx context.Context, in *RegistryPair, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/RegistryService/UpdateValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) GetValues(ctx context.Context, in *StringArray, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/RegistryService/GetValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Search(ctx context.Context, in *RegistrySearchRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/RegistryService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) FindRegistries(ctx context.Context, in *String, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/RegistryService/FindRegistries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) UpdateValues(ctx context.Context, in *StringMap, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/RegistryService/UpdateValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) SearchRegistry(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistriesResponse, error) {
	out := new(RegistriesResponse)
	err := c.cc.Invoke(ctx, "/RegistryService/SearchRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) CreateRegistry(ctx context.Context, in *RegistryCreateRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/RegistryService/CreateRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServiceServer is the server API for RegistryService service.
type RegistryServiceServer interface {
	// * 获取注册表键值,key
	GetValue(context.Context, *String) (*RegistryValueResponse, error)
	// * 更新注册表值
	UpdateValue(context.Context, *RegistryPair) (*Result, error)
	// * 获取键值存储数据字典,keys
	GetValues(context.Context, *StringArray) (*StringMap, error)
	// * 搜索键值
	Search(context.Context, *RegistrySearchRequest) (*StringMap, error)
	// * 按键前缀获取键数据,prefix
	FindRegistries(context.Context, *String) (*StringMap, error)
	// * 更新注册表键值
	UpdateValues(context.Context, *StringMap) (*Result, error)
	// * 搜索注册表,keyword
	SearchRegistry(context.Context, *String) (*RegistriesResponse, error)
	// * 创建自定义注册表项,@defaultValue 默认值,如需更改,使用UpdateRegistry方法
	CreateRegistry(context.Context, *RegistryCreateRequest) (*Result, error)
}

func RegisterRegistryServiceServer(s *grpc.Server, srv RegistryServiceServer) {
	s.RegisterService(&_RegistryService_serviceDesc, srv)
}

func _RegistryService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetValue(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_UpdateValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).UpdateValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/UpdateValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).UpdateValue(ctx, req.(*RegistryPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetValues(ctx, req.(*StringArray))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Search(ctx, req.(*RegistrySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_FindRegistries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).FindRegistries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/FindRegistries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).FindRegistries(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_UpdateValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).UpdateValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/UpdateValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).UpdateValues(ctx, req.(*StringMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_SearchRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).SearchRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/SearchRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).SearchRegistry(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_CreateRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).CreateRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/CreateRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).CreateRegistry(ctx, req.(*RegistryCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RegistryService",
	HandlerType: (*RegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _RegistryService_GetValue_Handler,
		},
		{
			MethodName: "UpdateValue",
			Handler:    _RegistryService_UpdateValue_Handler,
		},
		{
			MethodName: "GetValues",
			Handler:    _RegistryService_GetValues_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _RegistryService_Search_Handler,
		},
		{
			MethodName: "FindRegistries",
			Handler:    _RegistryService_FindRegistries_Handler,
		},
		{
			MethodName: "UpdateValues",
			Handler:    _RegistryService_UpdateValues_Handler,
		},
		{
			MethodName: "SearchRegistry",
			Handler:    _RegistryService_SearchRegistry_Handler,
		},
		{
			MethodName: "CreateRegistry",
			Handler:    _RegistryService_CreateRegistry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry_service.proto",
}

func init() {
	proto.RegisterFile("registry_service.proto", fileDescriptor_registry_service_d11359167f93674f)
}

var fileDescriptor_registry_service_d11359167f93674f = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0xad, 0x5d, 0x5f, 0xb3, 0x02, 0xde, 0x18, 0x51, 0x2f, 0x44, 0x41, 0x68, 0xe5,
	0x80, 0x41, 0x9d, 0xb4, 0x0b, 0x27, 0x40, 0x1a, 0xa7, 0x49, 0x28, 0x15, 0x1c, 0xb8, 0x54, 0x6e,
	0xfa, 0xe6, 0x59, 0x64, 0x75, 0xb0, 0x9d, 0x49, 0xf9, 0x8f, 0xf8, 0x17, 0xb9, 0xa1, 0x39, 0x4e,
	0x9d, 0xfe, 0x40, 0xe2, 0x94, 0xf8, 0xbd, 0xef, 0x7d, 0xdf, 0xf7, 0x9e, 0x9f, 0xe1, 0x5c, 0x21,
	0x17, 0xda, 0xa8, 0x6a, 0xae, 0x51, 0x3d, 0x88, 0x0c, 0x69, 0xa1, 0xa4, 0x91, 0xe3, 0x90, 0xe7,
	0x72, 0xc1, 0xf2, 0xfa, 0x94, 0x5c, 0x01, 0x49, 0x6b, 0x9c, 0x40, 0x9d, 0xa2, 0x2e, 0xe4, 0x4a,
	0x23, 0x89, 0xe1, 0xe8, 0x81, 0xe5, 0x25, 0x46, 0x41, 0xdc, 0x9d, 0x0c, 0xa7, 0x40, 0x67, 0x0e,
	0x54, 0xa5, 0x75, 0x22, 0xb9, 0x82, 0xb0, 0x09, 0x7d, 0x65, 0x42, 0x91, 0xa7, 0xd0, 0xfd, 0x89,
	0x55, 0x14, 0xc4, 0xc1, 0x64, 0x90, 0x3e, 0xfe, 0x92, 0xb3, 0x86, 0xe3, 0xc0, 0xc6, 0x5c, 0xdd,
	0x35, 0x3c, 0x6f, 0xea, 0xbe, 0x3f, 0x06, 0xd6, 0x92, 0x67, 0x5e, 0xd2, 0xc3, 0xc9, 0x0b, 0xe8,
	0xa3, 0x52, 0xf3, 0x7b, 0xcd, 0x1d, 0x4d, 0x0f, 0x95, 0xba, 0xd1, 0x3c, 0x29, 0x3c, 0xcf, 0x67,
	0x85, 0xcc, 0x60, 0x8a, 0xbf, 0x4a, 0xd4, 0x66, 0x8f, 0x91, 0x57, 0x70, 0xb2, 0xc4, 0x5b, 0x56,
	0xe6, 0x66, 0xde, 0x36, 0x14, 0xba, 0xa0, 0xb5, 0x41, 0x62, 0x18, 0x2e, 0x51, 0x67, 0x4a, 0x14,
	0x46, 0xc8, 0x55, 0xd4, 0xb5, 0x90, 0x76, 0x28, 0x79, 0xe3, 0x15, 0x67, 0xc8, 0x54, 0x76, 0xf7,
	0x4f, 0xc5, 0xe4, 0x77, 0x00, 0x83, 0xf5, 0xc4, 0xfe, 0x77, 0x34, 0xbb, 0x3e, 0xbb, 0x7b, 0x7c,
	0x46, 0xd0, 0x97, 0xd6, 0x8f, 0x8e, 0x0e, 0x6d, 0xba, 0x39, 0x12, 0x02, 0x87, 0xb7, 0x39, 0xe3,
	0xd1, 0x51, 0x1c, 0x4c, 0x9e, 0xa5, 0xf6, 0x7f, 0xbb, 0xab, 0xde, 0x4e, 0x57, 0xd3, 0x3f, 0x07,
	0xf0, 0xc4, 0xb7, 0x65, 0xf7, 0x84, 0xbc, 0x85, 0xe3, 0x2f, 0xe8, 0xf4, 0xfa, 0x74, 0x66, 0x94,
	0x58, 0xf1, 0xf1, 0x39, 0xdd, 0x7b, 0x6f, 0x49, 0x87, 0x5c, 0xc0, 0xf0, 0x5b, 0xb1, 0x64, 0x06,
	0xeb, 0x8a, 0x13, 0xda, 0x5e, 0x8c, 0x71, 0x9f, 0xa6, 0xa8, 0xcb, 0xdc, 0x58, 0xe0, 0xa0, 0xe1,
	0xd5, 0x24, 0x74, 0xc4, 0x1f, 0x95, 0x62, 0xd5, 0x18, 0xdc, 0xe9, 0x86, 0x15, 0x49, 0x87, 0x50,
	0xe8, 0xd5, 0x23, 0x26, 0x5e, 0x75, 0x63, 0xe6, 0x5b, 0xf8, 0x0b, 0x18, 0x5d, 0x8b, 0xd5, 0xd2,
	0x2f, 0xb2, 0xb7, 0xbd, 0x09, 0x7c, 0x0d, 0x61, 0xcb, 0xaa, 0x26, 0xad, 0x6c, 0xdb, 0xe8, 0x7b,
	0x18, 0x35, 0x72, 0xee, 0x0e, 0xd7, 0x7c, 0xa7, 0x74, 0xf7, 0xb9, 0x24, 0x1d, 0x72, 0x09, 0xa3,
	0x66, 0x0d, 0x5d, 0x85, 0x77, 0xbe, 0xb1, 0x9f, 0x2d, 0x99, 0x4f, 0x2f, 0xe1, 0x34, 0x93, 0xf7,
	0x94, 0x0b, 0x73, 0x57, 0x2e, 0x28, 0x97, 0x53, 0x49, 0x55, 0x91, 0xfd, 0x38, 0xa6, 0xef, 0x3e,
	0xd8, 0xc7, 0xb9, 0xe8, 0xd9, 0xcf, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x49, 0xd7,
	0x4f, 0xcb, 0x03, 0x00, 0x00,
}
