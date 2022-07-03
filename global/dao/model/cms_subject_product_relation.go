package model

//专题商品关系表
type CmsSubjectProductRelation struct {
	Id        int `gorm:"column:id" json:"id" `                 //是否可空:NO
	SubjectId int `gorm:"column:subject_id" json:"subject_id" ` //是否可空:YES
	ProductId int `gorm:"column:product_id" json:"product_id" ` //是否可空:YES
}

func (*CmsSubjectProductRelation) TableName() string {
	return "cms_subject_product_relation"
}
