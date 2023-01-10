/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2013-12-23 23:15
 * description :
 * history :
 */

package impl

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/ixre/go2o/core/domain/interface/domain/enum"
	"github.com/ixre/go2o/core/domain/interface/item"
	"github.com/ixre/go2o/core/domain/interface/merchant"
	promodel "github.com/ixre/go2o/core/domain/interface/pro_model"
	"github.com/ixre/go2o/core/domain/interface/product"
	"github.com/ixre/go2o/core/domain/interface/valueobject"
	"github.com/ixre/go2o/core/infrastructure/domain"
	"github.com/ixre/go2o/core/infrastructure/format"
	"github.com/ixre/go2o/core/query"
	"github.com/ixre/go2o/core/service/parser"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/storage"
	"github.com/ixre/gof/types"
)

var _ proto.ItemServiceServer = new(itemService)

type itemService struct {
	serviceUtil
	itemRepo  item.IItemRepo
	itemQuery *query.ItemQuery
	cateRepo  product.ICategoryRepo
	labelRepo item.ISaleLabelRepo
	promRepo  promodel.IProductModelRepo
	mchRepo   merchant.IMerchantRepo
	valueRepo valueobject.IValueRepo
	sto       storage.Interface
	proto.UnimplementedItemServiceServer
}

func NewSaleService(sto storage.Interface, cateRepo product.ICategoryRepo,
	goodsRepo item.IItemRepo,
	goodsQuery *query.ItemQuery,
	labelRepo item.ISaleLabelRepo, promRepo promodel.IProductModelRepo,
	mchRepo merchant.IMerchantRepo, valueRepo valueobject.IValueRepo) *itemService {
	return &itemService{
		sto:       sto,
		itemRepo:  goodsRepo,
		itemQuery: goodsQuery,
		cateRepo:  cateRepo,
		labelRepo: labelRepo,
		promRepo:  promRepo,
		mchRepo:   mchRepo,
		valueRepo: valueRepo,
	}
}

// GetItem 获取商品
func (s *itemService) GetItem(_ context.Context, req *proto.GetItemRequest) (*proto.SItemDataResponse, error) {
	it := s.itemRepo.GetItem(req.ItemId)
	if it != nil {
		ret := parser.ItemDataDto(it.GetValue())
		skuArr := it.SkuArray()
		ret.Images = it.Images()
		ret.AttrArray = parser.AttrArrayDto(it.Product().Attr())
		ret.SkuArray = parser.SkuArrayDto(skuArr)
		ret.LevelPrices = parser.PriceArrayDto(it.GetLevelPrices())
		ret.FlagData = s.getFlagData(it.GetValue().ItemFlag)
		return ret, nil
	}
	return nil, item.ErrNoSuchItem
}

// 获取商品的标志数据
func (s *itemService) getFlagData(flag int) *proto.SItemFlagData {
	return &proto.SItemFlagData{
		IsNewOnShelve:  domain.TestFlag(flag, item.FlagNewOnShelve),
		IsHotSales:     domain.TestFlag(flag, item.FlagHotSales),
		IsRecommend:    domain.TestFlag(flag, item.FlagRecommend),
		IsExchange:     domain.TestFlag(flag, item.FlagExchange),
		IsGift:         domain.TestFlag(flag, item.FlagGift),
		IsAffiliate:    domain.TestFlag(flag, item.FlagAffiliate),
		IsSelfSales:    domain.TestFlag(flag, item.FlagSelfSales),
		IsFreeDelivery: domain.TestFlag(flag, item.FlagFreeDelivery),
		IsSelfDelivery: domain.TestFlag(flag, item.FlagSelfDelivery),
	}
}

