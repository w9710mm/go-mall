package model

//会员积分成长规则表
type UmsMemberRuleSetting struct {
	Id                int      `gorm:"column:id" json:"id" `                                   //是否可空:NO
	ContinueSignDay   *int     `gorm:"column:continue_sign_day" json:"continue_sign_day" `     //是否可空:YES 连续签到天数
	ContinueSignPoint *int     `gorm:"column:continue_sign_point" json:"continue_sign_point" ` //是否可空:YES 连续签到赠送数量
	ConsumePerPoint   *float64 `gorm:"column:consume_per_point" json:"consume_per_point" `     //是否可空:YES 每消费多少元获取1个点
	LowOrderAmount    *float64 `gorm:"column:low_order_amount" json:"low_order_amount" `       //是否可空:YES 最低获取点数的订单金额
	MaxPointPerOrder  *int     `gorm:"column:max_point_per_order" json:"max_point_per_order" ` //是否可空:YES 每笔订单最高获取点数
	Type              *int     `gorm:"column:type" json:"type" `                               //是否可空:YES 类型：0->积分规则；1->成长值规则
}

func (*UmsMemberRuleSetting) TableName() string {
	return "ums_member_rule_setting"
}
