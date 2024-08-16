/**
 * This file is auto generated by tto v0.5.7 !
 * If you want to modify this code, please read the guide
 * to modify code template.
 *
 * Get started: https://github.com/ixre/tto
 *
 * Copyright (C) 2009-2024 56X.NET, All rights reserved.
 *
 * name : staff_repo.go
 * author : jarrysix
 * date : 2024/06/10 09:30:40
 * description :
 * history :
 */
package repos

import (
	"github.com/ixre/go2o/core/domain/interface/merchant/staff"
	"github.com/ixre/go2o/core/infrastructure/fw"
)

var _ staff.IStaffRepo = new(staffRepoImpl)

type staffRepoImpl struct {
	fw.BaseRepository[staff.Staff]
}

func NewStaffRepo(o fw.ORM) staff.IStaffRepo {
	s := &staffRepoImpl{}
	s.ORM = o
	return s
}

// SelectStaff Select 商户代理人坐席(员工)
func (m *staffRepoImpl) GetStaffByMemberId(memberId int) *staff.Staff {
	return m.FindBy("member_id = $1 and work_status <> $2", memberId,
		staff.WorkStatusOff)
}

var _ staff.IStaffTransferRepo = new(mchStaffTransferRepoImpl)

type mchStaffTransferRepoImpl struct {
	fw.BaseRepository[staff.StaffTransfer]
}

// NewStaffTransferRepo 创建员工转商户仓储
func NewStaffTransferRepo(o fw.ORM) staff.IStaffTransferRepo {
	r := &mchStaffTransferRepoImpl{}
	r.ORM = o
	return r
}
