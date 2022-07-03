package model

import "time"

//限时购通知记录
type SmsFlashPromotionLog struct {
	Id            int       `gorm:"column:id" json:"id" `                         //是否可空:NO
	MemberId      int       `gorm:"column:member_id" json:"member_id" `           //是否可空:YES
	ProductId     int       `gorm:"column:product_id" json:"product_id" `         //是否可空:YES
	MemberPhone   string    `gorm:"column:member_phone" json:"member_phone" `     //是否可空:YES
	ProductName   string    `gorm:"column:product_name" json:"product_name" `     //是否可空:YES
	SubscribeTime time.Time `gorm:"column:subscribe_time" json:"subscribe_time" ` //是否可空:YES 会员订阅时间
	SendTime      time.Time `gorm:"column:send_time" json:"send_time" `           //是否可空:YES
}

func (*SmsFlashPromotionLog) TableName() string {
	return "sms_flash_promotion_log"
}
