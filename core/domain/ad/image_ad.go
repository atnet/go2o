/**
 * Copyright 2015 @ 56x.net.
 * name : image_ad
 * author : jarryliu
 * date : 2016-05-25 11:29
 * description :
 * history :
 */
package ad

import (
	"github.com/ixre/go2o/core/domain/interface/ad"
)

var _ ad.IImageAd = new(ImageAdImpl)

type ImageAdImpl struct {
	extValue *ad.Data
	*adImpl
}

// 获取链接广告值
func (i *ImageAdImpl) getData() *ad.Data {
	if i.extValue == nil {
		gallery := i._rep.GetSwiperAd(int64(i.GetDomainId()))
		if gallery.Len() > 0 {
			i.extValue = gallery[0]
		}
	}
	return i.extValue
}

func (i *ImageAdImpl) SetData(d *ad.Data) error {
	v := i.getData()
	v.LinkUrl = d.LinkUrl
	v.Title = d.Title
	v.ImageUrl = d.ImageUrl
	v.Enabled = 1 // 图片广告的图片无启用和停用功能
	return nil
}

// 保存广告
func (i *ImageAdImpl) Save() (int64, error) {
	id, err := i.adImpl.Save()
	if err == nil {
		v := i.getData()
		if v == nil {
			v = &ad.Data{
				ImageUrl: "",
			}
		}
		v.AdId = int(id)
		_, err = i._rep.SaveImageAdData(v)
	}
	return id, err
}

// 转换为数据传输对象
func (i *ImageAdImpl) Dto() *ad.AdDto {
	return &ad.AdDto{
		Id:     i.adImpl.GetDomainId(),
		AdType: i.adImpl.Type(),
		Data:   i.getData(),
	}
}
