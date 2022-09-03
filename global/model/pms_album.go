package model

//相册表
type PmsAlbum struct {
	Id          int     `gorm:"column:id" json:"id" `                   //是否可空:NO
	Name        *string `gorm:"column:name" json:"name" `               //是否可空:YES
	CoverPic    *string `gorm:"column:cover_pic" json:"cover_pic" `     //是否可空:YES
	PicCount    *int    `gorm:"column:pic_count" json:"pic_count" `     //是否可空:YES
	Sort        *int    `gorm:"column:sort" json:"sort" `               //是否可空:YES
	Description *string `gorm:"column:description" json:"description" ` //是否可空:YES
}

func (*PmsAlbum) TableName() string {
	return "pms_album"
}
