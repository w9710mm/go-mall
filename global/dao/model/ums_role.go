package model

import "time"

//后台用户角色表
type UmsRole struct {
	Id          int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	Name        string    `gorm:"column:name" json:"name" `               //是否可空:YES 名称
	Description string    `gorm:"column:description" json:"description" ` //是否可空:YES 描述
	AdminCount  int       `gorm:"column:admin_count" json:"admin_count" ` //是否可空:YES 后台用户数量
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES 创建时间
	Status      int       `gorm:"column:status" json:"status" `           //是否可空:YES 启用状态：0->禁用；1->启用
	Sort        int       `gorm:"column:sort" json:"sort" `               //是否可空:YES
}

func (*UmsRole) TableName() string {
	return "ums_role"
}
