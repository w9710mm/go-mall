package model

import "time"

//后台用户权限表
type UmsPermission struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	Pid        int       `gorm:"column:pid" json:"pid" `                 //是否可空:YES 父级权限id
	Name       string    `gorm:"column:name" json:"name" `               //是否可空:YES 名称
	Value      string    `gorm:"column:value" json:"value" `             //是否可空:YES 权限值
	Icon       string    `gorm:"column:icon" json:"icon" `               //是否可空:YES 图标
	Type       int       `gorm:"column:type" json:"type" `               //是否可空:YES 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
	Uri        string    `gorm:"column:uri" json:"uri" `                 //是否可空:YES 前端资源路径
	Status     int       `gorm:"column:status" json:"status" `           //是否可空:YES 启用状态；0->禁用；1->启用
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES 创建时间
	Sort       int       `gorm:"column:sort" json:"sort" `               //是否可空:YES 排序
}

func (*UmsPermission) TableName() string {
	return "ums_permission"
}
