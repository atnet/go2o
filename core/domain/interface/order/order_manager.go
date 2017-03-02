/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-05 17:19
 * description :
 * history :
 */

package order

import (
	"go2o/core/domain/interface/cart"
	"go2o/core/domain/interface/member"
	"go2o/core/domain/interface/merchant/shop"
	"go2o/core/domain/interface/payment"
	"go2o/core/dto"
)

type (
	//　购物聚合根
	IOrderManager interface {
		// 创建订单
		CreateOrder(*Order) IOrder
		// 生成空白订单,并保存返回对象
		CreateSubOrder(*ValueSubOrder) ISubOrder

		// 将购物车转换为订单
		ParseToOrder(c cart.ICart) (IOrder, member.IMember, error)

		// 预生成订单及支付单
		PrepareOrder(c cart.ICart, addressId int32, subject string, couponCode string) (IOrder,
			payment.IPaymentOrder, error)

		// 提交订单
		SubmitOrder(c cart.ICart, addressId int32, subject string, couponCode string,
			balanceDiscount bool) (IOrder, payment.IPaymentOrder, error)

		// 获取可用的订单号, 系统直营传入vendor为0
		GetFreeOrderNo(vendor int32) string

		// 根据订单编号获取订单
		GetOrderById(orderId int64) IOrder

		// 根据订单号获取订单
		GetOrderByNo(orderNo string) IOrder

		// 接收在线交易支付的通知，不主动调用
		ReceiveNotifyOfOnlineTrade(orderId int64) error

		// 智能选择门店
		SmartChoiceShop(address string) (shop.IShop, error)

		// 根据父订单编号获取购买的商品项
		GetItemsByParentOrderId(orderId int64) []*OrderItem

		// 获取子订单
		GetSubOrder(id int64) ISubOrder
	}

	IOrderRepo interface {
		// 获取订单服务
		Manager() IOrderManager

		// 保存订单,返回订单编号
		SaveNormalOrder(v *NormalOrder) (int, error)

		// 保存订单优惠券绑定
		SaveOrderCouponBind(*OrderCoupon) error

		// 获取订单的促销绑定
		GetOrderPromotionBinds(orderNo string) []*OrderPromotionBind

		// 保存订单的促销绑定
		SavePromotionBindForOrder(*OrderPromotionBind) (int32, error)

		// 获取可用的订单号, 系统直营传入vendor为0
		GetFreeOrderNo(vendorId int32) string

		// 根据编号获取订单
		GetOrderById(id int64) *NormalOrder

		// 根据订单号获取订单
		GetValueOrderByNo(orderNo string) *NormalOrder

		// 获取等待处理的订单
		GetWaitingSetupOrders(vendorId int32) ([]*NormalOrder, error)

		// 保存订单日志
		SaveSubOrderLog(*OrderLog) error

		// 获取订单的所有子订单
		GetSubOrdersByParentId(orderId int64) []*ValueSubOrder

		// 获取订单编号
		GetOrderId(orderNo string, subOrder bool) int64

		// 获取子订单
		GetSubOrder(id int64) *ValueSubOrder

		// 根据订单号获取子订单
		GetSubOrderByNo(orderNo string) *ValueSubOrder

		// 保存子订单
		SaveSubOrder(value *ValueSubOrder) (int64, error)

		// 保存子订单的商品项,并返回编号和错误
		SaveOrderItem(subOrderId int64, value *OrderItem) (int32, error)

		// 获取订单项
		GetSubOrderItems(orderId int64) []*OrderItem

		// 根据父订单编号获取购买的商品项
		GetItemsByParentOrderId(orderId int64) []*OrderItem

		// 获取订单的操作记录
		GetSubOrderLogs(orderId int64) []*OrderLog

		// 根据商品快照获取订单项
		GetOrderItemBySnapshotId(orderId int32, snapshotId int32) *OrderItem

		// 根据商品快照获取订单项数据传输对象
		GetOrderItemDtoBySnapshotId(orderId int32, snapshotId int32) *dto.OrderItem

		// Get OrderList
		GetOrder(where string, arg ...interface{}) *Order
		// Save OrderList
		SaveOrderList(v *Order) (int, error)
		// Delete OrderList
		//DeleteOrderList(primary interface{}) error
	}
)
