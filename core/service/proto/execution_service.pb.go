//*
// This file is auto generated by tto v0.4.5 !
// If you want to modify this code, please read the guide
// to modify code template.
//
// Get started: https://github.com/ixre/tto
//
// Copyright (C) 2009-2022 56X.NET, All rights reserved.
//
// name : job_exec_data_service.proto
// author : jarrysix
// date : 2022/03/06 03:16:21
// description :
// history :

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.28.1
// source: execution_service.proto

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

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 任务名称
	JobName string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName"`
	// * 任务不存在时是否创建
	Create bool `protobuf:"varint,2,opt,name=create,proto3" json:"create"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execution_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_execution_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetJobRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *GetJobRequest) GetCreate() bool {
	if x != nil {
		return x.Create
	}
	return false
}

// * UpdateCursorRequest
type UpdateCursorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 任务名称
	JobName string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName"`
	// * 记录编号
	CursorId int64 `protobuf:"varint,2,opt,name=cursorId,proto3" json:"cursorId"`
}

func (x *UpdateCursorRequest) Reset() {
	*x = UpdateCursorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCursorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCursorRequest) ProtoMessage() {}

func (x *UpdateCursorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execution_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCursorRequest.ProtoReflect.Descriptor instead.
func (*UpdateCursorRequest) Descriptor() ([]byte, []int) {
	return file_execution_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCursorRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *UpdateCursorRequest) GetCursorId() int64 {
	if x != nil {
		return x.CursorId
	}
	return 0
}

type AddFailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 任务名称
	JobName string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName"`
	// * 记录编号
	CursorId int64 `protobuf:"varint,2,opt,name=cursorId,proto3" json:"cursorId"`
}

func (x *AddFailRequest) Reset() {
	*x = AddFailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFailRequest) ProtoMessage() {}

func (x *AddFailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execution_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFailRequest.ProtoReflect.Descriptor instead.
func (*AddFailRequest) Descriptor() ([]byte, []int) {
	return file_execution_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddFailRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *AddFailRequest) GetCursorId() int64 {
	if x != nil {
		return x.CursorId
	}
	return 0
}

// * JobExecData
type SExecutionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// * 任务名称
	JobName string `protobuf:"bytes,2,opt,name=jobName,proto3" json:"jobName"`
	// * 上次执行位置索引
	LastExecuteCursorId int64 `protobuf:"varint,3,opt,name=lastExecuteCursorId,proto3" json:"lastExecuteCursorId"`
	// * 最后执行时间
	LastExecuteTime int64 `protobuf:"varint,4,opt,name=lastExecuteTime,proto3" json:"lastExecuteTime"`
}

func (x *SExecutionData) Reset() {
	*x = SExecutionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SExecutionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SExecutionData) ProtoMessage() {}

func (x *SExecutionData) ProtoReflect() protoreflect.Message {
	mi := &file_execution_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SExecutionData.ProtoReflect.Descriptor instead.
func (*SExecutionData) Descriptor() ([]byte, []int) {
	return file_execution_service_proto_rawDescGZIP(), []int{3}
}

func (x *SExecutionData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SExecutionData) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *SExecutionData) GetLastExecuteCursorId() int64 {
	if x != nil {
		return x.LastExecuteCursorId
	}
	return 0
}

func (x *SExecutionData) GetLastExecuteTime() int64 {
	if x != nil {
		return x.LastExecuteTime
	}
	return 0
}

// * 重新加入队列请求
type RejoinQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 桶名称
	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName"`
	// * 关联数据编号
	RelateId int64 `protobuf:"varint,2,opt,name=relateId,proto3" json:"relateId"`
	// * 数据
	RelateData string `protobuf:"bytes,3,opt,name=relateData,proto3" json:"relateData"`
}

func (x *RejoinQueueRequest) Reset() {
	*x = RejoinQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejoinQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejoinQueueRequest) ProtoMessage() {}

