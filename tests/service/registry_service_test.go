package service

import (
	"context"
	"testing"

	"github.com/ixre/go2o/core/inject"
	"github.com/ixre/go2o/core/service"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/typeconv"
)

func TestRegistryUpateValueMap(t *testing.T) {
	mp := map[string]string{
		"order_push_sub_order_enabled":    "0",
		"member_withdrawal_push_enabled1": "1",
	}
	ret, _ := inject.GetRegistryService().UpdateValues(context.TODO(), &proto.StringMap{
		Value: mp,
	})
	if ret.Code > 0 {
		t.Error(ret.Message)
	}
	t.Log(typeconv.MustJson(ret))
}

func TestRegistryClientUpateValueMap(t *testing.T) {
	mp := map[string]string{
		"order_push_sub_order_enabled":    "0",
		"member_withdrawal_push_enabled1": "1",
	}
	service.ConfigureClient(nil, "go2o.dev:1427")
	trans, cli, _ := service.RegistryServiceClient()
	defer trans.Close()
	ret, err := cli.UpdateValues(context.TODO(), &proto.StringMap{
		Value: mp,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if ret.Code > 0 {
		t.Error(ret.Message)
	}
	t.Log(typeconv.MustJson(ret))
}