// SaveItem 保存商品
func (s *itemService) SaveItem(_ context.Context, r *proto.SaveItemRequest) (*proto.SaveItemResponse, error) {
	var gi item.IGoodsItem
	it := parser.ParseGoodsItem(r)
	var err error
	if it.Id > 0 {
		gi = s.itemRepo.GetItem(it.Id)
		if gi == nil || gi.GetValue().VendorId != r.VendorId {
			return &proto.SaveItemResponse{
				ErrCode: 1,
				ErrMsg:  item.ErrNoSuchItem.Error(),
			}, nil
		}
	} else {
		gi = s.itemRepo.CreateItem(it)
	}
	err = gi.SetValue(it)
	if err == nil {
		err = gi.SetSku(it.SkuArray)
		if err == nil {
			if r.Images != nil && len(r.Images) > 0 {
				_ = gi.SetImages(r.Images)
			}
			it.Id, err = gi.Save()
		}
	}
	if err != nil {
		return &proto.SaveItemResponse{
			ErrCode: 1,
			ErrMsg:  err.Error(),
		}, nil
	}
	if err == nil && r.FlagData != nil {
		s.saveItemFlag(gi, r)
	}
	return &proto.SaveItemResponse{
		ErrCode: 0,
		ItemId:  it.Id,
	}, nil
}

// 保存商品标志
func (*itemService) saveItemFlag(gi item.IGoodsItem, r *proto.SaveItemRequest) {
	f := func(flag int, b bool) {
		if b {
			gi.GrantFlag(flag)
		} else {
			gi.GrantFlag(-flag)
		}
	}
	f(item.FlagNewOnShelve, r.FlagData.IsNewOnShelve)
	f(item.FlagHotSales, r.FlagData.IsHotSales)
	f(item.FlagRecommend, r.FlagData.IsRecommend)
	f(item.FlagExchange, r.FlagData.IsExchange)
	f(item.FlagGift, r.FlagData.IsGift)
	f(item.FlagAffiliate, r.FlagData.IsAffiliate)
	f(item.FlagSelfDelivery, r.FlagData.IsSelfDelivery)
	gi.Save()
}

// 附加商品的信息
func (s *itemService) attachUnifiedItem(item item.IGoodsItem, extra bool) *proto.SUnifiedViewItem {
	ret := parser.ItemDtoV2(item.GetValue())
	skuService := s.itemRepo.SkuService()
	skuArr := item.SkuArray()
	ret.SkuArray = parser.SkuArrayDto(skuArr)
	ret.LevelPrices = parser.PriceArrayDto(item.GetLevelPrices())
	specArr := item.SpecArray()
	ret.FlagData = s.getFlagData(item.GetValue().ItemFlag)
	if extra {
		ret.ViewData = &proto.SItemViewData{
			Details: "",  //todo:??
			Thumbs:  nil, //todo:??
			Images:  nil, //todo:??
			SkuHtml: skuService.GetSpecHtml(specArr),
			SkuJson: string(skuService.GetItemSkuJson(skuArr)),
		}
	}
	return ret
}

// 根据SKU获取商品
func (s *itemService) GetItemBySku(_ context.Context, r *proto.ItemBySkuRequest) (*proto.SUnifiedViewItem, error) {
	v := s.itemRepo.GetValueGoodsBySku(r.ProductId, r.SkuId)
	if v != nil {
		item := s.itemRepo.CreateItem(v)
		return s.attachUnifiedItem(item, r.Extra), nil
	}
	return nil, item.ErrNoSuchSku
}

// GetItemAndSnapshot 获取商品用于销售的快照和信息
func (s *itemService) GetItemAndSnapshot(_ context.Context, r *proto.GetItemAndSnapshotRequest) (*proto.ItemSnapshotResponse, error) {
	it := s.itemRepo.GetItem(r.GetItemId())
	if it != nil {
		skuService := s.itemRepo.SkuService()
		sn := it.Snapshot()
		// 基础数据及其销售数量
		ret := parser.ParseItemSnapshotDto(sn)
		ret.SaleNum = it.GetValue().SaleNum
		ret.StockNum = it.GetValue().StockNum
		// 图片
		ret.Images = it.Images()
		if len(ret.Images) == 0 && len(sn.Image) > 0 {
			ret.Images = []string{sn.Image}
		}
		// 获取SKU和详情等
		skuArr := it.SkuArray()
		specArr := it.SpecArray()
		ret.SkuArray = parser.SkuArrayDto(skuArr)
		ret.SpecOptions = parser.SpecOptionsDto(specArr)
		// 产品详情
		prod := it.Product()
		ret.Description = prod.GetValue().Description
		// 返回SKU的HTML选择器
		if r.ReturnSkuHtml {
			ret.SkuHtml = skuService.GetSpecHtml(specArr)
		}
		if r.ReturnSkuJson {
			ret.SkuJson = string(skuService.GetItemSkuJson(skuArr)) //todo: 是否可以去掉
		}
		return ret, nil
	}
	return nil, item.ErrNoSuchItem
}

