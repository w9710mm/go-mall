package model

//
type SysCasbinRule struct {
	PType string `gorm:"column:p_type" json:"p_type" ` //是否可空:YES
	V0    string `gorm:"column:v0" json:"v0" `         //是否可空:YES
	V1    string `gorm:"column:v1" json:"v1" `         //是否可空:YES
	V2    string `gorm:"column:v2" json:"v2" `         //是否可空:YES
	V3    string `gorm:"column:v3" json:"v3" `         //是否可空:YES
	V4    string `gorm:"column:v4" json:"v4" `         //是否可空:YES
	V5    string `gorm:"column:v5" json:"v5" `         //是否可空:YES
}

func (*SysCasbinRule) TableName() string {
	return "sys_casbin_rule"
}
