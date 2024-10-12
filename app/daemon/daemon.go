/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2014-01-08 21:01
 * description :
 * history :
 */

package daemon

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/ixre/go2o/app/daemon/job"
	mss "github.com/ixre/go2o/core/domain/interface/message"
	"github.com/ixre/go2o/core/domain/interface/order"
	"github.com/ixre/go2o/core/infrastructure/locker"
	"github.com/ixre/go2o/core/initial"
	"github.com/ixre/go2o/core/repos/clickhouse"
	"github.com/ixre/go2o/core/service"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/go2o/core/variable"
	"github.com/ixre/gof"
	"github.com/ixre/gof/db"
	"github.com/ixre/gof/db/orm"
	"github.com/ixre/gof/types"
	"github.com/robfig/cron/v3"
)

// Func 守护进程执行的函数
type Func func(gof.App)

// Service 守护进程服务
type Service interface {
	// Name 服务名称
	Name() string
	// Start 启动服务,并传入APP上下文对象
	Start(gof.App)
	// OrderObs 处理订单,需根据订单不同的状态,作不同的业务,返回布尔值,如果返回false,则不继续执行
	OrderObs(singleOrder *proto.SSingleOrder) bool
	// MemberObs 监视会员修改,@create:是否为新注册会员,返回布尔值,如果返回false,则不继续执行
	MemberObs(m *proto.SMember, create bool) bool
	// PaymentOrderObs 通知支付单完成队列,返回布尔值,如果返回false,则不继续执行
	PaymentOrderObs(order *proto.SPaymentOrder) bool
	// HandleMailQueue 处理邮件队列,返回布尔值,如果返回false,则不继续执行
	HandleMailQueue([]*mss.MailTask) bool
}

var (
	appCtx           *initial.AppImpl
	conn             db.Connector
	_orm             orm.Orm
	services         []Service
	serviceNames     = make(map[string]int)
	tickerDuration   = 20 * time.Second // 间隔20秒执行
	tickerInvokeFunc []Func
	cronTab          = cron.New()
	ticker           = time.NewTicker(tickerDuration)
	mux              sync.Mutex
	ch               = make(chan int, 1)
)

// RegisterService 注册服务
func RegisterService(s Service) {
	mux.Lock()
	defer mux.Unlock()
	if s == nil {
		panic("service is nil")
	}
	name := s.Name()
	if _, ok := serviceNames[name]; ok {
		panic("service named " + name + " is registed!")
	}
	serviceNames[name] = len(services)
	services = append(services, s)
}

// AddTickerFunc 添加定时执行任务(默认5秒)
func AddTickerFunc(f Func) {
	tickerInvokeFunc = append(tickerInvokeFunc, f)
}

// Start 启动守护进程
func Start() {
	defer func() {
		cronTab.Stop()
		ticker.Stop()
	}()
	//运行自定义服务
	for i, s := range services {
		log.Println("** [ GO2O][ Daemon] - (", i, ")", s.Name(), "daemon running")
		go s.Start(appCtx)
	}
	startCronTab() // 运行计划任务
	<-ch
	//startTicker()  // 阻塞
}

func startTicker() {
	// 执行定时任务
	for {
		select {
		case <-ticker.C:
			for _, f := range tickerInvokeFunc {
				go f(appCtx)
			}
		}
	}
}

// 判断是否处理
func isHandled(key string, unix int64) bool {
	conn := initial.GetRedisConn()
	defer conn.Close()
	unix2, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		return false
	}
	return unix == int64(unix2)
}

// 标记最后处理时间
func signHandled(key string, unix int64) {
	conn := initial.GetRedisConn()
	defer conn.Close()
	conn.Do("SET", key, unix)
}

// CompareLastUnix 比较最后运行的时间戳
func CompareLastUnix(key string, unix int64) bool {
	return isHandled(key, unix)
}

// SetLastUnix 设置最后运行的时间戳
func SetLastUnix(key string, unix int64) {
	signHandled(key, unix)
}

// 添加定时任务
func AddCron(spec string, cmd func()) {
	mux.Lock()
	defer mux.Unlock()
	cronTab.AddFunc(spec, cmd)
}

