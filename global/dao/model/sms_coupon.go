package model

import "time"

//优惠卷表
type SmsCoupon struct {
	Id           int        `gorm:"column:id" json:"id" `                       //是否可空:NO
	Type         *int       `gorm:"column:type" json:"type" `                   //是否可空:YES 优惠卷类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	Name         *string    `gorm:"column:name" json:"name" `                   //是否可空:YES
	Platform     *int       `gorm:"column:platform" json:"platform" `           //是否可空:YES 使用平台：0->全部；1->移动；2->PC
	Count        *int       `gorm:"column:count" json:"count" `                 //是否可空:YES 数量
	Amount       *float64   `gorm:"column:amount" json:"amount" `               //是否可空:YES 金额
	PerLimit     *int       `gorm:"column:per_limit" json:"per_limit" `         //是否可空:YES 每人限领张数
	MinPoint     *float64   `gorm:"column:min_point" json:"min_point" `         //是否可空:YES 使用门槛；0表示无门槛
	StartTime    *time.Time `gorm:"column:start_time" json:"start_time" `       //是否可空:YES
	EndTime      *time.Time `gorm:"column:end_time" json:"end_time" `           //是否可空:YES
	UseType      *int       `gorm:"column:use_type" json:"use_type" `           //是否可空:YES 使用类型：0->全场通用；1->指定分类；2->指定商品
	Note         *string    `gorm:"column:note" json:"note" `                   //是否可空:YES 备注
	PublishCount *int       `gorm:"column:publish_count" json:"publish_count" ` //是否可空:YES 发行数量
	UseCount     *int       `gorm:"column:use_count" json:"use_count" `         //是否可空:YES 已使用数量
	ReceiveCount *int       `gorm:"column:receive_count" json:"receive_count" ` //是否可空:YES 领取数量
	EnableTime   *time.Time `gorm:"column:enable_time" json:"enable_time" `     //是否可空:YES 可以领取的日期
	Code         *string    `gorm:"column:code" json:"code" `                   //是否可空:YES 优惠码
	MemberLevel  *int       `gorm:"column:member_level" json:"member_level" `   //是否可空:YES 可领取的会员类型：0->无限时
}

func (*SmsCoupon) TableName() string {
	return "sms_coupon"
}
