package dto

import "mall/global/model"

type AdminUserDetails struct {
	UmsAdmin       model.UmsAdmin
	PermissionList []model.UmsPermission
}
