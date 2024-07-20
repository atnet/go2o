// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: workorder_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WorkorderService_SubmitWorkorder_FullMethodName = "/WorkorderService/SubmitWorkorder"
	WorkorderService_GetWorkorder_FullMethodName    = "/WorkorderService/GetWorkorder"
	WorkorderService_DeleteWorkorder_FullMethodName = "/WorkorderService/DeleteWorkorder"
	WorkorderService_AllocateAgentId_FullMethodName = "/WorkorderService/AllocateAgentId"
	WorkorderService_Finish_FullMethodName          = "/WorkorderService/Finish"
	WorkorderService_Close_FullMethodName           = "/WorkorderService/Close"
	WorkorderService_Apprise_FullMethodName         = "/WorkorderService/Apprise"
	WorkorderService_SubmitComment_FullMethodName   = "/WorkorderService/SubmitComment"
)

// WorkorderServiceClient is the client API for WorkorderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkorderServiceClient interface {
	// 保存工单
	SubmitWorkorder(ctx context.Context, in *SubmitWorkorderRequest, opts ...grpc.CallOption) (*SubmitWorkorderResponse, error)
	// 获取工单
	GetWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*SWorkorder, error)
	// 删除工单
	DeleteWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*ResultV2, error)
	// 分配客服
	AllocateAgentId(ctx context.Context, in *AllocateWorkorderAgentRequest, opts ...grpc.CallOption) (*ResultV2, error)
	// 完结
	Finish(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*ResultV2, error)
	// 用户关闭工单
	Close(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*ResultV2, error)
	// 评价
	Apprise(ctx context.Context, in *WorkorderAppriseRequest, opts ...grpc.CallOption) (*ResultV2, error)
	// 提交回复
	SubmitComment(ctx context.Context, in *SubmitWorkorderCommentRequest, opts ...grpc.CallOption) (*ResultV2, error)
}

type workorderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkorderServiceClient(cc grpc.ClientConnInterface) WorkorderServiceClient {
	return &workorderServiceClient{cc}
}

