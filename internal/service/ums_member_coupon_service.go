package service

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/wxnacy/wgo/arrays"
	"gorm.io/gorm"
	"mall/global/dao/mapper"
	"mall/global/domain"
	"mall/global/model"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/**
 *@author:
 *@date:2022/8/23
**/

type umsMemberCouponService struct {
	db               *gorm.DB
	couponHistoryDao mapper.SmsCouponHistoryDao
}

func NewUmsMemberCouponService(db *gorm.DB) UmsMemberCouponService {
	return &umsMemberCouponService{db: db}
}

func (u *umsMemberCouponService) Add(couponId int64, member model.UmsMember) (err error) {
	var coupon model.SmsCoupon
	u.db.Model(&coupon).Where(couponId).First(&coupon)
	if coupon.Id == 0 {
		err = errors.New("coupons don't exist")
		return
	}
	if coupon.Count <= 0 {
		err = errors.New("insufficient coupons")
		return
	}
	now := time.Now()
	if now.Before(*coupon.EnableTime) {
		err = errors.New("the coupon is not due yet")
		return
	}
	var count int64
	u.db.Model(&model.SmsCouponHistory{}).Where(map[string]interface{}{"coupon_id": couponId, "member_id": member.Id}).Count(&count)
	if count >= int64(*coupon.PerLimit) {
		err = errors.New("You've already picked up  coupon")
		return
	}
	var couponHistory model.SmsCouponHistory
	couponHistory.CouponId = couponId
	code := u.GenerateCouponCode(member.Id)
	couponHistory.CouponCode = code
	t := time.Now()
	couponHistory.CreateTime = t
	couponHistory.MemberId = member.Id
	couponHistory.UseStatus = 0
	couponHistory.GetType = 1
	u.db.Create(&couponHistory)
	coupon.Count = coupon.Count - 1
	if coupon.ReceiveCount == nil {
		c := 1
		coupon.ReceiveCount = &c
	} else {
		c := *coupon.ReceiveCount - 1
		coupon.ReceiveCount = &c
	}
	u.db.Save(&coupon)
	return nil
}

func (u *umsMemberCouponService) GenerateCouponCode(memberId int64) string {
	milli := []byte(strconv.FormatInt(time.Now().UnixMilli(), 10))
	var str strings.Builder
	str.Write(milli[len(milli)-8:])
	r := []byte(strconv.FormatInt(rand.Int63n(10000), 10))
	str.Write(r)
	m := []byte(strconv.FormatInt(memberId, 10))
	str.Write(m[len(m)-4:])
	return str.String()
}
func (u *umsMemberCouponService) ListHistory(useStatus int, memberId int64) (list []model.SmsCouponHistory) {
	m := map[string]interface{}{}
	if useStatus != -1 {
		m["status"] = useStatus
	}
	m["member_id"] = memberId
	u.db.Model(&model.SmsCouponHistory{}).Where(m).Find(&list)
	return
}

