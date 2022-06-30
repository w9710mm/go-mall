package service

import (
	"github.com/jinzhu/copier"
	"mall/dto"
	"mall/model"
)

type umsAdminService struct {
}

var UmsAdminService = new(umsAdminService)

func register(adminDto dto.UmsAdminParam) (admin model.UmsAdmin, err error) {

	copier.Copy(&adminDto, &admin)

	return
}
