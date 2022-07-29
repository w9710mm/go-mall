package service

import (
	"mall/global/config"
	"mall/global/dao"
	"mall/global/dao/mapper"
	"mall/global/dao/repository"
)

var db = dao.GetMysqlDB()

var esProductMapper = mapper.EsProductMapper
var portalOrderMapper = mapper.PortalOrderMapper

var esProductRepository = repository.EsProductRepository

var redisDB, redisCtx = dao.GetRedis()

var redisConfig = config.GetConfig().Redis