func (c *workorderServiceClient) SubmitWorkorder(ctx context.Context, in *SubmitWorkorderRequest, opts ...grpc.CallOption) (*SubmitWorkorderResponse, error) {
	out := new(SubmitWorkorderResponse)
	err := c.cc.Invoke(ctx, WorkorderService_SubmitWorkorder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) GetWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*SWorkorder, error) {
	out := new(SWorkorder)
	err := c.cc.Invoke(ctx, WorkorderService_GetWorkorder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) DeleteWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*ResultV2, error) {
	out := new(ResultV2)
	err := c.cc.Invoke(ctx, WorkorderService_DeleteWorkorder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) AllocateAgentId(ctx context.Context, in *AllocateWorkorderAgentRequest, opts ...grpc.CallOption) (*ResultV2, error) {
	out := new(ResultV2)
	err := c.cc.Invoke(ctx, WorkorderService_AllocateAgentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) Finish(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*ResultV2, error) {
	out := new(ResultV2)
	err := c.cc.Invoke(ctx, WorkorderService_Finish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) Close(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*ResultV2, error) {
	out := new(ResultV2)
	err := c.cc.Invoke(ctx, WorkorderService_Close_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) Apprise(ctx context.Context, in *WorkorderAppriseRequest, opts ...grpc.CallOption) (*ResultV2, error) {
	out := new(ResultV2)
	err := c.cc.Invoke(ctx, WorkorderService_Apprise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) SubmitComment(ctx context.Context, in *SubmitWorkorderCommentRequest, opts ...grpc.CallOption) (*ResultV2, error) {
	out := new(ResultV2)
	err := c.cc.Invoke(ctx, WorkorderService_SubmitComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkorderServiceServer is the server API for WorkorderService service.
// All implementations must embed UnimplementedWorkorderServiceServer
// for forward compatibility
type WorkorderServiceServer interface {
	// 保存工单
	SubmitWorkorder(context.Context, *SubmitWorkorderRequest) (*SubmitWorkorderResponse, error)
	// 获取工单
	GetWorkorder(context.Context, *WorkorderId) (*SWorkorder, error)
	// 删除工单
	DeleteWorkorder(context.Context, *WorkorderId) (*ResultV2, error)
	// 分配客服
	AllocateAgentId(context.Context, *AllocateWorkorderAgentRequest) (*ResultV2, error)
	// 完结
	Finish(context.Context, *WorkorderId) (*ResultV2, error)
	// 用户关闭工单
	Close(context.Context, *WorkorderId) (*ResultV2, error)
	// 评价
	Apprise(context.Context, *WorkorderAppriseRequest) (*ResultV2, error)
	// 提交回复
	SubmitComment(context.Context, *SubmitWorkorderCommentRequest) (*ResultV2, error)
	mustEmbedUnimplementedWorkorderServiceServer()
}

// UnimplementedWorkorderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkorderServiceServer struct {
}

func (UnimplementedWorkorderServiceServer) SubmitWorkorder(context.Context, *SubmitWorkorderRequest) (*SubmitWorkorderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) GetWorkorder(context.Context, *WorkorderId) (*SWorkorder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) DeleteWorkorder(context.Context, *WorkorderId) (*ResultV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) AllocateAgentId(context.Context, *AllocateWorkorderAgentRequest) (*ResultV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateAgentId not implemented")
}
func (UnimplementedWorkorderServiceServer) Finish(context.Context, *WorkorderId) (*ResultV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}
func (UnimplementedWorkorderServiceServer) Close(context.Context, *WorkorderId) (*ResultV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedWorkorderServiceServer) Apprise(context.Context, *WorkorderAppriseRequest) (*ResultV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apprise not implemented")
}
func (UnimplementedWorkorderServiceServer) SubmitComment(context.Context, *SubmitWorkorderCommentRequest) (*ResultV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitComment not implemented")
}
func (UnimplementedWorkorderServiceServer) mustEmbedUnimplementedWorkorderServiceServer() {}

// UnsafeWorkorderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkorderServiceServer will
// result in compilation errors.
type UnsafeWorkorderServiceServer interface {
	mustEmbedUnimplementedWorkorderServiceServer()
}

func RegisterWorkorderServiceServer(s grpc.ServiceRegistrar, srv WorkorderServiceServer) {
	s.RegisterService(&WorkorderService_ServiceDesc, srv)
}

func _WorkorderService_SubmitWorkorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitWorkorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).SubmitWorkorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_SubmitWorkorder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).SubmitWorkorder(ctx, req.(*SubmitWorkorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_GetWorkorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).GetWorkorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_GetWorkorder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).GetWorkorder(ctx, req.(*WorkorderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_DeleteWorkorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).DeleteWorkorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_DeleteWorkorder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).DeleteWorkorder(ctx, req.(*WorkorderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_AllocateAgentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateWorkorderAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).AllocateAgentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_AllocateAgentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).AllocateAgentId(ctx, req.(*AllocateWorkorderAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_Finish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).Finish(ctx, req.(*WorkorderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_Close_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).Close(ctx, req.(*WorkorderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_Apprise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderAppriseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).Apprise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_Apprise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).Apprise(ctx, req.(*WorkorderAppriseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_SubmitComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitWorkorderCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).SubmitComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkorderService_SubmitComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).SubmitComment(ctx, req.(*SubmitWorkorderCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkorderService_ServiceDesc is the grpc.ServiceDesc for WorkorderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkorderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WorkorderService",
	HandlerType: (*WorkorderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitWorkorder",
			Handler:    _WorkorderService_SubmitWorkorder_Handler,
		},
		{
			MethodName: "GetWorkorder",
			Handler:    _WorkorderService_GetWorkorder_Handler,
		},
		{
			MethodName: "DeleteWorkorder",
			Handler:    _WorkorderService_DeleteWorkorder_Handler,
		},
		{
			MethodName: "AllocateAgentId",
			Handler:    _WorkorderService_AllocateAgentId_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _WorkorderService_Finish_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _WorkorderService_Close_Handler,
		},
		{
			MethodName: "Apprise",
			Handler:    _WorkorderService_Apprise_Handler,
		},
		{
			MethodName: "SubmitComment",
			Handler:    _WorkorderService_SubmitComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workorder_service.proto",
}
