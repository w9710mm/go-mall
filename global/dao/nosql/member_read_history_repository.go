package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mall/global/domain"
)

type memberReadHistoryRepository struct {
	Name       string
	collection *mongo.Collection
}

func NewMemberReadHistoryRepository(name string, collection *mongo.Collection) MemberReadHistoryRepository {
	return &memberReadHistoryRepository{
		Name:       name,
		collection: collection,
	}
}

func (r memberReadHistoryRepository) Save(ctx context.Context, h domain.MemberReadHistory) (result *mongo.InsertOneResult, err error) {
	getCtx(&ctx)

	return r.collection.InsertOne(context.TODO(), h)
}

func (r memberReadHistoryRepository) Delete(ctx context.Context, m interface{}) (
	count int64, err error) {
	getCtx(&ctx)

	result, err := r.collection.DeleteMany(
		ctx,
		m,
	)
	count = result.DeletedCount
	return
}
func (r memberReadHistoryRepository) List(query *PagingQuery[domain.MemberReadHistory]) (*PaginatedData[domain.MemberReadHistory], error) {
	return query.SetCollection(r.collection).Find()
}
