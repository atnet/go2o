// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.1
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

// WorkorderServiceClient is the client API for WorkorderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkorderServiceClient interface {
	// 保存工单
	SubmitWorkorder(ctx context.Context, in *SubmitWorkorderRequest, opts ...grpc.CallOption) (*SubmitWorkorderResponse, error)
	// 获取工单
	GetWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*SWorkorder, error)
	// 删除工单
	DeleteWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*TxResult, error)
	// 分配客服
	AllocateAgentId(ctx context.Context, in *AllocateWorkorderAgentRequest, opts ...grpc.CallOption) (*TxResult, error)
	// 完结
	FinishWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*TxResult, error)
	// 用户关闭工单
	CloseWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*TxResult, error)
	// 评价
	AppriseWorkorder(ctx context.Context, in *WorkorderAppriseRequest, opts ...grpc.CallOption) (*TxResult, error)
	// 提交回复
	SubmitComment(ctx context.Context, in *SubmitWorkorderCommentRequest, opts ...grpc.CallOption) (*TxResult, error)
}

type workorderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkorderServiceClient(cc grpc.ClientConnInterface) WorkorderServiceClient {
	return &workorderServiceClient{cc}
}

func (c *workorderServiceClient) SubmitWorkorder(ctx context.Context, in *SubmitWorkorderRequest, opts ...grpc.CallOption) (*SubmitWorkorderResponse, error) {
	out := new(SubmitWorkorderResponse)
	err := c.cc.Invoke(ctx, "/WorkorderService/SubmitWorkorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) GetWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*SWorkorder, error) {
	out := new(SWorkorder)
	err := c.cc.Invoke(ctx, "/WorkorderService/GetWorkorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) DeleteWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/WorkorderService/DeleteWorkorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) AllocateAgentId(ctx context.Context, in *AllocateWorkorderAgentRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/WorkorderService/AllocateAgentId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) FinishWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/WorkorderService/FinishWorkorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) CloseWorkorder(ctx context.Context, in *WorkorderId, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/WorkorderService/CloseWorkorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) AppriseWorkorder(ctx context.Context, in *WorkorderAppriseRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/WorkorderService/AppriseWorkorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workorderServiceClient) SubmitComment(ctx context.Context, in *SubmitWorkorderCommentRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/WorkorderService/SubmitComment", in, out, opts...)
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
	DeleteWorkorder(context.Context, *WorkorderId) (*TxResult, error)
	// 分配客服
	AllocateAgentId(context.Context, *AllocateWorkorderAgentRequest) (*TxResult, error)
	// 完结
	FinishWorkorder(context.Context, *WorkorderId) (*TxResult, error)
	// 用户关闭工单
	CloseWorkorder(context.Context, *WorkorderId) (*TxResult, error)
	// 评价
	AppriseWorkorder(context.Context, *WorkorderAppriseRequest) (*TxResult, error)
	// 提交回复
	SubmitComment(context.Context, *SubmitWorkorderCommentRequest) (*TxResult, error)
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
func (UnimplementedWorkorderServiceServer) DeleteWorkorder(context.Context, *WorkorderId) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) AllocateAgentId(context.Context, *AllocateWorkorderAgentRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateAgentId not implemented")
}
func (UnimplementedWorkorderServiceServer) FinishWorkorder(context.Context, *WorkorderId) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) CloseWorkorder(context.Context, *WorkorderId) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) AppriseWorkorder(context.Context, *WorkorderAppriseRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppriseWorkorder not implemented")
}
func (UnimplementedWorkorderServiceServer) SubmitComment(context.Context, *SubmitWorkorderCommentRequest) (*TxResult, error) {
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
		FullMethod: "/WorkorderService/SubmitWorkorder",
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
		FullMethod: "/WorkorderService/GetWorkorder",
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
		FullMethod: "/WorkorderService/DeleteWorkorder",
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
		FullMethod: "/WorkorderService/AllocateAgentId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).AllocateAgentId(ctx, req.(*AllocateWorkorderAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_FinishWorkorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).FinishWorkorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkorderService/FinishWorkorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).FinishWorkorder(ctx, req.(*WorkorderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_CloseWorkorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).CloseWorkorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkorderService/CloseWorkorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).CloseWorkorder(ctx, req.(*WorkorderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkorderService_AppriseWorkorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkorderAppriseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkorderServiceServer).AppriseWorkorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkorderService/AppriseWorkorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkorderServiceServer).AppriseWorkorder(ctx, req.(*WorkorderAppriseRequest))
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
		FullMethod: "/WorkorderService/SubmitComment",
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
			MethodName: "FinishWorkorder",
			Handler:    _WorkorderService_FinishWorkorder_Handler,
		},
		{
			MethodName: "CloseWorkorder",
			Handler:    _WorkorderService_CloseWorkorder_Handler,
		},
		{
			MethodName: "AppriseWorkorder",
			Handler:    _WorkorderService_AppriseWorkorder_Handler,
		},
		{
			MethodName: "SubmitComment",
			Handler:    _WorkorderService_SubmitComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workorder_service.proto",
}