// 获取商品交易快照
func (s *itemService) GetTradeSnapshot(_ context.Context, id *proto.Int64) (*proto.STradeSnapshot, error) {
	sn := s.itemRepo.GetSalesSnapshot(id.Value)
	if sn != nil {
		return parser.ParseTradeSnapshot(sn), nil
	}
	return nil, item.ErrNoSuchSnapshot
}

// 获取SKU
func (s *itemService) GetSku(_ context.Context, request *proto.SkuId) (*proto.SSku, error) {
	it := s.itemRepo.GetItem(request.ItemId)
	if it != nil {
		sku := it.GetSku(request.SkuId)
		if sku != nil {
			return s.parseSkuDto(sku), nil
		}
	}
	return nil, item.ErrNoSuchItem
}

// 获取商品详细数据
func (s *itemService) GetItemDetailData(_ context.Context, request *proto.ItemDetailRequest) (*proto.String, error) {
	it := s.itemRepo.CreateItem(&item.GoodsItem{Id: request.ItemId})
	switch request.ItemType {
	case item.ItemWholesale:
		data := it.Wholesale().GetJsonDetailData()
		return &proto.String{Value: string(data)}, nil
	}
	return &proto.String{Value: "不支持的商品类型"}, nil
}

// 获取上架商品数据（分页）
func (s *itemService) GetPagedOnShelvesItem(_ context.Context, r *proto.PagingGoodsRequest) (*proto.PagingGoodsResponse, error) {
	ret := &proto.PagingGoodsResponse{
		Total: 0,
		Data:  make([]*proto.SUnifiedViewItem, 0),
	}
	where := r.Params.Where
	if wd := strings.TrimSpace(r.Keyword); len(wd) > 0 {
		where += " AND item_info.title LIKE '%" + r.Keyword + "%'"
	}

	var total int32
	var list []*proto.SUnifiedViewItem
	switch r.ItemType {
	case proto.EItemSalesType_IT_NORMAL:
		total, list = s.getPagedOnShelvesItem(
			int32(r.CategoryId),
			int32(r.Params.Begin),
			int32(r.Params.End),
			where,
			r.Params.SortBy)
	case proto.EItemSalesType_IT_WHOLESALE:
		total, list = s.getPagedOnShelvesItemForWholesale(
			int32(r.CategoryId),
			int32(r.Params.Begin),
			int32(r.Params.End),
			where,
			r.Params.SortBy)
	}
	ret.Total = int64(total)
	ret.Data = list
	return ret, nil
}
func (s *itemService) getPagedOnShelvesItem(catId int32, start,
	end int32, where, sortBy string) (int32, []*proto.SUnifiedViewItem) {

	total, list := s.itemQuery.GetPagedOnShelvesItem(catId,
		start, end, where, sortBy)
	arr := make([]*proto.SUnifiedViewItem, len(list))
	for i, v := range list {
		v.Image = format.GetGoodsImageUrl(v.Image)
		arr[i] = parser.ItemDtoV2(v)
	}
	return total, arr
}

