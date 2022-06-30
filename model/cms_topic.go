package model

import "time"

//话题表
type CmsTopic struct {
	Id             int       `gorm:"column:id" json:"id" `                           //是否可空:NO
	CategoryId     int       `gorm:"column:category_id" json:"category_id" `         //是否可空:YES
	Name           string    `gorm:"column:name" json:"name" `                       //是否可空:YES
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time" `         //是否可空:YES
	StartTime      time.Time `gorm:"column:start_time" json:"start_time" `           //是否可空:YES
	EndTime        time.Time `gorm:"column:end_time" json:"end_time" `               //是否可空:YES
	AttendCount    int       `gorm:"column:attend_count" json:"attend_count" `       //是否可空:YES 参与人数
	AttentionCount int       `gorm:"column:attention_count" json:"attention_count" ` //是否可空:YES 关注人数
	ReadCount      int       `gorm:"column:read_count" json:"read_count" `           //是否可空:YES
	AwardName      string    `gorm:"column:award_name" json:"award_name" `           //是否可空:YES 奖品名称
	AttendType     string    `gorm:"column:attend_type" json:"attend_type" `         //是否可空:YES 参与方式
	Content        string    `gorm:"column:content" json:"content" `                 //是否可空:YES 话题内容
}

func (*CmsTopic) TableName() string {
	return "cms_topic"
}
