package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mall/global/dao/nosql"
	"mall/global/domain"
	"mall/global/model"
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

func (s *memberReadHistoryService) Create(member model.UmsMember, h domain.MemberReadHistory) (err error) {

	h.MemberId = member.Id
	if member.Nickname != nil {
		h.MemberNickname = *member.Nickname
	}

	h.Id = primitive.NewObjectID().Hex()
	h.CreateTime = time.Now()

	_, err = s.memberReadHistoryRepository.Save(context.TODO(), h)

	return
}

func (s *memberReadHistoryService) Delete(ids []string) (count int64, err error) {
	Newids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		Newids[i], _ = primitive.ObjectIDFromHex(id)
	}
	d := bson.D{{"_id", bson.D{{"$in", ids}}}}
	count, err = s.memberReadHistoryRepository.Delete(context.TODO(), d)
	return
}

func (s *memberReadHistoryService) List(memberID int64, pageNum int, pageSize int) (
	p *nosql.PaginatedData[domain.MemberReadHistory], err error) {
	page := &nosql.PagingQuery[domain.MemberReadHistory]{}

	page.Page(int64(pageNum)).Limit(int64(pageSize)).Context(context.TODO()).
		Sort("createTime", -1).Filter(bson.D{{"memberId", memberID}})

	return s.memberReadHistoryRepository.List(page)

}

func (s *memberReadHistoryService) Clear(memberID int64) (count int64, err error) {
	d := bson.D{{"memberId", memberID}}
	return s.memberReadHistoryRepository.Delete(context.TODO(), d)

}
