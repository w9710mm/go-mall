package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mall/global/dao"
	"mall/global/dao/domain"
)

type memberReadHistoryRepository struct {
	name string
}

var MemberReadHistoryRepository = memberReadHistoryRepository{
	name: "Member_read_history"}

var historyCollection = dao.GetMongoDB().Collection(MemberReadHistoryRepository.name)

func (r memberReadHistoryRepository) Save(ctx context.Context, h domain.MemberReadHistory) (result *mongo.InsertOneResult, err error) {
	getCtx(&ctx)

	return historyCollection.InsertOne(context.TODO(), h)
}

func (r memberReadHistoryRepository) Delete(ctx context.Context, d interface{}) (
	count int64, err error) {
	getCtx(&ctx)

	result, err := historyCollection.DeleteMany(
		ctx,
		d,
	)
	count = result.DeletedCount
	return
}

func getCtx(ctx *context.Context) {
	if ctx == nil {
		*ctx = context.TODO()
	}
}
