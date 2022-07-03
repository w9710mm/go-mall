package model

//首页推荐专题表
type SmsHomeRecommendSubject struct {
	Id              int    `gorm:"column:id" json:"id" `                             //是否可空:NO
	SubjectId       int    `gorm:"column:subject_id" json:"subject_id" `             //是否可空:YES
	SubjectName     string `gorm:"column:subject_name" json:"subject_name" `         //是否可空:YES
	RecommendStatus int    `gorm:"column:recommend_status" json:"recommend_status" ` //是否可空:YES
	Sort            int    `gorm:"column:sort" json:"sort" `                         //是否可空:YES
}

func (*SmsHomeRecommendSubject) TableName() string {
	return "sms_home_recommend_subject"
}
