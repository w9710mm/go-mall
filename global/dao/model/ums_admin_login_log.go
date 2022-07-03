package model

import "time"

//后台用户登录日志表
type UmsAdminLoginLog struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	AdminId    int       `gorm:"column:admin_id" json:"admin_id" `       //是否可空:YES
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES
	Ip         string    `gorm:"column:ip" json:"ip" `                   //是否可空:YES
	Address    string    `gorm:"column:address" json:"address" `         //是否可空:YES
	UserAgent  string    `gorm:"column:user_agent" json:"user_agent" `   //是否可空:YES 浏览器登录类型
}

func (*UmsAdminLoginLog) TableName() string {
	return "ums_admin_login_log"
}
