package model

import "time"

//会员登录记录
type UmsMemberLoginLog struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	MemberId   int       `gorm:"column:member_id" json:"member_id" `     //是否可空:YES
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES
	Ip         string    `gorm:"column:ip" json:"ip" `                   //是否可空:YES
	City       string    `gorm:"column:city" json:"city" `               //是否可空:YES
	LoginType  int       `gorm:"column:login_type" json:"login_type" `   //是否可空:YES 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string    `gorm:"column:province" json:"province" `       //是否可空:YES
}

func (*UmsMemberLoginLog) TableName() string {
	return "ums_member_login_log"
}
