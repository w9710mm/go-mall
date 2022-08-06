package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mall/common/util"
	"mall/global/dao/model"
	"mall/global/log"
	"math/rand"
	"strconv"
	"time"
)

type umsMemberService struct {
	umsMemberCacheService UmsMemberCacheService
	db                    *gorm.DB
}

func NewUmsMemberService(s UmsMemberCacheService, db *gorm.DB) UmsMemberService {
	return &umsMemberService{
		umsMemberCacheService: s,
		db:                    db,
	}
}

func (s *umsMemberService) GetByUserName(username string) (member model.UmsMember, err error) {
	member, err = s.umsMemberCacheService.GetMember(username)
	if err == nil && member.Id != 0 {
		return
	}
	s.db.Where(&model.UmsMember{Username: &username}).First(&member)
	return
}

func (s *umsMemberService) GetById(id int64) (member model.UmsMember) {
	s.db.First(&member, id)
	return
}
func (s *umsMemberService) Register(username string, password, telephone string, authCode string) (err error) {
	flag, err := s.VerifyAuthCode(telephone, authCode)
	if err != nil && !flag {
		return
	}

	rows := s.db.Where(&model.UmsMember{Username: &username}).RowsAffected
	if rows != 0 {
		err = errors.New("this username is exists")
		return
	}
	t := time.Now()
	status := 1
	passw, _ := util.ScryptPassword(password)
	member := model.UmsMember{
		Username:   &username,
		Phone:      &telephone,
		Password:   &passw,
		CreateTime: &t,
		Status:     &status,
	}
	level := model.UmsMemberLevel{}
	s.db.Where(&model.UmsMemberLevel{DefaultStatus: &status}).First(&level)
	if level.Id != 0 {
		member.MemberLevelId = &level.Id
	}
	s.db.Save(&member)
	member.Password = nil
	return
}
func (s *umsMemberService) GenerateAuthCode(telephone string) (vcode string) {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode = strconv.Itoa(rnd.Intn(1000000))
	log.Logger.Debug(fmt.Sprintf("generate authCode:%s", vcode))
	s.umsMemberCacheService.SetAuthCode(telephone, vcode)

	return vcode
}

func (s *umsMemberService) UpdatePassword(telephone string, password string, authCode string) (err error) {
	member := model.UmsMember{}
	s.db.Where(&model.UmsMember{Phone: &telephone}).First(&member)
	if member.Id == 0 {
		err = errors.New("this account is not exists")
		return
	}
	flag, _ := s.VerifyAuthCode(telephone, authCode)
	if !flag {
		err = errors.New("verify auth code is failed")
		return
	}
	newPassword, _ := util.ScryptPassword(password)
	member.Password = &newPassword
	s.db.Save(&member)
	s.umsMemberCacheService.DelMember(member.Id)
	return
}

func (s *umsMemberService) GetCurrentMember(tokenString string) (member model.UmsMember, err error) {
	_, claims, err := util.ParseToken(tokenString)
	if err != nil {
		return
	}
	return s.umsMemberCacheService.GetMember(claims.Username)
}

func (s *umsMemberService) UpdateIntegration(id int64, integration int) {
	member := model.UmsMember{
		Id:          id,
		Integration: &integration,
	}
	s.db.Save(&member)
	s.umsMemberCacheService.DelMember(id)
}

func (s *umsMemberService) LoadUserByUsername(username string) (user model.UmsMember, err error) {
	user, err = s.GetByUserName(username)
	//TODO 权限控制
	return
}

func (s *umsMemberService) Login(username string, password string) (tokenString string, err error) {

	member, err := s.LoadUserByUsername(username)
	if err != nil {
		return
	}
	tokenString, err = util.GenerateToken(*member.Username)
	return
}

func (s *umsMemberService) RefreshToken(tokenString string) (newTokenString string, err error) {
	newTokenString, err = util.RefreshToken(tokenString)
	return
}
func (s *umsMemberService) VerifyAuthCode(telephone string, authCode string) (bool, error) {
	realAuthCode, err := s.umsMemberCacheService.GetAuthCode(telephone)
	if err != nil {
		return false, err
	}

	if realAuthCode == "" || realAuthCode != authCode {
		return false, nil
	} else {
		return true, nil
	}

}
