package service

import (
	"fmt"
	"mall/global/dao/model"
	"time"
)

type umsMemberCacheService struct {
}

var UmsMemberCacheService = new(umsMemberCacheService)

func (s umsMemberCacheService) SetMember(member model.UmsMember) {
	key := fmt.Sprintf("%d:%s:%s", redisConfig.DB, redisConfig.Key.Member, member.Username)
	redisDB.Set(redisCtx, key, member, time.Duration(redisConfig.Expire.AuthCode)*time.Second)
}
func (s umsMemberCacheService) DelMember(memberId int64) {
	member := model.UmsMember{}
	db.Where(&model.UmsMember{Id: memberId}).First(&member)
	if member.Username == nil {
		key := fmt.Sprintf("%d:%s:%s", redisConfig.DB, redisConfig.Key.Member, member.Username)
		redisDB.Del(redisCtx, key)
	}
}

func (s umsMemberCacheService) GetMember(username string) (member model.UmsMember, err error) {
	key := fmt.Sprintf("%d:%s:%s", redisConfig.DB, redisConfig.Key.Member, username)
	err = redisDB.Get(redisCtx, key).Scan(&member)
	return
}

func (s umsMemberCacheService) SetAuthCode(telephone string, authCode string) {
	key := fmt.Sprintf("%d:%s:%s", redisConfig.DB, redisConfig.Key.AuthCode, telephone)
	redisDB.Set(redisCtx, key, authCode, time.Duration(redisConfig.Expire.AuthCode)*time.Second)
}

func (s umsMemberCacheService) GetAuthCode(telephone string) (code string, err error) {
	key := fmt.Sprintf("%d:%s:%s", redisConfig.DB, redisConfig.Key.AuthCode, telephone)
	return redisDB.Get(redisCtx, key).Result()
}
