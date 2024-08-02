/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2014-02-14 15:42
 * description :
 * history :
 */
package delivery

import (
	"errors"

	"github.com/ixre/go2o/core/domain/interface/delivery"
	"github.com/ixre/go2o/core/infrastructure/util/lbs"
)

var _ delivery.ICoverageArea = new(CoverageArea)

type CoverageArea struct {
	value *delivery.CoverageValue
	rep   delivery.IDeliveryRepo
}

func newCoverageArea(v *delivery.CoverageValue, rep delivery.IDeliveryRepo) delivery.ICoverageArea {
	return &CoverageArea{
		value: v,
		rep:   rep,
	}
}

// 是否可以配送
// 返回是否可以配送，以及距离(米)
func (c *CoverageArea) CanDeliver(lng, lat float64) (bool, int) {
	distance := lbs.GetLocDistance(
		c.value.Lng, c.value.Lat, lng, lat)
	i := int(distance)
	return i <= c.value.Radius*1000, i
}

// 是否可以配送
// 返回是否可以配送，以及距离(米)
func (c *CoverageArea) CanDeliverTo(address string) (bool, int) {
	lng, lat, err := lbs.GetLocation(address)
	if err != nil {
		return false, -1
	}
	return c.CanDeliver(lng, lat)
}

func (c *CoverageArea) GetDomainId() int32 {
	return c.value.Id
}

func (c *CoverageArea) GetValue() delivery.CoverageValue {
	return *c.value
}

func (c *CoverageArea) SetValue(v *delivery.CoverageValue) error {
	if v.Id == c.value.Id && v.Id > 0 {
		c.value = v
		return nil
	}
	return errors.New("no such value")
}

func (c *CoverageArea) Save() (int32, error) {
	return c.rep.SaveCoverageArea(c.value)
}
