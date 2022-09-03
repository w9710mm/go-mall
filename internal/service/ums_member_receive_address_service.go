package service

import (
	"gorm.io/gorm"
	"mall/global/model"
)

/**
 *@author:
 *@date:2022/8/22
**/

type umsMemberReceiveAddressService struct {
	db *gorm.DB
}

func NewUmsMemberReceiveAddressService(db *gorm.DB) UmsMemberReceiveAddressService {
	return &umsMemberReceiveAddressService{db: db}
}
func (u *umsMemberReceiveAddressService) Add(address model.UmsMemberReceiveAddress) {
	u.db.Save(&address)
}

func (u *umsMemberReceiveAddressService) Delete(id int64, memberId int64) int64 {
	return u.db.Delete(&model.UmsMemberReceiveAddress{Id: id}).
		Where(map[string]interface{}{"member_id": memberId}).RowsAffected
}

func (u *umsMemberReceiveAddressService) Update(id int64, address model.UmsMemberReceiveAddress, memberId int64) (count int64) {
	//var a model.UmsMemberReceiveAddress
	//u.db.Model(&model.UmsMemberReceiveAddress{Id: id}).Where(map[string]interface{}{"member_id": memberId}).First(&a)
	//count:=0
	if address.DefaultStatus != nil && *address.DefaultStatus != 1 {
		count += u.db.Model(&model.UmsMemberReceiveAddress{}).Where(map[string]interface{}{"member_id": memberId, "default_status": 1}).
			Update("default_status", 0).RowsAffected
	}
	count += u.db.Create(&address).RowsAffected
	return
}

func (u *umsMemberReceiveAddressService) List(memberId int64) (address []model.UmsMemberReceiveAddress) {
	u.db.Model(&model.UmsMemberReceiveAddress{}).Where(map[string]interface{}{"member_id": memberId}).Find(&address)
	return
}

func (u *umsMemberReceiveAddressService) GetItem(memberId int64, id int64) (add model.UmsMemberReceiveAddress) {
	add.Id = id
	u.db.Model(&add).Where(map[string]interface{}{"member_id": memberId}).First(&add)
	return
}
