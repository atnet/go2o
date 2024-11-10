// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.28.1
// source: cart_service.proto

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

type ECartItemOp int32

const (
	// * 加入购物车
	ECartItemOp_PUT ECartItemOp = 0
	// * 更改数量
	ECartItemOp_CHANGE ECartItemOp = 1
	// * 结账
	ECartItemOp_CHECKOUT ECartItemOp = 2
)

// Enum value maps for ECartItemOp.
var (
	ECartItemOp_name = map[int32]string{
		0: "PUT",
		1: "CHANGE",
		2: "CHECKOUT",
	}
	ECartItemOp_value = map[string]int32{
		"PUT":      0,
		"CHANGE":   1,
		"CHECKOUT": 2,
	}
)

func (x ECartItemOp) Enum() *ECartItemOp {
	p := new(ECartItemOp)
	*p = x
	return p
}

func (x ECartItemOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECartItemOp) Descriptor() protoreflect.EnumDescriptor {
	return file_cart_service_proto_enumTypes[0].Descriptor()
}

func (ECartItemOp) Type() protoreflect.EnumType {
	return &file_cart_service_proto_enumTypes[0]
}

func (x ECartItemOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECartItemOp.Descriptor instead.
func (ECartItemOp) EnumDescriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{0}
}

// 购物车加入商品操作请求
type CartItemOpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 购物车编号
	CartId *ShoppingCartId `protobuf:"bytes,1,opt,name=cartId,proto3" json:"cartId"`
	// 商品项
	Item *RCartItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item"`
	// 操作
	Op ECartItemOp `protobuf:"varint,3,opt,name=op,proto3,enum=ECartItemOp" json:"op"`
}

func (x *CartItemOpRequest) Reset() {
	*x = CartItemOpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemOpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemOpRequest) ProtoMessage() {}

func (x *CartItemOpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemOpRequest.ProtoReflect.Descriptor instead.
func (*CartItemOpRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{0}
}

func (x *CartItemOpRequest) GetCartId() *ShoppingCartId {
	if x != nil {
		return x.CartId
	}
	return nil
}

func (x *CartItemOpRequest) GetItem() *RCartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CartItemOpRequest) GetOp() ECartItemOp {
	if x != nil {
		return x.Op
	}
	return ECartItemOp_PUT
}

// 商品项请求参数
type RCartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品编号
	ItemId int64 `protobuf:"zigzag64,1,opt,name=itemId,proto3" json:"itemId"`
	// SKU编号
	SkuId int64 `protobuf:"zigzag64,2,opt,name=skuId,proto3" json:"skuId"`
	// 数量
	Quantity int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity"`
}

func (x *RCartItem) Reset() {
	*x = RCartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RCartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RCartItem) ProtoMessage() {}

func (x *RCartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RCartItem.ProtoReflect.Descriptor instead.
func (*RCartItem) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{1}
}

func (x *RCartItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *RCartItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *RCartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CheckCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *ShoppingCartId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Items []*SCheckCartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *CheckCartRequest) Reset() {
	*x = CheckCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCartRequest) ProtoMessage() {}

func (x *CheckCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCartRequest.ProtoReflect.Descriptor instead.
func (*CheckCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckCartRequest) GetId() *ShoppingCartId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CheckCartRequest) GetItems() []*SCheckCartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 购物车编号
type ShoppingCartId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 会员/用户编号
	UserId int64 `protobuf:"zigzag64,1,opt,name=userId,proto3" json:"userId"`
	// 购物车标识,当未指定用户时候使用
	CartCode string `protobuf:"bytes,2,opt,name=cartCode,proto3" json:"cartCode"`
	// 是否为批发销售的购物车
	IsWholesale bool `protobuf:"varint,3,opt,name=isWholesale,proto3" json:"isWholesale"`
}

func (x *ShoppingCartId) Reset() {
	*x = ShoppingCartId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingCartId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingCartId) ProtoMessage() {}

