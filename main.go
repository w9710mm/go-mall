package main

import (
	"mall/global/log"
	"mall/internal/server"
)

// @title mall
// @version 1.0
// @description mall
// @termsofservice https://github.com/18211167516/Go-Gin-Api
//// @securityDefinitions.apikey ApiKeyAuth
//// @in header
//// @name Authorization
// @contact.name meme
// @contact.email 962349367@qq.com
// @host 127.0.0.1:8080
func main() {

	log.Logger.Info("starting server...")
	s, err := server.New()

	if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
	s.NewServer()

}
