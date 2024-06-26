// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: query_service.proto

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
	QueryService_SummaryStatistics_FullMethodName        = "/QueryService/SummaryStatistics"
	QueryService_MemberStatistics_FullMethodName         = "/QueryService/MemberStatistics"
	QueryService_PagingShops_FullMethodName              = "/QueryService/PagingShops"
	QueryService_MemberNormalOrders_FullMethodName       = "/QueryService/MemberNormalOrders"
	QueryService_QueryWholesaleOrders_FullMethodName     = "/QueryService/QueryWholesaleOrders"
	QueryService_QueryTradeOrders_FullMethodName         = "/QueryService/QueryTradeOrders"
	QueryService_QueryMemberList_FullMethodName          = "/QueryService/QueryMemberList"
	QueryService_SearchMembers_FullMethodName            = "/QueryService/SearchMembers"
	QueryService_QueryMemberFavoriteShops_FullMethodName = "/QueryService/QueryMemberFavoriteShops"
	QueryService_QueryMemberFavoriteGoods_FullMethodName = "/QueryService/QueryMemberFavoriteGoods"
	QueryService_PagingMemberAccountLog_FullMethodName   = "/QueryService/PagingMemberAccountLog"
	QueryService_PagingOnShelvesGoods_FullMethodName     = "/QueryService/PagingOnShelvesGoods"
	QueryService_QueryItemSalesHistory_FullMethodName    = "/QueryService/QueryItemSalesHistory"
	QueryService_SearchItem_FullMethodName               = "/QueryService/SearchItem"
)

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceClient interface {
	// 汇总统计
	SummaryStatistics(ctx context.Context, in *SummaryStatisticsRequest, opts ...grpc.CallOption) (*SummaryStatisticsResponse, error)
	// * 获取会员的订单状态及其数量
	MemberStatistics(ctx context.Context, in *MemberStatisticsRequest, opts ...grpc.CallOption) (*MemberStatisticsResponse, error)
	// 获取分页店铺数据
	PagingShops(ctx context.Context, in *QueryPagingShopRequest, opts ...grpc.CallOption) (*QueryPagingShopsResponse, error)
	// 查询分页普通订单
	MemberNormalOrders(ctx context.Context, in *MemberOrderPagingRequest, opts ...grpc.CallOption) (*MemberOrderPagingResponse, error)
	// 查询分页批发订单
	QueryWholesaleOrders(ctx context.Context, in *MemberOrderPagingRequest, opts ...grpc.CallOption) (*MemberOrderPagingResponse, error)
	// 查询分页交易/服务类订单
	QueryTradeOrders(ctx context.Context, in *MemberOrderPagingRequest, opts ...grpc.CallOption) (*MemberOrderPagingResponse, error)
	// 查询会员列表
	QueryMemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error)
	// 查询会员列表
	SearchMembers(ctx context.Context, in *MemberSearchRequest, opts ...grpc.CallOption) (*MemberListResponse, error)
	// 获取分页店铺收藏
	QueryMemberFavoriteShops(ctx context.Context, in *FavoriteQueryRequest, opts ...grpc.CallOption) (*PagingShopFavoriteResponse, error)
	// 获取分页店铺收藏
	QueryMemberFavoriteGoods(ctx context.Context, in *FavoriteQueryRequest, opts ...grpc.CallOption) (*PagingGoodsFavoriteResponse, error)
	// 获取指定账户的流水记录
	PagingMemberAccountLog(ctx context.Context, in *PagingAccountLogRequest, opts ...grpc.CallOption) (*MemberAccountPagingLogResponse, error)
	// 获取店铺分页上架的商品
	PagingOnShelvesGoods(ctx context.Context, in *PagingShopGoodsRequest, opts ...grpc.CallOption) (*PagingShopGoodsResponse, error)
	// 查询商品销售记录
	QueryItemSalesHistory(ctx context.Context, in *QueryItemSalesHistoryRequest, opts ...grpc.CallOption) (*QueryItemSalesHistoryResponse, error)
	// 搜索商品
	SearchItem(ctx context.Context, in *SearchItemRequest, opts ...grpc.CallOption) (*SearchItemResponse, error)
}

type queryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceClient(cc grpc.ClientConnInterface) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) SummaryStatistics(ctx context.Context, in *SummaryStatisticsRequest, opts ...grpc.CallOption) (*SummaryStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SummaryStatisticsResponse)
	err := c.cc.Invoke(ctx, QueryService_SummaryStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) MemberStatistics(ctx context.Context, in *MemberStatisticsRequest, opts ...grpc.CallOption) (*MemberStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberStatisticsResponse)
	err := c.cc.Invoke(ctx, QueryService_MemberStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) PagingShops(ctx context.Context, in *QueryPagingShopRequest, opts ...grpc.CallOption) (*QueryPagingShopsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPagingShopsResponse)
	err := c.cc.Invoke(ctx, QueryService_PagingShops_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) MemberNormalOrders(ctx context.Context, in *MemberOrderPagingRequest, opts ...grpc.CallOption) (*MemberOrderPagingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberOrderPagingResponse)
	err := c.cc.Invoke(ctx, QueryService_MemberNormalOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryWholesaleOrders(ctx context.Context, in *MemberOrderPagingRequest, opts ...grpc.CallOption) (*MemberOrderPagingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberOrderPagingResponse)
	err := c.cc.Invoke(ctx, QueryService_QueryWholesaleOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryTradeOrders(ctx context.Context, in *MemberOrderPagingRequest, opts ...grpc.CallOption) (*MemberOrderPagingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberOrderPagingResponse)
	err := c.cc.Invoke(ctx, QueryService_QueryTradeOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryMemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberListResponse)
	err := c.cc.Invoke(ctx, QueryService_QueryMemberList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) SearchMembers(ctx context.Context, in *MemberSearchRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberListResponse)
	err := c.cc.Invoke(ctx, QueryService_SearchMembers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryMemberFavoriteShops(ctx context.Context, in *FavoriteQueryRequest, opts ...grpc.CallOption) (*PagingShopFavoriteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PagingShopFavoriteResponse)
	err := c.cc.Invoke(ctx, QueryService_QueryMemberFavoriteShops_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryMemberFavoriteGoods(ctx context.Context, in *FavoriteQueryRequest, opts ...grpc.CallOption) (*PagingGoodsFavoriteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PagingGoodsFavoriteResponse)
	err := c.cc.Invoke(ctx, QueryService_QueryMemberFavoriteGoods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) PagingMemberAccountLog(ctx context.Context, in *PagingAccountLogRequest, opts ...grpc.CallOption) (*MemberAccountPagingLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberAccountPagingLogResponse)
	err := c.cc.Invoke(ctx, QueryService_PagingMemberAccountLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) PagingOnShelvesGoods(ctx context.Context, in *PagingShopGoodsRequest, opts ...grpc.CallOption) (*PagingShopGoodsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PagingShopGoodsResponse)
	err := c.cc.Invoke(ctx, QueryService_PagingOnShelvesGoods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryItemSalesHistory(ctx context.Context, in *QueryItemSalesHistoryRequest, opts ...grpc.CallOption) (*QueryItemSalesHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryItemSalesHistoryResponse)
	err := c.cc.Invoke(ctx, QueryService_QueryItemSalesHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) SearchItem(ctx context.Context, in *SearchItemRequest, opts ...grpc.CallOption) (*SearchItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchItemResponse)
	err := c.cc.Invoke(ctx, QueryService_SearchItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
// All implementations must embed UnimplementedQueryServiceServer
// for forward compatibility
type QueryServiceServer interface {
	// 汇总统计
	SummaryStatistics(context.Context, *SummaryStatisticsRequest) (*SummaryStatisticsResponse, error)
	// * 获取会员的订单状态及其数量
	MemberStatistics(context.Context, *MemberStatisticsRequest) (*MemberStatisticsResponse, error)
	// 获取分页店铺数据
	PagingShops(context.Context, *QueryPagingShopRequest) (*QueryPagingShopsResponse, error)
	// 查询分页普通订单
	MemberNormalOrders(context.Context, *MemberOrderPagingRequest) (*MemberOrderPagingResponse, error)
	// 查询分页批发订单
	QueryWholesaleOrders(context.Context, *MemberOrderPagingRequest) (*MemberOrderPagingResponse, error)
	// 查询分页交易/服务类订单
	QueryTradeOrders(context.Context, *MemberOrderPagingRequest) (*MemberOrderPagingResponse, error)
	// 查询会员列表
	QueryMemberList(context.Context, *MemberListRequest) (*MemberListResponse, error)
	// 查询会员列表
	SearchMembers(context.Context, *MemberSearchRequest) (*MemberListResponse, error)
	// 获取分页店铺收藏
	QueryMemberFavoriteShops(context.Context, *FavoriteQueryRequest) (*PagingShopFavoriteResponse, error)
	// 获取分页店铺收藏
	QueryMemberFavoriteGoods(context.Context, *FavoriteQueryRequest) (*PagingGoodsFavoriteResponse, error)
	// 获取指定账户的流水记录
	PagingMemberAccountLog(context.Context, *PagingAccountLogRequest) (*MemberAccountPagingLogResponse, error)
	// 获取店铺分页上架的商品
	PagingOnShelvesGoods(context.Context, *PagingShopGoodsRequest) (*PagingShopGoodsResponse, error)
	// 查询商品销售记录
	QueryItemSalesHistory(context.Context, *QueryItemSalesHistoryRequest) (*QueryItemSalesHistoryResponse, error)
	// 搜索商品
	SearchItem(context.Context, *SearchItemRequest) (*SearchItemResponse, error)
	mustEmbedUnimplementedQueryServiceServer()
}

// UnimplementedQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (UnimplementedQueryServiceServer) SummaryStatistics(context.Context, *SummaryStatisticsRequest) (*SummaryStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryStatistics not implemented")
}
func (UnimplementedQueryServiceServer) MemberStatistics(context.Context, *MemberStatisticsRequest) (*MemberStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberStatistics not implemented")
}
func (UnimplementedQueryServiceServer) PagingShops(context.Context, *QueryPagingShopRequest) (*QueryPagingShopsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PagingShops not implemented")
}
func (UnimplementedQueryServiceServer) MemberNormalOrders(context.Context, *MemberOrderPagingRequest) (*MemberOrderPagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberNormalOrders not implemented")
}
func (UnimplementedQueryServiceServer) QueryWholesaleOrders(context.Context, *MemberOrderPagingRequest) (*MemberOrderPagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWholesaleOrders not implemented")
}
func (UnimplementedQueryServiceServer) QueryTradeOrders(context.Context, *MemberOrderPagingRequest) (*MemberOrderPagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTradeOrders not implemented")
}
func (UnimplementedQueryServiceServer) QueryMemberList(context.Context, *MemberListRequest) (*MemberListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMemberList not implemented")
}
func (UnimplementedQueryServiceServer) SearchMembers(context.Context, *MemberSearchRequest) (*MemberListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMembers not implemented")
}
func (UnimplementedQueryServiceServer) QueryMemberFavoriteShops(context.Context, *FavoriteQueryRequest) (*PagingShopFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMemberFavoriteShops not implemented")
}
func (UnimplementedQueryServiceServer) QueryMemberFavoriteGoods(context.Context, *FavoriteQueryRequest) (*PagingGoodsFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMemberFavoriteGoods not implemented")
}
func (UnimplementedQueryServiceServer) PagingMemberAccountLog(context.Context, *PagingAccountLogRequest) (*MemberAccountPagingLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PagingMemberAccountLog not implemented")
}
func (UnimplementedQueryServiceServer) PagingOnShelvesGoods(context.Context, *PagingShopGoodsRequest) (*PagingShopGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PagingOnShelvesGoods not implemented")
}
func (UnimplementedQueryServiceServer) QueryItemSalesHistory(context.Context, *QueryItemSalesHistoryRequest) (*QueryItemSalesHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryItemSalesHistory not implemented")
}
func (UnimplementedQueryServiceServer) SearchItem(context.Context, *SearchItemRequest) (*SearchItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchItem not implemented")
}
func (UnimplementedQueryServiceServer) mustEmbedUnimplementedQueryServiceServer() {}

// UnsafeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceServer will
// result in compilation errors.
type UnsafeQueryServiceServer interface {
	mustEmbedUnimplementedQueryServiceServer()
}

func RegisterQueryServiceServer(s grpc.ServiceRegistrar, srv QueryServiceServer) {
	s.RegisterService(&QueryService_ServiceDesc, srv)
}

func _QueryService_SummaryStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).SummaryStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_SummaryStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).SummaryStatistics(ctx, req.(*SummaryStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_MemberStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).MemberStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_MemberStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).MemberStatistics(ctx, req.(*MemberStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_PagingShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPagingShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PagingShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_PagingShops_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PagingShops(ctx, req.(*QueryPagingShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_MemberNormalOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberOrderPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).MemberNormalOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_MemberNormalOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).MemberNormalOrders(ctx, req.(*MemberOrderPagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryWholesaleOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberOrderPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryWholesaleOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_QueryWholesaleOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryWholesaleOrders(ctx, req.(*MemberOrderPagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryTradeOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberOrderPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryTradeOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_QueryTradeOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryTradeOrders(ctx, req.(*MemberOrderPagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_QueryMemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryMemberList(ctx, req.(*MemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_SearchMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).SearchMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_SearchMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).SearchMembers(ctx, req.(*MemberSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryMemberFavoriteShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryMemberFavoriteShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_QueryMemberFavoriteShops_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryMemberFavoriteShops(ctx, req.(*FavoriteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryMemberFavoriteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryMemberFavoriteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_QueryMemberFavoriteGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryMemberFavoriteGoods(ctx, req.(*FavoriteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_PagingMemberAccountLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingAccountLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PagingMemberAccountLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_PagingMemberAccountLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PagingMemberAccountLog(ctx, req.(*PagingAccountLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_PagingOnShelvesGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingShopGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PagingOnShelvesGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_PagingOnShelvesGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PagingOnShelvesGoods(ctx, req.(*PagingShopGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryItemSalesHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryItemSalesHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryItemSalesHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_QueryItemSalesHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryItemSalesHistory(ctx, req.(*QueryItemSalesHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_SearchItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).SearchItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_SearchItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).SearchItem(ctx, req.(*SearchItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryService_ServiceDesc is the grpc.ServiceDesc for QueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SummaryStatistics",
			Handler:    _QueryService_SummaryStatistics_Handler,
		},
		{
			MethodName: "MemberStatistics",
			Handler:    _QueryService_MemberStatistics_Handler,
		},
		{
			MethodName: "PagingShops",
			Handler:    _QueryService_PagingShops_Handler,
		},
		{
			MethodName: "MemberNormalOrders",
			Handler:    _QueryService_MemberNormalOrders_Handler,
		},
		{
			MethodName: "QueryWholesaleOrders",
			Handler:    _QueryService_QueryWholesaleOrders_Handler,
		},
		{
			MethodName: "QueryTradeOrders",
			Handler:    _QueryService_QueryTradeOrders_Handler,
		},
		{
			MethodName: "QueryMemberList",
			Handler:    _QueryService_QueryMemberList_Handler,
		},
		{
			MethodName: "SearchMembers",
			Handler:    _QueryService_SearchMembers_Handler,
		},
		{
			MethodName: "QueryMemberFavoriteShops",
			Handler:    _QueryService_QueryMemberFavoriteShops_Handler,
		},
		{
			MethodName: "QueryMemberFavoriteGoods",
			Handler:    _QueryService_QueryMemberFavoriteGoods_Handler,
		},
		{
			MethodName: "PagingMemberAccountLog",
			Handler:    _QueryService_PagingMemberAccountLog_Handler,
		},
		{
			MethodName: "PagingOnShelvesGoods",
			Handler:    _QueryService_PagingOnShelvesGoods_Handler,
		},
		{
			MethodName: "QueryItemSalesHistory",
			Handler:    _QueryService_QueryItemSalesHistory_Handler,
		},
		{
			MethodName: "SearchItem",
			Handler:    _QueryService_SearchItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query_service.proto",
}
