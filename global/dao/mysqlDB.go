package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mall/global/config"
	_ "mall/global/config"
	"mall/global/log"
	"net/url"
)

var DB *gorm.DB // 全局db句柄

func init() {
	dbconfig := config.GetConfig().MySQL

	host := dbconfig.Host
	port := dbconfig.Port
	database := dbconfig.Name
	username := dbconfig.User
	password := dbconfig.Password
	charset := dbconfig.Charset
	loc := dbconfig.Loc
	time := dbconfig.Time
	moc := dbconfig.MaxOpenConns
	mic := dbconfig.MaxIdelConns

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=%s",
		username, password, host, port, database, charset, url.QueryEscape(loc))

	if time != 0 {
		args = args + fmt.Sprintf("&timeout=%ds", time)
	}

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Logger.Fatal("fail to connect database, err: " + err.Error())
	}

	sqlDB, _ := db.DB()
	if mic != 0 {
		sqlDB.SetMaxIdleConns(mic)
	} else {
		sqlDB.SetMaxIdleConns(2)
	}

	if moc != 0 {
		sqlDB.SetMaxOpenConns(moc)
	} else {
		sqlDB.SetMaxOpenConns(20)
	}
	log.Logger.Info("success to connect database")

	DB = db
}

// 获取db句柄
func GetDB() *gorm.DB {
	return DB
}
