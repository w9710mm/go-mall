package dao

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

// 这里需要注意的是：page和pageSize 的数据类型需要指定
// 这里根据我之前封装的：获取全部请求参数的一个map指定的返回类型后，决定使用string类型进行定义
func Paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageUp, _ := strconv.Atoi(page)
		if pageUp == 0 {
			pageUp = 1
		}
		pageSizeInt, _ := strconv.Atoi(pageSize)
		switch {
		case pageSizeInt <= 0:
			pageSizeInt = 10
		}
		offset := (pageUp - 1) * pageSizeInt
		return db.Offset(offset).Limit(pageSizeInt)
	}
}
