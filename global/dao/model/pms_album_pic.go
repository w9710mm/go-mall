package model

//画册图片表
type PmsAlbumPic struct {
	Id      int     `gorm:"column:id" json:"id" `             //是否可空:NO
	AlbumId *int    `gorm:"column:album_id" json:"album_id" ` //是否可空:YES
	Pic     *string `gorm:"column:pic" json:"pic" `           //是否可空:YES
}

func (*PmsAlbumPic) TableName() string {
	return "pms_album_pic"
}
