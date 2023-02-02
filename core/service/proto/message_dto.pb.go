// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: message/message_dto.proto

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

// * 消息方式
type EMessageChannel int32

const (
	// * 站内信
	EMessageChannel_SITE_MEMSSAGE EMessageChannel = 0 //* 邮件
	EMessageChannel_EMAIL_MESSAGE EMessageChannel = 1 //* 短信
	EMessageChannel_SMS_MESSAGE   EMessageChannel = 2
)

// Enum value maps for EMessageChannel.
var (
	EMessageChannel_name = map[int32]string{
		0: "SITE_MEMSSAGE",
		1: "EMAIL_MESSAGE",
		2: "SMS_MESSAGE",
	}
	EMessageChannel_value = map[string]int32{
		"SITE_MEMSSAGE": 0,
		"EMAIL_MESSAGE": 1,
		"SMS_MESSAGE":   2,
	}
)

func (x EMessageChannel) Enum() *EMessageChannel {
	p := new(EMessageChannel)
	*p = x
	return p
}

func (x EMessageChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EMessageChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_message_message_dto_proto_enumTypes[0].Descriptor()
}

func (EMessageChannel) Type() protoreflect.EnumType {
	return &file_message_message_dto_proto_enumTypes[0]
}

func (x EMessageChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EMessageChannel.Descriptor instead.
func (EMessageChannel) EnumDescriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{0}
}

// 站内信用户类型
type EMessageUserType int32

const (
	EMessageUserType_ALL      EMessageUserType = 0
	EMessageUserType_MEMBER   EMessageUserType = 1
	EMessageUserType_MERCHANT EMessageUserType = 2
)

// Enum value maps for EMessageUserType.
var (
	EMessageUserType_name = map[int32]string{
		0: "ALL",
		1: "MEMBER",
		2: "MERCHANT",
	}
	EMessageUserType_value = map[string]int32{
		"ALL":      0,
		"MEMBER":   1,
		"MERCHANT": 2,
	}
)

func (x EMessageUserType) Enum() *EMessageUserType {
	p := new(EMessageUserType)
	*p = x
	return p
}

func (x EMessageUserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EMessageUserType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_message_dto_proto_enumTypes[1].Descriptor()
}

func (EMessageUserType) Type() protoreflect.EnumType {
	return &file_message_message_dto_proto_enumTypes[1]
}

func (x EMessageUserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EMessageUserType.Descriptor instead.
func (EMessageUserType) EnumDescriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{1}
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 接收信息的账户
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account"`
	// 模板编号
	TemplateId string `protobuf:"bytes,3,opt,name=templateId,proto3" json:"templateId"`
	// 消息内容,如果isTemplateId为true,则message传入模板编号
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// 模板数据参数
	Data []string `protobuf:"bytes,4,rep,name=data,proto3" json:"data"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SendMessageRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMessageRequest) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

type NotifyItemListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*SNotifyItem `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
}

func (x *NotifyItemListResponse) Reset() {
	*x = NotifyItemListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyItemListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyItemListResponse) ProtoMessage() {}

func (x *NotifyItemListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyItemListResponse.ProtoReflect.Descriptor instead.
func (*NotifyItemListResponse) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyItemListResponse) GetValue() []*SNotifyItem {
	if x != nil {
		return x.Value
	}
	return nil
}

// * 通知项
type SNotifyItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 键
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	// * 发送方式
	NotifyBy int32 `protobuf:"zigzag32,2,opt,name=notifyBy,proto3" json:"notifyBy"`
	// * 不允许修改发送方式
	ReadonlyBy bool `protobuf:"varint,3,opt,name=readonlyBy,proto3" json:"readonlyBy"`
	// * 模板编号
	TplId int32 `protobuf:"zigzag32,4,opt,name=tplId,proto3" json:"tplId"`
	// * 内容
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`
	// * 模板包含的标签
	Tags map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SNotifyItem) Reset() {
	*x = SNotifyItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SNotifyItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SNotifyItem) ProtoMessage() {}

func (x *SNotifyItem) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SNotifyItem.ProtoReflect.Descriptor instead.
func (*SNotifyItem) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{2}
}

func (x *SNotifyItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SNotifyItem) GetNotifyBy() int32 {
	if x != nil {
		return x.NotifyBy
	}
	return 0
}

func (x *SNotifyItem) GetReadonlyBy() bool {
	if x != nil {
		return x.ReadonlyBy
	}
	return false
}

func (x *SNotifyItem) GetTplId() int32 {
	if x != nil {
		return x.TplId
	}
	return 0
}

func (x *SNotifyItem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SNotifyItem) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// 邮件模版
type SMailTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 商户编号
	MerchantId int64 `protobuf:"varint,2,opt,name=merchantId,proto3" json:"merchantId"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 主题
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject"`
	// 内容
	Body string `protobuf:"bytes,5,opt,name=body,proto3" json:"body"`
	// 是否启用
	Enabled bool `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled"`
}

func (x *SMailTemplate) Reset() {
	*x = SMailTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMailTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMailTemplate) ProtoMessage() {}

func (x *SMailTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMailTemplate.ProtoReflect.Descriptor instead.
func (*SMailTemplate) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{3}
}

