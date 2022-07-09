package model

import "time"

//首页轮播广告表
type SmsHomeAdvertise struct {
	Id         int        `gorm:"column:id" json:"id" `                   //是否可空:NO
	Name       *string    `gorm:"column:name" json:"name" `               //是否可空:YES
	Type       *int       `gorm:"column:type" json:"type" `               //是否可空:YES 轮播位置：0->PC首页轮播；1->app首页轮播
	Pic        *string    `gorm:"column:pic" json:"pic" `                 //是否可空:YES
	StartTime  *time.Time `gorm:"column:start_time" json:"start_time" `   //是否可空:YES
	EndTime    *time.Time `gorm:"column:end_time" json:"end_time" `       //是否可空:YES
	Status     *int       `gorm:"column:status" json:"status" `           //是否可空:YES 上下线状态：0->下线；1->上线
	ClickCount *int       `gorm:"column:click_count" json:"click_count" ` //是否可空:YES 点击数
	OrderCount *int       `gorm:"column:order_count" json:"order_count" ` //是否可空:YES 下单数
	Url        *string    `gorm:"column:url" json:"url" `                 //是否可空:YES 链接地址
	Note       *string    `gorm:"column:note" json:"note" `               //是否可空:YES 备注
	Sort       *int       `gorm:"column:sort" json:"sort" `               //是否可空:YES 排序
}

func (*SmsHomeAdvertise) TableName() string {
	return "sms_home_advertise"
}
