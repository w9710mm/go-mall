package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"mall/global/config"
	"mall/global/dao/model"
	"time"
)

type umsMemberCacheService struct {
	redisDB *redis.Client
	config  config.RedisConfig
	mysqlDB *gorm.DB
}

func NewUmsMemberCacheService(mysqlDB *gorm.DB,
	client *redis.Client) UmsMemberCacheService {
	return &umsMemberCacheService{
		mysqlDB: mysqlDB,
		redisDB: client,
		config:  config.GetConfig().Redis,
	}

}

func (s *umsMemberCacheService) SetMember(member model.UmsMember) {
	key := fmt.Sprintf("%d:%s:%s", s.config.DB, s.config.Key.Member, member.Username)
	s.redisDB.Set(context.TODO(), key, member, time.Duration(s.config.Expire.AuthCode)*time.Second)
}
func (s umsMemberCacheService) DelMember(memberId int64) {
	member := model.UmsMember{}
	s.mysqlDB.Where(&model.UmsMember{Id: memberId}).First(&member)
	if member.Username == nil {
		key := fmt.Sprintf("%d:%s:%s", s.config.DB, s.config.Key.Member, member.Username)
		s.redisDB.Del(context.TODO(), key)
	}
}

func (s umsMemberCacheService) GetMember(username string) (member model.UmsMember, err error) {
	key := fmt.Sprintf("%d:%s:%s", s.config.DB, s.config.Key.Member, member.Username)
	err = s.redisDB.Get(context.TODO(), key).Scan(&member)
	return
}

func (s umsMemberCacheService) SetAuthCode(telephone string, authCode string) {
	key := fmt.Sprintf("%d:%s:%s", s.config.DB, s.config.Key.AuthCode, telephone)
	s.redisDB.Set(context.TODO(), key, authCode, time.Duration(s.config.Expire.AuthCode)*time.Second)
}

func (s umsMemberCacheService) GetAuthCode(telephone string) (code string, err error) {
	key := fmt.Sprintf("%d:%s:%s", s.config.DB, s.config.Key.AuthCode, telephone)
	return s.redisDB.Get(context.TODO(), key).Result()
}
