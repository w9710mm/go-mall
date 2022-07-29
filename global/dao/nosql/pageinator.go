package nosql

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

// Error constants
const (
	PageLimitError         = "page or limit cannot be less than 0"
	DecodeEmptyError       = "struct should be provide to decode data"
	DecodeNotAvail         = "this feature is not available for aggregate query"
	FilterInAggregateError = "you cannot use filter in aggregate query but you can pass multiple filter as param in aggregate function"
	NilFilterError         = "filter query cannot be nil"
)

// PagingQuery struct for holding mongo
// connection, filter needed to apply
// filter data with page, limit, sort key
// and sort value
type PagingQuery[T any] struct {
	collection  *mongo.Collection
	sortFields  bson.D
	ctx         context.Context
	data        paginatedData[T]
	project     interface{}
	filterQuery interface{}
	limitCount  int64
	pageCount   int64
	collation   *options.Collation
}

// AutoGenerated is to bind Aggregate query result data
type autoGenerated struct {
	Total []struct {
		Count int64 `json:"count"`
	} `json:"total"`
	Data []bson.Raw `json:"data"`
}

// Paginator struct for holding pagination info
type Paginator struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int64 `json:"total_page"`
	Offset      int64 `json:"offset"`
	Limit       int64 `json:"limit"`
	Page        int64 `json:"page"`
	PrevPage    int64 `json:"prev_page"`
	NextPage    int64 `json:"next_page"`
}

// PaginationData struct for returning pagination stat
type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

// SetCollation is function to set collation for mongo
func (paging *PagingQuery[T]) SetCollation(collation *options.Collation) *PagingQuery[T] {
	paging.collation = collation
	return paging
}

// SetCollation is function to set collation for mongo
func (paging *PagingQuery[T]) SetCollection(collection *mongo.Collection) *PagingQuery[T] {
	paging.collection = collection
	return paging
}

func (paging *PagingQuery[T]) Context(ctx context.Context) *PagingQuery[T] {
	paging.ctx = ctx
	return paging
}

// Select helps you to add projection on query
func (paging *PagingQuery[T]) Select(selector interface{}) *PagingQuery[T] {
	paging.project = selector
	return paging
}

// Filter function is to add filter for mongo query
func (paging *PagingQuery[T]) Filter(criteria interface{}) *PagingQuery[T] {
	paging.filterQuery = criteria
	return paging
}

// Limit is to add limit for pagination
func (paging *PagingQuery[T]) Limit(limit int64) *PagingQuery[T] {
	if limit < 1 {
		paging.limitCount = 10
	} else {
		paging.limitCount = limit
	}
	return paging
}

// Page is to specify which page to serve in mongo paginated result
func (paging *PagingQuery[T]) Page(page int64) *PagingQuery[T] {
	if page < 1 {
		paging.pageCount = 1
	} else {
		paging.pageCount = page
	}
	return paging
}

// Sort is to sor mongo result by certain key
func (paging *PagingQuery[T]) Sort(sortField string, sortValue interface{}) *PagingQuery[T] {
	sortQuery := bson.E{}
	sortQuery.Key = sortField
	sortQuery.Value = sortValue
	paging.sortFields = append(paging.sortFields, sortQuery)
	return paging
}

// validateQuery query is to check if user has added certain required params or not
func (paging *PagingQuery[T]) validateQuery(isNormal bool) error {
	if paging.limitCount <= 0 || paging.pageCount <= 0 {
		return errors.New(PageLimitError)
	}
	return nil
}

func (paging *PagingQuery[T]) getContext() context.Context {
	if paging.ctx != nil {
		return paging.ctx
	} else {
		return context.Background()
	}
}

