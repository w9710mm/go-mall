package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
	"mall/global/dao/document"
)

type EsProductRepository interface {
	SaveAll([]document.EsProduct, context.Context) (*elastic.BulkResponse, error)
	DeleteById(int, context.Context) (*elastic.DeleteResponse, error)
	DeleteAll([]int, context.Context) (*elastic.BulkResponse, error)
	Save(document.EsProduct, context.Context) (*elastic.IndexResponse, error)
	SearchPage(*elastic.SearchSource, *Page[document.EsProduct], context.Context) error
	SearchAll(*elastic.SearchSource, context.Context) (*elastic.SearchResult, error)
}

func getContext(ctx context.Context) {
	if ctx == nil {
		ctx = context.TODO()
	}
}
