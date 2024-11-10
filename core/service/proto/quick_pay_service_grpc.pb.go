// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.1
// source: quick_pay_service.proto

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

// QuickPayServiceClient is the client API for QuickPayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuickPayServiceClient interface {
	// 银行卡信息查询
	QueryCardBin(ctx context.Context, in *BankCardNo, opts ...grpc.CallOption) (*CardBinQueryResponse, error)
	// 检查签名,用于回调时判断参数正确
	CheckSign(ctx context.Context, in *CheckQPaySignRequest, opts ...grpc.CallOption) (*CheckQPaySignResponse, error)
	// 申请银行侧认证授权(某些银行需跳转到银行页面进行授权)
	RequestBankSideAuth(ctx context.Context, in *BankAuthRequest, opts ...grpc.CallOption) (*BankAuthResponse, error)
	// 查询银行认证授权结果(根据申请认证的随机Id或者银行卡号)
	QueryBankAuthResult(ctx context.Context, in *BankAuthQueryRequest, opts ...grpc.CallOption) (*BankAuthQueryResponse, error)
	// 直接支付
	DirectPayment(ctx context.Context, in *QPaymentRequest, opts ...grpc.CallOption) (*QPaymentResponse, error)
	// 查询支付状态
	QueryPaymentStatus(ctx context.Context, in *QPaymentQueryRequest, opts ...grpc.CallOption) (*QPaymentQueryResponse, error)
	// 批量付款
	BatchTransfer(ctx context.Context, in *BatchTransferRequest, opts ...grpc.CallOption) (*BatchTransferResponse, error)
}

type quickPayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuickPayServiceClient(cc grpc.ClientConnInterface) QuickPayServiceClient {
	return &quickPayServiceClient{cc}
}

