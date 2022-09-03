package domain

import (
	"time"
)

type MemberReadHistory struct {
	Id              string    `bson:"_id"`
	MemberId        int64     `bson:"memberId"`
	MemberNickname  string    `bson:"memberNickname"`
	MemberIcon      string    `bson:"memberIcon"`
	ProductId       int       `bson:"productId"`
	ProductName     string    `bson:"productName"`
	ProductPic      string    `bson:"productPic"`
	ProductSubTitle string    `bson:"productSubTitle"`
	ProductPrice    string    `bson:"productPrice"`
	CreateTime      time.Time `bson:"createTime"`
}
