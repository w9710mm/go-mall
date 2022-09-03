package model

//公司收发货地址表
type OmsCompanyAddress struct {
	Id            int     `gorm:"column:id" json:"id" `                         //是否可空:NO
	AddressName   *string `gorm:"column:address_name" json:"address_name" `     //是否可空:YES 地址名称
	SendStatus    *int    `gorm:"column:send_status" json:"send_status" `       //是否可空:YES 默认发货地址：0->否；1->是
	ReceiveStatus *int    `gorm:"column:receive_status" json:"receive_status" ` //是否可空:YES 是否默认收货地址：0->否；1->是
	Name          *string `gorm:"column:name" json:"name" `                     //是否可空:YES 收发货人姓名
	Phone         *string `gorm:"column:phone" json:"phone" `                   //是否可空:YES 收货人电话
	Province      *string `gorm:"column:province" json:"province" `             //是否可空:YES 省/直辖市
	City          *string `gorm:"column:city" json:"city" `                     //是否可空:YES 市
	Region        *string `gorm:"column:region" json:"region" `                 //是否可空:YES 区
	DetailAddress *string `gorm:"column:detail_address" json:"detail_address" ` //是否可空:YES 详细地址
}

func (*OmsCompanyAddress) TableName() string {
	return "oms_company_address"
}
