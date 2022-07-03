package model

//后台用户角色和权限关系表
type UmsRolePermissionRelation struct {
	Id           int `gorm:"column:id" json:"id" `                       //是否可空:NO
	RoleId       int `gorm:"column:role_id" json:"role_id" `             //是否可空:YES
	PermissionId int `gorm:"column:permission_id" json:"permission_id" ` //是否可空:YES
}

func (*UmsRolePermissionRelation) TableName() string {
	return "ums_role_permission_relation"
}
