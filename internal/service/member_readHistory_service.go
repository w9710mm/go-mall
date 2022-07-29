package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"mall/global/dao/domain"
	"mall/global/dao/nosql"
	"time"
)

type memberReadHistoryService struct {
}

var MemberReadHistoryService = new(memberReadHistoryService)
var memberReadHistoryRepository = nosql.MemberReadHistoryRepository

func (s memberReadHistoryService) Create(tokenString string, h domain.MemberReadHistory) (id string, err error) {

	member, err := UmsMemberService.GetCurrentMember(tokenString)
	if err != nil {
		return
	}

	h.MemberId = member.Id
	h.MemberNick = *member.Nickname
	h.MemberIcon = *member.Icon
	h.Id = ""
	h.CreateTime = time.Now()
	result, err := memberReadHistoryRepository.Save(context.TODO(), h)
	if err != nil {
		return "", err
	}
	id = result.InsertedID.(string)
	return
}

func (s memberReadHistoryService) Delete(ids []string) (count int64, err error) {
	d := bson.D{{Key: "_id", Value: bson.M{"$in": ids}}}
	count, err = memberReadHistoryRepository.Delete(context.TODO(), d)
	return
}

func (s memberReadHistoryService) List(pageNum int, pageSize int) {
	nosql.PaginatedData[]{}
}
