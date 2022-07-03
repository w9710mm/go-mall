package model

//用户和标签关系表
type UmsMemberMemberTagRelation struct {
	Id       int `gorm:"column:id" json:"id" `               //是否可空:NO
	MemberId int `gorm:"column:member_id" json:"member_id" ` //是否可空:YES
	TagId    int `gorm:"column:tag_id" json:"tag_id" `       //是否可空:YES
}

func (*UmsMemberMemberTagRelation) TableName() string {
	return "ums_member_member_tag_relation"
}
