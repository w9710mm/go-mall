package dao

import (
	"github.com/olivere/elastic/v7"
	"mall/global/config"
	_ "mall/global/config"
)

var ESClient *elastic.Client

func init() {
	esconfig := config.GetConfig().ElasticSearch
	url := "http://" + esconfig.ClusterNodes
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		panic("connect to eslasticSearch failed:" + err.Error())
	}

	ESClient = client
}
