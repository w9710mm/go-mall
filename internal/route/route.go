package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	middleware "mall/internal/middleware"
	v12 "mall/internal/route/api/v1"

	_ "mall/docs"
	"mall/global/log"
	"net/http"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(Cors())
	r.Use(middleware.Recovery(true))
	r.Use(middleware.GinLogger())

	router = r
	newRouter()
	// server.Use(gin.Recovery())

	initSwagger()
}

func newRouter() {

	brand := router.Group("/brand")
	{
		brand.POST("/create", v12.CreateBrand)
		brand.GET("/list", v12.ListBrand)
		brand.GET("/:id", v12.Brand)
		brand.GET("/delete/:id", v12.DeleteBrand)
		brand.POST("/update/:id", v12.UpdateBrand)

	}
	brand.Use(middleware.JWTAuth())
	sso := router.Group("/sso")
	{
		sso.GET("/getAuthCode", v12.GetAuthCode)
		sso.POST("/verifyAuthCode", v12.UpdatePassword)
	}
	sso.Use(middleware.JWTAuth())
}

func GetRoute() *gin.Engine {
	return router
}

func initSwagger() {
	router.Static("/html", "./public")

	// 设置 swagger 访问路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Logger.Info("look at swagger: \n http://localhost:8080/swagger/index.html")
	go router.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("HttpError", zap.Any("HttpError", err))
			}
		}()

		c.Next()
	}
}
