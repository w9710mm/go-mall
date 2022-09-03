package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
	"strconv"
)

type HomeController struct {
	homeService service.HomeService
}

func NewHomeController(homeService service.HomeService) api.Controller {
	return &HomeController{homeService: homeService}
}

// content godoc
// @Summary 首页内容信息展示
// @Description  首页展示
// @Tags 首页内容管理
// @ID v1/HomeController/content
// @Accept  json
// @Produce  json
// @Security JWT
// @Success 200 {object} response.ResponseMsg{data=domain.HomeContentResult} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /home/content [get]
func (C *HomeController) content(c *gin.Context) {
	content := C.homeService.Content()
	c.JSON(http.StatusOK, response.SuccessMsg(content))

}

// recommendProductList godoc
// @Summary 推荐内容展示
// @Description  分页获取推荐商品
// @Tags 首页内容管理
// @ID v1/HomeController/recommendProductList
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pageNum query int false "page number" default(0)
// @Param pageSize query int false "page size"  default(5)
// @Success 200 {object} response.ResponseMsg{data=model.PmsProduct} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /home/recommendProductList [get]
func (C *HomeController) recommendProductList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	_, productList := C.homeService.RecommendProductList(pageSize, pageNum)
	c.JSON(http.StatusOK, response.SuccessMsg(productList))
}

// getSubjectList godoc
// @Summary 根据分类获取专题
// @Description  根据分类获取专题
// @Tags 首页内容管理
// @ID v1/HomeController/getSubjectList
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pageNum query int false "page number" default(0)
// @Param pageSize query int false "page size"  default(5)
// @Param cateId query int false "cateId"
// @Success 200 {object} response.ResponseMsg{data=model.CmsSubject} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /home/subjectList [get]
func (C *HomeController) getSubjectList(c *gin.Context) {
	cateId, _ := strconv.ParseInt(c.Query("cateId"), 10, 64)
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	_, subjectList := C.homeService.GetSubjectList(cateId, pageNum, pageSize)
	c.JSON(http.StatusOK, response.SuccessMsg(subjectList))
}

// getProductCateList godoc
// @Summary 获取首页商品分类
// @Description  获取首页商品分类
// @Tags 首页内容管理
// @ID v1/HomeController/getProductCateList
// @Accept  json
// @Produce  json
// @Security JWT
// @Param parentId path int false "parentId"
// @Success 200 {object} response.ResponseMsg{data=model.PmsProductCategory} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /home/productCateList/{parentId}  [get]
func (C *HomeController) getProductCateList(c *gin.Context) {
	param, _ := strconv.ParseInt(c.Param("parentId"), 10, 64)
	productCateList := C.homeService.GetProductCateList(param)
	c.JSON(http.StatusOK, productCateList)
}

// hotProductList godoc
// @Summary 分页获取人气推荐商品
// @Description  分页获取人气推荐商品
// @Tags 首页内容管理
// @ID v1/HomeController/hotProductList
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pageNum query int false "page number" default(0)
// @Param pageSize query int false "page size"  default(5)
//@Success 200 {object} response.ResponseMsg{data=model.PmsProduct} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /home/PmsProduct  [get]
func (C *HomeController) hotProductList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	hotProductList := C.homeService.HotProductList(pageNum, pageSize)
	c.JSON(http.StatusOK, response.SuccessMsg(hotProductList))
}

// newProductList godoc
// @Summary 分页获取新品推荐商品
// @Description  分页获取新品推荐商品
// @Tags 首页内容管理
// @ID v1/HomeController/newProductList
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pageNum query int false "page number" default(0)
// @Param pageSize query int false "page size"  default(5)
//@Success 200 {object} response.ResponseMsg{data=model.PmsProduct} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /home/newProductList  [get]
func (C *HomeController) newProductList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	newProductList := C.homeService.NewProductList(pageNum, pageSize)
	c.JSON(http.StatusOK, response.SuccessMsg(newProductList))
}

func (C *HomeController) Name() string {
	return "HomeController"
}

func (C *HomeController) RegisterRoute(group *gin.RouterGroup) {
	group.GET("/content", C.content)
	group.GET("/recommendProductList", C.recommendProductList)
	group.GET("/productCateList/:parentId", C.getProductCateList)
	group.GET("/subjectList", C.getSubjectList)
	group.GET("/hotProductList", C.hotProductList)
	group.GET("newProductList", C.newProductList)

}
