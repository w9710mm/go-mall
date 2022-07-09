package service

import (
	"github.com/olivere/elastic/v7"
	"mall/global/dao"
	"mall/global/dao/mapper"
)

type esProductService struct {
}

var EsProductService = new(esProductService)
var esProductMapper = mapper.EsProductMapper
var esDB = dao.ESClient

func (s esProductService) importAll() (err error) {
	esProducts, err := esProductMapper.GetAllEsProductList(0)
	if err != nil {
		return
	}
	index := "pms"
	bulk := esDB.Bulk().Index(index).Type("product").Refresh("true")
	for i, product := range esProducts {
		bulk.Add(elastic.NewBulkCreateRequest().Index(index).
			Id(string(i)).
			Doc(product))
	}
	_, err = bulk.Do(context)
	if err != nil {
		return
	}

	return

}
