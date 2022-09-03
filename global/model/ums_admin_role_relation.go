package model

//后台用户和角色关系表
type UmsAdminRoleRelation struct {
	Id      int  `gorm:"column:id" json:"id" `             //是否可空:NO
	AdminId *int `gorm:"column:admin_id" json:"admin_id" ` //是否可空:YES
	RoleId  *int `gorm:"column:role_id" json:"role_id" `   //是否可空:YES
}

func (*UmsAdminRoleRelation) TableName() string {
	return "ums_admin_role_relation"
}
