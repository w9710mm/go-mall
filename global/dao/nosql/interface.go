package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mall/global/dao/domain"
)

type MemberReadHistoryRepository interface {
	Save(context.Context, domain.MemberReadHistory) (*mongo.InsertOneResult, error)
	Delete(context.Context, interface{}) (int64, error)
	List(*PagingQuery[domain.MemberReadHistory]) (*PaginatedData[domain.MemberReadHistory], error)
}

func getCtx(ctx *context.Context) {
	if ctx == nil {
		*ctx = context.TODO()
	}
}
