package dao

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"mall/global/config"
	_ "mall/global/config"
	"mall/global/log"
	"sync"
)

var (
	esctx    = context.Background()
	esonce   sync.Once
	esClient *elastic.Client
)

func init() {

	esconfig := config.GetConfig().ElasticSearch
	url := "http://" + esconfig.ClusterNodes
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	if err != nil {
		panic("connect to eslasticSearch failed:" + err.Error())
	}

	info, code, err := client.Ping(url).Do(esctx)
	if err != nil {
		log.Logger.Error("ping es failed" + err.Error())
	}
	log.Logger.Info(fmt.Sprintf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number))

	esClient = client
}
func GetESDB() (*elastic.Client, context.Context, sync.Once) {
	return esClient, esctx, esonce
}
