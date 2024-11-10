// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.1
// source: finance_service.proto

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

// FinanceServiceClient is the client API for FinanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceServiceClient interface {
	// 获取用户的账户信息
	GetRiseInfo(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*SRiseInfo, error)
	// 转入
	RiseTransferIn(ctx context.Context, in *TransferInRequest, opts ...grpc.CallOption) (*Result, error)
	// 转出
	RiseTransferOut(ctx context.Context, in *RiseTransferOutRequest, opts ...grpc.CallOption) (*Result, error)
	// 结算收益(按日期每天结息)
	RiseSettleByDay(ctx context.Context, in *RiseSettleRequest, opts ...grpc.CallOption) (*Result, error)
	// 提交转入/转出日志
	CommitTransfer(ctx context.Context, in *CommitTransferRequest, opts ...grpc.CallOption) (*Result, error)
	// 开通增利服务
	OpenRiseService(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Result, error)
}

type financeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceServiceClient(cc grpc.ClientConnInterface) FinanceServiceClient {
	return &financeServiceClient{cc}
}

func (c *financeServiceClient) GetRiseInfo(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*SRiseInfo, error) {
	out := new(SRiseInfo)
	err := c.cc.Invoke(ctx, "/FinanceService/GetRiseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) RiseTransferIn(ctx context.Context, in *TransferInRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseTransferIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) RiseTransferOut(ctx context.Context, in *RiseTransferOutRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseTransferOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) RiseSettleByDay(ctx context.Context, in *RiseSettleRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseSettleByDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) CommitTransfer(ctx context.Context, in *CommitTransferRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/CommitTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) OpenRiseService(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/OpenRiseService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServiceServer is the server API for FinanceService service.
// All implementations must embed UnimplementedFinanceServiceServer
// for forward compatibility
type FinanceServiceServer interface {
	// 获取用户的账户信息
	GetRiseInfo(context.Context, *PersonId) (*SRiseInfo, error)
	// 转入
	RiseTransferIn(context.Context, *TransferInRequest) (*Result, error)
	// 转出
	RiseTransferOut(context.Context, *RiseTransferOutRequest) (*Result, error)
	// 结算收益(按日期每天结息)
	RiseSettleByDay(context.Context, *RiseSettleRequest) (*Result, error)
	// 提交转入/转出日志
	CommitTransfer(context.Context, *CommitTransferRequest) (*Result, error)
	// 开通增利服务
	OpenRiseService(context.Context, *PersonId) (*Result, error)
	mustEmbedUnimplementedFinanceServiceServer()
}

// UnimplementedFinanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceServiceServer struct {
}

func (UnimplementedFinanceServiceServer) GetRiseInfo(context.Context, *PersonId) (*SRiseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRiseInfo not implemented")
}
func (UnimplementedFinanceServiceServer) RiseTransferIn(context.Context, *TransferInRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RiseTransferIn not implemented")
}
func (UnimplementedFinanceServiceServer) RiseTransferOut(context.Context, *RiseTransferOutRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RiseTransferOut not implemented")
}
func (UnimplementedFinanceServiceServer) RiseSettleByDay(context.Context, *RiseSettleRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RiseSettleByDay not implemented")
}
func (UnimplementedFinanceServiceServer) CommitTransfer(context.Context, *CommitTransferRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitTransfer not implemented")
}
func (UnimplementedFinanceServiceServer) OpenRiseService(context.Context, *PersonId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenRiseService not implemented")
}
func (UnimplementedFinanceServiceServer) mustEmbedUnimplementedFinanceServiceServer() {}

// UnsafeFinanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceServiceServer will
// result in compilation errors.
type UnsafeFinanceServiceServer interface {
	mustEmbedUnimplementedFinanceServiceServer()
}

func RegisterFinanceServiceServer(s grpc.ServiceRegistrar, srv FinanceServiceServer) {
	s.RegisterService(&FinanceService_ServiceDesc, srv)
}

func _FinanceService_GetRiseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).GetRiseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/GetRiseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).GetRiseInfo(ctx, req.(*PersonId))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_RiseTransferIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseTransferIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseTransferIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseTransferIn(ctx, req.(*TransferInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_RiseTransferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiseTransferOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseTransferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseTransferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseTransferOut(ctx, req.(*RiseTransferOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_RiseSettleByDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiseSettleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseSettleByDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseSettleByDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseSettleByDay(ctx, req.(*RiseSettleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_CommitTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).CommitTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/CommitTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).CommitTransfer(ctx, req.(*CommitTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_OpenRiseService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).OpenRiseService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/OpenRiseService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).OpenRiseService(ctx, req.(*PersonId))
	}
	return interceptor(ctx, in, info, handler)
}

// FinanceService_ServiceDesc is the grpc.ServiceDesc for FinanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FinanceService",
	HandlerType: (*FinanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRiseInfo",
			Handler:    _FinanceService_GetRiseInfo_Handler,
		},
		{
			MethodName: "RiseTransferIn",
			Handler:    _FinanceService_RiseTransferIn_Handler,
		},
		{
			MethodName: "RiseTransferOut",
			Handler:    _FinanceService_RiseTransferOut_Handler,
		},
		{
			MethodName: "RiseSettleByDay",
			Handler:    _FinanceService_RiseSettleByDay_Handler,
		},
		{
			MethodName: "CommitTransfer",
			Handler:    _FinanceService_CommitTransfer_Handler,
		},
		{
			MethodName: "OpenRiseService",
			Handler:    _FinanceService_OpenRiseService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance_service.proto",
}
