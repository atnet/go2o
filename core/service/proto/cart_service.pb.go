// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart_service.proto

package proto // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SettlePersistRequest struct {
	// 买家编号
	BuyerId int64 `protobuf:"varint,1,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id"`
	// 店铺编号
	ShopId               int64    `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`
	PaymentOpt           int64    `protobuf:"varint,3,opt,name=payment_opt,json=paymentOpt,proto3" json:"payment_opt"`
	DeliverOpt           int64    `protobuf:"varint,4,opt,name=deliver_opt,json=deliverOpt,proto3" json:"deliver_opt"`
	AddressId            int64    `protobuf:"varint,5,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettlePersistRequest) Reset()         { *m = SettlePersistRequest{} }
func (m *SettlePersistRequest) String() string { return proto.CompactTextString(m) }
func (*SettlePersistRequest) ProtoMessage()    {}
func (*SettlePersistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{0}
}
func (m *SettlePersistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettlePersistRequest.Unmarshal(m, b)
}
func (m *SettlePersistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettlePersistRequest.Marshal(b, m, deterministic)
}
func (dst *SettlePersistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlePersistRequest.Merge(dst, src)
}
func (m *SettlePersistRequest) XXX_Size() int {
	return xxx_messageInfo_SettlePersistRequest.Size(m)
}
func (m *SettlePersistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlePersistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SettlePersistRequest proto.InternalMessageInfo

func (m *SettlePersistRequest) GetBuyerId() int64 {
	if m != nil {
		return m.BuyerId
	}
	return 0
}

func (m *SettlePersistRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SettlePersistRequest) GetPaymentOpt() int64 {
	if m != nil {
		return m.PaymentOpt
	}
	return 0
}

func (m *SettlePersistRequest) GetDeliverOpt() int64 {
	if m != nil {
		return m.DeliverOpt
	}
	return 0
}

func (m *SettlePersistRequest) GetAddressId() int64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

type CartItemRequest struct {
	// 买家编号
	BuyerId int64 `protobuf:"varint,1,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id"`
	// 购物车编码
	CartCode string `protobuf:"bytes,2,opt,name=cart_code,json=cartCode,proto3" json:"cart_code"`
	// 商品编号
	ItemId int64 `protobuf:"zigzag64,3,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	// SKU编号
	SkuId int64 `protobuf:"zigzag64,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	// 数量
	Quantity             int32    `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartItemRequest) Reset()         { *m = CartItemRequest{} }