func (c *quickPayServiceClient) QueryCardBin(ctx context.Context, in *BankCardNo, opts ...grpc.CallOption) (*CardBinQueryResponse, error) {
	out := new(CardBinQueryResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/QueryCardBin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickPayServiceClient) CheckSign(ctx context.Context, in *CheckQPaySignRequest, opts ...grpc.CallOption) (*CheckQPaySignResponse, error) {
	out := new(CheckQPaySignResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/CheckSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickPayServiceClient) RequestBankSideAuth(ctx context.Context, in *BankAuthRequest, opts ...grpc.CallOption) (*BankAuthResponse, error) {
	out := new(BankAuthResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/RequestBankSideAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickPayServiceClient) QueryBankAuthResult(ctx context.Context, in *BankAuthQueryRequest, opts ...grpc.CallOption) (*BankAuthQueryResponse, error) {
	out := new(BankAuthQueryResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/QueryBankAuthResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickPayServiceClient) DirectPayment(ctx context.Context, in *QPaymentRequest, opts ...grpc.CallOption) (*QPaymentResponse, error) {
	out := new(QPaymentResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/DirectPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickPayServiceClient) QueryPaymentStatus(ctx context.Context, in *QPaymentQueryRequest, opts ...grpc.CallOption) (*QPaymentQueryResponse, error) {
	out := new(QPaymentQueryResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/QueryPaymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickPayServiceClient) BatchTransfer(ctx context.Context, in *BatchTransferRequest, opts ...grpc.CallOption) (*BatchTransferResponse, error) {
	out := new(BatchTransferResponse)
	err := c.cc.Invoke(ctx, "/QuickPayService/BatchTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuickPayServiceServer is the server API for QuickPayService service.
// All implementations must embed UnimplementedQuickPayServiceServer
// for forward compatibility
type QuickPayServiceServer interface {
	// 银行卡信息查询
	QueryCardBin(context.Context, *BankCardNo) (*CardBinQueryResponse, error)
	// 检查签名,用于回调时判断参数正确
	CheckSign(context.Context, *CheckQPaySignRequest) (*CheckQPaySignResponse, error)
	// 申请银行侧认证授权(某些银行需跳转到银行页面进行授权)
	RequestBankSideAuth(context.Context, *BankAuthRequest) (*BankAuthResponse, error)
	// 查询银行认证授权结果(根据申请认证的随机Id或者银行卡号)
	QueryBankAuthResult(context.Context, *BankAuthQueryRequest) (*BankAuthQueryResponse, error)
	// 直接支付
	DirectPayment(context.Context, *QPaymentRequest) (*QPaymentResponse, error)
	// 查询支付状态
	QueryPaymentStatus(context.Context, *QPaymentQueryRequest) (*QPaymentQueryResponse, error)
	// 批量付款
	BatchTransfer(context.Context, *BatchTransferRequest) (*BatchTransferResponse, error)
	mustEmbedUnimplementedQuickPayServiceServer()
}

// UnimplementedQuickPayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuickPayServiceServer struct {
}

func (UnimplementedQuickPayServiceServer) QueryCardBin(context.Context, *BankCardNo) (*CardBinQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCardBin not implemented")
}
func (UnimplementedQuickPayServiceServer) CheckSign(context.Context, *CheckQPaySignRequest) (*CheckQPaySignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSign not implemented")
}
func (UnimplementedQuickPayServiceServer) RequestBankSideAuth(context.Context, *BankAuthRequest) (*BankAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestBankSideAuth not implemented")
}
func (UnimplementedQuickPayServiceServer) QueryBankAuthResult(context.Context, *BankAuthQueryRequest) (*BankAuthQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBankAuthResult not implemented")
}
func (UnimplementedQuickPayServiceServer) DirectPayment(context.Context, *QPaymentRequest) (*QPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectPayment not implemented")
}
func (UnimplementedQuickPayServiceServer) QueryPaymentStatus(context.Context, *QPaymentQueryRequest) (*QPaymentQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPaymentStatus not implemented")
}
func (UnimplementedQuickPayServiceServer) BatchTransfer(context.Context, *BatchTransferRequest) (*BatchTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchTransfer not implemented")
}
func (UnimplementedQuickPayServiceServer) mustEmbedUnimplementedQuickPayServiceServer() {}

// UnsafeQuickPayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuickPayServiceServer will
// result in compilation errors.
type UnsafeQuickPayServiceServer interface {
	mustEmbedUnimplementedQuickPayServiceServer()
}

func RegisterQuickPayServiceServer(s grpc.ServiceRegistrar, srv QuickPayServiceServer) {
	s.RegisterService(&QuickPayService_ServiceDesc, srv)
}

func _QuickPayService_QueryCardBin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankCardNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).QueryCardBin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/QueryCardBin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).QueryCardBin(ctx, req.(*BankCardNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickPayService_CheckSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckQPaySignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).CheckSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/CheckSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).CheckSign(ctx, req.(*CheckQPaySignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickPayService_RequestBankSideAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).RequestBankSideAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/RequestBankSideAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).RequestBankSideAuth(ctx, req.(*BankAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickPayService_QueryBankAuthResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankAuthQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).QueryBankAuthResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/QueryBankAuthResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).QueryBankAuthResult(ctx, req.(*BankAuthQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickPayService_DirectPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).DirectPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/DirectPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).DirectPayment(ctx, req.(*QPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickPayService_QueryPaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPaymentQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).QueryPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/QueryPaymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).QueryPaymentStatus(ctx, req.(*QPaymentQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickPayService_BatchTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickPayServiceServer).BatchTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuickPayService/BatchTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickPayServiceServer).BatchTransfer(ctx, req.(*BatchTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuickPayService_ServiceDesc is the grpc.ServiceDesc for QuickPayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuickPayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QuickPayService",
	HandlerType: (*QuickPayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryCardBin",
			Handler:    _QuickPayService_QueryCardBin_Handler,
		},
		{
			MethodName: "CheckSign",
			Handler:    _QuickPayService_CheckSign_Handler,
		},
		{
			MethodName: "RequestBankSideAuth",
			Handler:    _QuickPayService_RequestBankSideAuth_Handler,
		},
		{
			MethodName: "QueryBankAuthResult",
			Handler:    _QuickPayService_QueryBankAuthResult_Handler,
		},
		{
			MethodName: "DirectPayment",
			Handler:    _QuickPayService_DirectPayment_Handler,
		},
		{
			MethodName: "QueryPaymentStatus",
			Handler:    _QuickPayService_QueryPaymentStatus_Handler,
		},
		{
			MethodName: "BatchTransfer",
			Handler:    _QuickPayService_BatchTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quick_pay_service.proto",
}
