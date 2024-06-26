// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: shipment_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ShipmentService_CreateCoverageArea__FullMethodName    = "/ShipmentService/CreateCoverageArea_"
	ShipmentService_GetOrderShipments_FullMethodName      = "/ShipmentService/GetOrderShipments"
	ShipmentService_GetLogisticFlowTrack_FullMethodName   = "/ShipmentService/GetLogisticFlowTrack"
	ShipmentService_ShipOrderLogisticTrack_FullMethodName = "/ShipmentService/ShipOrderLogisticTrack"
)

// ShipmentServiceClient is the client API for ShipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 发货服务
type ShipmentServiceClient interface {
	// 创建一个配送覆盖的区域
	CreateCoverageArea_(ctx context.Context, in *SCoverageValue, opts ...grpc.CallOption) (*Result, error)
	// 获取订单的发货单信息
	GetOrderShipments(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*ShipmentOrderListResponse, error)
	// * 物流追踪
	GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
}

type shipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShipmentServiceClient(cc grpc.ClientConnInterface) ShipmentServiceClient {
	return &shipmentServiceClient{cc}
}

func (c *shipmentServiceClient) CreateCoverageArea_(ctx context.Context, in *SCoverageValue, opts ...grpc.CallOption) (*Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Result)
	err := c.cc.Invoke(ctx, ShipmentService_CreateCoverageArea__FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetOrderShipments(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*ShipmentOrderListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShipmentOrderListResponse)
	err := c.cc.Invoke(ctx, ShipmentService_GetOrderShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, ShipmentService_GetLogisticFlowTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, ShipmentService_ShipOrderLogisticTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentServiceServer is the server API for ShipmentService service.
// All implementations must embed UnimplementedShipmentServiceServer
// for forward compatibility
//
// 发货服务
type ShipmentServiceServer interface {
	// 创建一个配送覆盖的区域
	CreateCoverageArea_(context.Context, *SCoverageValue) (*Result, error)
	// 获取订单的发货单信息
	GetOrderShipments(context.Context, *OrderId) (*ShipmentOrderListResponse, error)
	// * 物流追踪
	GetLogisticFlowTrack(context.Context, *LogisticFlowTrackRequest) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(context.Context, *OrderLogisticTrackRequest) (*SShipOrderTrack, error)
	mustEmbedUnimplementedShipmentServiceServer()
}

// UnimplementedShipmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShipmentServiceServer struct {
}

func (UnimplementedShipmentServiceServer) CreateCoverageArea_(context.Context, *SCoverageValue) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoverageArea_ not implemented")
}
func (UnimplementedShipmentServiceServer) GetOrderShipments(context.Context, *OrderId) (*ShipmentOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderShipments not implemented")
}
func (UnimplementedShipmentServiceServer) GetLogisticFlowTrack(context.Context, *LogisticFlowTrackRequest) (*SShipOrderTrack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogisticFlowTrack not implemented")
}
func (UnimplementedShipmentServiceServer) ShipOrderLogisticTrack(context.Context, *OrderLogisticTrackRequest) (*SShipOrderTrack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipOrderLogisticTrack not implemented")
}
func (UnimplementedShipmentServiceServer) mustEmbedUnimplementedShipmentServiceServer() {}

// UnsafeShipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShipmentServiceServer will
// result in compilation errors.
type UnsafeShipmentServiceServer interface {
	mustEmbedUnimplementedShipmentServiceServer()
}

func RegisterShipmentServiceServer(s grpc.ServiceRegistrar, srv ShipmentServiceServer) {
	s.RegisterService(&ShipmentService_ServiceDesc, srv)
}

func _ShipmentService_CreateCoverageArea__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SCoverageValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).CreateCoverageArea_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_CreateCoverageArea__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).CreateCoverageArea_(ctx, req.(*SCoverageValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetOrderShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetOrderShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_GetOrderShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetOrderShipments(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetLogisticFlowTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogisticFlowTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_GetLogisticFlowTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, req.(*LogisticFlowTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ShipOrderLogisticTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderLogisticTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_ShipOrderLogisticTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, req.(*OrderLogisticTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShipmentService_ServiceDesc is the grpc.ServiceDesc for ShipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ShipmentService",
	HandlerType: (*ShipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoverageArea_",
			Handler:    _ShipmentService_CreateCoverageArea__Handler,
		},
		{
			MethodName: "GetOrderShipments",
			Handler:    _ShipmentService_GetOrderShipments_Handler,
		},
		{
			MethodName: "GetLogisticFlowTrack",
			Handler:    _ShipmentService_GetLogisticFlowTrack_Handler,
		},
		{
			MethodName: "ShipOrderLogisticTrack",
			Handler:    _ShipmentService_ShipOrderLogisticTrack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipment_service.proto",
}
