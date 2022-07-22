package service

import (
	"encoding/json"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"mall/global/dao/document"
	"mall/global/dao/domain"
	"mall/global/dao/mapper"
	"mall/global/dao/repository"
	"strconv"
)

type esProductService struct {
}

var EsProductService = new(esProductService)
var esProductMapper = mapper.EsProductMapper

var esProductRepository = repository.EsProductRepository

func (s esProductService) ImportAll() (num int, err error) {
	esProducts, err := esProductMapper.GetAllEsProductList(0)
	if err != nil {
		return
	}
	res, err := esProductRepository.SaveAll(esProducts)
	if err != nil {
		return
	}
	num = len(res.Succeeded())
	return

}

func (s esProductService) DeleteById(id int) (err error) {
	_, err = esProductRepository.DeleteById(id)
	return
}

func (s esProductService) DeleteByList(ids []int) (count int, err error) {
	res, err := esProductRepository.DeleteAll(ids)
	count = len(res.Succeeded())
	return
}

func (s esProductService) Create(id int64) (esProduct document.EsProduct,
	err error) {
	esProductList, err := esProductMapper.GetAllEsProductList(id)
	if err != nil {
		return
	}
	if len(esProductList) > 0 {
		esProduct = esProductList[0]
		_, err = esProductRepository.Save(esProduct)
		if err != nil {
			return
		}
	}
	return
}

func (s esProductService) SearchByKeyword(keyword string, pageNum int, pageSize int) (
	repository.Page[document.EsProduct], error) {

	boolQuery := elastic.NewBoolQuery()
	page := &repository.Page[document.EsProduct]{Pages: pageNum, PageSize: pageSize}
	searchSource := elastic.NewSearchSource()
	if keyword != "" {
		boolQuery.Should(elastic.NewMatchQuery("subTitle", keyword))
		boolQuery.Should(elastic.NewMatchQuery("keywords", keyword))
		boolQuery.Should(elastic.NewMatchQuery("name", keyword))
	}

	searchSource.Query(boolQuery)

	err := esProductRepository.SearchPage(searchSource, page)
	return *page, err

}

func (s esProductService) SearchByDetail(keyword string, brandId int64,
	productCategoryId int64, pageNum int, pageSize int, sort int) (
	repository.Page[document.EsProduct], error) {

	page := &repository.Page[document.EsProduct]{Pages: pageNum, PageSize: pageSize}
	searchSource := elastic.NewSearchSource()
	sortInfo := elastic.SortInfo{}
	query := elastic.NewBoolQuery()

	if brandId != 0 || productCategoryId != 0 {
		q := elastic.NewBoolQuery()

		if brandId != 0 {
			q.Must(elastic.NewTermsQuery("brandId", brandId))
		}
		if productCategoryId != 0 {
			q.Must(elastic.NewTermsQuery("productCategoryId", productCategoryId))
		}
		query.Filter(q)
	}

	if keyword == "" {
		q := elastic.NewMatchAllQuery().Boost(1)
		query.Must(q)
	} else {
		q := elastic.NewFunctionScoreQuery()

		q.Add(elastic.NewMatchQuery("name", keyword),
			elastic.NewWeightFactorFunction(10))
		q.Add(elastic.NewMatchQuery("subTitle", keyword),
			elastic.NewWeightFactorFunction(5))
		q.Add(elastic.NewMatchQuery("keywords", keyword),
			elastic.NewWeightFactorFunction(2))

		q.MinScore(2).ScoreMode("sum")
		query.Should(q)
	}
	if sort != 0 {
		switch sort {
		case 1:
			sortInfo.Field = "id"
			sortInfo.Ascending = false
		case 2:
			sortInfo.Field = "sale"
			sortInfo.Ascending = true
		case 3:
			sortInfo.Field = "price"
			sortInfo.Ascending = true
		case 4:
			sortInfo.Field = "price"
			sortInfo.Ascending = false
		default:
			sortInfo.Field = "id"
			sortInfo.Ascending = false
		}
		searchSource.SortBy(sortInfo)
	}

	searchSource.Query(query)

	err := esProductRepository.SearchPage(searchSource, page)
	return *page, err
}

