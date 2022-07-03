package model

import "time"

//商品评价表
type PmsComment struct {
	Id               int       `gorm:"column:id" json:"id" `                               //是否可空:NO
	ProductId        int       `gorm:"column:product_id" json:"product_id" `               //是否可空:YES
	MemberNickName   string    `gorm:"column:member_nick_name" json:"member_nick_name" `   //是否可空:YES
	ProductName      string    `gorm:"column:product_name" json:"product_name" `           //是否可空:YES
	Star             int       `gorm:"column:star" json:"star" `                           //是否可空:YES 评价星数：0->5
	MemberIp         string    `gorm:"column:member_ip" json:"member_ip" `                 //是否可空:YES 评价的ip
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time" `             //是否可空:YES
	ShowStatus       int       `gorm:"column:show_status" json:"show_status" `             //是否可空:YES
	ProductAttribute string    `gorm:"column:product_attribute" json:"product_attribute" ` //是否可空:YES 购买时的商品属性
	CollectCouont    int       `gorm:"column:collect_couont" json:"collect_couont" `       //是否可空:YES
	ReadCount        int       `gorm:"column:read_count" json:"read_count" `               //是否可空:YES
	Content          string    `gorm:"column:content" json:"content" `                     //是否可空:YES
	Pics             string    `gorm:"column:pics" json:"pics" `                           //是否可空:YES 上传图片地址，以逗号隔开
	MemberIcon       string    `gorm:"column:member_icon" json:"member_icon" `             //是否可空:YES 评论用户头像
	ReplayCount      int       `gorm:"column:replay_count" json:"replay_count" `           //是否可空:YES
}

func (*PmsComment) TableName() string {
	return "pms_comment"
}
