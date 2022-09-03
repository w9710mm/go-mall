package model

import (
	"encoding/json"
	"time"
)

//会员表
type UmsMember struct {
	Id                    int64      `gorm:"column:id" json:"id" `                                         //是否可空:NO
	MemberLevelId         *int       `gorm:"column:member_level_id" json:"member_level_id" `               //是否可空:YES
	Username              *string    `gorm:"column:username" json:"username" `                             //是否可空:YES 用户名
	Password              *string    `gorm:"column:password" json:"password" `                             //是否可空:YES 密码
	Nickname              *string    `gorm:"column:nickname" json:"nickname" `                             //是否可空:YES 昵称
	Phone                 *string    `gorm:"column:phone" json:"phone" `                                   //是否可空:YES 手机号码
	Status                *int       `gorm:"column:status" json:"status" `                                 //是否可空:YES 帐号启用状态:0->禁用；1->启用
	CreateTime            *time.Time `gorm:"column:create_time" json:"create_time" `                       //是否可空:YES 注册时间
	Icon                  *string    `gorm:"column:icon" json:"icon" `                                     //是否可空:YES 头像
	Gender                *int       `gorm:"column:gender" json:"gender" `                                 //是否可空:YES 性别：0->未知；1->男；2->女
	Birthday              *time.Time `gorm:"column:birthday" json:"birthday" `                             //是否可空:YES 生日
	City                  *string    `gorm:"column:city" json:"city" `                                     //是否可空:YES 所做城市
	Job                   *string    `gorm:"column:job" json:"job" `                                       //是否可空:YES 职业
	PersonalizedSignature *string    `gorm:"column:personalized_signature" json:"personalized_signature" ` //是否可空:YES 个性签名
	SourceType            *int       `gorm:"column:source_type" json:"source_type" `                       //是否可空:YES 用户来源
	Integration           int64      `gorm:"column:integration" json:"integration" `                       //是否可空:YES 积分
	Growth                *int       `gorm:"column:growth" json:"growth" `                                 //是否可空:YES 成长值
	LuckeyCount           *int       `gorm:"column:luckey_count" json:"luckey_count" `                     //是否可空:YES 剩余抽奖次数
	HistoryIntegration    *int       `gorm:"column:history_integration" json:"history_integration" `       //是否可空:YES 历史积分数量
}

func (*UmsMember) TableName() string {
	return "ums_member"
}
func (m *UmsMember) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
func (m *UmsMember) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}