func (x *ShoppingCartId) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingCartId.ProtoReflect.Descriptor instead.
func (*ShoppingCartId) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{3}
}

func (x *ShoppingCartId) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ShoppingCartId) GetCartCode() string {
	if x != nil {
		return x.CartCode
	}
	return ""
}

func (x *ShoppingCartId) GetIsWholesale() bool {
	if x != nil {
		return x.IsWholesale
	}
	return false
}

var File_cart_service_proto protoreflect.FileDescriptor

var file_cart_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x11, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64,
	0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x55, 0x0a, 0x09, 0x52, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x6b, 0x75, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x5a, 0x0a,
	0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x66, 0x0a, 0x0e, 0x53, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x73, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c,
	0x65, 0x2a, 0x30, 0x0a, 0x0b, 0x45, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x4f, 0x55,
	0x54, 0x10, 0x02, 0x32, 0xd2, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x56, 0x31, 0x12, 0x0e, 0x2e, 0x57, 0x73, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x53, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x43, 0x61, 0x72, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cart_service_proto_rawDescOnce sync.Once
	file_cart_service_proto_rawDescData = file_cart_service_proto_rawDesc
)

func file_cart_service_proto_rawDescGZIP() []byte {
	file_cart_service_proto_rawDescOnce.Do(func() {
		file_cart_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_service_proto_rawDescData)
	})
	return file_cart_service_proto_rawDescData
}

var file_cart_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cart_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cart_service_proto_goTypes = []interface{}{
	(ECartItemOp)(0),          // 0: ECartItemOp
	(*CartItemOpRequest)(nil), // 1: CartItemOpRequest
	(*RCartItem)(nil),         // 2: RCartItem
	(*CheckCartRequest)(nil),  // 3: CheckCartRequest
	(*ShoppingCartId)(nil),    // 4: ShoppingCartId
	(*SCheckCartItem)(nil),    // 5: SCheckCartItem
	(*WsCartRequest)(nil),     // 6: WsCartRequest
	(*Result)(nil),            // 7: Result
	(*SShoppingCart)(nil),     // 8: SShoppingCart
	(*CartItemResponse)(nil),  // 9: CartItemResponse
}
var file_cart_service_proto_depIdxs = []int32{
	4, // 0: CartItemOpRequest.cartId:type_name -> ShoppingCartId
	2, // 1: CartItemOpRequest.item:type_name -> RCartItem
	0, // 2: CartItemOpRequest.op:type_name -> ECartItemOp
	4, // 3: CheckCartRequest.id:type_name -> ShoppingCartId
	5, // 4: CheckCartRequest.items:type_name -> SCheckCartItem
	6, // 5: CartService.WholesaleCartV1:input_type -> WsCartRequest
	4, // 6: CartService.GetShoppingCart:input_type -> ShoppingCartId
	1, // 7: CartService.ApplyItem:input_type -> CartItemOpRequest
	3, // 8: CartService.CheckCart:input_type -> CheckCartRequest
	7, // 9: CartService.WholesaleCartV1:output_type -> Result
	8, // 10: CartService.GetShoppingCart:output_type -> SShoppingCart
	9, // 11: CartService.ApplyItem:output_type -> CartItemResponse
	7, // 12: CartService.CheckCart:output_type -> Result
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cart_service_proto_init() }
func file_cart_service_proto_init() {
	if File_cart_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_cart_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cart_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemOpRequest); i {
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
		file_cart_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RCartItem); i {
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
		file_cart_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCartRequest); i {
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
		file_cart_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingCartId); i {
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
			RawDescriptor: file_cart_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_service_proto_goTypes,
		DependencyIndexes: file_cart_service_proto_depIdxs,
		EnumInfos:         file_cart_service_proto_enumTypes,
		MessageInfos:      file_cart_service_proto_msgTypes,
	}.Build()
	File_cart_service_proto = out.File
	file_cart_service_proto_rawDesc = nil
	file_cart_service_proto_goTypes = nil
	file_cart_service_proto_depIdxs = nil
}