func (u *umsMemberCouponService) ListCart(cartItem []domain.CartPromotionItem, memberId int64, t int) []domain.SmsCouponHistoryDetail {
	now := time.Now()
	detailList := u.couponHistoryDao.GetDetailList(memberId)
	var enableList []domain.SmsCouponHistoryDetail
	var disableList []domain.SmsCouponHistoryDetail
	for _, detail := range detailList {
		useType := *detail.Coupon.UseType
		point := detail.Coupon.MinPoint
		endTime := *detail.Coupon.EndTime
		if useType == 0 {
			totalAmount := u.CalcTotalAmount(cartItem)
			if (now.Before(endTime)) && totalAmount.Sub(point).IsPositive() {
				enableList = append(enableList, detail)
			} else {
				disableList = append(disableList, detail)
			}
		} else if useType == 1 {
			var productCategoryIds []int64
			for _, relation := range detail.CategoryRelationList {
				productCategoryIds = append(productCategoryIds, relation.Id)
			}
			totalAmount := u.CalcTotalAmountByProductCategoryId(cartItem, productCategoryIds)
			if now.Before(endTime) && totalAmount.IsPositive() && totalAmount.Sub(point).IsPositive() {
				enableList = append(enableList, detail)
			} else {
				disableList = append(disableList, detail)
			}
		} else if useType == 2 {
			var productIds []int64
			for _, relation := range detail.ProductRelationList {
				productIds = append(productIds, relation.ProductId)
			}
			totalAmount := u.CalcTotalAmountByProductId(cartItem, productIds)
			if now.Before(endTime) && totalAmount.IsPositive() && totalAmount.Sub(point).IsPositive() {
				enableList = append(enableList, detail)
			} else {
				disableList = append(disableList, detail)
			}
		}

	}
	if t == 1 {
		return enableList
	} else {
		return disableList
	}
}
func (u *umsMemberCouponService) CalcTotalAmount(cartItemList []domain.CartPromotionItem) decimal.Decimal {
	total := decimal.NewFromInt(0)
	for _, item := range cartItemList {
		realPrice := item.Price.Sub(item.ReduceAmount)
		total = total.Add(realPrice.Mul(decimal.NewFromInt32(int32(*item.Quantity))))
	}
	return total
}

func (u *umsMemberCouponService) CalcTotalAmountByProductCategoryId(cartItemList []domain.CartPromotionItem,
	productCategoryIds []int64) decimal.Decimal {
	total := decimal.NewFromInt(0)
	for _, item := range cartItemList {
		if -1 != arrays.ContainsInt(productCategoryIds, int64(*item.ProductCategoryId)) {
			realPrice := item.Price.Sub(item.ReduceAmount)
			total = total.Add(realPrice.Mul(decimal.NewFromInt32(int32(*item.Quantity))))

		}
	}
	return total
}
func (u *umsMemberCouponService) CalcTotalAmountByProductId(cartItemList []domain.CartPromotionItem,
	productIds []int64) decimal.Decimal {
	total := decimal.NewFromInt(0)
	for _, item := range cartItemList {
		if -1 != arrays.ContainsInt(productIds, item.ProductId) {
			realPrice := item.Price.Sub(item.ReduceAmount)
			total = total.Add(realPrice.Mul(decimal.NewFromInt32(int32(*item.Quantity))))
		}
	}
	return total
}
func (u *umsMemberCouponService) ListByProduct(productId int64) (coupons []model.SmsCoupon) {
	var allCouponIds []int64
	var cpr []model.SmsCouponProductRelation
	u.db.Model(&model.SmsCouponProductRelation{}).Where(map[string]interface{}{"product_id": productId}).
		Find(&cpr)
	if len(cpr) != 0 {
		for _, relation := range cpr {
			allCouponIds = append(allCouponIds, relation.ProductId)
		}
	}

	var product model.PmsProduct
	u.db.Model(&product).Where(product).First(&product)
	var cpcrList []model.SmsCouponProductCategoryRelation
	u.db.Model(&model.SmsCouponProductCategoryRelation{}).Where(map[string]interface{}{"product_category_id": product.ProductCategoryId}).
		Find(&cpcrList)
	if len(cpcrList) != 0 {
		for _, relation := range cpcrList {
			allCouponIds = append(allCouponIds, relation.CouponId)
		}
	}
	if len(allCouponIds) == 0 {
		return
	}
	u.db.Model(&model.SmsCoupon{}).Where(" (end_time>now() and start_time<now() and use_type=0) or"+
		" (end_time>now() and start_time<now() and use_type!=0 and  id in ?) ", allCouponIds).Find(&coupons)
	return
}

func (u *umsMemberCouponService) List(useStatus int, memberId int64) []model.SmsCoupon {
	list := u.couponHistoryDao.GetCouponList(memberId, useStatus)
	return list
}
