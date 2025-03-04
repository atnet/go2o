/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2013-12-03 23:20
 * description :
 * history :
 */

package impl

import (
	"encoding/json"
	"strings"

	"github.com/ixre/go2o/core/repos/clickhouse"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof"
)

// InitializeClickhouse 初始化clickhouse查询连接
func initializeClickhouse(app gof.App) {
	cfg := app.Config()
	server := cfg.GetString("clickhouse_server")
	servers := strings.Split(server, ",")
	database := cfg.GetString("clickhouse_database")
	password := cfg.GetString("clickhouse_password")
	clickhouse.Configure(servers, database, password)
}

// 服务工具类，实现的服务组合此类,可直接调用其方法
type serviceUtil struct{}

// 返回失败的结果
func (s serviceUtil) failResult(msg string) *proto.Result {
	return s.failCodeResult(1, msg)
}

// 返回错误的结果
func (s serviceUtil) error(err error) *proto.Result {
	if err == nil {
		return s.success(nil)
	}
	return s.failResult(err.Error())
}

// 返回错误的结果
func (s serviceUtil) errorV2(err error) *proto.TxResult {
	if err == nil {
		return &proto.TxResult{Data: map[string]string{}}
	}
	return &proto.TxResult{Code: 1, Message: err.Error(), Data: map[string]string{}}
}

// 返回结果
func (s serviceUtil) result(err error) *proto.Result {
	if err == nil {
		return s.success(nil)
	}
	return s.error(err)
}

// 返回自定义编码的结果
func (s serviceUtil) resultWithCodeV2(code int, message string) *proto.TxResult {
	return &proto.TxResult{Code: int32(code), Message: message, Data: map[string]string{}}
}

// 返回自定义编码的结果
func (s serviceUtil) resultWithCode(code int, message string) *proto.Result {
	return &proto.Result{ErrCode: int32(code), ErrMsg: message, Data: map[string]string{}}
}

// 返回失败的结果
func (s serviceUtil) errorCodeResult(code int, err error) *proto.Result {
	return &proto.Result{ErrCode: int32(code), ErrMsg: err.Error(), Data: map[string]string{}}
}

// 返回失败的结果
func (s serviceUtil) failCodeResult(code int, msg string) *proto.Result {
	return &proto.Result{ErrCode: int32(code), ErrMsg: msg, Data: map[string]string{}}
}

// 返回成功的结果
func (s serviceUtil) success(data map[string]string) *proto.Result {
	if data == nil {
		data = map[string]string{}
	}
	return &proto.Result{ErrCode: 0, ErrMsg: "", Data: data}
}

// 返回成功的结果
func (s serviceUtil) txResult(txId int, data map[string]string) *proto.TxResult {
	return &proto.TxResult{
		Code:    0,
		Message: "",
		TxId:    int64(txId),
		Data:    data,
	}
}

// 返回成功的结果
func (s serviceUtil) successV2(data map[string]string) *proto.TxResult {
	if data == nil {
		data = map[string]string{}
	}
	return &proto.TxResult{Code: 0, Message: "", Data: data}
}

// 将int32数组装换为int数组
func (s serviceUtil) intArray(values []int32) []int {
	arr := make([]int, len(values))
	for i, v := range values {
		arr[i] = int(v)
	}
	return arr
}

// 转换为JSON
func (s serviceUtil) json(data interface{}) string {
	if data == nil {
		return "{}"
	}
	r, err := json.Marshal(data)
	if err != nil {
		return "{\"error\":\"parse error:" + err.Error() + "\"}"
	}
	return string(r)
}

// 分页响应结果
func (s serviceUtil) pagingResult(total int, data interface{}) *proto.SPagingResult {
	switch data.(type) {
	case string:
		return &proto.SPagingResult{
			Count:  int32(total),
			Data:   data.(string),
			Extras: map[string]string{},
		}
	}
	r, _ := json.Marshal(data)
	return &proto.SPagingResult{
		ErrCode: 0,
		ErrMsg:  "",
		Count:   int32(total),
		Data:    string(r),
		Extras:  map[string]string{},
	}
}