func (s esProductService) Recommend(id int64, pageNum int, pageSize int) (
	repository.Page[document.EsProduct], error) {

	query := elastic.NewBoolQuery()
	page := &repository.Page[document.EsProduct]{Pages: pageNum, PageSize: pageSize}
	searchSource := elastic.NewSearchSource()

	esProducts, err := esProductMapper.GetAllEsProductList(id)
	if err != nil {
		return *page, err
	}
	if len(esProducts) > 0 {
		esProduct := esProducts[0]
		keyword := esProduct.Name
		brandId := esProduct.BrandId
		productCategoryId := esProduct.ProductCategoryId
		q := elastic.NewFunctionScoreQuery()
		q.Add(elastic.NewMatchQuery("name", keyword),
			elastic.NewWeightFactorFunction(8))
		q.Add(elastic.NewMatchQuery("subTitle", keyword),
			elastic.NewWeightFactorFunction(2))
		q.Add(elastic.NewMatchQuery("keywords", keyword),
			elastic.NewWeightFactorFunction(2))

		q.Add(elastic.NewMatchQuery("brandId", brandId),
			elastic.NewWeightFactorFunction(5))

		q.Add(elastic.NewMatchQuery("productCategoryId", productCategoryId),
			elastic.NewWeightFactorFunction(3))
		q.MinScore(2).ScoreMode("sum")

		query.MustNot(elastic.NewBoolQuery().MustNot(elastic.NewTermsQuery("id", id)))
		query.Should(q)
		searchSource.Query(query)
		err = esProductRepository.SearchPage(searchSource, page)
		if err != nil {
			return *page, err
		}

	}

	return *page, err

}

func (s esProductService) SearchRelatedInfo(keyword string) (domain.EsProductRelatedInfo,
	error) {

	searchSource := elastic.NewSearchSource()

	if keyword == "" {
		searchSource.Query(elastic.NewMatchAllQuery())
	} else {
		searchSource.Query(elastic.NewMultiMatchQuery(
			keyword, "name", "subTitle", "keywords"))
	}

	//聚合搜索
	searchSource.Aggregation("brandNames",
		elastic.NewTermsAggregation().Field("brandName"))

	searchSource.Aggregation("productCategoryNames",
		elastic.NewTermsAggregation().Field("productCategoryName"))
	//
	searchSource.Aggregation("allAttrValues", elastic.NewNestedAggregation().Path("attrValueList").
		SubAggregation("productAttrs", elastic.NewFilterAggregation().Filter(elastic.NewTermsQuery("attrValueList.type", 1)).
			SubAggregation("attrIds", elastic.NewTermsAggregation().Field("attrValueList.productAttributeId").
				SubAggregation("attrValues", elastic.NewTermsAggregation().Field("attrValueList.value")).
				SubAggregation("attrNames", elastic.NewTermsAggregation().Field("attrValueList.name")))))
	//
	source, _ := searchSource.Source()

	data, _ := json.Marshal(source)
	fmt.Println(string(data))
	res, err := esProductRepository.SearchAll(searchSource)
	if err != nil {
		return domain.EsProductRelatedInfo{}, err
	}
	return convertProductRelatedInfo(*res), nil

}

func convertProductRelatedInfo(res elastic.SearchResult) (productRelatedInfo domain.EsProductRelatedInfo) {

	brandNames, b := res.Aggregations.Terms("brandNames")
	brandNameList := make([]string, len(brandNames.Buckets))
	if b {
		for i, bucket := range brandNames.Buckets {
			brandNameList[i] = bucket.Key.(string)
		}
	}
	productRelatedInfo.BrandNames = brandNameList

	productCategoryNames, b := res.Aggregations.Terms("productCategoryNames")
	productCategoryList := make([]string, len(productCategoryNames.Buckets))
	if b {
		for i, bucket := range productCategoryNames.Buckets {
			productCategoryList[i] = bucket.Key.(string)
		}
	}
	productRelatedInfo.ProductCategoryNames = productCategoryList

	nested, b := res.Aggregations.Nested("allAttrValues")
	if b {
		filter, b := nested.Aggregations.Filter("productAttrs")
		if b {
			attrIds, b := filter.Aggregations.Terms("attrIds")
			if b {
				attrList := make([]domain.ProductAttr, len(attrIds.Buckets))
				for i, bucket := range attrIds.Buckets {
					attr := domain.ProductAttr{}
					attr.AttrId, _ = strconv.Atoi(bucket.KeyNumber.String())

					attrValues, b := bucket.Aggregations.Terms("attrValues")
					if b {
						attrValueList := make([]string, len(attrValues.Buckets))
						for i, item := range attrValues.Buckets {
							attrValueList[i] = item.Key.(string)
						}
						attr.AttrValues = attrValueList
					}

					attrNames, b := bucket.Aggregations.Terms("attrNames")
					if b && len(attrNames.Buckets) > 0 {
						attr.AttrName = attrNames.Buckets[0].Key.(string)
					}
					attrList[i] = attr
				}
				productRelatedInfo.ProductAttrs = attrList
			}
		}
	}
	return
}
