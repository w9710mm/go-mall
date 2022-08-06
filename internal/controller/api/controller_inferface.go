package api

import "github.com/gin-gonic/gin"

type Controller interface {
	Name() string
	RegisterRoute(*gin.RouterGroup)
}
