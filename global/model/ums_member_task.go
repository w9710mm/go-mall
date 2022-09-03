package model

//会员任务表
type UmsMemberTask struct {
	Id           int     `gorm:"column:id" json:"id" `                     //是否可空:NO
	Name         *string `gorm:"column:name" json:"name" `                 //是否可空:YES
	Growth       *int    `gorm:"column:growth" json:"growth" `             //是否可空:YES 赠送成长值
	Intergration *int    `gorm:"column:intergration" json:"intergration" ` //是否可空:YES 赠送积分
	Type         *int    `gorm:"column:type" json:"type" `                 //是否可空:YES 任务类型：0->新手任务；1->日常任务
}

func (*UmsMemberTask) TableName() string {
	return "ums_member_task"
}