func (m *CartItemRequest) String() string { return proto.CompactTextString(m) }
func (*CartItemRequest) ProtoMessage()    {}
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{1}
}
func (m *CartItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItemRequest.Unmarshal(m, b)
}
func (m *CartItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItemRequest.Marshal(b, m, deterministic)
}
func (dst *CartItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItemRequest.Merge(dst, src)
}
func (m *CartItemRequest) XXX_Size() int {
	return xxx_messageInfo_CartItemRequest.Size(m)
}
func (m *CartItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CartItemRequest proto.InternalMessageInfo

func (m *CartItemRequest) GetBuyerId() int64 {
	if m != nil {
		return m.BuyerId
	}
	return 0
}

func (m *CartItemRequest) GetCartCode() string {
	if m != nil {
		return m.CartCode
	}
	return ""
}

func (m *CartItemRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *CartItemRequest) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *CartItemRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type CheckSignRequest struct {
	BuyerId              int64             `protobuf:"varint,1,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id"`
	CartCode             string            `protobuf:"bytes,2,opt,name=cart_code,json=cartCode,proto3" json:"cart_code"`
	Items                []*SCheckCartItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSignRequest) Reset()         { *m = CheckSignRequest{} }
func (m *CheckSignRequest) String() string { return proto.CompactTextString(m) }
func (*CheckSignRequest) ProtoMessage()    {}
func (*CheckSignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{2}
}
func (m *CheckSignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSignRequest.Unmarshal(m, b)
}
func (m *CheckSignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSignRequest.Marshal(b, m, deterministic)
}
func (dst *CheckSignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSignRequest.Merge(dst, src)
}
func (m *CheckSignRequest) XXX_Size() int {
	return xxx_messageInfo_CheckSignRequest.Size(m)
}
func (m *CheckSignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSignRequest proto.InternalMessageInfo

func (m *CheckSignRequest) GetBuyerId() int64 {
	if m != nil {
		return m.BuyerId
	}
	return 0
}

func (m *CheckSignRequest) GetCartCode() string {
	if m != nil {
		return m.CartCode
	}
	return ""
}

func (m *CheckSignRequest) GetItems() []*SCheckCartItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type CartCode struct {
	BuyerId              int64    `protobuf:"varint,1,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id"`
	CartCode             string   `protobuf:"bytes,2,opt,name=cart_code,json=cartCode,proto3" json:"cart_code"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartCode) Reset()         { *m = CartCode{} }
func (m *CartCode) String() string { return proto.CompactTextString(m) }
func (*CartCode) ProtoMessage()    {}
func (*CartCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{3}
}
func (m *CartCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartCode.Unmarshal(m, b)
}
func (m *CartCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartCode.Marshal(b, m, deterministic)
}
func (dst *CartCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartCode.Merge(dst, src)
}
func (m *CartCode) XXX_Size() int {
	return xxx_messageInfo_CartCode.Size(m)
}
func (m *CartCode) XXX_DiscardUnknown() {
	xxx_messageInfo_CartCode.DiscardUnknown(m)
}

var xxx_messageInfo_CartCode proto.InternalMessageInfo

func (m *CartCode) GetBuyerId() int64 {
	if m != nil {
		return m.BuyerId
	}
	return 0
}

func (m *CartCode) GetCartCode() string {
	if m != nil {
		return m.CartCode
	}
	return ""
}

type SettleMeta_ struct {
	PaymentOpt           int64               `protobuf:"varint,1,opt,name=payment_opt,json=paymentOpt,proto3" json:"payment_opt"`
	DeliverOpt           int64               `protobuf:"varint,2,opt,name=deliver_opt,json=deliverOpt,proto3" json:"deliver_opt"`
	Shop                 *SettleShopMeta_    `protobuf:"bytes,3,opt,name=shop,proto3" json:"shop"`
	Deliver              *SettleDeliverMeta_ `protobuf:"bytes,4,opt,name=deliver,proto3" json:"deliver"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SettleMeta_) Reset()         { *m = SettleMeta_{} }
func (m *SettleMeta_) String() string { return proto.CompactTextString(m) }
func (*SettleMeta_) ProtoMessage()    {}
func (*SettleMeta_) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{4}
}
func (m *SettleMeta_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleMeta_.Unmarshal(m, b)
}
func (m *SettleMeta_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleMeta_.Marshal(b, m, deterministic)
}
func (dst *SettleMeta_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleMeta_.Merge(dst, src)
}
func (m *SettleMeta_) XXX_Size() int {
	return xxx_messageInfo_SettleMeta_.Size(m)
}
func (m *SettleMeta_) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleMeta_.DiscardUnknown(m)
}

var xxx_messageInfo_SettleMeta_ proto.InternalMessageInfo

func (m *SettleMeta_) GetPaymentOpt() int64 {
	if m != nil {
		return m.PaymentOpt
	}
	return 0
}

func (m *SettleMeta_) GetDeliverOpt() int64 {
	if m != nil {
		return m.DeliverOpt
	}
	return 0
}

func (m *SettleMeta_) GetShop() *SettleShopMeta_ {
	if m != nil {
		return m.Shop
	}
	return nil
}

func (m *SettleMeta_) GetDeliver() *SettleDeliverMeta_ {
	if m != nil {
		return m.Deliver
	}
	return nil
}

