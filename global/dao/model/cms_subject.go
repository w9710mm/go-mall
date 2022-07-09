package model

import "time"

//专题表
type CmsSubject struct {
	Id              int        `gorm:"column:id" json:"id" `                             //是否可空:NO
	CategoryId      *int       `gorm:"column:category_id" json:"category_id" `           //是否可空:YES
	Title           *string    `gorm:"column:title" json:"title" `                       //是否可空:YES
	Pic             *string    `gorm:"column:pic" json:"pic" `                           //是否可空:YES 专题主图
	ProductCount    *int       `gorm:"column:product_count" json:"product_count" `       //是否可空:YES 关联产品数量
	RecommendStatus *int       `gorm:"column:recommend_status" json:"recommend_status" ` //是否可空:YES
	CreateTime      *time.Time `gorm:"column:create_time" json:"create_time" `           //是否可空:YES
	CollectCount    *int       `gorm:"column:collect_count" json:"collect_count" `       //是否可空:YES
	ReadCount       *int       `gorm:"column:read_count" json:"read_count" `             //是否可空:YES
	CommentCount    *int       `gorm:"column:comment_count" json:"comment_count" `       //是否可空:YES
	AlbumPics       *string    `gorm:"column:album_pics" json:"album_pics" `             //是否可空:YES 画册图片用逗号分割
	Description     *string    `gorm:"column:description" json:"description" `           //是否可空:YES
	ShowStatus      *int       `gorm:"column:show_status" json:"show_status" `           //是否可空:YES 显示状态：0->不显示；1->显示
	Content         *string    `gorm:"column:content" json:"content" `                   //是否可空:YES
	ForwardCount    *int       `gorm:"column:forward_count" json:"forward_count" `       //是否可空:YES 转发数
	CategoryName    *string    `gorm:"column:category_name" json:"category_name" `       //是否可空:YES 专题分类名称
}

func (*CmsSubject) TableName() string {
	return "cms_subject"
}
