// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.28.1
// source: content_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 栏目
type SArticleCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 父类编号,如为一级栏目则为0
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid"`
	// 浏览权限
	Flag int32 `protobuf:"varint,3,opt,name=flag,proto3" json:"flag"`
	// 名称(唯一)
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// 别名
	Alias string `protobuf:"bytes,5,opt,name=alias,proto3" json:"alias"`
	// 排序编号
	SortNo int32 `protobuf:"varint,6,opt,name=sortNo,proto3" json:"sortNo"`
	// 定位路径（打开栏目页定位到的路径）
	Location string `protobuf:"bytes,7,opt,name=location,proto3" json:"location"`
	// 页面标题
	Title string `protobuf:"bytes,8,opt,name=title,proto3" json:"title"`
	// 关键字
	Keywords string `protobuf:"bytes,9,opt,name=keywords,proto3" json:"keywords"`
	// 描述
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description"`
}

func (x *SArticleCategory) Reset() {
	*x = SArticleCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SArticleCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SArticleCategory) ProtoMessage() {}

func (x *SArticleCategory) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SArticleCategory.ProtoReflect.Descriptor instead.
func (*SArticleCategory) Descriptor() ([]byte, []int) {
	return file_content_service_proto_rawDescGZIP(), []int{0}
}

func (x *SArticleCategory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SArticleCategory) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SArticleCategory) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *SArticleCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SArticleCategory) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SArticleCategory) GetSortNo() int32 {
	if x != nil {
		return x.SortNo
	}
	return 0
}

func (x *SArticleCategory) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SArticleCategory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SArticleCategory) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SArticleCategory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// * 文章
type SArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 编号
	Id int64 `protobuf:"zigzag64,1,opt,name=id,proto3" json:"id"`
	// * 栏目编号
	CategoryId int64 `protobuf:"zigzag64,2,opt,name=categoryId,proto3" json:"categoryId"`
	// * 标题
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	// * 标志
	Flag int32 `protobuf:"zigzag32,4,opt,name=flag,proto3" json:"flag"`
	// * 小标题
	ShortTitle string `protobuf:"bytes,5,opt,name=shortTitle,proto3" json:"shortTitle"`
	// * 文章附图
	Thumbnail string `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail"`
	// * 商户编号
	MchId int32 `protobuf:"zigzag32,7,opt,name=mchId,proto3" json:"mchId"`
	// * 发布人用户编号
	PublisherId int64 `protobuf:"zigzag64,8,opt,name=publisherId,proto3" json:"publisherId"`
	// * 重定向URL
	Location string `protobuf:"bytes,9,opt,name=location,proto3" json:"location"`
	// * 优先级,优先级越高，则置顶
	Priority int32 `protobuf:"zigzag32,10,opt,name=priority,proto3" json:"priority"`
	// * 浏览钥匙
	AccessKey string `protobuf:"bytes,11,opt,name=accessKey,proto3" json:"accessKey"`
	// * 文档内容
	Content string `protobuf:"bytes,12,opt,name=content,proto3" json:"content"`
	// * 标签（关键词）
	Tags string `protobuf:"bytes,13,opt,name=tags,proto3" json:"tags"`
	// * 显示次数
	ViewCount int32 `protobuf:"zigzag32,14,opt,name=viewCount,proto3" json:"viewCount"`
	// * 点赞次数
	LikeCount int32 `protobuf:"zigzag32,15,opt,name=likeCount,proto3" json:"likeCount"`
	// * 不喜欢次数
	DislikeCount int32 `protobuf:"zigzag32,16,opt,name=dislikeCount,proto3" json:"dislikeCount"`
	// * 排序序号
	SortNum int32 `protobuf:"zigzag32,17,opt,name=sortNum,proto3" json:"sortNum"`
	// * 创建时间
	CreateTime int64 `protobuf:"zigzag64,18,opt,name=createTime,proto3" json:"createTime"`
	// * 最后修改时间
	UpdateTime int64 `protobuf:"zigzag64,19,opt,name=updateTime,proto3" json:"updateTime"`
	// 扩展信息
	Ext map[string]string `protobuf:"bytes,20,rep,name=ext,proto3" json:"ext" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SArticle) Reset() {
	*x = SArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SArticle) ProtoMessage() {}

func (x *SArticle) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SArticle.ProtoReflect.Descriptor instead.
func (*SArticle) Descriptor() ([]byte, []int) {
	return file_content_service_proto_rawDescGZIP(), []int{1}
}

func (x *SArticle) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SArticle) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *SArticle) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SArticle) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *SArticle) GetShortTitle() string {
	if x != nil {
		return x.ShortTitle
	}
	return ""
}

func (x *SArticle) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *SArticle) GetMchId() int32 {
	if x != nil {
		return x.MchId
	}
	return 0
}

func (x *SArticle) GetPublisherId() int64 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

func (x *SArticle) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SArticle) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *SArticle) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *SArticle) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SArticle) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *SArticle) GetViewCount() int32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *SArticle) GetLikeCount() int32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *SArticle) GetDislikeCount() int32 {
	if x != nil {
		return x.DislikeCount
	}
	return 0
}

func (x *SArticle) GetSortNum() int32 {
	if x != nil {
		return x.SortNum
	}
	return 0
}

func (x *SArticle) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *SArticle) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SArticle) GetExt() map[string]string {
	if x != nil {
		return x.Ext
	}
	return nil
}

type SPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 商户编号
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId"`
	// 标题
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	// 字符标识
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code"`
	// 浏览权限
	Flag int32 `protobuf:"varint,5,opt,name=flag,proto3" json:"flag"`
	// 浏览钥匙
	AccessKey string `protobuf:"bytes,6,opt,name=accessKey,proto3" json:"accessKey"`
	// 关键词
	Keyword string `protobuf:"bytes,7,opt,name=keyword,proto3" json:"keyword"`
	// 描述
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description"`
	// 样式表地址
	CssPath string `protobuf:"bytes,9,opt,name=cssPath,proto3" json:"cssPath"`
	// 内容
	Content string `protobuf:"bytes,10,opt,name=content,proto3" json:"content"`
	// 修改时间
	UpdateTime int64 `protobuf:"varint,11,opt,name=updateTime,proto3" json:"updateTime"`
	// 是否启用
	Enabled bool `protobuf:"varint,12,opt,name=enabled,proto3" json:"enabled"`
}

