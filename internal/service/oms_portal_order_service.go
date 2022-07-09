package service

import (
	"mall/global/dao"
	"mall/global/dao/mapper"
	"mall/global/dao/model"
	time "time"
)

type omsPortalOrderService struct {
}

var OmsPortalOrderService = new(omsPortalOrderService)

var portalOrderMapper = mapper.PortalOrderMapper

func (s *omsPortalOrderService) CancelTimeOutOrder() (count int, err error) {

	var orderSetting model.OmsOrderSetting
	dao.DB.First(&orderSetting)
	timeOutOrders, err := portalOrderMapper.GetTimeOutOrders(*orderSetting.NormalOrderOvertime)

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

	err = portalOrderMapper.UpdateOrderStatus(ids, 4)
	if err != nil {
		return 0, err
	}
	for _, order := range timeOutOrders {
		portalOrderMapper.ReleaseSkuStockLock(order.OrderItemList)
		updateCouponStatus(*order.CouponId, order.MemberId, 0)
		if *order.UseIntegration != 0 {
			var user model.UmsMember
			dao.DB.First(&user, order.MemberId)
			*user.Integration = *user.Integration + *order.Integration
			dao.DB.Save(&user)
		}
	}
	count = len(timeOutOrders)
	return
}

func updateCouponStatus(couponId int, memberId int, useStatus int) {
	if couponId == 0 {
		return
	}
	if useStatus == 0 {
		useStatus = 1
	} else {
		useStatus = 0
	}

	var coupons []model.SmsCouponHistory
	dao.DB.Model(&model.SmsCouponHistory{}).Where(&model.SmsCouponHistory{MemberId: &memberId,
		CouponId: &couponId, UseStatus: &useStatus,
	}).Find(&coupons)

	if coupons != nil && len(coupons) != 0 {
		coupon := coupons[0]
		t := time.Now()
		coupon.UseTime = &t
		coupon.UseStatus = &useStatus
		dao.DB.Save(&coupon)
	}

}
