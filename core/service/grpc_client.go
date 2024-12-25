/**
 * Copyright 2015 @ at3.net.
 * name : restful_client.go
 * author : jarryliu
 * date : 2016-11-13 12:32
 * description :
 * history :
 */
package service

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/ixre/go2o/core/etcd"
	"github.com/ixre/go2o/core/service/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var staticAddr string
var selector etcd.Selector

// ConfigureClient 设置RPC地址,defaultAddr为默认地址,当未指定clientv3.Config时使用
func ConfigureClient(c *clientv3.Config, defaultAddr string) {
	staticAddr = defaultAddr
	if c != nil && len(staticAddr) == 0 {
		log.Println("[ GO2O][ INFO]: connecting go2o rpc server...")
		s, err := etcd.NewSelector(service, *c, etcd.AlgRoundRobin)
		if err != nil {
			log.Println("[ GO2O][ ERROR]: can't connect go2o rpc server! ", err.Error())
			os.Exit(1)
		}
		selector = s
	} else if len(staticAddr) > 0 {
		log.Printf("[ GO2O][ INFO]: connecting static rpc server (node:%s)...\n", staticAddr)
	}
	tryConnect(30)
}

// 尝试连接服务,如果连接不成功,则退出
func tryConnect(retryTimes int) {
	for i := 0; i < retryTimes; i++ {
		trans, _, err := StatusServiceClient()
		if err == nil {
			trans.Close()
			break
		}
		if i > 0 {
			log.Printf("[ GO2O][ INFO]: the %dth retry connecting rpc server \n", i)
		}
		time.Sleep(time.Second)
		if i >= retryTimes-1 {
			log.Println("[ GO2O][ Fatal]: Can not connect go2o rpc server")
			os.Exit(1)
		}
	}
}

// 获取连接
func getConn(selector etcd.Selector) (*grpc.ClientConn, error) {
	addr := staticAddr
	if len(addr) == 0 && selector != nil {
		next, err := selector.Next()
		if err != nil {
			return nil, err
		}
		addr = next.Addr
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             5 * time.Second,
			PermitWithoutStream: true}))
	//defer cancel()
	if err != nil {
		log.Printf("[ GO2O][ ERROR]: %s addr:%s \n", err.Error(), addr)
		return conn, fmt.Errorf("%s addr:%s", err.Error(), addr)
	}
	return conn, nil
}

// StatusServiceClient 状态客户端
func StatusServiceClient() (*grpc.ClientConn, proto.StatusServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewStatusServiceClient(conn), err
	}
	return conn, nil, err
}

// RegistryServiceClient 基础服务
func RegistryServiceClient() (*grpc.ClientConn, proto.RegistryServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewRegistryServiceClient(conn), err
	}
	return conn, nil, err
}

// MerchantServiceClient 商户客户端
func MerchantServiceClient() (*grpc.ClientConn, proto.MerchantServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewMerchantServiceClient(conn), err
	}
	return conn, nil, err
}

// MemberServiceClient 会员客户端
func MemberServiceClient() (*grpc.ClientConn, proto.MemberServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewMemberServiceClient(conn), err
	}
	return conn, nil, err
}

// SystemServiceClient 基础服务
func SystemServiceClient() (*grpc.ClientConn, proto.SystemServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewSystemServiceClient(conn), err
	}
	return conn, nil, err
}

// MessageServiceClient 消息客户端
func MessageServiceClient() (*grpc.ClientConn, proto.MessageServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewMessageServiceClient(conn), err
	}
	return conn, nil, err
}

// ContentServiceClient 消息客户端
func ContentServiceClient() (*grpc.ClientConn, proto.ContentServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewContentServiceClient(conn), err
	}
	return conn, nil, err
}

// PaymentServiceClient 支付服务
func PaymentServiceClient() (*grpc.ClientConn, proto.PaymentServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewPaymentServiceClient(conn), err
	}
	return conn, nil, err
}

// QuickPaymentServiceClient 快捷支付服务
func QuickPaymentServiceClient() (*grpc.ClientConn, proto.QuickPayServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewQuickPayServiceClient(conn), err
	}
	return conn, nil, err
}

// WalletClient 钱包服务
func WalletClient() (*grpc.ClientConn, proto.WalletServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewWalletServiceClient(conn), err
	}
	return conn, nil, err
}

