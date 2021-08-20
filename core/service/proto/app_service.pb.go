// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app_service.proto

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

type AppId struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppId) Reset()         { *m = AppId{} }
func (m *AppId) String() string { return proto.CompactTextString(m) }
func (*AppId) ProtoMessage()    {}
func (*AppId) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{0}
}
func (m *AppId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppId.Unmarshal(m, b)
}
func (m *AppId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppId.Marshal(b, m, deterministic)
}
func (dst *AppId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppId.Merge(dst, src)
}
func (m *AppId) XXX_Size() int {
	return xxx_messageInfo_AppId.Size(m)
}
func (m *AppId) XXX_DiscardUnknown() {
	xxx_messageInfo_AppId.DiscardUnknown(m)
}

var xxx_messageInfo_AppId proto.InternalMessageInfo

func (m *AppId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AppVersionId struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppVersionId) Reset()         { *m = AppVersionId{} }
func (m *AppVersionId) String() string { return proto.CompactTextString(m) }
func (*AppVersionId) ProtoMessage()    {}
func (*AppVersionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{1}
}
func (m *AppVersionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppVersionId.Unmarshal(m, b)
}
func (m *AppVersionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppVersionId.Marshal(b, m, deterministic)
}
func (dst *AppVersionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppVersionId.Merge(dst, src)
}
func (m *AppVersionId) XXX_Size() int {
	return xxx_messageInfo_AppVersionId.Size(m)
}
func (m *AppVersionId) XXX_DiscardUnknown() {
	xxx_messageInfo_AppVersionId.DiscardUnknown(m)
}

var xxx_messageInfo_AppVersionId proto.InternalMessageInfo

