package dto

import "mall/model"

type AdminUserDetails struct {
	UmsAdmin       model.UmsAdmin
	PermissionList []model.UmsPermission
}
