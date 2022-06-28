package service

import (
	"fmt"
	"mall/global/config"
	"mall/global/log"
	"mall/global/redis"
	"math/rand"
	"time"
)

type umsMemberService struct {
}

var UmsMemberService = new(umsMemberService)

var redisConfig = config.GetConfig().Redis

var redis_prefix_authcode = redisConfig.Prefix.AuthCode

var redis_expire_authcode = redisConfig.Expire.AuthCode

var redishelper, context = redis.GetRedis()

func (s umsMemberService) GetAuthCode(telephone string) (int, error) {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := rnd.Intn(1000000)
	log.Logger.Debug(fmt.Sprintf("generate authCode:%d", vcode))
	_, err := redishelper.Set(context, redis_prefix_authcode+telephone, vcode, time.Duration(redis_expire_authcode)*time.Second).Result()
	if err != nil {
		return 0, err
	}

	return vcode, nil
}

func (s umsMemberService) VerifyAuthCode(telephone string, authCode string) (bool, error) {
	realAuthCode, err := redishelper.Get(context, redis_prefix_authcode+telephone).Result()
	if err != nil {
		return false, err
	}

	if realAuthCode == "" || realAuthCode != authCode {
		return false, nil
	} else {
		return true, nil
	}

}
