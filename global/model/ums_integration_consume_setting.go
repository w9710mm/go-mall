package model

//积分消费设置
type UmsIntegrationConsumeSetting struct {
	Id                 int   `gorm:"column:id" json:"id" `                                       //是否可空:NO
	DeductionPerAmount *int  `gorm:"column:deduction_per_amount" json:"deduction_per_amount" `   //是否可空:YES 每一元需要抵扣的积分数量
	MaxPercentPerOrder *int  `gorm:"column:max_percent_per_order" json:"max_percent_per_order" ` //是否可空:YES 每笔订单最高抵用百分比
	UseUnit            int64 `gorm:"column:use_unit" json:"use_unit" `                           //是否可空:YES 每次使用积分最小单位100
	CouponStatus       *int  `gorm:"column:coupon_status" json:"coupon_status" `                 //是否可空:YES 是否可以和优惠券同用；0->不可以；1->可以
}

func (*UmsIntegrationConsumeSetting) TableName() string {
	return "ums_integration_consume_setting"
}
