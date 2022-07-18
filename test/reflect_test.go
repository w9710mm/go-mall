package test

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"mall/global/dao"
	"mall/global/dao/document"
	"mall/global/dao/repository"
	"testing"
)

func TestRe(t *testing.T) {

	esdb, ctx, _ := dao.GetESDB()
	boolQuery := elastic.NewBoolQuery()
	keywords := make(map[string]string)
	keywords["subtitle"] = "智慧"
	keywords["keywords"] = "智慧"
	keywords["name"] = "智慧"
	for s, s2 := range keywords {
		boolQuery.Should(elastic.NewTermsQuery(s2, s))
	}
	boolQuery.MinimumNumberShouldMatch(1)
	ss := esdb.Search().Index("pms").Type("product").
		Sort("_id", false)

	//value := reflect.ValueOf(ss).Elem().Field(0)
	r := repository.Page[document.EsProduct]{Pages: 0, PageSize: 10}
	err := r.SelectPages(ss, []string{}, ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("__________")
	fmt.Println(r)
	fmt.Println("__________")
}

func TestName(t *testing.T) {

	c1 := make(chan int)
	c1 <- 0
	data, ok := <-c1
	fmt.Println(data, ok)

}
