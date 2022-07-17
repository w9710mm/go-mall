package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"mall/common/util"
	"mall/global/dao"
	"mall/global/dao/dto"
	"mall/global/dao/model"
	"time"
)

type umsAdminService struct {
}

var UmsAdminService = new(umsAdminService)

func (s *umsAdminService) Register(adminDto dto.UmsAdminParam) (admin model.UmsAdmin, err error) {

	copier.Copy(&adminDto, &admin)
	var ct = time.Now()
	admin.CreateTime = &ct
	var status = 1
	admin.Status = &status
	var dbAdmin model.UmsAdmin
	row := dao.DB.Where(&model.UmsAdmin{Username: admin.Username}).First(&dbAdmin).RowsAffected
	if row > 0 {
		return model.UmsAdmin{}, errors.New("same username exists in database")
	}
	password, err := util.ScryptPassword(adminDto.Password)
	if err != nil {
		return model.UmsAdmin{}, errors.New("scrypt password error")
	}
	admin.Password = &password
	dao.DB.Save(&admin)
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
