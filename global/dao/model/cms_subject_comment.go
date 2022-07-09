package model

import "time"

//专题评论表
type CmsSubjectComment struct {
	Id             int        `gorm:"column:id" json:"id" `                             //是否可空:NO
	SubjectId      *int       `gorm:"column:subject_id" json:"subject_id" `             //是否可空:YES
	MemberNickName *string    `gorm:"column:member_nick_name" json:"member_nick_name" ` //是否可空:YES
	MemberIcon     *string    `gorm:"column:member_icon" json:"member_icon" `           //是否可空:YES
	Content        *string    `gorm:"column:content" json:"content" `                   //是否可空:YES
	CreateTime     *time.Time `gorm:"column:create_time" json:"create_time" `           //是否可空:YES
	ShowStatus     *int       `gorm:"column:show_status" json:"show_status" `           //是否可空:YES
}

func (*CmsSubjectComment) TableName() string {
	return "cms_subject_comment"
}