func (x *SPage) Reset() {
	*x = SPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPage) ProtoMessage() {}

func (x *SPage) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPage.ProtoReflect.Descriptor instead.
func (*SPage) Descriptor() ([]byte, []int) {
	return file_content_service_proto_rawDescGZIP(), []int{2}
}

func (x *SPage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SPage) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SPage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SPage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SPage) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *SPage) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *SPage) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SPage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SPage) GetCssPath() string {
	if x != nil {
		return x.CssPath
	}
	return ""
}

func (x *SPage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SPage) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SPage) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type ArticleCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*SArticleCategory `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
}

func (x *ArticleCategoriesResponse) Reset() {
	*x = ArticleCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCategoriesResponse) ProtoMessage() {}

func (x *ArticleCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCategoriesResponse.ProtoReflect.Descriptor instead.
func (*ArticleCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_content_service_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleCategoriesResponse) GetValue() []*SArticleCategory {
	if x != nil {
		return x.Value
	}
	return nil
}

// 文章浏览数更新请求
type ArticleViewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文章编号
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 次数
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
}

func (x *ArticleViewsRequest) Reset() {
	*x = ArticleViewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleViewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleViewsRequest) ProtoMessage() {}

func (x *ArticleViewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleViewsRequest.ProtoReflect.Descriptor instead.
func (*ArticleViewsRequest) Descriptor() ([]byte, []int) {
	return file_content_service_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleViewsRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleViewsRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 文章点赞/不喜欢请求
type ArticleLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文章编号
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 会员编号
	MemberId uint32 `protobuf:"varint,2,opt,name=memberId,proto3" json:"memberId"`
	// 是否不喜欢
	IsDislike bool `protobuf:"varint,3,opt,name=isDislike,proto3" json:"isDislike"`
}

func (x *ArticleLikeRequest) Reset() {
	*x = ArticleLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleLikeRequest) ProtoMessage() {}

func (x *ArticleLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleLikeRequest.ProtoReflect.Descriptor instead.
func (*ArticleLikeRequest) Descriptor() ([]byte, []int) {
	return file_content_service_proto_rawDescGZIP(), []int{5}
}

func (x *ArticleLikeRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleLikeRequest) GetMemberId() uint32 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *ArticleLikeRequest) GetIsDislike() bool {
	if x != nil {
		return x.IsDislike
	}
	return false
}

var File_content_service_proto protoreflect.FileDescriptor

var file_content_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x10, 0x53, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x72, 0x74, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74,
	0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xf6, 0x04, 0x0a, 0x08, 0x53, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x63, 0x68, 0x49, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x6d, 0x63, 0x68, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x03, 0x65, 0x78, 0x74, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03,
	0x65, 0x78, 0x74, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x02, 0x0a, 0x05,
	0x53, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x73, 0x73, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x73, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x19, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x53, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a, 0x13, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x69,
	0x73, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44,
	0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x32, 0x96, 0x04, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x09, 0x2e, 0x49, 0x64, 0x4f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x06, 0x2e, 0x53, 0x50, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x08, 0x73, 0x61, 0x76,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x06, 0x2e, 0x53, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x09, 0x2e,
	0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x09,
	0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3a, 0x0a, 0x14, 0x67, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x09, 0x2e, 0x49, 0x64,
	0x4f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x11, 0x2e, 0x53, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x13, 0x73, 0x61, 0x76,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x11, 0x2e, 0x53, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x1a, 0x09, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a,
	0x0a, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a,
	0x09, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x09, 0x2e, 0x49, 0x64, 0x4f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x53, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x37,
	0x0a, 0x14, 0x61, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x54, 0x78,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a,
	0x09, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0b, 0x73, 0x61,
	0x76, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x09, 0x2e, 0x53, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x1a, 0x09, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f,
	0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_service_proto_rawDescOnce sync.Once
	file_content_service_proto_rawDescData = file_content_service_proto_rawDesc
)

func file_content_service_proto_rawDescGZIP() []byte {
	file_content_service_proto_rawDescOnce.Do(func() {
		file_content_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_service_proto_rawDescData)
	})
	return file_content_service_proto_rawDescData
}

var file_content_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_content_service_proto_goTypes = []interface{}{
	(*SArticleCategory)(nil),          // 0: SArticleCategory
	(*SArticle)(nil),                  // 1: SArticle
	(*SPage)(nil),                     // 2: SPage
	(*ArticleCategoriesResponse)(nil), // 3: ArticleCategoriesResponse
	(*ArticleViewsRequest)(nil),       // 4: ArticleViewsRequest
	(*ArticleLikeRequest)(nil),        // 5: ArticleLikeRequest
	nil,                               // 6: SArticle.ExtEntry
	(*IdOrName)(nil),                  // 7: IdOrName
	(*Int64)(nil),                     // 8: Int64
	(*Empty)(nil),                     // 9: Empty
	(*TxResult)(nil),                  // 10: TxResult
}
var file_content_service_proto_depIdxs = []int32{
	6,  // 0: SArticle.ext:type_name -> SArticle.ExtEntry
	0,  // 1: ArticleCategoriesResponse.value:type_name -> SArticleCategory
	7,  // 2: ContentService.getPage:input_type -> IdOrName
	2,  // 3: ContentService.savePage:input_type -> SPage
	8,  // 4: ContentService.deletePage:input_type -> Int64
	9,  // 5: ContentService.getArticleCategories:input_type -> Empty
	7,  // 6: ContentService.getArticleCategory:input_type -> IdOrName
	0,  // 7: ContentService.saveArticleCategory:input_type -> SArticleCategory
	8,  // 8: ContentService.deleteArticleCategory:input_type -> Int64
	7,  // 9: ContentService.getArticle:input_type -> IdOrName
	4,  // 10: ContentService.addArticleViewsCount:input_type -> ArticleViewsRequest
	5,  // 11: ContentService.likeArticle:input_type -> ArticleLikeRequest
	8,  // 12: ContentService.deleteArticle:input_type -> Int64
	1,  // 13: ContentService.saveArticle:input_type -> SArticle
	2,  // 14: ContentService.getPage:output_type -> SPage
	10, // 15: ContentService.savePage:output_type -> TxResult
	10, // 16: ContentService.deletePage:output_type -> TxResult
	3,  // 17: ContentService.getArticleCategories:output_type -> ArticleCategoriesResponse
	0,  // 18: ContentService.getArticleCategory:output_type -> SArticleCategory
	10, // 19: ContentService.saveArticleCategory:output_type -> TxResult
	10, // 20: ContentService.deleteArticleCategory:output_type -> TxResult
	1,  // 21: ContentService.getArticle:output_type -> SArticle
	10, // 22: ContentService.addArticleViewsCount:output_type -> TxResult
	10, // 23: ContentService.likeArticle:output_type -> TxResult
	10, // 24: ContentService.deleteArticle:output_type -> TxResult
	10, // 25: ContentService.saveArticle:output_type -> TxResult
	14, // [14:26] is the sub-list for method output_type
	2,  // [2:14] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_content_service_proto_init() }
func file_content_service_proto_init() {
	if File_content_service_proto != nil {
		return
	}
	file_global_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_content_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SArticleCategory); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SArticle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCategoriesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleViewsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleLikeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_service_proto_goTypes,
		DependencyIndexes: file_content_service_proto_depIdxs,
		MessageInfos:      file_content_service_proto_msgTypes,
	}.Build()
	File_content_service_proto = out.File
	file_content_service_proto_rawDesc = nil
	file_content_service_proto_goTypes = nil
	file_content_service_proto_depIdxs = nil
}
