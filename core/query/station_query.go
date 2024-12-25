package query

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ixre/go2o/core/domain/interface/sys"
	"github.com/ixre/go2o/core/infrastructure/fw/collections"

	"github.com/ixre/gof/db"
	"github.com/ixre/gof/db/orm"
)

type StationQuery struct {
	db.Connector
	_orm orm.Orm
}

// 站点区域
type StationArea struct {
	// 编号
	Id int `db:"id" pk:"yes" auto:"yes" json:"id" bson:"id"`
	// Name
	Name string `db:"name" json:"name" bson:"name"`
	// 状态: 0: 待开通  1: 已开通  2: 已关闭
	Stations []SubStation `db:"stations" json:"stations" bson:"stations"`
}

// SubStation 地区子站
type SubStation struct {
	// 编号
	Id int `db:"id" pk:"yes" auto:"yes" json:"id" bson:"id"`
	// Name
	Name string `db:"name" json:"name" bson:"name"`
	// 城市编号
	CityCode int `db:"city_code" json:"cityCode" bson:"cityCode"`
	// 状态: 0: 待开通  1: 已开通  2: 已关闭
	Status int `db:"status" json:"status" bson:"status"`
	// 首字母
	Letter string `db:"letter" json:"-" bson:"letter"`
	// 是否热门
	IsHot int `db:"is_hot" json:"isHot" bson:"isHot"`
	// 上级
	Parent int `db:"parent" json:"-" bson:"-"`
}

// 站点查询
func NewStationQuery(o orm.Orm) *StationQuery {
	return &StationQuery{
		Connector: o.Connector(),
		_orm:      o,
	}
}

// QueryStations 查询站点列表
func (s *StationQuery) QueryStations(status int) []*StationArea {
	list := make([]*sys.District, 0)
	_ = s._orm.Select(&list, "parent = 0 and code <> 0")

	province := collections.MapList(list, func(v *sys.District) *StationArea {
		return &StationArea{
			Id:       v.Code,
			Name:     v.Name,
			Stations: []SubStation{},
		}
	})

	provinceIdList := collections.MapList(list, func(v *sys.District) string {
		return strconv.Itoa(v.Code)
	})
	cities := make([]*SubStation, 0)
	err := s._orm.SelectByQuery(&cities, fmt.Sprintf(`
		SELECT s.id,a.name,s.status,s.letter,s.is_hot,a.parent FROM sys_sub_station s
		INNER JOIN sys_district a ON a.code = s.city_code
		WHERE a.parent IN (%s)`,
		strings.Join(provinceIdList, ",")))
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[ Orm][ Error]: %s; Entity:Area\n", err.Error())
	}
	for _, c := range cities {
		pid := c.Parent
		if status == 1 && c.Status != 1 {
			// 如果只查询已开通的城市
			continue
		}
		for _, p := range province {
			if p.Id == pid {
				p.Stations = append(p.Stations, *c)
			}
		}
	}
	return province
}

// QueryGroupStations 按字母分组查询站点列表
func (s *StationQuery) QueryGroupStations(status int) map[string][]SubStation {
	cities := make([]*SubStation, 0)
	err := s._orm.SelectByQuery(&cities, `
	SELECT s.id,a.name,s.city_code,s.status,s.letter,s.is_hot,a.parent FROM sys_sub_station s
	LEFT JOIN sys_district a ON a.code = s.city_code`)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[ Orm][ Error]: %s; Entity:Area\n", err.Error())
	}
	ret := make(map[string][]SubStation, 0)
	for _, c := range cities {
		arr, ok := ret[c.Letter]
		if !ok {
			// 初始化
			arr = make([]SubStation, 0)
		}
		arr = append(arr, *c)
		ret[c.Letter] = arr
	}
	return ret
}
