package dto

type UmsMemberRole struct {
	Name string
	Role []UmsRolePermission
}

type UmsRolePermission struct {
	Role       string
	Permission []string
}
