package model

import "time"

//产品评价回复表
type PmsCommentReplay struct {
	Id             int       `gorm:"column:id" json:"id" `                             //是否可空:NO
	CommentId      int       `gorm:"column:comment_id" json:"comment_id" `             //是否可空:YES
	MemberNickName string    `gorm:"column:member_nick_name" json:"member_nick_name" ` //是否可空:YES
	MemberIcon     string    `gorm:"column:member_icon" json:"member_icon" `           //是否可空:YES
	Content        string    `gorm:"column:content" json:"content" `                   //是否可空:YES
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time" `           //是否可空:YES
	Type           int       `gorm:"column:type" json:"type" `                         //是否可空:YES 评论人员类型；0->会员；1->管理员
}

func (*PmsCommentReplay) TableName() string {
	return "pms_comment_replay"
}
