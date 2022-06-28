package model

import "time"

//后台用户表
type UmsAdmin struct {
	Id         int        `gorm:"column:id" json:"id" `                   //是否可空:NO
	Username   *string    `gorm:"column:username" json:"username" `       //是否可空:YES
	Password   *string    `gorm:"column:password" json:"password" `       //是否可空:YES
	Icon       *string    `gorm:"column:icon" json:"icon" `               //是否可空:YES 头像
	Email      *string    `gorm:"column:email" json:"email" `             //是否可空:YES 邮箱
	NickName   *string    `gorm:"column:nick_name" json:"nick_name" `     //是否可空:YES 昵称
	Note       *string    `gorm:"column:note" json:"note" `               //是否可空:YES 备注信息
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES 创建时间
	LoginTime  *time.Time `gorm:"column:login_time" json:"login_time" `   //是否可空:YES 最后登录时间
	Status     *int       `gorm:"column:status" json:"status" `           //是否可空:YES 帐号启用状态：0->禁用；1->启用
}

func (*UmsAdmin) TableName() string {
	return "ums_admin"
}