func (s *itemService) getPagedOnShelvesItemForWholesale(catId int32, start,
	end int32, where, sortBy string) (int32, []*proto.SUnifiedViewItem) {

	total, list := s.itemQuery.GetPagedOnShelvesItemForWholesale(catId,
		start, end, where, sortBy)
	arr := make([]*proto.SUnifiedViewItem, len(list))
	for i, v := range list {
		v.Image = format.GetGoodsImageUrl(v.Image)
		dto := parser.ItemDtoV2(v)
		s.attachWholesaleItemDataV2(dto)
		arr[i] = dto
	}
	return total, arr
}

// 附加批发商品的信息
func (s *itemService) attachWholesaleItemDataV2(dto *proto.SUnifiedViewItem) {
	dto.Data = make(map[string]string)
	vendor := s.mchRepo.GetMerchant(int(dto.VendorId))
	if vendor != nil {
		vv := vendor.GetValue()
		pStr := s.valueRepo.GetAreaName(int32(vv.Province))
		cStr := s.valueRepo.GetAreaName(int32(vv.City))
		dto.Data["VendorName"] = vv.CompanyName
		dto.Data["ShipArea"] = pStr + cStr
		// 认证信息
		ei := vendor.ProfileManager().GetEnterpriseInfo()
		if ei != nil && ei.Reviewed == enum.ReviewPass {
			dto.Data["Authorized"] = "true"
		} else {
			dto.Data["Authorized"] = "false"
		}
		// 品牌
		b := s.promRepo.BrandService().Get(int32(dto.BrandId))
		if b != nil {
			dto.Data["BrandName"] = b.Name
			dto.Data["BrandImage"] = b.Image
			dto.Data["BrandId"] = strconv.Itoa(int(b.ID))
		}
	}
}

// 获取上架商品数据
func (s *itemService) GetItems(_ context.Context, r *proto.GetItemsRequest) (*proto.PagingGoodsResponse, error) {
	/*
		hash := fmt.Sprintf("%d-%d-%_s", catId, quantity, where)
			hash = crypto.Md5([]byte(hash))
			key := "go2o:shopQuery:cache:rd-item:" + hash
			var arr []*proto.SOldItem

			fn := func() interface{} {
				list := _s.itemQuery.GetRandomItem(catId, quantity, where)
				for _, v := range list {
					v.Image = format.GetGoodsImageUrl(v.Image)
					arr = append(arr, parser.ItemDto(v))
				}
				return arr
			}
			_s.sto.RWJson(key, &arr, fn, 600)
			return arr
	*/
	c := s.cateRepo.GlobCatService().GetCategory(int(r.CategoryId))
	var idList []int = nil
	var list []*item.GoodsItem
	if c != nil {
		idList = c.GetChildes()
	}
	if r.Random {
		list = s.itemQuery.GetRandomItem(idList, int(r.Begin), int(r.End), r.Where)

	} else {
		list = s.itemQuery.GetOnShelvesItem(idList, int(r.Begin), int(r.End), r.Where)
	}
	arr := make([]*proto.SUnifiedViewItem, len(list))
	for i, v := range list {
		v.Image = format.GetGoodsImageUrl(v.Image)
		arr[i] = parser.ItemDtoV2(v)
	}
	return &proto.PagingGoodsResponse{
		Total: 0,
		Data:  arr,
	}, nil
}

func (s *itemService) GetAllSaleLabels(_ context.Context, _ *proto.Empty) (*proto.ItemLabelListResponse, error) {
	tags := s.labelRepo.LabelService().GetAllSaleLabels()
	ret := &proto.ItemLabelListResponse{
		Value: make([]*proto.SItemLabel, len(tags)),
	}
	for i, v := range tags {
		ret.Value[i] = parser.ParseSaleLabelDto(v.GetValue())
	}
	return ret, nil
}

func (s *itemService) GetSaleLabel(_ context.Context, id *proto.IdOrName) (*proto.SItemLabel, error) {
	var tag item.ISaleLabel
	if id.Id <= 0 {
		tag = s.labelRepo.LabelService().GetSaleLabelByCode(id.Name)
	} else {
		tag = s.labelRepo.LabelService().GetSaleLabel(int32(id.Id))
	}
	if tag != nil {
		return parser.ParseSaleLabelDto(tag.GetValue()), nil
	}
	return nil, errors.New("no such sale label")
}

