// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: aftersales_service.proto

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
	AfterSalesService_SubmitAfterSalesOrder_FullMethodName             = "/AfterSalesService/SubmitAfterSalesOrder"
	AfterSalesService_GetAllAfterSalesOrderOfSaleOrder_FullMethodName  = "/AfterSalesService/GetAllAfterSalesOrderOfSaleOrder"
	AfterSalesService_QueryPagerAfterSalesOrderOfMember_FullMethodName = "/AfterSalesService/QueryPagerAfterSalesOrderOfMember"
	AfterSalesService_QueryPagerAfterSalesOrderOfVendor_FullMethodName = "/AfterSalesService/QueryPagerAfterSalesOrderOfVendor"
	AfterSalesService_GetAfterSaleOrder_FullMethodName                 = "/AfterSalesService/GetAfterSaleOrder"
	AfterSalesService_AgreeAfterSales_FullMethodName                   = "/AfterSalesService/AgreeAfterSales"
	AfterSalesService_DeclineAfterSales_FullMethodName                 = "/AfterSalesService/DeclineAfterSales"
	AfterSalesService_RequestIntercede_FullMethodName                  = "/AfterSalesService/RequestIntercede"
	AfterSalesService_ConfirmAfterSales_FullMethodName                 = "/AfterSalesService/ConfirmAfterSales"
	AfterSalesService_RejectAfterSales_FullMethodName                  = "/AfterSalesService/RejectAfterSales"
	AfterSalesService_ProcessAfterSalesOrder_FullMethodName            = "/AfterSalesService/ProcessAfterSalesOrder"
	AfterSalesService_ReturnShipment_FullMethodName                    = "/AfterSalesService/ReturnShipment"
	AfterSalesService_ReceiveItem_FullMethodName                       = "/AfterSalesService/ReceiveItem"
)

// AfterSalesServiceClient is the client API for AfterSalesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AfterSalesServiceClient interface {
	// 提交售后单
	SubmitAfterSalesOrder(ctx context.Context, in *SubmitAfterSalesOrderRequest, opts ...grpc.CallOption) (*SubmitAfterSalesOrderResponse, error)
	// 获取订单的所有售后单,orderId
	GetAllAfterSalesOrderOfSaleOrder(ctx context.Context, in *OriginOrderIdRequest, opts ...grpc.CallOption) (*AfterSalesOrderListResponse, error)
	// 获取会员的分页售后单
	QueryPagerAfterSalesOrderOfMember(ctx context.Context, in *PagingBuyerOrdersRequest, opts ...grpc.CallOption) (*PagingBuyerAfterSalesOrderListResponse, error)
	// 获取商户的分页售后单
	QueryPagerAfterSalesOrderOfVendor(ctx context.Context, in *PagingSellerOrdersRequest, opts ...grpc.CallOption) (*PagingSellerAfterSalesOrderListResponse, error)
	// 获取售后单
	GetAfterSaleOrder(ctx context.Context, in *AfterSalesOrderNo, opts ...grpc.CallOption) (*SAfterSalesOrder, error)
	// 同意售后
	AgreeAfterSales(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error)
	// 拒绝售后
	DeclineAfterSales(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error)
	// 申请调解
	RequestIntercede(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error)
	// 系统确认
	ConfirmAfterSales(ctx context.Context, in *AfterSalesOrderNo, opts ...grpc.CallOption) (*Result, error)
	// 系统退回
	RejectAfterSales(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error)
	// 处理退款/退货完成,一般是系统自动调用
	ProcessAfterSalesOrder(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error)
	// 换货发货
	ReturnShipment(ctx context.Context, in *ReturnShipmentRequest, opts ...grpc.CallOption) (*Result, error)
	// 换货收货
	ReceiveItem(ctx context.Context, in *AfterSalesOrderNo, opts ...grpc.CallOption) (*Result, error)
}

type afterSalesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAfterSalesServiceClient(cc grpc.ClientConnInterface) AfterSalesServiceClient {
	return &afterSalesServiceClient{cc}
}