func (x *RejoinQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execution_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejoinQueueRequest.ProtoReflect.Descriptor instead.
func (*RejoinQueueRequest) Descriptor() ([]byte, []int) {
	return file_execution_service_proto_rawDescGZIP(), []int{4}
}

func (x *RejoinQueueRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *RejoinQueueRequest) GetRelateId() int64 {
	if x != nil {
		return x.RelateId
	}
	return 0
}

func (x *RejoinQueueRequest) GetRelateData() string {
	if x != nil {
		return x.RelateData
	}
	return ""
}

// * 保存重新加入队列响应
type RejoinQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg"`
	QueueId int64  `protobuf:"varint,3,opt,name=queueId,proto3" json:"queueId"`
}

func (x *RejoinQueueResponse) Reset() {
	*x = RejoinQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejoinQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejoinQueueResponse) ProtoMessage() {}

func (x *RejoinQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_execution_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejoinQueueResponse.ProtoReflect.Descriptor instead.
func (*RejoinQueueResponse) Descriptor() ([]byte, []int) {
	return file_execution_service_proto_rawDescGZIP(), []int{5}
}

func (x *RejoinQueueResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *RejoinQueueResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *RejoinQueueResponse) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

var File_execution_service_proto protoreflect.FileDescriptor

var file_execution_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x22,
	0x96, 0x01, 0x0a, 0x0e, 0x53, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13,
	0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x12, 0x52, 0x65, 0x6a, 0x6f,
	0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x13, 0x52, 0x65,
	0x6a, 0x6f, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x32, 0xda, 0x01,
	0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x2e, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x61,
	0x69, 0x6c, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0b, 0x52, 0x65, 0x6a, 0x6f, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x13, 0x2e,
	0x52, 0x65, 0x6a, 0x6f, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x6a, 0x6f, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70,
	0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_execution_service_proto_rawDescOnce sync.Once
	file_execution_service_proto_rawDescData = file_execution_service_proto_rawDesc
)

func file_execution_service_proto_rawDescGZIP() []byte {
	file_execution_service_proto_rawDescOnce.Do(func() {
		file_execution_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_execution_service_proto_rawDescData)
	})
	return file_execution_service_proto_rawDescData
}

var file_execution_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_execution_service_proto_goTypes = []interface{}{
	(*GetJobRequest)(nil),       // 0: GetJobRequest
	(*UpdateCursorRequest)(nil), // 1: UpdateCursorRequest
	(*AddFailRequest)(nil),      // 2: AddFailRequest
	(*SExecutionData)(nil),      // 3: SExecutionData
	(*RejoinQueueRequest)(nil),  // 4: RejoinQueueRequest
	(*RejoinQueueResponse)(nil), // 5: RejoinQueueResponse
	(*Result)(nil),              // 6: Result
}
var file_execution_service_proto_depIdxs = []int32{
	0, // 0: ExecutionService.GetJob:input_type -> GetJobRequest
	1, // 1: ExecutionService.UpdateExecuteCursor:input_type -> UpdateCursorRequest
	2, // 2: ExecutionService.AddFail:input_type -> AddFailRequest
	4, // 3: ExecutionService.RejoinQueue:input_type -> RejoinQueueRequest
	3, // 4: ExecutionService.GetJob:output_type -> SExecutionData
	6, // 5: ExecutionService.UpdateExecuteCursor:output_type -> Result
	6, // 6: ExecutionService.AddFail:output_type -> Result
	5, // 7: ExecutionService.RejoinQueue:output_type -> RejoinQueueResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_execution_service_proto_init() }
func file_execution_service_proto_init() {
	if File_execution_service_proto != nil {
		return
	}
	file_global_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_execution_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
		file_execution_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCursorRequest); i {
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
		file_execution_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFailRequest); i {
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
		file_execution_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SExecutionData); i {
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
		file_execution_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejoinQueueRequest); i {
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
		file_execution_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejoinQueueResponse); i {
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
			RawDescriptor: file_execution_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_execution_service_proto_goTypes,
		DependencyIndexes: file_execution_service_proto_depIdxs,
		MessageInfos:      file_execution_service_proto_msgTypes,
	}.Build()
	File_execution_service_proto = out.File
	file_execution_service_proto_rawDesc = nil
	file_execution_service_proto_goTypes = nil
	file_execution_service_proto_depIdxs = nil
}