// 运行定时任务
func startCronTab() {

	// 删除分布式锁,会导致重启一直不执行任务
	locker.Unlock("/SyncWalletLogToClickHouse")
	locker.Unlock("/CheckExpiresPaymentOrder")
	locker.Unlock("/SubmitDivideJob")
	// 注册任务
	startClickhouseJob(cronTab)
	// 注册定时任务
	for _, j := range job.GetJobs() {
		cronTab.AddFunc(j.Spec, j.Cmd)
	}
	//商户每日报表
	//cronTab.AddFunc("0 0 0 * * *", mchDayChart)
	//个人金融结算,每天00:20更新数据
	//cronTab.AddFunc("0 20 0 * * *", personFinanceSettle)
	//订单自动收货,2分钟检测一次
	//cronTab.AddFunc("0 */2 * * * *", orderAutoReceive)
	// 自动解锁会员
	//cronTab.AddFunc("0 * * * * *", memberAutoUnlock)
	cronTab.Start()
}

func startClickhouseJob(tab *cron.Cron) {
	clickhouseSupport := clickhouse.GetClickhouseConn() != nil
	if !clickhouseSupport {
		return
	}
	tab.AddFunc("@every 2s", job.SyncWalletLogToClickHouse)
}

type defaultService struct {
	app     gof.App
	sOrder  bool
	sMember bool
	sMail   bool
}

// 注册系统服务
func (d *defaultService) init() {
	if len(services) == 0 {
		RegisterService(d)
	} else {
		services = append([]Service{d}, services...)
	}
}

// 服务名称
func (d *defaultService) Name() string {
	return "sys"
}

// 启动服务
func (d *defaultService) Start(a gof.App) {
	d.app = a
	go superviseMemberUpdate(services)
	go superviseOrder(services)
	go supervisePaymentOrderFinish(services)
	go startMailQueue(services)
	go personFinanceSettle() //启动时结算
	go mchDayChart()         //商户每日报表

	//go func() {
	//    time.Sleep(time.Second * 6)
	//    o, _ := impl.ShoppingService.GetSubOrderByNo(context.TODO(),"100000021289")
	//    d.OrderObs(o)
	//    detectOrderExpires()
	//   orderAutoRecive()
	//}()
}

// 处理订单,需根据订单不同的状态,作不同的业务
// 返回布尔值,如果返回false,则不继续执行
func (d *defaultService) OrderObs(o *proto.SSingleOrder) bool {
	if d.app.Debug() {
		d.app.Log().Println("-- 订单", o.OrderNo, "状态:", o.Status)
	}
	if d.sOrder {
		conn := initial.GetRedisConn()
		defer conn.Close()
		defer Recover()

		switch o.Status {
		//订单未支付，则超时自动取消
		case order.StatAwaitingPayment:
			d.updateOrderExpires(conn, o)
			//自动确认订单
		case order.StatAwaitingConfirm:
			d.orderAutoConfirm(conn, o)
			//订单自动收货
		case order.StatShipped:
			d.orderAutoReceive(conn, o)
			//订单已经收货
		case order.StatCompleted:
			d.orderReceived(conn, o)
		}
	}
	return true
}

// 监视会员修改,@create:是否为新注册会员
// 返回布尔值,如果返回false,则不继续执行
func (d *defaultService) MemberObs(m *proto.SMember, create bool) bool {
	defer Recover()
	if d.sMember {
		//todo: 执行会员逻辑
	}
	return true
}

// 通知支付单完成队列,返回布尔值,如果返回false,则不继续执行
func (d *defaultService) PaymentOrderObs(order *proto.SPaymentOrder) bool {
	if order == nil {
		return false
	}
	if d.app.Debug() {
		d.app.Log().Println("---支付单", order.TradeNo, "支付完成")
	}
	return true
}

// 测试是否为子订单,并返回编号
func (d *defaultService) testSubId(o *proto.SSingleOrder) (string, bool) {
	// if o.ParentOrderId <= 0 {
	// 	return o.OrderNo, true
	// }
	return o.OrderNo, false
}

// 批量删除REDIS KEY
func (d *defaultService) batchDelKeys(conn redis.Conn, key string) {
	list, err := redis.Strings(conn.Do("KEYS", key))
	if err == nil {
		for _, oKey := range list {
			conn.Do("DEL", oKey)
		}
	}
}

// 设置订单过期时间
func (d *defaultService) updateOrderExpires(conn redis.Conn, o *proto.SSingleOrder) {
	// //订单刚创建时,设置过期时间
	// if o.Status == order.StatAwaitingPayment {
	// 	trans, cli, _ := service.SystemServiceClient()
	// 	defer trans.Close()
	// 	ss, _ := cli.GetGlobMchSaleConf_(context.TODO(), &proto.Empty{})
	// 	unix := o.UpdateTime + int64(ss.OrderTimeOutMinute)*60
	// 	t := time.Unix(unix, 0)
	// 	tk := getTick(t)
	// 	orderNo, sub := d.testSubId(o)
	// 	prefix := types.StringCond(sub, "sub!", "")
	// 	key := fmt.Sprintf("%s:%s%s:%s", variable.KvOrderExpiresTime, prefix, orderNo, tk)
	// 	//log.Println(" [Daemon][Exprire][ Key]:", key)
	// 	conn.Do("SET", key, unix)
	// }
}

