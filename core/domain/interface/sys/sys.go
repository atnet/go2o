package sys

import (
	"github.com/ixre/go2o/core/infrastructure/domain"
	"github.com/ixre/go2o/core/infrastructure/fw"
)

// ISystemAggregateRoot 系统聚合根
type (
	ISystemAggregateRoot interface {
		domain.IAggregateRoot

		// 获取地址管理器
		Address() IAddressManager

		// 获取选项管理器
		Options() IOptionManager

		// 标记更新状态,通常监听数据变更或事件后调用
		FlushUpdateStatus()

		// 获取最后更新时间,用于对比系统设置是否已经变更
		LastUpdateTime() int64
	}

	// IAddressManager 地址管理器
	IAddressManager interface {
		// GetAllCities 获取所有城市
		GetAllCities() []*District
		// GetDistrictList 获取区域信息
		GetDistrictList(parentId int) []*District
		// GetDistricts 获取区域名称
		GetDistrictNames(code ...int) map[int]string
	}

	// IOptionManager 选项管理器
	IOptionManager interface {
		// GetOptionNames 获取选项名称
		GetOptionNames(code ...int) map[int]string
		// 获取下级选项
		GetChildOptions(parentId int, typeName string) []*GeneralOption
		// 是否为叶子节点
		IsLeaf(n *GeneralOption) bool
	}
)

// ISystemRepo 系统仓储
type ISystemRepo interface {
	// GetSystemAggregateRoot 获取系统聚合根
	GetSystemAggregateRoot() ISystemAggregateRoot
	// 标记已更新
	FlushUpdateStatus()
	// 获取最后更新时间
	LastUpdateTime() int64
	// District 获取区域仓储
	District() fw.Repository[District]
	// Option 获取选项仓储
	Option() fw.Repository[GeneralOption]
}

type (
	// District ChinaArea
	District struct {
		// Code
		Code int `db:"code" pk:"yes" json:"code" bson:"code"`
		// Name
		Name string `db:"name" json:"name" bson:"name"`
		// Parent
		Parent int `db:"parent" json:"parent" bson:"parent"`
	}
)

func (a District) TableName() string {
	return "sys_district"
}

// GeneralOption 系统通用选项(用于存储分类和选项等数据)
type GeneralOption struct {
	// 编号
	Id int `db:"id" pk:"yes" auto:"yes" json:"id" bson:"id"`
	// 类型
	Type string `db:"type" json:"type" bson:"type"`
	// 上级编号
	Pid int `db:"pid" json:"pid" bson:"pid"`
	// 名称
	Name string `db:"name" json:"name" bson:"name"`
	// 值
	Value string `db:"value" json:"value" bson:"value"`
	// 排列序号
	SortNum int `db:"sort_num" json:"sortNum" bson:"sortNum"`
	// 是否启用
	Enabled int `db:"enabled" json:"enabled" bson:"enabled"`
	// 创建时间
	CreateTime int `db:"create_time" json:"createTime" bson:"createTime"`
}

func (s GeneralOption) TableName() string {
	return "sys_general_option"
}
