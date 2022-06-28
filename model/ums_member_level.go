package model

//会员等级表
type UmsMemberLevel struct {
	Id                    int      `gorm:"column:id" json:"id" `                                           //是否可空:NO
	Name                  *string  `gorm:"column:name" json:"name" `                                       //是否可空:YES
	GrowthPoint           *int     `gorm:"column:growth_point" json:"growth_point" `                       //是否可空:YES
	DefaultStatus         *int     `gorm:"column:default_status" json:"default_status" `                   //是否可空:YES 是否为默认等级：0->不是；1->是
	FreeFreightPoint      *float64 `gorm:"column:free_freight_point" json:"free_freight_point" `           //是否可空:YES 免运费标准
	CommentGrowthPoint    *int     `gorm:"column:comment_growth_point" json:"comment_growth_point" `       //是否可空:YES 每次评价获取的成长值
	PriviledgeFreeFreight *int     `gorm:"column:priviledge_free_freight" json:"priviledge_free_freight" ` //是否可空:YES 是否有免邮特权
	PriviledgeSignIn      *int     `gorm:"column:priviledge_sign_in" json:"priviledge_sign_in" `           //是否可空:YES 是否有签到特权
	PriviledgeComment     *int     `gorm:"column:priviledge_comment" json:"priviledge_comment" `           //是否可空:YES 是否有评论获奖励特权
	PriviledgePromotion   *int     `gorm:"column:priviledge_promotion" json:"priviledge_promotion" `       //是否可空:YES 是否有专享活动特权
	PriviledgeMemberPrice *int     `gorm:"column:priviledge_member_price" json:"priviledge_member_price" ` //是否可空:YES 是否有会员价格特权
	PriviledgeBirthday    *int     `gorm:"column:priviledge_birthday" json:"priviledge_birthday" `         //是否可空:YES 是否有生日特权
	Note                  *string  `gorm:"column:note" json:"note" `                                       //是否可空:YES
}

func (*UmsMemberLevel) TableName() string {
	return "ums_member_level"
}
