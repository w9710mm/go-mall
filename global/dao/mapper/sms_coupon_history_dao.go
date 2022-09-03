package mapper

import (
	"gorm.io/gorm"
	"mall/global/domain"
	"mall/global/model"
)

/**
 *@author:
 *@date:2022/8/23
**/

type smsCouponHistoryDao struct {
	db *gorm.DB
}

func NewSmsCouponHistoryDao(db *gorm.DB) SmsCouponHistoryDao {
	return &smsCouponHistoryDao{db: db}
}

func (s *smsCouponHistoryDao) GetDetailList(memberId int64) (history []domain.SmsCouponHistoryDetail) {
	s.db.Model(&model.SmsCouponHistory{}).Preload("Coupon", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name ,amount,platform ,start_time,end_time,note,use_type,type")
	}).Preload("ProductRelationList").Preload("CategoryRelationList").
		Where(map[string]interface{}{"member_id": memberId, "use_status": 0}).Find(&history)
	return
}

func (s *smsCouponHistoryDao) GetCouponList(memberId int64, useStatus int) (coupon []model.SmsCoupon) {
	tx := s.db.Model(&model.SmsCouponHistory{}).Joins("LEFT JOIN sms_coupon c ON sms_coupon_history.coupon_id = c.id").
		Where("member_id =? ", memberId)
	if useStatus != -1 && useStatus != 2 {
		tx = tx.Where("use_status=? and now() >start_time and c.end_time>now()", useStatus)
	}
	if useStatus == 2 {
		tx = tx.Where("now()>end_time")
	}
	tx.Find(&coupon)
	return
}
