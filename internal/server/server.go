package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/olivere/elastic/v7"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall/global/config"
	"mall/global/dao"
	"mall/global/dao/mapper"
	"mall/global/dao/nosql"
	"mall/global/dao/repository"
	"mall/internal/componet/time_task"
	"mall/internal/controller/api"
	v1 "mall/internal/controller/api/v1"
	"mall/internal/middleware"
	"mall/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "mall/docs"
	"mall/global/log"
)

func New() (*Server, error) {

	redisDB := dao.GetRedis()
	esDB := dao.GetESDB()
	mysqlDb := dao.GetMysqlDB()
	mongoDB := dao.GetMongoDB()

	productMapper := mapper.NewEsProductMapper(mysqlDb)
	orderMapper := mapper.NewPortalOrderMapper(mysqlDb)

	esProductRepository := repository.NewEsProductRepository(esDB)

	historyRepository := nosql.NewMemberReadHistoryRepository("memberReadHistory", mongoDB.Collection("memberReadHistory"))

	productService := service.NewEsProductService(esProductRepository, productMapper)
	historyService := service.NewMemberReadHistoryService(historyRepository)
	orderService := service.NewOmsPortalOrderService(orderMapper, mysqlDb)
	brandService := service.NewPmsBrandService(mysqlDb)
	adminService := service.NewUmsAdminService(mysqlDb)
	cacheService := service.NewUmsMemberCacheService(mysqlDb, redisDB)
	memberService := service.NewUmsMemberService(cacheService, mysqlDb)

	esProductController := v1.NewEsProductController(productService)
	memberReadHistoryController := v1.NewMemberReadHistoryController(historyService, memberService, cacheService)
	pmsBrandController := v1.NewPmsBrandController(brandService)
	umsAdminController := v1.NewUmsAdminController(adminService)
	umsMemberController := v1.NewUmsMemberController(memberService, cacheService)

	controllers := &controllers{
		esProductController,
		memberReadHistoryController,
		pmsBrandController,
		umsMemberController,
		umsAdminController,
	}

	orderTimeOutCancelTask := time_task.NewOrderTimeOutCancelTask(orderService)

	timeTasks := &timeTasks{
		orderTimeOutCancelTask,
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Recovery(true))
	r.Use(middleware.GinLogger())
	//r.Use(middleware.TraceMiddleware()

	// server.Use(gin.Recovery())
	conf := config.GetConfig()
	s := &Server{
		engine:  r,
		config:  conf,
		logger:  log.Logger,
		mysqldb: mysqlDb,
		rdb:     redisDB,
		esdb:    esDB,
		mongodb: mongoDB,
		//containerClient: nil,
		//kubeClient:      nil,
		controllers: controllers,
		timeTasks:   timeTasks,
	}
	return s, nil
}

type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *zap.Logger

	mysqldb *gorm.DB
	rdb     *redis.Client
	mongodb *mongo.Database
	esdb    *elastic.Client

	//containerClient *docker.Client
	//kubeClient      *kubernetes.KubeClient

	controllers *controllers
	timeTasks   *timeTasks
}
type controllers struct {
	esProductController         api.Controller
	memberReadHistoryController api.Controller
	pmsBrandController          api.Controller
	umsMemberController         api.Controller
	umsAdminController          api.Controller
}
type timeTasks struct {
	orderTimeOutCancelTask time_task.TimeTask
}

func (S *Server) NewServer() {
	defer S.Close()
	S.initSwagger()
	S.initRouter()
	addr := fmt.Sprintf("%s:%d", S.config.Server.Host, S.config.Server.Port)
	S.logger.Info("Start server on: addr", zap.String("addr", addr))
	server := &http.Server{
		Addr:    addr,
		Handler: S.engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			S.logger.Error("Failed to start server, err", zap.Error(err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	ch := <-sig
	S.logger.Info("Receive signal: ch", zap.Any("ch", ch))

	server.Shutdown(ctx)
}

func (S *Server) initRouter() {
	controllers := make([]string, 0)

	router := S.engine
	brand := router.Group("/brand")
	{
		brand.Use(middleware.JWTAuth())
		S.controllers.pmsBrandController.RegisterRoute(brand)
		controllers = append(controllers, S.controllers.pmsBrandController.Name())

	}

	sso := router.Group("/sso")
	{
		//sso.Use(middleware.JWTAuth())
		S.controllers.umsMemberController.RegisterRoute(sso)
		controllers = append(controllers, S.controllers.umsMemberController.Name())

	}

	esProduct := router.Group("/esProduct")
	{
		esProduct.Use(middleware.JWTAuth())
		S.controllers.esProductController.RegisterRoute(esProduct)
		controllers = append(controllers, S.controllers.esProductController.Name())

	}

	memberReadHistory := router.Group("/member/readHistory")
	{
		memberReadHistory.Use(middleware.JWTAuth())
		S.controllers.memberReadHistoryController.RegisterRoute(memberReadHistory)
		controllers = append(controllers, S.controllers.memberReadHistoryController.Name())
	}
	log.Logger.Info("server enabled controllers: controllers ", zap.Any("controllers", controllers))

}

func (S *Server) Close() {
	ctx := context.TODO()
	S.logger.Info("close server....")
	S.mongodb.Client().Disconnect(ctx)
	db, _ := S.mysqldb.DB()
	if db != nil {
		db.Close()
	}
	S.rdb.Close()
	S.mongodb.Client().Disconnect(ctx)
	S.esdb.Stop()

	S.logger.Info("close server.... over")

}

func (S *Server) initSwagger() {
	S.engine.Static("/html", "./public")

	// 设置 swagger 访问路由
	S.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Logger.Info("look at swagger: \n http://localhost:8080/swagger/index.html")
	go S.engine.Run(":8080")
}