// OrderServiceClient 订单服务
func OrderServiceClient() (*grpc.ClientConn, proto.OrderServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewOrderServiceClient(conn), err
	}
	return conn, nil, err
}

// CartServiceClient 购物车服务
func CartServiceClient() (*grpc.ClientConn, proto.CartServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewCartServiceClient(conn), err
	}
	return conn, nil, err
}

// ExpressServiceClient 快递服务
func ExpressServiceClient() (*grpc.ClientConn, proto.ExpressServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewExpressServiceClient(conn), err
	}
	return conn, nil, err
}

// ShipmentServiceClient 物流服务
func ShipmentServiceClient() (*grpc.ClientConn, proto.ShipmentServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewShipmentServiceClient(conn), err
	}
	return conn, nil, err
}

// ItemServiceClient 商品服务
func ItemServiceClient() (*grpc.ClientConn, proto.ItemServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewItemServiceClient(conn), err
	}
	return conn, nil, err
}

// ProductServiceClient 产品服务
func ProductServiceClient() (*grpc.ClientConn, proto.ProductServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewProductServiceClient(conn), err
	}
	return conn, nil, err
}

// ShopServiceClient 商店服务
func ShopServiceClient() (*grpc.ClientConn, proto.ShopServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewShopServiceClient(conn), err
	}
	return conn, nil, err
}

// FinanceServiceClient 财务服务
func FinanceServiceClient() (*grpc.ClientConn, proto.FinanceServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewFinanceServiceClient(conn), err
	}
	return conn, nil, err
}

// QueryServiceClient 查询服务
func QueryServiceClient() (*grpc.ClientConn, proto.QueryServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewQueryServiceClient(conn), err
	}
	return conn, nil, err
}

// AfterSalesServiceClient 售后服务
func AfterSalesServiceClient() (*grpc.ClientConn, proto.AfterSalesServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewAfterSalesServiceClient(conn), err
	}
	return conn, nil, err
}

// AdvertisementServiceClient 广告服务
func AdvertisementServiceClient() (*grpc.ClientConn, proto.AdvertisementServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewAdvertisementServiceClient(conn), err
	}
	return conn, nil, err
}

// PortalServiceClient 门户服务
func PortalServiceClient() (*grpc.ClientConn, proto.PortalServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewPortalServiceClient(conn), err
	}
	return conn, nil, err
}

// ExecutionServiceClient 任务执行服务
func ExecutionServiceClient() (*grpc.ClientConn, proto.ExecutionServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewExecutionServiceClient(conn), err
	}
	return conn, nil, err
}

// AppServiceClient APP服务
func AppServiceClient() (*grpc.ClientConn, proto.AppServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewAppServiceClient(conn), err
	}
	return conn, nil, err
}

// RbacServiceClient RBAC服务
func RbacServiceClient() (*grpc.ClientConn, proto.RbacServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewRbacServiceClient(conn), err
	}
	return conn, nil, err
}

// CheckServiceClient 校验服务
func CheckServiceClient() (*grpc.ClientConn, proto.CheckServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewCheckServiceClient(conn), err
	}
	return conn, nil, err
}

// InvoiceServiceClient 发票服务
func InvoiceServiceClient() (*grpc.ClientConn, proto.InvoiceServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewInvoiceServiceClient(conn), err
	}
	return conn, nil, err
}

// ChatServiceClient 聊天服务
func ChatServiceClient() (*grpc.ClientConn, proto.ChatServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewChatServiceClient(conn), err
	}
	return conn, nil, err
}

// WorkorderServiceClient 工单服务
func WorkorderServiceClient() (*grpc.ClientConn, proto.WorkorderServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewWorkorderServiceClient(conn), err
	}
	return conn, nil, err
}

// ApprovalServiceClient 审批服务
func ApprovalServiceClient() (*grpc.ClientConn, proto.ApprovalServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewApprovalServiceClient(conn), err
	}
	return conn, nil, err
}

// ProviderServiceClient 服务商服务
func ProviderServiceClient() (*grpc.ClientConn, proto.ServiceProviderServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewServiceProviderServiceClient(conn), err
	}
	return conn, nil, err
}
