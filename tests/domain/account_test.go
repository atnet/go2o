package domain

import (
	"testing"

	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/domain/interface/registry"
	"github.com/ixre/go2o/core/domain/interface/wallet"
	"github.com/ixre/go2o/core/inject"
)

/**
 * Copyright 2009-2019 @ 56x.net
 * name : account_test.go.go
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2019-06-25 07:25
 * description :
 * history :
 */

func TestFlowAccount(t *testing.T) {
	repo := inject.GetMemberRepo()
	m := repo.GetMember(1)
	acc := m.GetAccount()
	balance := acc.GetValue().FlowBalance
	err := acc.Adjust(member.AccountFlow, "系统赠送", 10000, "系统", 1)
	if err == nil {
		err = acc.Charge(member.AccountFlow, "用户充值50元", 5000, "-", "")
		if err == nil {
			err = acc.Consume(member.AccountFlow, "消费150", 15000, "-", "")
		}
	}
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	nowBalance := repo.GetMember(int64(m.GetAggregateRootId())).GetAccount().GetValue().FlowBalance
	if nowBalance != balance {
		t.Logf("before:%.2f  now:%.2f", float64(balance)/100, float64(nowBalance)/100)
		t.FailNow()
	}
}

func TestMemberWalletOperate(t *testing.T) {
	var memberId int64 = 1
	inject.GetRegistryRepo().UpdateValue(registry.MemberWithdrawalMustVerification, "false")
	m := inject.GetMemberRepo().GetMember(memberId)
	ic := m.GetAccount()
	iw := ic.Wallet()
	amount := iw.Get().Balance

	// 获取第一张银行卡
	cards := m.Profile().GetBankCards()
	bankCardNo := cards[0].BankAccount
	assertError(t, ic.Charge(member.AccountWallet, "钱包充值",
		100000, "-", "测试"))
	id, _, err := ic.RequestWithdrawal(&wallet.WithdrawTransaction{
		Amount:           70000,
		TransactionFee:   0,
		Kind:             0,
		TransactionTitle: "提现到银行卡",
		BankName:         "",
		AccountNo:        bankCardNo,
		AccountName:      "",
	})
	assertError(t, err)
	ic.ReviewWithdrawal(id, true, "")
	id, _, err = ic.RequestWithdrawal(&wallet.WithdrawTransaction{
		Amount:           30000,
		TransactionFee:   0,
		Kind:             0,
		TransactionTitle: "提现到银行卡",
		BankName:         "",
		AccountNo:        bankCardNo,
		AccountName:      "",
	})
	assertError(t, err)
	assertError(t, ic.ReviewWithdrawal(id, false, "退回提现"))
	assertError(t, ic.Discount(member.AccountWallet, "钱包抵扣",
		30000, "-", "测试"))
	if final := int(ic.GetValue().WalletBalance); final != amount {
		t.Log("want ", amount, " final ", final)
		t.FailNow()
	}
}

func TestMemberFreeWallet(t *testing.T) {
	var memberId int64 = 1
	m := inject.GetMemberRepo().GetMember(memberId)
	ic := m.GetAccount()
	_, err := ic.CarryTo(member.AccountWallet, member.AccountOperateData{
		TransactionTitle:   "测试冻结1元",
		Amount:             100,
		OuterTransactionNo: "-",
		TransactionRemark:  "",
	}, true, 0)
	if err == nil {
		err = ic.Unfreeze(member.AccountWallet, member.AccountOperateData{
			TransactionTitle:   "解冻",
			Amount:             100,
			OuterTransactionNo: "",
			TransactionRemark:  "",
		}, true, 0)
	}
	if err != nil {
		t.Error(err)
	}

}

func TestMemberRedPack(t *testing.T) {
	//var memberId int64 = 1
	//m := inject.GetMemberRepo().GetMember(memberId)
	//ic := m.GetAccount()
	//iw := ic.Wallet()
	//amount := iw.Get().Balance
	//err := ic.Freeze("发送红包", "-", 100, 0)
	//if err ==nil{
	//}
}

func TestAccountAdjust(t *testing.T) {
	var memberId int64 = 723
	m := inject.GetMemberRepo().GetMember(memberId)
	ic := m.GetAccount()
	err := ic.Adjust(member.AccountWallet, "[KF]客服调整",
		8990,
		"-",
		1)
	if err != nil {
		t.Error(err)
	}
}

// 测试解冻钱包
func TestMemberUnfreezeWallet(t *testing.T) {
	var memberId int64 = 848

	m := inject.GetMemberRepo().GetMember(memberId)
	ic := m.GetAccount()
	err := ic.Unfreeze(member.AccountWallet, member.AccountOperateData{
		TransactionTitle:   "解冻",
		Amount:             8833,
		OuterTransactionNo: "",
		TransactionRemark:  "",
	}, true, 1)
	if err != nil {
		t.Error(err)
	}
}
