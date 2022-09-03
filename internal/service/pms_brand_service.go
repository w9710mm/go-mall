package service

import (
	"errors"
	paginator "github.com/yafeng-Soong/gorm-paginator" // 导入包
	"gorm.io/gorm"
	"mall/global/model"
)

type pmsBrandService struct {
	db *gorm.DB
}

func NewPmsBrandService(db *gorm.DB) PmsBrandService {
	return &pmsBrandService{db: db}
}

func (s *pmsBrandService) CrateBrand(brand model.PmsBrand) int64 {
	return s.db.Create(&brand).RowsAffected
}

func (s *pmsBrandService) UpdateBrand(id int, brand model.PmsBrand) int64 {
	brand.Id = id
	return s.db.Save(&brand).RowsAffected
}

func (s *pmsBrandService) DeleteBrand(id int) int64 {
	return s.db.Delete(&model.PmsBrand{}, id).RowsAffected
}

func (s *pmsBrandService) GetBrand(id int) (model.PmsBrand, error) {
	var brand model.PmsBrand
	row := s.db.First(&brand, id).RowsAffected
	if row != 1 {
		return model.PmsBrand{}, errors.New("have not pamsbrand")
	}
	return brand, nil
}

func (s *pmsBrandService) ListBrand(num int, size int) (paginator.Page[model.PmsBrand], error) {
	page := paginator.Page[model.PmsBrand]{CurrentPage: int64(num), PageSize: int64(size)}

	query := s.db.Model(&model.PmsBrand{})

	err := page.SelectPages(query)
	if err != nil {
		return paginator.Page[model.PmsBrand]{}, err
	}
	return page, nil

}
