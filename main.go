package main

import (
	_ "mall/global/dao"
	"mall/global/log"
	_ "mall/internal/componet"
	"mall/internal/route"

	"net/http"
	"time"
)

// @title mall
// @version 1.0
// @description mall
// @termsofservice https://github.com/18211167516/Go-Gin-Api
// @contact.name meme
// @contact.email 962349367@qq.com
// @host 127.0.0.1:8080
func main() {
	log.Logger.Info("starting server...")

	newRouter := route.GetRoute()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