// Aggregate help you to paginate mongo pipeline query
// it returns PaginatedData struct and  error if any error
// occurs during document query
func (paging *PagingQuery[T]) Aggregate(filters ...interface{}) (paginatedData *paginatedData[T], err error) {
	// checking if user added required params
	if err := paging.validateQuery(false); err != nil {
		return nil, err
	}
	if paging.filterQuery != nil {
		return nil, errors.New(FilterInAggregateError)
	}

	var aggregationFilter []bson.M
	// combining user sent queries
	for _, filter := range filters {
		aggregationFilter = append(aggregationFilter, filter.(bson.M))
	}
	skip := getSkip(paging.pageCount, paging.limitCount)
	var facetData []bson.M
	if len(paging.sortFields) > 0 {
		facetData = append(facetData, bson.M{"$sort": paging.sortFields})
	}
	facetData = append(facetData, bson.M{"$skip": skip})
	facetData = append(facetData, bson.M{"$limit": paging.limitCount})

	//if paging.SortField != "" {
	//	facetData = append(facetData, bson.M{"$sort": bson.M{paging.SortField: paging.SortValue}})
	//}
	// making facet aggregation pipeline for result and total document count
	facet := bson.M{"$facet": bson.M{
		"data":  facetData,
		"total": []bson.M{{"$count": "count"}},
	},
	}
	aggregationFilter = append(aggregationFilter, facet)
	diskUse := true
	opt := &options.AggregateOptions{
		AllowDiskUse: &diskUse,
	}
	ctx := paging.getContext()
	cursor, err := paging.collection.Aggregate(ctx, aggregationFilter, opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var docs []autoGenerated
	for cursor.Next(ctx) {
		var document *autoGenerated
		if err := cursor.Decode(&document); err == nil {
			docs = append(docs, *document)
		}
	}

	var data []bson.Raw
	var aggCount int64

	if len(docs) > 0 && len(docs[0].Data) > 0 {
		aggCount = docs[0].Total[0].Count
		data = docs[0].Data
	}

	paginationInfoChan := make(chan *Paginator, 1)
	paging.Paging(paginationInfoChan, true, aggCount)
	paginationInfo := <-paginationInfoChan
	dec := paging.data.Data
	var agg []T
	for i, raw := range data {

		if marshallErr := bson.Unmarshal(raw, &dec[i]); marshallErr == nil {
			agg = append(agg, agg[i])
		}

	}
	paging.data.Pagination = *paginationInfo.PaginationData()

	return &paging.data, nil
}

// Find returns two value pagination data with document queried from mongodb and
// error if any error occurs during document query
func (paging *PagingQuery[T]) Find() (paginatedData *paginatedData[T], err error) {
	if err := paging.validateQuery(true); err != nil {
		return nil, err
	}
	if paging.filterQuery == nil {
		return nil, errors.New(NilFilterError)
	}
	// get Pagination Info
	paginationInfoChan := make(chan *Paginator, 1)
	paging.Paging(paginationInfoChan, false, 0)

	// set options for sorting and skipping
	skip := getSkip(paging.pageCount, paging.limitCount)
	opt := &options.FindOptions{
		Skip:  &skip,
		Limit: &paging.limitCount,
	}
	if paging.project != nil {
		opt.SetProjection(paging.project)
	}
	if len(paging.sortFields) > 0 {
		opt.SetSort(paging.sortFields)
	}
	if paging.collation != nil {
		opt.SetCollation(paging.collation)
	}

	ctx := paging.getContext()
	cursor, err := paging.collection.Find(ctx, paging.filterQuery, opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	docs := paging.data
	err = cursor.All(ctx, docs)
	if err != nil {
		return nil, err
	}

	paginationInfo := <-paginationInfoChan

	paging.data.Pagination = *paginationInfo.PaginationData()

	return &paging.data, nil
}

// PaginatedData struct holds data and
// pagination detail
type paginatedData[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationData `json:"pagination"`
}

// getSkip return calculated skip value for query
func getSkip(page, limit int64) int64 {
	page--
	skip := page * limit

	if skip <= 0 {
		skip = 0
	}

	return skip
}

// PaginationData returns PaginationData struct which
// holds information of all stats needed for pagination
func (p *Paginator) PaginationData() *PaginationData {
	data := PaginationData{
		Total:     p.TotalRecord,
		Page:      p.Page,
		PerPage:   p.Limit,
		Prev:      0,
		Next:      0,
		TotalPage: p.TotalPage,
	}
	if p.Page != p.PrevPage && p.TotalRecord > 0 {
		data.Prev = p.PrevPage
	}
	if p.Page != p.NextPage && p.TotalRecord > 0 && p.Page <= p.TotalPage {
		data.Next = p.NextPage
	}

	return &data
}

// Paging returns Paginator struct which hold pagination
// stats
func (p *PagingQuery[T]) Paging(paginationInfo chan<- *Paginator, aggregate bool, aggCount int64) {
	var paginator Paginator
	var offset int64
	var count int64
	ctx := p.getContext()
	if !aggregate {
		count, _ = p.collection.CountDocuments(ctx, p.filterQuery)
	} else {
		count = aggCount
	}

	if p.pageCount > 0 {
		offset = (p.pageCount - 1) * p.limitCount
	} else {
		offset = 0
	}
	paginator.TotalRecord = count
	paginator.Page = p.pageCount
	paginator.Offset = offset
	paginator.Limit = p.limitCount
	paginator.TotalPage = int64(math.Ceil(float64(count) / float64(p.limitCount)))
	if p.pageCount > 1 {
		paginator.PrevPage = p.pageCount - 1
	} else {
		paginator.PrevPage = p.pageCount
	}
	if p.pageCount == paginator.TotalPage {
		paginator.NextPage = p.pageCount
	} else {
		paginator.NextPage = p.pageCount + 1
	}
	paginationInfo <- &paginator
}