// 取消订单过期时间
func (d *defaultService) cancelOrderExpires(conn redis.Conn, o *proto.SSingleOrder) {
	orderNo, sub := d.testSubId(o)
	prefix := types.StringCond(sub, "sub!", "")
	key := fmt.Sprintf("%s:%s%s:*", variable.KvOrderExpiresTime, prefix, orderNo)
	d.batchDelKeys(conn, key)
}

// 确认订单
func (d *defaultService) orderAutoConfirm(conn redis.Conn, o *proto.SSingleOrder) {
	trans, cli, _ := service.OrderServiceClient()
	defer trans.Close()
	orderNo, sub := d.testSubId(o)
	ret, _ := cli.ConfirmOrder(context.TODO(), &proto.OrderNo{
		OrderNo: orderNo,
		Sub:     sub,
	})
	if ret.ErrCode > 0 {
		log.Println(fmt.Sprintf("[ GO2O][ ERROR]: confirm order failed, %s", ret.ErrMsg))
	}
	d.cancelOrderExpires(conn, o) //付款后取消自动取消
}

// 订单自动收货
func (d *defaultService) orderAutoReceive(conn redis.Conn, o *proto.SSingleOrder) {
	// if o.Status == order.StatShipped {
	// 	trans, cli, _ := service.SystemServiceClient()
	// 	defer trans.Close()
	// 	ss, _ := cli.GetGlobMchSaleConf_(context.TODO(), &proto.Empty{})
	// 	unix := o.UpdateTime + int64(ss.OrderTimeOutReceiveHour)*60*60
	// 	t := time.Unix(unix, 0)
	// 	tk := getTick(t)
	// 	orderNo, sub := d.testSubId(o)
	// 	prefix := types.StringCond(sub, "sub!", "")
	// 	key := fmt.Sprintf("%s:%s%s:%s", variable.KvOrderAutoReceive, prefix, orderNo, tk)
	// 	//log.Println(" [Daemon][AutoReceive][ Key]:", key)
	// 	conn.Do("SET", key, unix)
	// }
}

// 完成订单自动收货
func (d *defaultService) orderReceived(conn redis.Conn, o *proto.SSingleOrder) {
	if o.Status == order.StatCompleted {
		orderNo, sub := d.testSubId(o)
		prefix := types.StringCond(sub, "sub!", "")
		key := fmt.Sprintf("%s:%s%s:*", variable.KvOrderAutoReceive, prefix, orderNo)
		d.batchDelKeys(conn, key)
	}
}

// 处理邮件队列
// 返回布尔值,如果返回false,则不继续执行
func (d *defaultService) HandleMailQueue(list []*mss.MailTask) bool {
	defer Recover()
	if !d.sMail {
		handleMailQueue(list)
	}
	return true
}

// 运行
func Run(ctx gof.App) {
	if ctx != nil {
		appCtx = ctx.(*initial.AppImpl)
	} else {
		appCtx = initial.NewApp("app.conf", nil)
	}
	conn = appCtx.Db()
	//sMail := true // appCtx.Config().GetString(variable.SystemMailQueueOff) != "1" //是否关闭系统邮件队列
	//sMail := cnf.GetString(variable.)

	//s := &defaultService{
	//	sMember: true,
	//	sOrder:  true,
	//	sMail:   sMail,
	//}
	//s.init()
	Start()
}

// 自定义参数运行
func FlagRun() {
	var conf string
	var debug bool
	var trace bool
	var service string
	var serviceArr = []string{"mail", "order"}
	var ch = make(chan bool)
	flag.StringVar(&conf, "conf", "app.conf", "")
	flag.BoolVar(&debug, "debug", true, "")
	flag.BoolVar(&trace, "trace", true, "")
	flag.StringVar(&service, "service", strings.Join(serviceArr, ","), "")

	flag.Parse()

	appCtx = initial.NewApp(conf, nil)
	initial.Init(appCtx, debug, trace)

	conn = appCtx.Db()

	//todo: daemon 应不依赖于service
	//impl.Init(appCtx, app.FlagDaemon)

	//todo:???
	//	if service != "all" {
	//		serviceArr = strings.Split(service, ",")
	//	}
	// RegisterByName(serviceArr)

	//s := &defaultService{
	//	sMember: true,
	//	sOrder:  true,
	//	sMail:   true,
	//}
	//s.init()
	Start()

	<-ch
}
