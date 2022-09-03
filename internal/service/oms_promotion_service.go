package service

import (
	"gorm.io/gorm"
	"mall/global/domain"
	"mall/global/model"
)

type omsPromotionService struct {
	db *gorm.DB
}

func NewOmsPromotionService(db *gorm.DB) OmsPromotionService {
	return &omsPromotionService{db: db}
}

func (o *omsPromotionService) CalcCartPromotion(cartItemList []model.OmsCartItem) []domain.CartPromotionItem {
	//1.先根据productId对CartItem进行分组，以spu为单位进行计算优惠
	productCartMap := o.groupCartItemBySpu(cartItemList)
	o.g
}

/**
 * @Author 96234
 * @Description 按照商品id给商品分组
 * @Date 11:37 2022/8/17
 * @Param
 * @return
 **/
func (s *omsPromotionService) groupCartItemBySpu(cartItemList []model.OmsCartItem) (productCartMap map[int64][]model.OmsCartItem) {
	for _, item := range cartItemList {
		if items, ok := productCartMap[*item.ProductId]; ok {
			items = make([]model.OmsCartItem, 0)
			items = append(items, item)
			productCartMap[*item.ProductId] = items
		} else {
			items = append(items, item)
		}
	}
	return
}

func (s *omsPortalOrderService) getPromotionProductList(cartItemList []model.OmsCartItem) {
	ids := make([]int64, len(cartItemList))
	for i, item := range cartItemList {
		ids[i] = item.Id
	}
	s.db.Model()
}
