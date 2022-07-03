package model

//专题分类表
type CmsSubjectCategory struct {
	Id           int    `gorm:"column:id" json:"id" `                       //是否可空:NO
	Name         string `gorm:"column:name" json:"name" `                   //是否可空:YES
	Icon         string `gorm:"column:icon" json:"icon" `                   //是否可空:YES 分类图标
	SubjectCount int    `gorm:"column:subject_count" json:"subject_count" ` //是否可空:YES 专题数量
	ShowStatus   int    `gorm:"column:show_status" json:"show_status" `     //是否可空:YES
	Sort         int    `gorm:"column:sort" json:"sort" `                   //是否可空:YES
}

func (*CmsSubjectCategory) TableName() string {
	return "cms_subject_category"
}
