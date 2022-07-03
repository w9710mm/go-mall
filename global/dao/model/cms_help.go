package model

import "time"

//帮助表
type CmsHelp struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	CategoryId int       `gorm:"column:category_id" json:"category_id" ` //是否可空:YES
	Icon       string    `gorm:"column:icon" json:"icon" `               //是否可空:YES
	Title      string    `gorm:"column:title" json:"title" `             //是否可空:YES
	ShowStatus int       `gorm:"column:show_status" json:"show_status" ` //是否可空:YES
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES
	ReadCount  int       `gorm:"column:read_count" json:"read_count" `   //是否可空:YES
	Content    string    `gorm:"column:content" json:"content" `         //是否可空:YES
}

func (*CmsHelp) TableName() string {
	return "cms_help"
}
