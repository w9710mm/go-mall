package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	v1 "mall/api/v1"

	_ "mall/docs"
	"mall/global/log"
	"mall/middleware"
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
		brand.POST("/create", v1.CreateBrand)
		brand.GET("/list", v1.ListBrand)
		brand.GET("/:id", v1.Brand)
		brand.GET("/delete/:id", v1.DeleteBrand)
		brand.POST("/update/:id", v1.UpdateBrand)

	}

	sso := router.Group("/sso")
	{
		sso.GET("/getAuthCode", v1.GetAuthCode)
		sso.POST("/verifyAuthCode", v1.UpdatePassword)
	}
	InitApi(r)

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
