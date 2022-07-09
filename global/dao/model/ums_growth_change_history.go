package model

import "time"

//成长值变化历史记录表
type UmsGrowthChangeHistory struct {
	Id          int        `gorm:"column:id" json:"id" `                     //是否可空:NO
	MemberId    *int       `gorm:"column:member_id" json:"member_id" `       //是否可空:YES
	CreateTime  *time.Time `gorm:"column:create_time" json:"create_time" `   //是否可空:YES
	ChangeType  *int       `gorm:"column:change_type" json:"change_type" `   //是否可空:YES 改变类型：0->增加；1->减少
	ChangeCount *int       `gorm:"column:change_count" json:"change_count" ` //是否可空:YES 积分改变数量
	OperateMan  *string    `gorm:"column:operate_man" json:"operate_man" `   //是否可空:YES 操作人员
	OperateNote *string    `gorm:"column:operate_note" json:"operate_note" ` //是否可空:YES 操作备注
	SourceType  *int       `gorm:"column:source_type" json:"source_type" `   //是否可空:YES 积分来源：0->购物；1->管理员修改
}

func (*UmsGrowthChangeHistory) TableName() string {
	return "ums_growth_change_history"
}
