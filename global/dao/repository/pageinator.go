package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// 标准分页结构体，接收最原始的DO
// 建议在外部再建一个字段一样的结构体，用以将DO转换成DTO或VO
type Page[T any] struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	Total       int64 `json:"total"`
	Pages       int   `json:"pages"`
	Data        []T   `json:"data"`
}

// 各种查询条件先在query设置好后再放进来
func (page *Page[T]) SelectPages(ss *elastic.SearchService, search []string, ctx context.Context) (e error) {
	paginate(page)

	if len(search) != 0 {
		//TODO searchAfeter的详细实现
		for _, s := range search {
			ss = ss.SearchAfter(s)
		}
	} else {
		ss = ss.From(page.Pages)
	}
	ss = ss.Size(page.PageSize)
	re, err := ss.Do(ctx)
	if err != nil {
		return err
	}
	page.Total = re.Hits.TotalHits.Value
	page.Data = []T{}
	if page.Total == 0 || re.Hits.TotalHits.Value == 0 {

		return
	} else {
		page.Data = make([]T, len(re.Hits.Hits))
	}
	for i, hit := range re.Hits.Hits {
		var model T
		err := json.Unmarshal(hit.Source, &model)
		fmt.Println(string(hit.Source))
		if err != nil {
			return err
		}
		page.Data[i] = model
	}
	return
}

func paginate[T any](page *Page[T]) {

	if page.CurrentPage <= 0 {
		page.CurrentPage = 0
	}
	switch {
	case page.PageSize > 100:
		page.PageSize = 100 // 限制一下分页大小
	case page.PageSize <= 0:
		page.PageSize = 5
	}
}
