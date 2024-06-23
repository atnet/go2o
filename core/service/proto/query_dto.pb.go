// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: message/query_dto.proto

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

// 会员账户分页明细
type MemberAccountPagingLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总条数,如不查询总数则返回0
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 数据
	Data []*SMemberAccountLog `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *MemberAccountPagingLogResponse) Reset() {
	*x = MemberAccountPagingLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_query_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberAccountPagingLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberAccountPagingLogResponse) ProtoMessage() {}

func (x *MemberAccountPagingLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_query_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberAccountPagingLogResponse.ProtoReflect.Descriptor instead.
func (*MemberAccountPagingLogResponse) Descriptor() ([]byte, []int) {
	return file_message_query_dto_proto_rawDescGZIP(), []int{0}
}

func (x *MemberAccountPagingLogResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemberAccountPagingLogResponse) GetData() []*SMemberAccountLog {
	if x != nil {
		return x.Data
	}
	return nil
}

// * 钱包日志
type SMemberAccountLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// * 业务类型
	Kind int32 `protobuf:"varint,2,opt,name=kind,proto3" json:"kind"`
	// * 流水号
	TradeNo string `protobuf:"bytes,3,opt,name=tradeNo,proto3" json:"tradeNo"`
	// * 标题
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject"`
	// * 变动金额
	Value int64 `protobuf:"varint,5,opt,name=value,proto3" json:"value"`
	// * 余额
	Balance int64 `protobuf:"varint,6,opt,name=balance,proto3" json:"balance"`
	// * 交易手续费
	ProcedureFee int64 `protobuf:"varint,7,opt,name=procedureFee,proto3" json:"procedureFee"`
	// * 外部订单号
	OuterNo string `protobuf:"bytes,8,opt,name=outerNo,proto3" json:"outerNo"`
	// * 备注
	Remark string `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark"`
	// * 审核状态
	ReviewStatus int32 `protobuf:"varint,10,opt,name=reviewStatus,proto3" json:"reviewStatus"`
	// * 创建时间
	CreateTime int64 `protobuf:"varint,11,opt,name=createTime,proto3" json:"createTime"`
}

func (x *SMemberAccountLog) Reset() {
	*x = SMemberAccountLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_query_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMemberAccountLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMemberAccountLog) ProtoMessage() {}

func (x *SMemberAccountLog) ProtoReflect() protoreflect.Message {
	mi := &file_message_query_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMemberAccountLog.ProtoReflect.Descriptor instead.
func (*SMemberAccountLog) Descriptor() ([]byte, []int) {
	return file_message_query_dto_proto_rawDescGZIP(), []int{1}
}

func (x *SMemberAccountLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SMemberAccountLog) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *SMemberAccountLog) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *SMemberAccountLog) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SMemberAccountLog) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SMemberAccountLog) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *SMemberAccountLog) GetProcedureFee() int64 {
	if x != nil {
		return x.ProcedureFee
	}
	return 0
}

func (x *SMemberAccountLog) GetOuterNo() string {
	if x != nil {
		return x.OuterNo
	}
	return ""
}

func (x *SMemberAccountLog) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SMemberAccountLog) GetReviewStatus() int32 {
	if x != nil {
		return x.ReviewStatus
	}
	return 0
}

func (x *SMemberAccountLog) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_message_query_dto_proto protoreflect.FileDescriptor

var file_message_query_dto_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x1e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x53, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb5, 0x02, 0x0a, 0x11, 0x53, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x46, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x46, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_query_dto_proto_rawDescOnce sync.Once
	file_message_query_dto_proto_rawDescData = file_message_query_dto_proto_rawDesc
)

func file_message_query_dto_proto_rawDescGZIP() []byte {
	file_message_query_dto_proto_rawDescOnce.Do(func() {
		file_message_query_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_query_dto_proto_rawDescData)
	})
	return file_message_query_dto_proto_rawDescData
}

var file_message_query_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_query_dto_proto_goTypes = []interface{}{
	(*MemberAccountPagingLogResponse)(nil), // 0: MemberAccountPagingLogResponse
	(*SMemberAccountLog)(nil),              // 1: SMemberAccountLog
}
var file_message_query_dto_proto_depIdxs = []int32{
	1, // 0: MemberAccountPagingLogResponse.data:type_name -> SMemberAccountLog
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_query_dto_proto_init() }
func file_message_query_dto_proto_init() {
	if File_message_query_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_query_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberAccountPagingLogResponse); i {
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
		file_message_query_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMemberAccountLog); i {
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
			RawDescriptor: file_message_query_dto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_query_dto_proto_goTypes,
		DependencyIndexes: file_message_query_dto_proto_depIdxs,
		MessageInfos:      file_message_query_dto_proto_msgTypes,
	}.Build()
	File_message_query_dto_proto = out.File
	file_message_query_dto_proto_rawDesc = nil
	file_message_query_dto_proto_goTypes = nil
	file_message_query_dto_proto_depIdxs = nil
}
