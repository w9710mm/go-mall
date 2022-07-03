package model

import "time"

//优惠券使用、领取历史表
type SmsCouponHistory struct {
	Id             int       `gorm:"column:id" json:"id" `                           //是否可空:NO
	CouponId       int       `gorm:"column:coupon_id" json:"coupon_id" `             //是否可空:YES
	MemberId       int       `gorm:"column:member_id" json:"member_id" `             //是否可空:YES
	CouponCode     string    `gorm:"column:coupon_code" json:"coupon_code" `         //是否可空:YES
	MemberNickname string    `gorm:"column:member_nickname" json:"member_nickname" ` //是否可空:YES 领取人昵称
	GetType        int       `gorm:"column:get_type" json:"get_type" `               //是否可空:YES 获取类型：0->后台赠送；1->主动获取
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time" `         //是否可空:YES
	UseStatus      int       `gorm:"column:use_status" json:"use_status" `           //是否可空:YES 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        time.Time `gorm:"column:use_time" json:"use_time" `               //是否可空:YES 使用时间
	OrderId        int       `gorm:"column:order_id" json:"order_id" `               //是否可空:YES 订单编号
	OrderSn        string    `gorm:"column:order_sn" json:"order_sn" `               //是否可空:YES 订单号码
}

func (*SmsCouponHistory) TableName() string {
	return "sms_coupon_history"
}