func (c *afterSalesServiceClient) SubmitAfterSalesOrder(ctx context.Context, in *SubmitAfterSalesOrderRequest, opts ...grpc.CallOption) (*SubmitAfterSalesOrderResponse, error) {
	out := new(SubmitAfterSalesOrderResponse)
	err := c.cc.Invoke(ctx, AfterSalesService_SubmitAfterSalesOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) GetAllAfterSalesOrderOfSaleOrder(ctx context.Context, in *OriginOrderIdRequest, opts ...grpc.CallOption) (*AfterSalesOrderListResponse, error) {
	out := new(AfterSalesOrderListResponse)
	err := c.cc.Invoke(ctx, AfterSalesService_GetAllAfterSalesOrderOfSaleOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) QueryPagerAfterSalesOrderOfMember(ctx context.Context, in *PagingBuyerOrdersRequest, opts ...grpc.CallOption) (*PagingBuyerAfterSalesOrderListResponse, error) {
	out := new(PagingBuyerAfterSalesOrderListResponse)
	err := c.cc.Invoke(ctx, AfterSalesService_QueryPagerAfterSalesOrderOfMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) QueryPagerAfterSalesOrderOfVendor(ctx context.Context, in *PagingSellerOrdersRequest, opts ...grpc.CallOption) (*PagingSellerAfterSalesOrderListResponse, error) {
	out := new(PagingSellerAfterSalesOrderListResponse)
	err := c.cc.Invoke(ctx, AfterSalesService_QueryPagerAfterSalesOrderOfVendor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) GetAfterSaleOrder(ctx context.Context, in *AfterSalesOrderNo, opts ...grpc.CallOption) (*SAfterSalesOrder, error) {
	out := new(SAfterSalesOrder)
	err := c.cc.Invoke(ctx, AfterSalesService_GetAfterSaleOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) AgreeAfterSales(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_AgreeAfterSales_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) DeclineAfterSales(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_DeclineAfterSales_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) RequestIntercede(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_RequestIntercede_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) ConfirmAfterSales(ctx context.Context, in *AfterSalesOrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_ConfirmAfterSales_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) RejectAfterSales(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_RejectAfterSales_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) ProcessAfterSalesOrder(ctx context.Context, in *AfterSalesProcessRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_ProcessAfterSalesOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) ReturnShipment(ctx context.Context, in *ReturnShipmentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_ReturnShipment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterSalesServiceClient) ReceiveItem(ctx context.Context, in *AfterSalesOrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AfterSalesService_ReceiveItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AfterSalesServiceServer is the server API for AfterSalesService service.
// All implementations must embed UnimplementedAfterSalesServiceServer
// for forward compatibility
type AfterSalesServiceServer interface {
	// 提交售后单
	SubmitAfterSalesOrder(context.Context, *SubmitAfterSalesOrderRequest) (*SubmitAfterSalesOrderResponse, error)
	// 获取订单的所有售后单,orderId
	GetAllAfterSalesOrderOfSaleOrder(context.Context, *OriginOrderIdRequest) (*AfterSalesOrderListResponse, error)
	// 获取会员的分页售后单
	QueryPagerAfterSalesOrderOfMember(context.Context, *PagingBuyerOrdersRequest) (*PagingBuyerAfterSalesOrderListResponse, error)
	// 获取商户的分页售后单
	QueryPagerAfterSalesOrderOfVendor(context.Context, *PagingSellerOrdersRequest) (*PagingSellerAfterSalesOrderListResponse, error)
	// 获取售后单
	GetAfterSaleOrder(context.Context, *AfterSalesOrderNo) (*SAfterSalesOrder, error)
	// 同意售后
	AgreeAfterSales(context.Context, *AfterSalesProcessRequest) (*Result, error)
	// 拒绝售后
	DeclineAfterSales(context.Context, *AfterSalesProcessRequest) (*Result, error)
	// 申请调解
	RequestIntercede(context.Context, *AfterSalesProcessRequest) (*Result, error)
	// 系统确认
	ConfirmAfterSales(context.Context, *AfterSalesOrderNo) (*Result, error)
	// 系统退回
	RejectAfterSales(context.Context, *AfterSalesProcessRequest) (*Result, error)
	// 处理退款/退货完成,一般是系统自动调用
	ProcessAfterSalesOrder(context.Context, *AfterSalesProcessRequest) (*Result, error)
	// 换货发货
	ReturnShipment(context.Context, *ReturnShipmentRequest) (*Result, error)
	// 换货收货
	ReceiveItem(context.Context, *AfterSalesOrderNo) (*Result, error)
	mustEmbedUnimplementedAfterSalesServiceServer()
}

// UnimplementedAfterSalesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAfterSalesServiceServer struct {
}

func (UnimplementedAfterSalesServiceServer) SubmitAfterSalesOrder(context.Context, *SubmitAfterSalesOrderRequest) (*SubmitAfterSalesOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAfterSalesOrder not implemented")
}
func (UnimplementedAfterSalesServiceServer) GetAllAfterSalesOrderOfSaleOrder(context.Context, *OriginOrderIdRequest) (*AfterSalesOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAfterSalesOrderOfSaleOrder not implemented")
}
func (UnimplementedAfterSalesServiceServer) QueryPagerAfterSalesOrderOfMember(context.Context, *PagingBuyerOrdersRequest) (*PagingBuyerAfterSalesOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPagerAfterSalesOrderOfMember not implemented")
}
func (UnimplementedAfterSalesServiceServer) QueryPagerAfterSalesOrderOfVendor(context.Context, *PagingSellerOrdersRequest) (*PagingSellerAfterSalesOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPagerAfterSalesOrderOfVendor not implemented")
}
func (UnimplementedAfterSalesServiceServer) GetAfterSaleOrder(context.Context, *AfterSalesOrderNo) (*SAfterSalesOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAfterSaleOrder not implemented")
}
func (UnimplementedAfterSalesServiceServer) AgreeAfterSales(context.Context, *AfterSalesProcessRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreeAfterSales not implemented")
}
func (UnimplementedAfterSalesServiceServer) DeclineAfterSales(context.Context, *AfterSalesProcessRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineAfterSales not implemented")
}
func (UnimplementedAfterSalesServiceServer) RequestIntercede(context.Context, *AfterSalesProcessRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestIntercede not implemented")
}
func (UnimplementedAfterSalesServiceServer) ConfirmAfterSales(context.Context, *AfterSalesOrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmAfterSales not implemented")
}
func (UnimplementedAfterSalesServiceServer) RejectAfterSales(context.Context, *AfterSalesProcessRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectAfterSales not implemented")
}
func (UnimplementedAfterSalesServiceServer) ProcessAfterSalesOrder(context.Context, *AfterSalesProcessRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessAfterSalesOrder not implemented")
}
func (UnimplementedAfterSalesServiceServer) ReturnShipment(context.Context, *ReturnShipmentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnShipment not implemented")
}
func (UnimplementedAfterSalesServiceServer) ReceiveItem(context.Context, *AfterSalesOrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveItem not implemented")
}
func (UnimplementedAfterSalesServiceServer) mustEmbedUnimplementedAfterSalesServiceServer() {}

// UnsafeAfterSalesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AfterSalesServiceServer will
// result in compilation errors.
type UnsafeAfterSalesServiceServer interface {
	mustEmbedUnimplementedAfterSalesServiceServer()
}

func RegisterAfterSalesServiceServer(s grpc.ServiceRegistrar, srv AfterSalesServiceServer) {
	s.RegisterService(&AfterSalesService_ServiceDesc, srv)
}

func _AfterSalesService_SubmitAfterSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAfterSalesOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).SubmitAfterSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_SubmitAfterSalesOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).SubmitAfterSalesOrder(ctx, req.(*SubmitAfterSalesOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_GetAllAfterSalesOrderOfSaleOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginOrderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).GetAllAfterSalesOrderOfSaleOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_GetAllAfterSalesOrderOfSaleOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).GetAllAfterSalesOrderOfSaleOrder(ctx, req.(*OriginOrderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_QueryPagerAfterSalesOrderOfMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingBuyerOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).QueryPagerAfterSalesOrderOfMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_QueryPagerAfterSalesOrderOfMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).QueryPagerAfterSalesOrderOfMember(ctx, req.(*PagingBuyerOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_QueryPagerAfterSalesOrderOfVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingSellerOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).QueryPagerAfterSalesOrderOfVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_QueryPagerAfterSalesOrderOfVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).QueryPagerAfterSalesOrderOfVendor(ctx, req.(*PagingSellerOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_GetAfterSaleOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesOrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).GetAfterSaleOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_GetAfterSaleOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).GetAfterSaleOrder(ctx, req.(*AfterSalesOrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_AgreeAfterSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).AgreeAfterSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_AgreeAfterSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).AgreeAfterSales(ctx, req.(*AfterSalesProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_DeclineAfterSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).DeclineAfterSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_DeclineAfterSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).DeclineAfterSales(ctx, req.(*AfterSalesProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_RequestIntercede_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).RequestIntercede(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_RequestIntercede_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).RequestIntercede(ctx, req.(*AfterSalesProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_ConfirmAfterSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesOrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).ConfirmAfterSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_ConfirmAfterSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).ConfirmAfterSales(ctx, req.(*AfterSalesOrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_RejectAfterSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).RejectAfterSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_RejectAfterSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).RejectAfterSales(ctx, req.(*AfterSalesProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_ProcessAfterSalesOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).ProcessAfterSalesOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_ProcessAfterSalesOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).ProcessAfterSalesOrder(ctx, req.(*AfterSalesProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_ReturnShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).ReturnShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_ReturnShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).ReturnShipment(ctx, req.(*ReturnShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AfterSalesService_ReceiveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterSalesOrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterSalesServiceServer).ReceiveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AfterSalesService_ReceiveItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterSalesServiceServer).ReceiveItem(ctx, req.(*AfterSalesOrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

// AfterSalesService_ServiceDesc is the grpc.ServiceDesc for AfterSalesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AfterSalesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AfterSalesService",
	HandlerType: (*AfterSalesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitAfterSalesOrder",
			Handler:    _AfterSalesService_SubmitAfterSalesOrder_Handler,
		},
		{
			MethodName: "GetAllAfterSalesOrderOfSaleOrder",
			Handler:    _AfterSalesService_GetAllAfterSalesOrderOfSaleOrder_Handler,
		},
		{
			MethodName: "QueryPagerAfterSalesOrderOfMember",
			Handler:    _AfterSalesService_QueryPagerAfterSalesOrderOfMember_Handler,
		},
		{
			MethodName: "QueryPagerAfterSalesOrderOfVendor",
			Handler:    _AfterSalesService_QueryPagerAfterSalesOrderOfVendor_Handler,
		},
		{
			MethodName: "GetAfterSaleOrder",
			Handler:    _AfterSalesService_GetAfterSaleOrder_Handler,
		},
		{
			MethodName: "AgreeAfterSales",
			Handler:    _AfterSalesService_AgreeAfterSales_Handler,
		},
		{
			MethodName: "DeclineAfterSales",
			Handler:    _AfterSalesService_DeclineAfterSales_Handler,
		},
		{
			MethodName: "RequestIntercede",
			Handler:    _AfterSalesService_RequestIntercede_Handler,
		},
		{
			MethodName: "ConfirmAfterSales",
			Handler:    _AfterSalesService_ConfirmAfterSales_Handler,
		},
		{
			MethodName: "RejectAfterSales",
			Handler:    _AfterSalesService_RejectAfterSales_Handler,
		},
		{
			MethodName: "ProcessAfterSalesOrder",
			Handler:    _AfterSalesService_ProcessAfterSalesOrder_Handler,
		},
		{
			MethodName: "ReturnShipment",
			Handler:    _AfterSalesService_ReturnShipment_Handler,
		},
		{
			MethodName: "ReceiveItem",
			Handler:    _AfterSalesService_ReceiveItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aftersales_service.proto",
}
