package model

//后台用户和权限关系表(除角色中定义的权限以外的加减权限)
type UmsAdminPermissionRelation struct {
	Id           int `gorm:"column:id" json:"id" `                       //是否可空:NO
	AdminId      int `gorm:"column:admin_id" json:"admin_id" `           //是否可空:YES
	PermissionId int `gorm:"column:permission_id" json:"permission_id" ` //是否可空:YES
	Type         int `gorm:"column:type" json:"type" `                   //是否可空:YES
}

func (*UmsAdminPermissionRelation) TableName() string {
	return "ums_admin_permission_relation"
}
