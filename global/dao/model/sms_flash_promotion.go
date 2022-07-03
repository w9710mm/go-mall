package model

import "time"

//限时购表
type SmsFlashPromotion struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	Title      string    `gorm:"column:title" json:"title" `             //是否可空:YES
	StartDate  time.Time `gorm:"column:start_date" json:"start_date" `   //是否可空:YES 开始日期
	EndDate    time.Time `gorm:"column:end_date" json:"end_date" `       //是否可空:YES 结束日期
	Status     int       `gorm:"column:status" json:"status" `           //是否可空:YES 上下线状态
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES 秒杀时间段名称
}

func (*SmsFlashPromotion) TableName() string {
	return "sms_flash_promotion"
}
