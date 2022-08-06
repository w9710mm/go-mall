package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"mall/global/dao/domain"
	"mall/global/dao/model"
	"mall/global/dao/nosql"
	"time"
)

type memberReadHistoryService struct {
	memberReadHistoryRepository nosql.MemberReadHistoryRepository
}

func NewMemberReadHistoryService(service nosql.MemberReadHistoryRepository) MemberReadHistoryService {
	return &memberReadHistoryService{
		memberReadHistoryRepository: service,
	}
}

func (s *memberReadHistoryService) Create(member model.UmsMember, h domain.MemberReadHistory) (id string, err error) {

	h.MemberId = member.Id
	h.MemberNick = *member.Nickname
	h.MemberIcon = *member.Icon
	h.Id = ""
	h.CreateTime = time.Now()

	result, err := s.memberReadHistoryRepository.Save(context.TODO(), h)
	if err != nil {
		return "", err
	}
	id = result.InsertedID.(string)
	return
}

func (s *memberReadHistoryService) Delete(ids []string) (count int64, err error) {
	d := bson.D{{Key: "_id", Value: bson.M{"$in": ids}}}
	count, err = s.memberReadHistoryRepository.Delete(context.TODO(), d)
	return
}

func (s *memberReadHistoryService) List(memberID int64, pageNum int, pageSize int) (
	p *nosql.PaginatedData[domain.MemberReadHistory], err error) {
	page := &nosql.PagingQuery[domain.MemberReadHistory]{}
	if err != nil {
		return
	}
	page.Page(int64(pageNum)).Limit(int64(pageSize)).Context(context.TODO()).
		Sort("createTime", -1).Filter(bson.D{{"memberId", memberID}})

	return s.memberReadHistoryRepository.List(page)

}

func (s *memberReadHistoryService) Clear(memberID int64) (count int64, err error) {
	d := bson.D{{"memberId", memberID}}
	return s.memberReadHistoryRepository.Delete(context.TODO(), d)

}