// SaveSaleLabel 保存销售标签
func (s *itemService) SaveSaleLabel(_ context.Context, v *proto.SItemLabel) (*proto.Result, error) {
	ls := s.labelRepo.LabelService()
	var value = parser.FromSaleLabelDto(v)
	var lb item.ISaleLabel
	if v.Id > 0 {
		if lb = ls.GetSaleLabel(v.Id); lb == nil {
			return s.error(errors.New("没有销售标签")), nil
		}
	} else {
		lb = ls.CreateSaleLabel(value)
	}
	err := lb.SetValue(value)
	if err == nil {
		_, err = lb.Save()
	}
	return s.error(err), nil
}

// 删除销售标签
func (s *itemService) DeleteSaleLabel(_ context.Context, i *proto.Int64) (*proto.Result, error) {
	err := s.labelRepo.LabelService().DeleteSaleLabel(int32(i.Value))
	return s.error(err), nil
}

// // 根据销售标签获取指定数目的商品
// func (s *itemService) GetValueGoodsBySaleLabel(_ context.Context, r *proto.GetItemsByLabelRequest) (*proto.PagingShopGoodsResponse, error) {
// 	tag := s.labelRepo.LabelService().GetSaleLabelByCode(r.Label)
// 	ret := &proto.PagingShopGoodsResponse{
// 		Data: make([]*proto.SGoods, 0),
// 	}
// 	if tag != nil {
// 		list := tag.GetValueGoods(r.SortBy, int(r.Begin), int(r.End))
// 		for _, v := range list {
// 			v.Image = format.GetGoodsImageUrl(v.Image)
// 			ret.Data = append(ret.Data, s.parseGoods(v))
// 		}
// 	}
// 	return ret, nil
// }

// // 根据分页销售标签获取指定数目的商品
// func (s *itemService) GetPagedValueGoodsBySaleLabel_(_ context.Context, r *proto.SaleLabelItemsRequest_) (*proto.PagingGoodsResponse, error) {
// 	tag := s.labelRepo.LabelService().CreateSaleLabel(&item.Label{
// 		Id: r.LabelId,
// 	})
// 	total, list := tag.GetPagedValueGoods(r.Params.SortBy, int(r.Params.Begin), int(r.Params.End))
// 	arr := make([]*proto.SUnifiedViewItem, len(list))
// 	for i, v := range list {
// 		v.Image = format.GetGoodsImageUrl(v.Image)
// 		arr[i] = parser.ParseGoodsDto_(v)
// 	}
// 	return &proto.PagingGoodsResponse{
// 		Total: int64(total),
// 		Data:  arr,
// 	}, nil
// }

// 保存商品的会员价
func (s *itemService) SaveLevelPrices(_ context.Context, r *proto.SaveLevelPriceRequest) (*proto.Result, error) {
	it := s.itemRepo.GetItem(r.ItemId)
	var err error
	if it != nil {
		for _, v := range r.Value {
			e := parser.ParseLevelPrice(v)
			e.GoodsId = r.ItemId
			if _, err := it.SaveLevelPrice(e); err != nil {
				break
			}
		}
	}
	return s.error(err), nil
}

// 设置商品货架状态
func (s *itemService) SetShelveState(_ context.Context, r *proto.ShelveStateRequest) (*proto.Result, error) {
	it := s.itemRepo.GetItem(r.ItemId)
	var err error
	if it == nil || it.GetValue().VendorId != r.SellerId {
		err = item.ErrNoSuchItem
	} else {
		state := int32(types.ElseInt(r.ShelveOn,
			int(item.ShelvesOn),
			int(item.ShelvesDown)))
		switch r.ItemType {
		case proto.EItemSalesType_IT_WHOLESALE:
			err = it.Wholesale().SetShelve(state, r.Remark)
		case proto.EItemSalesType_IT_NORMAL:
			err = it.SetShelve(state, r.Remark)
		}
	}
	return s.result(err), nil
}

