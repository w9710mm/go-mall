package model

import "time"

//用户举报表
type CmsMemberReport struct {
	Id               int       `gorm:"column:id" json:"id" `                                 //是否可空:YES
	ReportType       int       `gorm:"column:report_type" json:"report_type" `               //是否可空:YES 举报类型：0->商品评价；1->话题内容；2->用户评论
	ReportMemberName string    `gorm:"column:report_member_name" json:"report_member_name" ` //是否可空:YES 举报人
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time" `               //是否可空:YES
	ReportObject     string    `gorm:"column:report_object" json:"report_object" `           //是否可空:YES
	ReportStatus     int       `gorm:"column:report_status" json:"report_status" `           //是否可空:YES 举报状态：0->未处理；1->已处理
	HandleStatus     int       `gorm:"column:handle_status" json:"handle_status" `           //是否可空:YES 处理结果：0->无效；1->有效；2->恶意
	Note             string    `gorm:"column:note" json:"note" `                             //是否可空:YES
}

func (*CmsMemberReport) TableName() string {
	return "cms_member_report"
}
