package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator" // 导入包
	"mall/common/response"
	"mall/global/log"
	"mall/global/model"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
	"strconv"
)

type PmsBrandController struct {
	pmsBrandService service.PmsBrandService
}

func NewPmsBrandController(brandService service.PmsBrandService) api.Controller {

	return &PmsBrandController{pmsBrandService: brandService}
}

// Create godoc
// @Summary 创建品牌
// @Description 创建一个品牌
// @Tags 品牌接口
// @ID v1/PmsBrandController/Create
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pmsBrand body model.PmsBrand true "pmsBrand"
// @Success 200 {object} response.ResponseMsg{data=model.PmsBrand} "success"
// @Failure 500 {object} response.ResponseMsg{data=model.PmsBrand} "failure"
// @Router /brand/create [post]
func (C *PmsBrandController) Create(c *gin.Context) {

	var brand model.PmsBrand
	c.ShouldBind(&brand)

	count := C.pmsBrandService.CrateBrand(brand)
	if count != 1 {
		log.Logger.Debug(fmt.Sprintf("createdBrand failed:%+v", brand))
		c.JSON(http.StatusOK, response.FailedMsg(brand))
		return
	}

	log.Logger.Info(fmt.Sprintf("createdBrand sucess :%+v", brand))
	c.JSON(http.StatusOK, response.SuccessMsg(brand))
}

// Update godoc
// @Summary 更新品牌
// @Description 更新品牌
// @Tags 品牌接口
// @ID v1/PmsBrandController/Update
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path int true "brand_id"
// @Param pmsBrand body model.PmsBrand true "pmsBrand"
// @Success 200 {object} response.ResponseMsg{data=model.PmsBrand} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /brand/update/{id} [post]
func (C *PmsBrandController) Update(c *gin.Context) {

	var brand model.PmsBrand
	c.ShouldBind(&brand)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		msg := "updateBrand param id error:" + err.Error()
		log.Logger.Debug(msg)
		c.JSON(http.StatusOK, response.FailedMsg(msg))
		return
	}
	count := C.pmsBrandService.UpdateBrand(id, brand)
	if count != 1 {
		log.Logger.Debug((fmt.Sprintf("updateBrand failed :id=%d", id)))
		c.JSON(http.StatusOK, response.FailedMsg("updateBrand failed"))
		return
	}

	log.Logger.Info("updateBrand success")
	c.JSON(http.StatusOK, response.SuccessMsg(brand))
}

// Delete godoc
// @Summary 删除品牌
// @Description 删除品牌
// @Tags 品牌接口
// @ID v1/PmsBrandController/Delete
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path int true "brand_id"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /brand/delete/{id} [get]
func (C *PmsBrandController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		msg := "deleteBrand param id error:" + err.Error()
		log.Logger.Debug(msg)
		c.JSON(http.StatusOK, response.FailedMsg(msg))
		return
	}
	count := C.pmsBrandService.DeleteBrand(id)
	if count != 1 {
		msg := fmt.Sprintf("deleteBrand failed :id=%d", id)
		log.Logger.Debug(msg)
		c.JSON(http.StatusOK, response.FailedMsg("delete failed"))
		return
	}

	log.Logger.Info(fmt.Sprintf("deleteBrand success :id=%d", id))
	c.JSON(http.StatusOK, response.SuccessMsg("delete success"))
}

// Brand godoc
// @Summary 获取一个品牌
// @Description  获取一个品牌
// @Tags 品牌接口
// @ID v1/PmsBrandController/Brand
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path int true "brand_id"
// @Success 200 {object} response.ResponseMsg{data=model.PmsBrand} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /brand/{id} [get]
func (C *PmsBrandController) Brand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		msg := "brand param id error:" + err.Error()
		log.Logger.Debug(msg)
		c.JSON(http.StatusOK, response.FailedMsg(msg))
		return
	}

	brand, err := C.pmsBrandService.GetBrand(id)
	if err != nil {
		log.Logger.Debug(fmt.Sprintf("brand failed :id=%d", id))
		c.JSON(http.StatusOK, response.FailedMsg("get brand failed"))
		return
	}
	log.Logger.Debug(fmt.Sprintf("brand sucess :id=%d", id))
	c.JSON(http.StatusOK, response.SuccessMsg(brand))
}

// ListBrand godoc
// @Summary 获取品牌列表
// @Description  获取品牌列表
// @Tags 品牌接口
// @ID v1/PmsBrandController/ListBrand
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pageNum query int false "page number" default(0)
// @Param pageSize query int false "page size"  default(3)
// @Success 200 {object} response.ResponseMsg{data=model.PmsBrand} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /brand/list [get]
func (C *PmsBrandController) ListBrand(c *gin.Context) {

	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		msg := "listBrand param pageNum error:" + err.Error()
		log.Logger.Debug(msg)
		c.JSON(http.StatusOK, response.FailedMsg(msg))
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "3"))

	if err != nil {
		msg := "listBrand param pageSize error:" + err.Error()

		log.Logger.Debug(msg)
		c.JSON(http.StatusOK, response.FailedMsg(msg))
		return
	}
	var page paginator.Page[model.PmsBrand]

	page, err = C.pmsBrandService.ListBrand(pageNum, pageSize)

	if err != nil {
		log.Logger.Debug("listBrand failed:" + err.Error())
		c.JSON(http.StatusOK, response.FailedMsg(err.Error()))
		return
	}
	log.Logger.Info("listBrand success")
	c.JSON(http.StatusOK, response.SuccessMsg(page))

}
func (C *PmsBrandController) Name() string {
	return "PmsBrandController"
}

func (C *PmsBrandController) RegisterRoute(group *gin.RouterGroup) {
	api := group.Group("/brand")
	{
		api.POST("/create", C.Create)
		api.GET("/list", C.ListBrand)
		api.GET("/:id", C.Brand)
		api.GET("/delete/:id", C.Delete)
		api.POST("/update/:id", C.Update)
	}
}
