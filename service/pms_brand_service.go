package service

import (
	"errors"
	paginator "github.com/yafeng-Soong/gorm-paginator" // 导入包
	common_DB "mall/global/dao"

	"mall/model"
)

type pmsBrandService struct {
}

var PmsBrandService = new(pmsBrandService)

var db = common_DB.GetDB()

func (s *pmsBrandService) CrateBrand(brand model.PmsBrand) int64 {
	return db.Create(&brand).RowsAffected
}

func (s *pmsBrandService) UpdateBrand(id int64, brand model.PmsBrand) int64 {
	brand.Id = id
	return db.Save(&brand).RowsAffected
}

func (s *pmsBrandService) DeleteBrand(id int64) int64 {
	return db.Delete(&model.PmsBrand{}, id).RowsAffected
}

func (s *pmsBrandService) GetBrand(id int64) (model.PmsBrand, error) {
	var brand model.PmsBrand
	row := db.First(&brand, id).RowsAffected
	if row != 1 {
		return model.PmsBrand{}, errors.New("have not pamsbrand")
	}
	return brand, nil
}

func (s *pmsBrandService) ListBrand(num int64, size int64) (paginator.Page[model.PmsBrand], error) {
	page := paginator.Page[model.PmsBrand]{CurrentPage: 1, PageSize: size}

	query := db.Model(&model.PmsBrand{})

	err := page.SelectPages(query)
	if err != nil {
		return paginator.Page[model.PmsBrand]{}, err
	}
	return page, nil

}
