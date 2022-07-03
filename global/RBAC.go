package global

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"mall/global/dao"
	"mall/global/log"
)

var E *casbin.Enforcer //定义一个casbin对象

func Casbin() *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(dao.DB)
	if err != nil {
		log.Logger.Error(err.Error())
		panic(err.Error())
	}
	e, err := casbin.NewEnforcer("rbac.conf", a, true)
	if err != nil {
		log.Logger.Error(err.Error())
		panic(err.Error())
	}

	e.LoadPolicy() // 从数据库载入配置
	E = e

	E.AddPolicy()
	return e
}

//func getRouterRoles()  {
//	var u dto.UmsMemberRole
//	dao.DB.Find()
//}
