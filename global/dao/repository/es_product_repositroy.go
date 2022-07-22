package repository

import (
	"encoding/json"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"mall/global/dao"
	"mall/global/dao/document"
	"strconv"
)

type esProductRepository struct {
}

var EsProductRepository = new(esProductRepository)
var esDB, esCtx, esSync = dao.GetESDB()

var p = document.EsProduct{}

func (r esProductRepository) SaveAll(esProducts []document.EsProduct) (
	*elastic.BulkResponse, error) {

	bulkService := esDB.Bulk().Index(p.GetIndex()).Refresh("true")
	for i, product := range esProducts {
		bulkService.Add(elastic.NewBulkCreateRequest().
			Id(strconv.Itoa(i)).
			Doc(product))
	}

	return bulkService.Do(esCtx)
}

func (r esProductRepository) DeleteById(id int) (*elastic.DeleteResponse, error) {

	return esDB.Delete().Index(p.GetIndex()).Id(strconv.Itoa(id)).
		Refresh("true").Do(esCtx)

}

func (r esProductRepository) DeleteAll(ids []int) (*elastic.BulkResponse, error) {
	bulkService := esDB.Bulk().Index(p.GetIndex()).Refresh("true")
	for _, id := range ids {
		req := elastic.NewBulkDeleteRequest().Id(strconv.Itoa(id))
		bulkService.Add(req)
	}
	return bulkService.Do(esCtx)

}

func (r esProductRepository) Save(esProduct document.EsProduct) (*elastic.IndexResponse, error) {

	return esDB.Index().Index(p.GetIndex()).
		Id(strconv.Itoa(esProduct.Id)).BodyJson(esProduct).
		Refresh("true").Do(esCtx)
}

func (r esProductRepository) SearchPage(searchSource *elastic.SearchSource,
	page *Page[document.EsProduct]) (
	err error) {
	elastic.NewSearchSource()
	ss := esDB.Search().Index(p.GetIndex()).
		SearchSource(searchSource)
	src, err := searchSource.Source()
	data, err := json.Marshal(src)

	got := string(data)
	fmt.Println(got)

	err = page.SelectPages(ss, []string{}, esCtx)
	return
}
func (r esProductRepository) SearchAll(searchSource *elastic.SearchSource) (
	res *elastic.SearchResult, err error) {
	ss := esDB.Search().Index(p.GetIndex()).SearchSource(searchSource)

	return ss.Do(esCtx)

}
