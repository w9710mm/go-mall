package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"mall/common/util"
	"mall/global/dao/dto"
	"mall/global/dao/model"
	"time"
)

type umsAdminService struct {
	db *gorm.DB
}

func NewUmsAdminService(db *gorm.DB) UmsAdminService {
	return &umsAdminService{db: db}
}
func (s *umsAdminService) Register(adminDto dto.UmsAdminParam) (admin model.UmsAdmin, err error) {

	copier.Copy(&adminDto, &admin)
	var ct = time.Now()
	admin.CreateTime = &ct
	var status = 1
	admin.Status = &status
	var dbAdmin model.UmsAdmin
	row := s.db.Where(&model.UmsAdmin{Username: admin.Username}).First(&dbAdmin).RowsAffected
	if row > 0 {
		return model.UmsAdmin{}, errors.New("same username exists in database")
	}
	password, err := util.ScryptPassword(adminDto.Password)
	if err != nil {
		return model.UmsAdmin{}, errors.New("scrypt password error")
	}
	admin.Password = &password
	s.db.Save(&admin)
	return admin, nil
}

func (s *umsAdminService) Login(username, password string) {
	//var token string

}

func (s *umsAdminService) LoadUserByUsername(username string) {

}

func (s *umsAdminService) RefreshToken(oldToken string) (string, error) {
	return util.RefreshToken(oldToken)
}
