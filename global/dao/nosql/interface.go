package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mall/global/domain"
)

type MemberReadHistoryRepository interface {
	Save(context.Context, domain.MemberReadHistory) (*mongo.InsertOneResult, error)
	Delete(ctx context.Context, D interface{}) (
		count int64, err error)
	List(*PagingQuery[domain.MemberReadHistory]) (*PaginatedData[domain.MemberReadHistory], error)
}

func getCtx(ctx *context.Context) {
	if ctx == nil {
		*ctx = context.TODO()
	}
}
