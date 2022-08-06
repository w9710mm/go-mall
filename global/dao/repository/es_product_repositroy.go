package repository

import (
	"context"
	"encoding/json"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"mall/global/dao/document"
	"strconv"
)

type esProductRepository struct {
	esDB      *elastic.Client
	esProduct document.EsProduct
}

func NewEsProductRepository(db *elastic.Client) EsProductRepository {
	return &esProductRepository{
		esDB: db,
	}
}

func (r *esProductRepository) SaveAll(esProducts []document.EsProduct, ctx context.Context) (
	*elastic.BulkResponse, error) {
	getContext(ctx)
	bulkService := r.esDB.Bulk().Index(r.esProduct.GetIndex()).Refresh("true")
	for i, product := range esProducts {
		bulkService.Add(elastic.NewBulkCreateRequest().
			Id(strconv.Itoa(i)).
			Doc(product))
	}

	return bulkService.Do(ctx)
}

func (r *esProductRepository) DeleteById(id int, ctx context.Context) (
	*elastic.DeleteResponse, error) {
	getContext(ctx)
	return r.esDB.Delete().Index(r.esProduct.GetIndex()).Id(strconv.Itoa(id)).
		Refresh("true").Do(ctx)

}

func (r *esProductRepository) DeleteAll(ids []int, ctx context.Context) (
	*elastic.BulkResponse, error) {
	getContext(ctx)
	bulkService := r.esDB.Bulk().Index(r.esProduct.GetIndex()).Refresh("true")
	for _, id := range ids {
		req := elastic.NewBulkDeleteRequest().Id(strconv.Itoa(id))
		bulkService.Add(req)
	}
	return bulkService.Do(ctx)

}

func (r *esProductRepository) Save(esProduct document.EsProduct, ctx context.Context) (
	*elastic.IndexResponse, error) {
	getContext(ctx)
	return r.esDB.Index().Index(r.esProduct.GetIndex()).
		Id(strconv.Itoa(esProduct.Id)).BodyJson(esProduct).
		Refresh("true").Do(ctx)
}

func (r *esProductRepository) SearchPage(searchSource *elastic.SearchSource,
	page *Page[document.EsProduct], ctx context.Context) (err error) {
	getContext(ctx)
	elastic.NewSearchSource()
	ss := r.esDB.Search().Index(r.esProduct.GetIndex()).
		SearchSource(searchSource)
	src, err := searchSource.Source()
	data, err := json.Marshal(src)

	got := string(data)
	fmt.Println(got)

	err = page.SelectPages(ss, []string{}, ctx)
	return
}
func (r *esProductRepository) SearchAll(searchSource *elastic.SearchSource, ctx context.Context) (
	res *elastic.SearchResult, err error) {
	getContext(ctx)
	ss := r.esDB.Search().Index(r.esProduct.GetIndex()).SearchSource(searchSource)

	return ss.Do(ctx)

}
