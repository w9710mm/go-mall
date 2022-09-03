package service

import (
	"github.com/wxnacy/wgo/arrays"
	"gorm.io/gorm"
	"mall/global/dao/mapper"
	"mall/global/domain"
	"mall/global/model"
	"time"
)

type omsCartItemService struct {
	db                  *gorm.DB
	omsPromotionService OmsPromotionService
	portalProductDao    mapper.PortalProductDao
}

func NewOmsCartItemService(db *gorm.DB) OmsCartItemService {
	return &omsCartItemService{db: db}
}
func (o *omsCartItemService) Add(cartItem model.OmsCartItem, member model.UmsMember) (count int64) {
	cartItem.MemberId = member.Id
	cartItem.MemberNickname = member.Nickname
	status := 0
	cartItem.DeleteStatus = &status
	item := o.getCartItem(cartItem)
	t := time.Now()
	if cartItem.Id != 0 {
		cartItem.CreateDate = &t
		count = o.db.Save(cartItem).RowsAffected
	} else {
		item.ModifyDate = &t
		modeify := *item.Quantity + *item.Quantity
		item.Quantity = &modeify
		count = o.db.Save(item).RowsAffected
	}
	return count

}

/**
 * 根据会员id,商品id和规格获取购物车中商品
 */
func (o *omsCartItemService) getCartItem(cartItem model.OmsCartItem) model.OmsCartItem {
	o.db.Model(&model.OmsCartItem{}).Where(cartItem).First(&cartItem)
	if cartItem.Id != 0 {
		return cartItem
	}
	return model.OmsCartItem{}
}

func (o *omsCartItemService) List(memberId int64) (cart []model.OmsCartItem) {
	status := 0
	o.db.Model(&model.OmsCartItem{}).Where(model.OmsCartItem{MemberId: memberId, DeleteStatus: &status}).
		Find(&cart)
	return
}

func (o *omsCartItemService) ListPromotion(memberId int64, cartIds []int64) []domain.CartPromotionItem {
	cartItems := o.List(memberId)
	cartFilterItem := make([]model.OmsCartItem, 0)
	if len(cartItems) != 0 {
		for _, item := range cartItems {
			if -1 != arrays.ContainsInt(cartIds, item.Id) {
				cartFilterItem = append(cartFilterItem, item)
			}
		}
	}
	var cartPromotionItemList []domain.CartPromotionItem
	if len(cartFilterItem) != 0 {
		cartPromotionItemList = o.omsPromotionService.CalcCartPromotion(cartFilterItem)
	}
	return cartPromotionItemList
}

func (o *omsCartItemService) UpdateQuantity(id int64, memberId int64, quantity int) int64 {
	item := model.OmsCartItem{Id: id}
	status := 0
	return o.db.Model(&item).
		Updates(&model.OmsCartItem{
			MemberId: memberId, Quantity: &quantity, DeleteStatus: &status}).RowsAffected
}

func (o *omsCartItemService) Delete(memberId int64, ids []int64) int64 {
	return o.db.Model(&model.OmsCartItem{}).Where("id in ? and member_id= ?", ids, memberId).
		Updates(map[string]interface{}{"delete_status": 1}).RowsAffected
}

func (o *omsCartItemService) GetCartProduct(productId int64) domain.CartProduct {
	return o.portalProductDao.GetCartProduct(productId)
}

func (o *omsCartItemService) UpdateAttr(item model.OmsCartItem) int64 {
	updateAttr := model.OmsCartItem{}
	updateAttr.Id = item.Id
	t := time.Now()
	updateAttr.ModifyDate = &t
	deleteStatus := 1
	return o.db.Model(&model.OmsCartItem{Id: item.Id}).Updates(&model.OmsCartItem{ModifyDate: &t,
		DeleteStatus: &deleteStatus}).RowsAffected

}

func (o *omsCartItemService) Clear(MemberId int64) int64 {
	return o.db.Model(&model.OmsCartItem{Id: MemberId}).Update("delete_status", 1).RowsAffected

}
