package model

import "time"

//限时购场次表
type SmsFlashPromotionSession struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO 编号
	Name       string    `gorm:"column:name" json:"name" `               //是否可空:YES 场次名称
	StartTime  time.Time `gorm:"column:start_time" json:"start_time" `   //是否可空:YES 每日开始时间
	EndTime    time.Time `gorm:"column:end_time" json:"end_time" `       //是否可空:YES 每日结束时间
	Status     int       `gorm:"column:status" json:"status" `           //是否可空:YES 启用状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES 创建时间
}

func (*SmsFlashPromotionSession) TableName() string {
	return "sms_flash_promotion_session"
}