func (x *SMailTemplate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SMailTemplate) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *SMailTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SMailTemplate) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SMailTemplate) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SMailTemplate) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type MailTemplateListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*SMailTemplate `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
}

func (x *MailTemplateListResponse) Reset() {
	*x = MailTemplateListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailTemplateListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailTemplateListResponse) ProtoMessage() {}

func (x *MailTemplateListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailTemplateListResponse.ProtoReflect.Descriptor instead.
func (*MailTemplateListResponse) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{4}
}

func (x *MailTemplateListResponse) GetValue() []*SMailTemplate {
	if x != nil {
		return x.Value
	}
	return nil
}

// 站内信
type SSiteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主题
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
	// 信息内容
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
}

func (x *SSiteMessage) Reset() {
	*x = SSiteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSiteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSiteMessage) ProtoMessage() {}

func (x *SSiteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSiteMessage.ProtoReflect.Descriptor instead.
func (*SSiteMessage) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{5}
}

func (x *SSiteMessage) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SSiteMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendSiteMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId     int64            `protobuf:"varint,1,opt,name=senderId,proto3" json:"senderId"`
	ReceiverType EMessageUserType `protobuf:"varint,2,opt,name=receiverType,proto3,enum=EMessageUserType" json:"receiverType"`
	ReceiverId   int64            `protobuf:"varint,3,opt,name=receiverId,proto3" json:"receiverId"`
	SendNow      bool             `protobuf:"varint,4,opt,name=sendNow,proto3" json:"sendNow"`
	Msg          *SSiteMessage    `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg"`
}

func (x *SendSiteMessageRequest) Reset() {
	*x = SendSiteMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_message_dto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSiteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSiteMessageRequest) ProtoMessage() {}

func (x *SendSiteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_message_dto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSiteMessageRequest.ProtoReflect.Descriptor instead.
func (*SendSiteMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_message_dto_proto_rawDescGZIP(), []int{6}
}

func (x *SendSiteMessageRequest) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *SendSiteMessageRequest) GetReceiverType() EMessageUserType {
	if x != nil {
		return x.ReceiverType
	}
	return EMessageUserType_ALL
}

func (x *SendSiteMessageRequest) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *SendSiteMessageRequest) GetSendNow() bool {
	if x != nil {
		return x.SendNow
	}
	return false
}

func (x *SendSiteMessageRequest) GetMsg() *SSiteMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_message_message_dto_proto protoreflect.FileDescriptor

var file_message_message_dto_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x12, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x16, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x0b, 0x53, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c,
	0x79, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x6f,
	0x6e, 0x6c, 0x79, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x70, 0x6c, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x74, 0x70, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x53,
	0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x4d, 0x61, 0x69, 0x6c,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x53, 0x53,
	0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc6,
	0x01, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x45, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x77, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x48, 0x0a, 0x0f, 0x45, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x49,
	0x54, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x4d, 0x53, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10,
	0x02, 0x2a, 0x35, 0x0a, 0x10, 0x45, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45,
	0x52, 0x43, 0x48, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_message_message_dto_proto_rawDescOnce sync.Once
	file_message_message_dto_proto_rawDescData = file_message_message_dto_proto_rawDesc
)

func file_message_message_dto_proto_rawDescGZIP() []byte {
	file_message_message_dto_proto_rawDescOnce.Do(func() {
		file_message_message_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_message_dto_proto_rawDescData)
	})
	return file_message_message_dto_proto_rawDescData
}

var file_message_message_dto_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_message_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_message_dto_proto_goTypes = []interface{}{
	(EMessageChannel)(0),             // 0: EMessageChannel
	(EMessageUserType)(0),            // 1: EMessageUserType
	(*SendMessageRequest)(nil),       // 2: SendMessageRequest
	(*NotifyItemListResponse)(nil),   // 3: NotifyItemListResponse
	(*SNotifyItem)(nil),              // 4: SNotifyItem
	(*SMailTemplate)(nil),            // 5: SMailTemplate
	(*MailTemplateListResponse)(nil), // 6: MailTemplateListResponse
	(*SSiteMessage)(nil),             // 7: SSiteMessage
	(*SendSiteMessageRequest)(nil),   // 8: SendSiteMessageRequest
	nil,                              // 9: SNotifyItem.TagsEntry
}
var file_message_message_dto_proto_depIdxs = []int32{
	4, // 0: NotifyItemListResponse.value:type_name -> SNotifyItem
	9, // 1: SNotifyItem.tags:type_name -> SNotifyItem.TagsEntry
	5, // 2: MailTemplateListResponse.value:type_name -> SMailTemplate
	1, // 3: SendSiteMessageRequest.receiverType:type_name -> EMessageUserType
	7, // 4: SendSiteMessageRequest.msg:type_name -> SSiteMessage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_message_dto_proto_init() }
func file_message_message_dto_proto_init() {
	if File_message_message_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_message_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_message_message_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyItemListResponse); i {
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
		file_message_message_dto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SNotifyItem); i {
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
		file_message_message_dto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMailTemplate); i {
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
		file_message_message_dto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailTemplateListResponse); i {
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
		file_message_message_dto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSiteMessage); i {
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
		file_message_message_dto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSiteMessageRequest); i {
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
			RawDescriptor: file_message_message_dto_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_message_dto_proto_goTypes,
		DependencyIndexes: file_message_message_dto_proto_depIdxs,
		EnumInfos:         file_message_message_dto_proto_enumTypes,
		MessageInfos:      file_message_message_dto_proto_msgTypes,
	}.Build()
	File_message_message_dto_proto = out.File
	file_message_message_dto_proto_rawDesc = nil
	file_message_message_dto_proto_goTypes = nil
	file_message_message_dto_proto_depIdxs = nil
}
