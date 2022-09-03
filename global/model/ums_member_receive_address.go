package model

//会员收货地址表
type UmsMemberReceiveAddress struct {
	Id            int64   `gorm:"column:id" json:"id" `                         //是否可空:NO
	MemberId      *int    `gorm:"column:member_id" json:"member_id" `           //是否可空:YES
	Name          *string `gorm:"column:name" json:"name" `                     //是否可空:YES 收货人名称
	PhoneNumber   *string `gorm:"column:phone_number" json:"phone_number" `     //是否可空:YES
	DefaultStatus *int    `gorm:"column:default_status" json:"default_status" ` //是否可空:YES 是否为默认
	PostCode      *string `gorm:"column:post_code" json:"post_code" `           //是否可空:YES 邮政编码
	Province      *string `gorm:"column:province" json:"province" `             //是否可空:YES 省份/直辖市
	City          *string `gorm:"column:city" json:"city" `                     //是否可空:YES 城市
	Region        *string `gorm:"column:region" json:"region" `                 //是否可空:YES 区
	DetailAddress *string `gorm:"column:detail_address" json:"detail_address" ` //是否可空:YES 详细地址(街道)
}

func (*UmsMemberReceiveAddress) TableName() string {
	return "ums_member_receive_address"
}