type SettleShopMeta_ struct {
	ShopId               int64    `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`
	ShopName             string   `protobuf:"bytes,2,opt,name=shop_name,json=shopName,proto3" json:"shop_name"`
	Telephone            string   `protobuf:"bytes,3,opt,name=telephone,proto3" json:"telephone"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleShopMeta_) Reset()         { *m = SettleShopMeta_{} }
func (m *SettleShopMeta_) String() string { return proto.CompactTextString(m) }
func (*SettleShopMeta_) ProtoMessage()    {}
func (*SettleShopMeta_) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{5}
}
func (m *SettleShopMeta_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleShopMeta_.Unmarshal(m, b)
}
func (m *SettleShopMeta_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleShopMeta_.Marshal(b, m, deterministic)
}
func (dst *SettleShopMeta_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleShopMeta_.Merge(dst, src)
}
func (m *SettleShopMeta_) XXX_Size() int {
	return xxx_messageInfo_SettleShopMeta_.Size(m)
}
func (m *SettleShopMeta_) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleShopMeta_.DiscardUnknown(m)
}

var xxx_messageInfo_SettleShopMeta_ proto.InternalMessageInfo

func (m *SettleShopMeta_) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SettleShopMeta_) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func (m *SettleShopMeta_) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

type SettleDeliverMeta_ struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ConsigneeName        string   `protobuf:"bytes,2,opt,name=consignee_name,json=consigneeName,proto3" json:"consignee_name"`
	ConsigneePhone       string   `protobuf:"bytes,3,opt,name=consignee_phone,json=consigneePhone,proto3" json:"consignee_phone"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleDeliverMeta_) Reset()         { *m = SettleDeliverMeta_{} }
func (m *SettleDeliverMeta_) String() string { return proto.CompactTextString(m) }
func (*SettleDeliverMeta_) ProtoMessage()    {}
func (*SettleDeliverMeta_) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_0317a332241fa71d, []int{6}
}
func (m *SettleDeliverMeta_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleDeliverMeta_.Unmarshal(m, b)
}
func (m *SettleDeliverMeta_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleDeliverMeta_.Marshal(b, m, deterministic)
}
func (dst *SettleDeliverMeta_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleDeliverMeta_.Merge(dst, src)
}
func (m *SettleDeliverMeta_) XXX_Size() int {
	return xxx_messageInfo_SettleDeliverMeta_.Size(m)
}
func (m *SettleDeliverMeta_) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleDeliverMeta_.DiscardUnknown(m)
}

var xxx_messageInfo_SettleDeliverMeta_ proto.InternalMessageInfo

func (m *SettleDeliverMeta_) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SettleDeliverMeta_) GetConsigneeName() string {
	if m != nil {
		return m.ConsigneeName
	}
	return ""
}

func (m *SettleDeliverMeta_) GetConsigneePhone() string {
	if m != nil {
		return m.ConsigneePhone
	}
	return ""
}

