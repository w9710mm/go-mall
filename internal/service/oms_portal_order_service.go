package service

import (
	"gorm.io/gorm"
	"mall/global/dao/mapper"
	"mall/global/dao/model"
	time "time"
)

type omsPortalOrderService struct {
	portalOrderMapper mapper.PortalOrderMapper
	db                *gorm.DB
}

func NewOmsPortalOrderService(m mapper.PortalOrderMapper,
	db *gorm.DB) OmsPortalOrderService {
	return &omsPortalOrderService{
		portalOrderMapper: m,
		db:                db,
	}
}

func (s *omsPortalOrderService) CancelTimeOutOrder() (count int, err error) {

	var orderSetting model.OmsOrderSetting
	s.db.First(&orderSetting)
	timeOutOrders, err := s.portalOrderMapper.GetTimeOutOrders(*orderSetting.NormalOrderOvertime)

	if err != nil {
		return 0, err
	}
	if timeOutOrders == nil && len(timeOutOrders) == 0 {
		return count, nil
	}
	ids := make([]int, len(timeOutOrders))
	for i, order := range timeOutOrders {
		ids[i] = order.Id
	}

	err = s.portalOrderMapper.UpdateOrderStatus(ids, 4)
	if err != nil {
		return 0, err
	}
	for _, order := range timeOutOrders {
		s.portalOrderMapper.ReleaseSkuStockLock(order.OrderItemList)
		s.UpdateCouponStatus(*order.CouponId, order.MemberId, 0)
		if *order.UseIntegration != 0 {
			var user model.UmsMember
			s.db.First(&user, order.MemberId)
			*user.Integration = *user.Integration + *order.Integration
			s.db.Save(&user)
		}
	}
	count = len(timeOutOrders)
	return
}

func (s omsPortalOrderService) UpdateCouponStatus(couponId int, memberId int, useStatus int) {
	if couponId == 0 {
		return
	}
	if useStatus == 0 {
		useStatus = 1
	} else {
		useStatus = 0
	}

	var coupons []model.SmsCouponHistory
	s.db.Model(&model.SmsCouponHistory{}).Where(&model.SmsCouponHistory{MemberId: &memberId,
		CouponId: &couponId, UseStatus: &useStatus,
	}).Find(&coupons)

	if coupons != nil && len(coupons) != 0 {
		coupon := coupons[0]
		t := time.Now()
		coupon.UseTime = &t
		coupon.UseStatus = &useStatus
		s.db.Save(&coupon)
	}

}
