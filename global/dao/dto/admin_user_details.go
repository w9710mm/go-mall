package dto

import "mall/global/dao/model"

type AdminUserDetails struct {
	UmsAdmin       model.UmsAdmin
	PermissionList []model.UmsPermission
}