func (m *AppVersionId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 检查版本请求
type CheckVersionRequest struct {
	// 版本号
	AppId int64 `protobuf:"varint,1,opt,name=App_id,json=AppId,proto3" json:"App_id"`
	// 更新通道, stable|beta|nightly
	Channel string `protobuf:"bytes,2,opt,name=Channel,proto3" json:"Channel"`
	// 当前版本
	Version              string   `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckVersionRequest) Reset()         { *m = CheckVersionRequest{} }
func (m *CheckVersionRequest) String() string { return proto.CompactTextString(m) }
func (*CheckVersionRequest) ProtoMessage()    {}
func (*CheckVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{2}
}
func (m *CheckVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckVersionRequest.Unmarshal(m, b)
}
func (m *CheckVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckVersionRequest.Marshal(b, m, deterministic)
}
func (dst *CheckVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckVersionRequest.Merge(dst, src)
}
func (m *CheckVersionRequest) XXX_Size() int {
	return xxx_messageInfo_CheckVersionRequest.Size(m)
}
func (m *CheckVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckVersionRequest proto.InternalMessageInfo

func (m *CheckVersionRequest) GetAppId() int64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *CheckVersionRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CheckVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// 检测版本响应结果
type CheckVersionResponse struct {
	// 最新版本号
	LatestVersion string `protobuf:"bytes,1,opt,name=LatestVersion,proto3" json:"LatestVersion"`
	// App更新资源地址
	AppPkgURL string `protobuf:"bytes,2,opt,name=AppPkgURL,proto3" json:"AppPkgURL"`
	// 版本信息
	VersionInfo string `protobuf:"bytes,3,opt,name=VersionInfo,proto3" json:"VersionInfo"`
	// 是否为最新版本
	IsLatest bool `protobuf:"varint,4,opt,name=IsLatest,proto3" json:"IsLatest"`
	// 是否强制升级
	ForceUpdate bool `protobuf:"varint,5,opt,name=ForceUpdate,proto3" json:"ForceUpdate"`
	// 更新文件类型,如APK,EXE,ZIP等
	UpdateType string `protobuf:"bytes,6,opt,name=Update_type,json=UpdateType,proto3" json:"Update_type"`
	// 发布时间
	ReleaseTime          int64    `protobuf:"varint,7,opt,name=ReleaseTime,proto3" json:"ReleaseTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckVersionResponse) Reset()         { *m = CheckVersionResponse{} }
func (m *CheckVersionResponse) String() string { return proto.CompactTextString(m) }
func (*CheckVersionResponse) ProtoMessage()    {}
func (*CheckVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{3}
}
func (m *CheckVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckVersionResponse.Unmarshal(m, b)
}
func (m *CheckVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckVersionResponse.Marshal(b, m, deterministic)
}
func (dst *CheckVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckVersionResponse.Merge(dst, src)
}
func (m *CheckVersionResponse) XXX_Size() int {
	return xxx_messageInfo_CheckVersionResponse.Size(m)
}
func (m *CheckVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckVersionResponse proto.InternalMessageInfo

func (m *CheckVersionResponse) GetLatestVersion() string {
	if m != nil {
		return m.LatestVersion
	}
	return ""
}

func (m *CheckVersionResponse) GetAppPkgURL() string {
	if m != nil {
		return m.AppPkgURL
	}
	return ""
}

func (m *CheckVersionResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *CheckVersionResponse) GetIsLatest() bool {
	if m != nil {
		return m.IsLatest
	}
	return false
}

func (m *CheckVersionResponse) GetForceUpdate() bool {
	if m != nil {
		return m.ForceUpdate
	}
	return false
}

func (m *CheckVersionResponse) GetUpdateType() string {
	if m != nil {
		return m.UpdateType
	}
	return ""
}

func (m *CheckVersionResponse) GetReleaseTime() int64 {
	if m != nil {
		return m.ReleaseTime
	}
	return 0
}

// APP产品
type AppProdRequest struct {
	// 产品编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 产品名称
	ProdName string `protobuf:"bytes,2,opt,name=ProdName,proto3" json:"ProdName"`
	// 产品描述
	ProdDes string `protobuf:"bytes,3,opt,name=ProdDes,proto3" json:"ProdDes"`
	// Icon
	Icon string `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon"`
	// 发布下载页面地址
	PublishUrl string `protobuf:"bytes,5,opt,name=PublishUrl,proto3" json:"PublishUrl"`
	// 正式版文件地址
	StableFileUrl string `protobuf:"bytes,6,opt,name=StableFileUrl,proto3" json:"StableFileUrl"`
	// 内测版文件地址
	AlphaFileUrl string `protobuf:"bytes,8,opt,name=AlphaFileUrl,proto3" json:"AlphaFileUrl"`
	// 每夜版文件地址
	NightlyFileUrl string `protobuf:"bytes,10,opt,name=NightlyFileUrl,proto3" json:"NightlyFileUrl"`
	// 更新方式,比如APK, EXE等
	UpdateType           string   `protobuf:"bytes,11,opt,name=Update_type,json=UpdateType,proto3" json:"Update_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppProdRequest) Reset()         { *m = AppProdRequest{} }
func (m *AppProdRequest) String() string { return proto.CompactTextString(m) }
func (*AppProdRequest) ProtoMessage()    {}
func (*AppProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{4}
}
func (m *AppProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppProdRequest.Unmarshal(m, b)
}
func (m *AppProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppProdRequest.Marshal(b, m, deterministic)
}
func (dst *AppProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProdRequest.Merge(dst, src)
}
func (m *AppProdRequest) XXX_Size() int {
	return xxx_messageInfo_AppProdRequest.Size(m)
}
func (m *AppProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppProdRequest proto.InternalMessageInfo

func (m *AppProdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppProdRequest) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func (m *AppProdRequest) GetProdDes() string {
	if m != nil {
		return m.ProdDes
	}
	return ""
}

func (m *AppProdRequest) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *AppProdRequest) GetPublishUrl() string {
	if m != nil {
		return m.PublishUrl
	}
	return ""
}

func (m *AppProdRequest) GetStableFileUrl() string {
	if m != nil {
		return m.StableFileUrl
	}
	return ""
}

func (m *AppProdRequest) GetAlphaFileUrl() string {
	if m != nil {
		return m.AlphaFileUrl
	}
	return ""
}

func (m *AppProdRequest) GetNightlyFileUrl() string {
	if m != nil {
		return m.NightlyFileUrl
	}
	return ""
}

func (m *AppProdRequest) GetUpdateType() string {
	if m != nil {
		return m.UpdateType
	}
	return ""
}

// APP版本
type AppVersionRequest struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 产品
	ProductId int64 `protobuf:"varint,2,opt,name=Product_id,json=ProductId,proto3" json:"Product_id"`
	// 更新通道, stable:0|alpha:1|nightly:2
	Channel int32 `protobuf:"varint,3,opt,name=Channel,proto3" json:"Channel"`
	// 版本号
	Version string `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version"`
	// 是否强制升级
	ForceUpdate bool `protobuf:"varint,5,opt,name=ForceUpdate,proto3" json:"ForceUpdate"`
	// 更新内容
	UpdateContent        string   `protobuf:"bytes,6,opt,name=UpdateContent,proto3" json:"UpdateContent"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppVersionRequest) Reset()         { *m = AppVersionRequest{} }
func (m *AppVersionRequest) String() string { return proto.CompactTextString(m) }
func (*AppVersionRequest) ProtoMessage()    {}
func (*AppVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{5}
}
func (m *AppVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppVersionRequest.Unmarshal(m, b)
}
func (m *AppVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppVersionRequest.Marshal(b, m, deterministic)
}
func (dst *AppVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppVersionRequest.Merge(dst, src)
}
func (m *AppVersionRequest) XXX_Size() int {
	return xxx_messageInfo_AppVersionRequest.Size(m)
}
func (m *AppVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppVersionRequest proto.InternalMessageInfo

func (m *AppVersionRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppVersionRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *AppVersionRequest) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *AppVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppVersionRequest) GetForceUpdate() bool {
	if m != nil {
		return m.ForceUpdate
	}
	return false
}

func (m *AppVersionRequest) GetUpdateContent() string {
	if m != nil {
		return m.UpdateContent
	}
	return ""
}

// APP产品
type SAppProd struct {
	// 产品编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 产品名称
	ProdName string `protobuf:"bytes,2,opt,name=ProdName,proto3" json:"ProdName"`
	// 产品描述
	ProdDes string `protobuf:"bytes,3,opt,name=ProdDes,proto3" json:"ProdDes"`
	// 最新的版本ID
	LatestVid int64 `protobuf:"varint,4,opt,name=LatestVid,proto3" json:"LatestVid"`
	// 正式版文件hash值
	Md5Hash string `protobuf:"bytes,5,opt,name=Md5Hash,proto3" json:"Md5Hash"`
	// 发布下载页面地址
	PublishUrl string `protobuf:"bytes,6,opt,name=PublishUrl,proto3" json:"PublishUrl"`
	// 正式版文件地址
	StableFileUrl string `protobuf:"bytes,7,opt,name=StableFileUrl,proto3" json:"StableFileUrl"`
	// 内测版文件地址
	AlphaFileUrl string `protobuf:"bytes,8,opt,name=AlphaFileUrl,proto3" json:"AlphaFileUrl"`
	// 每夜版文件地址
	NightlyFileUrl string `protobuf:"bytes,9,opt,name=NightlyFileUrl,proto3" json:"NightlyFileUrl"`
	// 更新方式,比如APK, EXE等
	UpdateType string `protobuf:"bytes,10,opt,name=Update_type,json=UpdateType,proto3" json:"Update_type"`
	// 更新时间
	UpdateTime           int64    `protobuf:"varint,11,opt,name=update_time,json=updateTime,proto3" json:"update_time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAppProd) Reset()         { *m = SAppProd{} }
func (m *SAppProd) String() string { return proto.CompactTextString(m) }
func (*SAppProd) ProtoMessage()    {}
func (*SAppProd) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{6}
}
func (m *SAppProd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAppProd.Unmarshal(m, b)
}
func (m *SAppProd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAppProd.Marshal(b, m, deterministic)
}
func (dst *SAppProd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAppProd.Merge(dst, src)
}
func (m *SAppProd) XXX_Size() int {
	return xxx_messageInfo_SAppProd.Size(m)
}
func (m *SAppProd) XXX_DiscardUnknown() {
	xxx_messageInfo_SAppProd.DiscardUnknown(m)
}

var xxx_messageInfo_SAppProd proto.InternalMessageInfo

func (m *SAppProd) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAppProd) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func (m *SAppProd) GetProdDes() string {
	if m != nil {
		return m.ProdDes
	}
	return ""
}

func (m *SAppProd) GetLatestVid() int64 {
	if m != nil {
		return m.LatestVid
	}
	return 0
}

func (m *SAppProd) GetMd5Hash() string {
	if m != nil {
		return m.Md5Hash
	}
	return ""
}

func (m *SAppProd) GetPublishUrl() string {
	if m != nil {
		return m.PublishUrl
	}
	return ""
}

func (m *SAppProd) GetStableFileUrl() string {
	if m != nil {
		return m.StableFileUrl
	}
	return ""
}

func (m *SAppProd) GetAlphaFileUrl() string {
	if m != nil {
		return m.AlphaFileUrl
	}
	return ""
}

func (m *SAppProd) GetNightlyFileUrl() string {
	if m != nil {
		return m.NightlyFileUrl
	}
	return ""
}

func (m *SAppProd) GetUpdateType() string {
	if m != nil {
		return m.UpdateType
	}
	return ""
}

func (m *SAppProd) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

// APP版本
type SAppVersion struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 产品
	ProductId int64 `protobuf:"varint,2,opt,name=Product_id,json=ProductId,proto3" json:"Product_id"`
	// 更新通道, 0:stable|1:beta|2:nightly
	Channel int32 `protobuf:"varint,3,opt,name=Channel,proto3" json:"Channel"`
	// 版本号
	Version string `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version"`
	// 数字版本
	VersionCode int32 `protobuf:"varint,5,opt,name=VersionCode,proto3" json:"VersionCode"`
	// 是否强制升级
	ForceUpdate bool `protobuf:"varint,6,opt,name=ForceUpdate,proto3" json:"ForceUpdate"`
	// 更新内容
	UpdateContent string `protobuf:"bytes,7,opt,name=UpdateContent,proto3" json:"UpdateContent"`
	// 发布时间
	CreateTime           int64    `protobuf:"varint,8,opt,name=CreateTime,proto3" json:"CreateTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAppVersion) Reset()         { *m = SAppVersion{} }
func (m *SAppVersion) String() string { return proto.CompactTextString(m) }
func (*SAppVersion) ProtoMessage()    {}
func (*SAppVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_service_cc298f28d18949f2, []int{7}
}
func (m *SAppVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAppVersion.Unmarshal(m, b)
}
func (m *SAppVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAppVersion.Marshal(b, m, deterministic)
}
func (dst *SAppVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAppVersion.Merge(dst, src)
}
func (m *SAppVersion) XXX_Size() int {
	return xxx_messageInfo_SAppVersion.Size(m)
}
func (m *SAppVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SAppVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SAppVersion proto.InternalMessageInfo

func (m *SAppVersion) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAppVersion) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SAppVersion) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *SAppVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SAppVersion) GetVersionCode() int32 {
	if m != nil {
		return m.VersionCode
	}
	return 0
}

func (m *SAppVersion) GetForceUpdate() bool {
	if m != nil {
		return m.ForceUpdate
	}
	return false
}

func (m *SAppVersion) GetUpdateContent() string {
	if m != nil {
		return m.UpdateContent
	}
	return ""
}

func (m *SAppVersion) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*AppId)(nil), "AppId")
	proto.RegisterType((*AppVersionId)(nil), "AppVersionId")
	proto.RegisterType((*CheckVersionRequest)(nil), "CheckVersionRequest")
	proto.RegisterType((*CheckVersionResponse)(nil), "CheckVersionResponse")
	proto.RegisterType((*AppProdRequest)(nil), "AppProdRequest")
	proto.RegisterType((*AppVersionRequest)(nil), "AppVersionRequest")
	proto.RegisterType((*SAppProd)(nil), "SAppProd")
	proto.RegisterType((*SAppVersion)(nil), "SAppVersion")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppServiceClient interface {
	// 保存APP产品
	SaveProd(ctx context.Context, in *AppProdRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存版本
	SaveVersion(ctx context.Context, in *AppVersionRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取产品信息
	GetProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*SAppProd, error)
	// 获取版本
	GetVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*SAppVersion, error)
	// 删除产品
	DeleteProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*Result, error)
	// 删除版本
	DeleteVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*Result, error)
	// 检测版本更新
	CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error)
}

type appServiceClient struct {
	cc *grpc.ClientConn
}

func NewAppServiceClient(cc *grpc.ClientConn) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) SaveProd(ctx context.Context, in *AppProdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AppService/SaveProd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) SaveVersion(ctx context.Context, in *AppVersionRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AppService/SaveVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*SAppProd, error) {
	out := new(SAppProd)
	err := c.cc.Invoke(ctx, "/AppService/GetProd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*SAppVersion, error) {
	out := new(SAppVersion)
	err := c.cc.Invoke(ctx, "/AppService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AppService/DeleteProd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AppService/DeleteVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error) {
	out := new(CheckVersionResponse)
	err := c.cc.Invoke(ctx, "/AppService/CheckVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
type AppServiceServer interface {
	// 保存APP产品
	SaveProd(context.Context, *AppProdRequest) (*Result, error)
	// 保存版本
	SaveVersion(context.Context, *AppVersionRequest) (*Result, error)
	// 获取产品信息
	GetProd(context.Context, *AppId) (*SAppProd, error)
	// 获取版本
	GetVersion(context.Context, *AppVersionId) (*SAppVersion, error)
	// 删除产品
	DeleteProd(context.Context, *AppId) (*Result, error)
	// 删除版本
	DeleteVersion(context.Context, *AppVersionId) (*Result, error)
	// 检测版本更新
	CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error)
}

func RegisterAppServiceServer(s *grpc.Server, srv AppServiceServer) {
	s.RegisterService(&_AppService_serviceDesc, srv)
}

func _AppService_SaveProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).SaveProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/SaveProd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).SaveProd(ctx, req.(*AppProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_SaveVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).SaveVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/SaveVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).SaveVersion(ctx, req.(*AppVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/GetProd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetProd(ctx, req.(*AppId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppVersionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetVersion(ctx, req.(*AppVersionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/DeleteProd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteProd(ctx, req.(*AppId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppVersionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/DeleteVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteVersion(ctx, req.(*AppVersionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CheckVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CheckVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AppService/CheckVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CheckVersion(ctx, req.(*CheckVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveProd",
			Handler:    _AppService_SaveProd_Handler,
		},
		{
			MethodName: "SaveVersion",
			Handler:    _AppService_SaveVersion_Handler,
		},
		{
			MethodName: "GetProd",
			Handler:    _AppService_GetProd_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _AppService_GetVersion_Handler,
		},
		{
			MethodName: "DeleteProd",
			Handler:    _AppService_DeleteProd_Handler,
		},
		{
			MethodName: "DeleteVersion",
			Handler:    _AppService_DeleteVersion_Handler,
		},
		{
			MethodName: "CheckVersion",
			Handler:    _AppService_CheckVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_service.proto",
}

func init() { proto.RegisterFile("app_service.proto", fileDescriptor_app_service_cc298f28d18949f2) }

var fileDescriptor_app_service_cc298f28d18949f2 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x4e, 0xdb, 0x30,
	0x14, 0x6e, 0x5a, 0xfa, 0x93, 0xd3, 0x96, 0x09, 0x03, 0x52, 0x54, 0x01, 0xad, 0x22, 0x84, 0x98,
	0x36, 0x79, 0x12, 0xd3, 0xae, 0xa6, 0x5d, 0x74, 0x45, 0x30, 0x24, 0x86, 0x50, 0x0a, 0xbb, 0xd8,
	0x0d, 0x4b, 0x9b, 0xb3, 0x36, 0x22, 0xc4, 0x5e, 0xe2, 0x20, 0xf1, 0x30, 0x7b, 0x91, 0x5d, 0xec,
	0x49, 0xf6, 0x24, 0x93, 0x26, 0x4d, 0x76, 0x9c, 0x36, 0x29, 0x45, 0xe3, 0x02, 0xed, 0x2a, 0x3e,
	0x9f, 0x3f, 0x1f, 0xe7, 0x7c, 0xe7, 0xb3, 0x0d, 0x6b, 0x2e, 0xe7, 0x57, 0x31, 0x46, 0xb7, 0xfe,
	0x18, 0x29, 0x8f, 0x98, 0x60, 0x9d, 0xd6, 0x24, 0x60, 0x23, 0x37, 0x48, 0x23, 0x7b, 0x1b, 0xaa,
	0x7d, 0xce, 0x4f, 0x3c, 0xb2, 0x01, 0xd5, 0x5b, 0x37, 0x48, 0xd0, 0x32, 0x7a, 0xc6, 0x7e, 0xc5,
	0x49, 0x03, 0x7b, 0x17, 0x5a, 0x7d, 0xce, 0x3f, 0x61, 0x14, 0xfb, 0x2c, 0x7c, 0x90, 0xf5, 0x05,
	0xd6, 0x07, 0x53, 0x1c, 0x5f, 0x6b, 0x9e, 0x83, 0xdf, 0x12, 0x8c, 0x05, 0xd9, 0x84, 0x5a, 0x9f,
	0xf3, 0x2b, 0xdf, 0xcb, 0xd8, 0xe9, 0x4e, 0x16, 0xd4, 0x07, 0x53, 0x37, 0x0c, 0x31, 0xb0, 0xca,
	0x3d, 0x63, 0xdf, 0x74, 0xb2, 0x50, 0xce, 0xe8, 0x14, 0x56, 0x25, 0x9d, 0xd1, 0xa1, 0xfd, 0xc7,
	0x80, 0x8d, 0xe2, 0x16, 0x31, 0x67, 0x61, 0x8c, 0x64, 0x17, 0xda, 0xa7, 0xae, 0xc0, 0x58, 0x64,
	0x0b, 0x0d, 0xb5, 0xb0, 0x08, 0x92, 0x2d, 0x30, 0xfb, 0x9c, 0x9f, 0x5f, 0x4f, 0x2e, 0x9d, 0x53,
	0xbd, 0xe9, 0x1c, 0x20, 0x3d, 0x68, 0x66, 0x15, 0x86, 0x5f, 0x99, 0xde, 0x3a, 0x0f, 0x91, 0x0e,
	0x34, 0x4e, 0xe2, 0x34, 0xa5, 0xb5, 0xd2, 0x33, 0xf6, 0x1b, 0xce, 0x2c, 0x96, 0xab, 0x8f, 0x58,
	0x34, 0xc6, 0x4b, 0xee, 0xb9, 0x02, 0xad, 0xaa, 0x9a, 0xce, 0x43, 0xa4, 0x0b, 0xcd, 0x74, 0x74,
	0x25, 0xee, 0x38, 0x5a, 0x35, 0x95, 0x1f, 0x52, 0xe8, 0xe2, 0x8e, 0xa3, 0x4c, 0xe1, 0x60, 0x80,
	0x6e, 0x8c, 0x17, 0xfe, 0x0d, 0x5a, 0x75, 0xa5, 0x56, 0x1e, 0xb2, 0xbf, 0x97, 0x61, 0x55, 0xfe,
	0x70, 0xc4, 0xbc, 0x4c, 0xdd, 0x55, 0x28, 0xcf, 0x94, 0x2d, 0xfb, 0x9e, 0xfc, 0x47, 0x39, 0x7d,
	0xe6, 0xde, 0xa0, 0x2e, 0x71, 0x16, 0x4b, 0x61, 0xe5, 0xf8, 0x10, 0xe3, 0x4c, 0x58, 0x1d, 0x12,
	0x02, 0x2b, 0x27, 0x63, 0x16, 0xaa, 0xaa, 0x4c, 0x47, 0x8d, 0xc9, 0x0e, 0xc0, 0x79, 0x32, 0x0a,
	0xfc, 0x78, 0x7a, 0x19, 0x05, 0xaa, 0x20, 0xd3, 0xc9, 0x21, 0x52, 0xf3, 0xa1, 0x70, 0x47, 0x01,
	0x1e, 0xf9, 0x01, 0x4a, 0x4a, 0x5a, 0x51, 0x11, 0x24, 0x36, 0xb4, 0xfa, 0x01, 0x9f, 0xba, 0x19,
	0xa9, 0xa1, 0x48, 0x05, 0x8c, 0xec, 0xc1, 0xea, 0x99, 0x3f, 0x99, 0x8a, 0xe0, 0x2e, 0x63, 0x81,
	0x62, 0x2d, 0xa0, 0x8b, 0x0a, 0x36, 0x17, 0x15, 0xb4, 0x7f, 0x1a, 0xb0, 0x36, 0x37, 0xea, 0x43,
	0x12, 0x6d, 0x03, 0xc8, 0xba, 0x93, 0xb1, 0x90, 0xa6, 0x2c, 0x2b, 0xdc, 0xd4, 0x48, 0xd1, 0x98,
	0x52, 0xa5, 0xea, 0x52, 0x63, 0xae, 0x14, 0x8c, 0xf9, 0x88, 0xee, 0xef, 0x42, 0x3b, 0x1d, 0x0d,
	0x58, 0x28, 0x30, 0x14, 0x99, 0x5a, 0x05, 0xd0, 0xfe, 0x55, 0x86, 0xc6, 0x50, 0x77, 0xf8, 0x89,
	0x5a, 0xbb, 0x05, 0xa6, 0x3e, 0x05, 0xbe, 0xa7, 0x7e, 0xbb, 0xe2, 0xcc, 0x01, 0xb9, 0xee, 0xa3,
	0xf7, 0xe6, 0x83, 0x1b, 0x4f, 0x75, 0x87, 0xb3, 0x70, 0xa1, 0xfd, 0xb5, 0x7f, 0xb7, 0xbf, 0xfe,
	0x74, 0xed, 0x37, 0x1f, 0xd3, 0x7e, 0xb8, 0x77, 0x80, 0xba, 0xd0, 0x4c, 0x34, 0x41, 0x1e, 0xa0,
	0xa6, 0x2a, 0x16, 0x52, 0x48, 0x9d, 0x9f, 0xdf, 0x06, 0x34, 0x87, 0x73, 0x83, 0xfc, 0x27, 0x67,
	0xe8, 0xe1, 0x80, 0x79, 0xa9, 0x33, 0xaa, 0x4e, 0x1e, 0x5a, 0xf4, 0x4e, 0xed, 0x11, 0xde, 0xa9,
	0x2f, 0xf1, 0x8e, 0x6c, 0xd8, 0x20, 0x42, 0x5d, 0xaa, 0x12, 0xba, 0xe2, 0xe4, 0x90, 0x83, 0x1f,
	0x65, 0x80, 0x3e, 0xe7, 0xc3, 0xf4, 0x19, 0x20, 0x7b, 0xd0, 0x18, 0xba, 0xb7, 0xa8, 0x9c, 0xf6,
	0x8c, 0x16, 0x6f, 0x95, 0x4e, 0x9d, 0x3a, 0x18, 0x27, 0x81, 0xb0, 0x4b, 0xe4, 0x25, 0x34, 0x25,
	0x2f, 0xab, 0x87, 0xd0, 0x7b, 0x07, 0x2c, 0xcf, 0xde, 0x81, 0xfa, 0x31, 0x0a, 0x95, 0xb4, 0x46,
	0xd5, 0x45, 0xdf, 0x31, 0x69, 0xe6, 0x68, 0xbb, 0x44, 0x5e, 0x00, 0x1c, 0xe3, 0xec, 0x42, 0x6e,
	0xd3, 0xfc, 0xb3, 0xd2, 0x69, 0xd1, 0x5c, 0x73, 0xec, 0x12, 0xe9, 0x02, 0x1c, 0x62, 0x80, 0x02,
	0x0b, 0xf9, 0x72, 0xbb, 0x3d, 0x87, 0x76, 0x4a, 0x78, 0x20, 0x61, 0x8e, 0xfa, 0x0e, 0x5a, 0xf9,
	0x97, 0x83, 0x6c, 0xd0, 0x25, 0x6f, 0x55, 0x67, 0x93, 0x2e, 0x7b, 0x5e, 0xec, 0xd2, 0xfb, 0x2e,
	0xac, 0x8f, 0xd9, 0x0d, 0x9d, 0xf8, 0x62, 0x9a, 0x8c, 0xe8, 0x84, 0x1d, 0x30, 0x1a, 0xf1, 0xf1,
	0xe7, 0x06, 0x7d, 0xf5, 0x56, 0xbd, 0xa0, 0xa3, 0x9a, 0xfa, 0xbc, 0xfe, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x86, 0xa7, 0x14, 0x41, 0x6b, 0x07, 0x00, 0x00,
}
