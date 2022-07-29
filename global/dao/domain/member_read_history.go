package domain

import "time"

type MemberReadHistory struct {
	Id               string
	MemberId         int64
	MemberNick       string
	MemberIcon       string
	ProductId        int
	ProductName      string
	ProductPic       string
	ProductSubTittle string
	ProductPrice     string
	CreateTime       time.Time
}
