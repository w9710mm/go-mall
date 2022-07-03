package model

//运费模版
type PmsFeightTemplate struct {
	Id             int     `gorm:"column:id" json:"id" `                           //是否可空:NO
	Name           string  `gorm:"column:name" json:"name" `                       //是否可空:YES
	ChargeType     int     `gorm:"column:charge_type" json:"charge_type" `         //是否可空:YES 计费类型:0->按重量；1->按件数
	FirstWeight    float64 `gorm:"column:first_weight" json:"first_weight" `       //是否可空:YES 首重kg
	FirstFee       float64 `gorm:"column:first_fee" json:"first_fee" `             //是否可空:YES 首费（元）
	ContinueWeight float64 `gorm:"column:continue_weight" json:"continue_weight" ` //是否可空:YES
	ContinmeFee    float64 `gorm:"column:continme_fee" json:"continme_fee" `       //是否可空:YES
	Dest           string  `gorm:"column:dest" json:"dest" `                       //是否可空:YES 目的地（省、市）
}

func (*PmsFeightTemplate) TableName() string {
	return "pms_feight_template"
}
