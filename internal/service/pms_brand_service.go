package service

import (
	"errors"
	paginator "github.com/yafeng-Soong/gorm-paginator" // 导入包
	"mall/global/dao"
	"mall/internal/model"
)

type pmsBrandService struct {
}

var PmsBrandService = new(pmsBrandService)

func (s *pmsBrandService) CrateBrand(brand model.PmsBrand) int64 {
	return dao.DB.Create(&brand).RowsAffected
}

func (s *pmsBrandService) UpdateBrand(id int, brand model.PmsBrand) int64 {
	brand.Id = id
	return dao.DB.Save(&brand).RowsAffected
}

func (s *pmsBrandService) DeleteBrand(id int64) int64 {
	return dao.DB.Delete(&model.PmsBrand{}, id).RowsAffected
}

func (s *pmsBrandService) GetBrand(id int64) (model.PmsBrand, error) {
	var brand model.PmsBrand
	row := dao.DB.First(&brand, id).RowsAffected
	if row != 1 {
		return model.PmsBrand{}, errors.New("have not pamsbrand")
	}
	return brand, nil
}

func (s *pmsBrandService) ListBrand(num int64, size int64) (paginator.Page[model.PmsBrand], error) {
	page := paginator.Page[model.PmsBrand]{CurrentPage: 1, PageSize: size}

	query := dao.DB.Model(&model.PmsBrand{})

	err := page.SelectPages(query)
	if err != nil {
		return paginator.Page[model.PmsBrand]{}, err
	}
	return page, nil

}