// 审核商品
func (s *itemService) ReviewItem(_ context.Context, r *proto.ItemReviewRequest) (*proto.Result, error) {
	it := s.itemRepo.GetItem(r.ItemId)
	var err error
	if it == nil {
		err = item.ErrNoSuchItem
	} else {
		err = it.Review(r.Pass, r.Remark)
	}
	return s.result(err), nil
}

// 标记为违规
func (s *itemService) SignAsIllegal(_ context.Context, r *proto.ItemIllegalRequest) (*proto.Result, error) {
	it := s.itemRepo.GetItem(r.ItemId)
	var err error
	if it == nil {
		err = item.ErrNoSuchItem
	} else {
		err = it.Incorrect(r.Remark)
	}
	return s.error(err), nil
}

// 获取批发价格数组
func (s *itemService) GetWholesalePriceArray(_ context.Context, id *proto.SkuId) (*proto.SWsSkuPriceListResponse, error) {
	it := s.itemRepo.GetItem(id.ItemId)
	arr := it.Wholesale().GetSkuPrice(id.SkuId)
	ret := make([]*proto.SWsSkuPrice, len(arr))
	for i, v := range arr {
		ret[i] = parser.WsSkuPriceDto(v)
	}
	return &proto.SWsSkuPriceListResponse{
		Value: ret,
	}, nil
}

// 保存批发价格
func (s *itemService) SaveWholesalePrice(_ context.Context, r *proto.SaveSkuPricesRequest) (*proto.Result, error) {
	it := s.itemRepo.GetItem(r.ItemId)
	arr := make([]*item.WsSkuPrice, len(r.Value))
	for i, v := range r.Value {
		e := parser.WsSkuPrice(v)
		e.ItemId = r.ItemId
		e.SkuId = r.SkuId
		arr[i] = e
	}
	err := it.Wholesale().SaveSkuPrice(r.SkuId, arr)
	return s.error(err), nil
}

// 获取批发折扣数组
func (s *itemService) GetWholesaleDiscountArray(_ context.Context, id *proto.GetWsDiscountRequest) (*proto.SWsItemDiscountListResponse, error) {

	it := s.itemRepo.GetItem(id.ItemId)
	arr := it.Wholesale().GetItemDiscount(int32(id.GroupId))
	ret := make([]*proto.SWsItemDiscount, len(arr))
	for i, v := range arr {
		ret[i] = parser.WsItemDiscountDto(v)
	}
	return &proto.SWsItemDiscountListResponse{
		Value: ret,
	}, nil
}

// 保存批发折扣
func (s *itemService) SaveWholesaleDiscount(_ context.Context, r *proto.SaveItemDiscountRequest) (*proto.Result, error) {
	it := s.itemRepo.GetItem(r.ItemId)
	arr := make([]*item.WsItemDiscount, len(r.Value))
	for i, v := range r.Value {
		e := parser.WsItemDiscount(v)
		e.ItemId = r.ItemId
		e.BuyerGid = int32(r.GroupId)
		arr[i] = e
	}
	err := it.Wholesale().SaveItemDiscount(int32(r.GroupId), arr)
	return s.error(err), nil
}

func (s *itemService) parseSkuDto(sku *item.Sku) *proto.SSku {
	return &proto.SSku{
		SkuId:       sku.Id,
		ProductId:   sku.ProductId,
		ItemId:      sku.ItemId,
		Title:       sku.Title,
		Image:       sku.Image,
		SpecData:    sku.SpecData,
		SpecWord:    sku.SpecWord,
		Code:        sku.Code,
		RetailPrice: sku.RetailPrice,
		Price:       sku.Price,
		Cost:        sku.Cost,
		Weight:      sku.Weight,
		Bulk:        sku.Bulk,
		Stock:       sku.Stock,
		SaleNum:     sku.SaleNum,
	}
}