func (m *SettleDeliverMeta_) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*SettlePersistRequest)(nil), "SettlePersistRequest")
	proto.RegisterType((*CartItemRequest)(nil), "CartItemRequest")
	proto.RegisterType((*CheckSignRequest)(nil), "CheckSignRequest")
	proto.RegisterType((*CartCode)(nil), "CartCode")
	proto.RegisterType((*SettleMeta_)(nil), "SettleMeta_")
	proto.RegisterType((*SettleShopMeta_)(nil), "SettleShopMeta_")
	proto.RegisterType((*SettleDeliverMeta_)(nil), "SettleDeliverMeta_")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	// 批发购物车接口
	WholesaleCartV1(ctx context.Context, in *WsCartRequest, opts ...grpc.CallOption) (*Result, error)
	// 普通购物车接口
	NormalCartV1(ctx context.Context, in *NormalCartRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取购物车,当购物车编号不存在时,将返回一个新的购物车
	GetShoppingCart_(ctx context.Context, in *CartCode, opts ...grpc.CallOption) (*SShoppingCart, error)
	// 放入购物车
	PutInCart_(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// 从购物车里删除指定数量商品
	SubCartItem_(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// 获取购物车结算信息
	GetCartSettle_(ctx context.Context, in *CartCode, opts ...grpc.CallOption) (*SettleMeta_, error)
	// 勾选商品结算
	CheckSign_(ctx context.Context, in *CheckSignRequest, opts ...grpc.CallOption) (*Result, error)
	// 更新购物车结算
	PrepareSettlePersist_(ctx context.Context, in *SettlePersistRequest, opts ...grpc.CallOption) (*Result, error)
}

type cartServiceClient struct {
	cc *grpc.ClientConn
}

func NewCartServiceClient(cc *grpc.ClientConn) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) WholesaleCartV1(ctx context.Context, in *WsCartRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/WholesaleCartV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) NormalCartV1(ctx context.Context, in *NormalCartRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/NormalCartV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetShoppingCart_(ctx context.Context, in *CartCode, opts ...grpc.CallOption) (*SShoppingCart, error) {
	out := new(SShoppingCart)
	err := c.cc.Invoke(ctx, "/CartService/GetShoppingCart_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) PutInCart_(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/PutInCart_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) SubCartItem_(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/SubCartItem_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartSettle_(ctx context.Context, in *CartCode, opts ...grpc.CallOption) (*SettleMeta_, error) {
	out := new(SettleMeta_)
	err := c.cc.Invoke(ctx, "/CartService/GetCartSettle_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CheckSign_(ctx context.Context, in *CheckSignRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/CheckSign_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) PrepareSettlePersist_(ctx context.Context, in *SettlePersistRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/PrepareSettlePersist_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	// 批发购物车接口
	WholesaleCartV1(context.Context, *WsCartRequest) (*Result, error)
	// 普通购物车接口
	NormalCartV1(context.Context, *NormalCartRequest) (*Result, error)
	// 获取购物车,当购物车编号不存在时,将返回一个新的购物车
	GetShoppingCart_(context.Context, *CartCode) (*SShoppingCart, error)
	// 放入购物车
	PutInCart_(context.Context, *CartItemRequest) (*CartItemResponse, error)
	// 从购物车里删除指定数量商品
	SubCartItem_(context.Context, *CartItemRequest) (*CartItemResponse, error)
	// 获取购物车结算信息
	GetCartSettle_(context.Context, *CartCode) (*SettleMeta_, error)
	// 勾选商品结算
	CheckSign_(context.Context, *CheckSignRequest) (*Result, error)
	// 更新购物车结算
	PrepareSettlePersist_(context.Context, *SettlePersistRequest) (*Result, error)
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_WholesaleCartV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WsCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).WholesaleCartV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/WholesaleCartV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).WholesaleCartV1(ctx, req.(*WsCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_NormalCartV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NormalCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).NormalCartV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/NormalCartV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).NormalCartV1(ctx, req.(*NormalCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetShoppingCart__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetShoppingCart_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/GetShoppingCart_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetShoppingCart_(ctx, req.(*CartCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_PutInCart__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).PutInCart_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/PutInCart_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).PutInCart_(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_SubCartItem__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).SubCartItem_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/SubCartItem_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).SubCartItem_(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartSettle__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartSettle_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/GetCartSettle_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartSettle_(ctx, req.(*CartCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CheckSign__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CheckSign_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/CheckSign_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CheckSign_(ctx, req.(*CheckSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_PrepareSettlePersist__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettlePersistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).PrepareSettlePersist_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/PrepareSettlePersist_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).PrepareSettlePersist_(ctx, req.(*SettlePersistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WholesaleCartV1",
			Handler:    _CartService_WholesaleCartV1_Handler,
		},
		{
			MethodName: "NormalCartV1",
			Handler:    _CartService_NormalCartV1_Handler,
		},
		{
			MethodName: "GetShoppingCart_",
			Handler:    _CartService_GetShoppingCart__Handler,
		},
		{
			MethodName: "PutInCart_",
			Handler:    _CartService_PutInCart__Handler,
		},
		{
			MethodName: "SubCartItem_",
			Handler:    _CartService_SubCartItem__Handler,
		},
		{
			MethodName: "GetCartSettle_",
			Handler:    _CartService_GetCartSettle__Handler,
		},
		{
			MethodName: "CheckSign_",
			Handler:    _CartService_CheckSign__Handler,
		},
		{
			MethodName: "PrepareSettlePersist_",
			Handler:    _CartService_PrepareSettlePersist__Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_service.proto",
}

func init() { proto.RegisterFile("cart_service.proto", fileDescriptor_cart_service_0317a332241fa71d) }

var fileDescriptor_cart_service_0317a332241fa71d = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x6f, 0xd3, 0x4a,
	0x10, 0x8e, 0x93, 0xe6, 0x87, 0x27, 0x79, 0x49, 0xba, 0x7d, 0x7d, 0x2f, 0xa4, 0xa0, 0x56, 0x16,
	0x15, 0x15, 0xd0, 0xad, 0x48, 0xc5, 0x01, 0x71, 0x6b, 0x90, 0x2a, 0x1f, 0x28, 0x91, 0x2d, 0x51,
	0x89, 0x8b, 0xe5, 0xd8, 0x23, 0xc7, 0xaa, 0xed, 0x75, 0xbd, 0xeb, 0x4a, 0xfd, 0x27, 0xb8, 0x71,
	0xe2, 0xca, 0x8d, 0x7f, 0x12, 0xed, 0xda, 0x49, 0xdc, 0xb4, 0x12, 0x15, 0x9c, 0xec, 0xf9, 0xe6,
	0x9b, 0x9d, 0x6f, 0x7e, 0xec, 0x02, 0xf1, 0xdc, 0x4c, 0x38, 0x1c, 0xb3, 0x9b, 0xd0, 0x43, 0x9a,
	0x66, 0x4c, 0xb0, 0x71, 0x2f, 0x88, 0xd8, 0xdc, 0x8d, 0x4a, 0xeb, 0xbf, 0x18, 0x39, 0x77, 0x03,
	0x3c, 0x51, 0x4c, 0x5f, 0xb0, 0x02, 0x37, 0x7e, 0x6a, 0xf0, 0xaf, 0x8d, 0x42, 0x44, 0x38, 0xc3,
	0x8c, 0x87, 0x5c, 0x58, 0x78, 0x9d, 0x23, 0x17, 0xe4, 0x09, 0x74, 0xe6, 0xf9, 0x2d, 0x66, 0x4e,
	0xe8, 0x8f, 0xb4, 0x03, 0xed, 0xa8, 0x61, 0xb5, 0x95, 0x6d, 0xfa, 0xe4, 0x7f, 0x68, 0xf3, 0x05,
	0x4b, 0xa5, 0xa7, 0xae, 0x3c, 0x2d, 0x69, 0x9a, 0x3e, 0xd9, 0x87, 0x6e, 0xea, 0xde, 0xc6, 0x98,
	0x08, 0x87, 0xa5, 0x62, 0xd4, 0x50, 0x4e, 0x28, 0xa1, 0x4f, 0xa9, 0x90, 0x04, 0x1f, 0xa3, 0xf0,
	0x06, 0x33, 0x45, 0xd8, 0x2a, 0x08, 0x25, 0x24, 0x09, 0xcf, 0x00, 0x5c, 0xdf, 0xcf, 0x90, 0x73,
	0x79, 0x7a, 0x53, 0xf9, 0xf5, 0x12, 0x31, 0x7d, 0xe3, 0x9b, 0x06, 0x83, 0xa9, 0x9b, 0x09, 0x53,
	0x60, 0xfc, 0x08, 0xa1, 0x7b, 0xa0, 0xab, 0x72, 0x3d, 0xe6, 0xa3, 0x92, 0xaa, 0x5b, 0x1d, 0x09,
	0x4c, 0x99, 0x8f, 0xb2, 0x8a, 0x50, 0x60, 0x2c, 0xc3, 0xa4, 0x50, 0x62, 0xb5, 0xa4, 0x69, 0xfa,
	0x64, 0x17, 0x5a, 0xfc, 0x2a, 0x97, 0xf8, 0x96, 0xc2, 0x9b, 0xfc, 0x2a, 0x37, 0x7d, 0x32, 0x86,
	0xce, 0x75, 0xee, 0x26, 0x22, 0x14, 0xb7, 0x4a, 0x58, 0xd3, 0x5a, 0xd9, 0xc6, 0x35, 0x0c, 0xa7,
	0x0b, 0xf4, 0xae, 0xec, 0x30, 0x48, 0xfe, 0x56, 0xd7, 0x21, 0x34, 0xa5, 0x10, 0x3e, 0x6a, 0x1c,
	0x34, 0x8e, 0xba, 0x93, 0x01, 0xb5, 0xd5, 0xd1, 0xab, 0xb2, 0x0b, 0xaf, 0x71, 0x06, 0x9d, 0xe9,
	0x32, 0xe4, 0x0f, 0x53, 0x19, 0x3f, 0x34, 0xe8, 0x16, 0xc3, 0xff, 0x88, 0xc2, 0x75, 0x36, 0xe7,
	0xa7, 0xfd, 0x6e, 0x7e, 0xf5, 0x7b, 0xf3, 0x7b, 0x0e, 0x5b, 0x72, 0x17, 0x54, 0x47, 0xbb, 0x93,
	0x21, 0x2d, 0x4e, 0xb7, 0x17, 0x2c, 0x55, 0x19, 0x2c, 0xe5, 0x25, 0xc7, 0xd0, 0x2e, 0x63, 0x54,
	0x8b, 0xbb, 0x93, 0x9d, 0x92, 0xf8, 0xa1, 0x40, 0x0b, 0xee, 0x92, 0x63, 0x20, 0x0c, 0x36, 0xce,
	0xa9, 0xae, 0xa0, 0x76, 0x67, 0x05, 0xf7, 0x40, 0x57, 0x8e, 0xc4, 0x8d, 0x57, 0xf5, 0x4a, 0xe0,
	0xc2, 0x8d, 0x91, 0x3c, 0x05, 0x5d, 0x60, 0x84, 0xe9, 0x82, 0x25, 0xa8, 0x24, 0xea, 0xd6, 0x1a,
	0x30, 0xbe, 0x6a, 0x40, 0xee, 0xcb, 0x20, 0x7d, 0xa8, 0xaf, 0xb2, 0xd4, 0x43, 0x9f, 0x1c, 0x42,
	0xdf, 0x63, 0x09, 0x0f, 0x83, 0x04, 0xb1, 0x9a, 0xe6, 0x9f, 0x15, 0xaa, 0x72, 0xbd, 0x80, 0xc1,
	0x9a, 0x56, 0xcd, 0xb8, 0x8e, 0x9e, 0x49, 0x94, 0x8c, 0xa0, 0x5d, 0x2e, 0xb8, 0x6a, 0x86, 0x6e,
	0x2d, 0xcd, 0xc9, 0xf7, 0x06, 0x74, 0xe5, 0x8c, 0xed, 0xe2, 0x5e, 0x93, 0xd7, 0x30, 0xb8, 0x5c,
	0xb0, 0x08, 0xb9, 0x1b, 0xa1, 0xc4, 0x3f, 0xbf, 0x21, 0x7d, 0x7a, 0xc9, 0xe5, 0x6f, 0xb9, 0x74,
	0xe3, 0x36, 0xb5, 0x90, 0xe7, 0x91, 0x30, 0x6a, 0xe4, 0x18, 0x7a, 0x17, 0x2c, 0x8b, 0xdd, 0xa8,
	0xa4, 0x12, 0xba, 0x36, 0x1f, 0xa0, 0x9f, 0xc0, 0xf0, 0x1c, 0x85, 0xec, 0x70, 0x1a, 0x26, 0x81,
	0x24, 0x39, 0x44, 0xa7, 0xcb, 0x15, 0x1b, 0xf7, 0xa9, 0x5d, 0xf5, 0x19, 0x35, 0x72, 0x0a, 0x30,
	0xcb, 0x85, 0x99, 0x14, 0xd4, 0x21, 0xdd, 0xb8, 0x97, 0xe3, 0xed, 0x0a, 0xc2, 0x53, 0x96, 0x70,
	0x34, 0x6a, 0xe4, 0x2d, 0xf4, 0xec, 0x7c, 0xbe, 0x74, 0x3c, 0x3a, 0xec, 0x15, 0xf4, 0xcf, 0x51,
	0x14, 0xbd, 0x90, 0x03, 0xba, 0x23, 0xad, 0x47, 0x2b, 0x3b, 0x6c, 0xd4, 0xc8, 0x4b, 0x80, 0xd5,
	0x65, 0x74, 0xc8, 0x36, 0xdd, 0xbc, 0x99, 0xd5, 0xaa, 0xdf, 0xc1, 0xee, 0x2c, 0xc3, 0xd4, 0xcd,
	0xf0, 0xce, 0x23, 0xe8, 0x90, 0x5d, 0xfa, 0xd0, 0xab, 0x58, 0x09, 0x3d, 0xdb, 0x87, 0x1d, 0x8f,
	0xc5, 0x34, 0x08, 0xc5, 0x22, 0x9f, 0xd3, 0x80, 0x4d, 0x18, 0xcd, 0x52, 0xef, 0x4b, 0x87, 0x9e,
	0xbc, 0x57, 0x4f, 0xeb, 0xbc, 0xa5, 0x3e, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x44, 0xcf,
	0x13, 0x8a, 0x9d, 0x05, 0x00, 0x00,
}
