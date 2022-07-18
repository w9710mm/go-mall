package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/internal/service"
	"net/http"
	"strconv"
)

type esProductController struct {
}

var EsProductController = new(esProductController)
var esProductService = service.EsProductService

// ImportAllList godoc
// @Summary 导入所有数据库中商品到ES
// @Description  导入所有数据库中商品到ES
// @Tags 搜索商品管理
// @ID v1/esProduct/importAll
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /esProduct/importAll [get]
func (e esProductController) ImportAllList(c *gin.Context) {
	count, err := esProductService.ImportAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(count))

}

// Delete godoc
// @Summary 根据id删除商品
// @Description  根据id删除商品
// @Tags 搜索商品管理
// @ID v1/esProduct/Delete
// @Accept  json
// @Produce  json
// @Param id path int true "product_id"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /delete/{id} [get]
func (e esProductController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	err = esProductService.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))

}

// DeleteBatch godoc
// @Summary 根据id批量删除商品
// @Description  根据id批量删除商品
// @Tags 搜索商品管理
// @ID v1/esProduct/DeleteBatch
// @Accept  json
// @Produce  json
// @Param ids query []int true "product_ids"
// @Success 200 {object} response.ResponseMsg{data=int} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /delete/batch [post]
func (e esProductController) DeleteBatch(c *gin.Context) {
	ids := c.QueryArray("ids")
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i], _ = strconv.Atoi(id)
	}
	count, err := esProductService.DeleteByList(idsInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(count))

}

// Create godoc
// @Summary 根据id创建商品
// @Description  根据id创建商品
// @Tags 搜索商品管理
// @ID v1/esProduct/Create
// @Accept  json
// @Produce  json
// @Param id path int true "product_id"
// @Success 200 {object} response.ResponseMsg{data=document.EsProduct} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /create/{id} [post]
func (e esProductController) Create(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product, err := esProductService.Create(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(product))
}

// SearchSimple godoc
// @Summary 简单搜索
// @Description  简单搜索
// @Tags 搜索商品管理
// @ID v1/esProduct/SearchSimple
// @Accept  json
// @Produce  json
// @Param keyword query string false "page number" defualt("")
// @Param pageNum query int false "page number" defualt(0)
// @Param pageSize query int false "page size"  defualt(5)
// @Success 200 {object} response.ResponseMsg{data=document.EsProduct} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /search/simple [get]
func (e esProductController) SearchSimple(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))

	page, err := esProductService.SearchByKeyword(keyword, pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(page))

}

// SearchDetail godoc
// @Summary 综合搜索、筛选、排序
// @Description  综合搜索、筛选、排序
// @Tags 搜索商品管理
// @ID v1/esProduct/SearchDetail
// @Accept  json
// @Produce  json
// @Param keyword query string false "keyword" defualt("")
// @Param brandId query int false "brandId" defualt(0)
// @Param productCategoryId query int false "page number" defualt(0)
// @Param pageNum query int false "page number" defualt(0)
// @Param pageSize query int false "page size"  defualt(5)
// @Param sort query int false "sort 排序字段:0->按相关度；1->按新品；2->按销量；3->价格从低到高；4->价格从高到低"  defualt(5)
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /search [get]
func (e esProductController) SearchDetail(c *gin.Context) {
	keyword := c.Query("keyword")
	brandId, _ := strconv.ParseInt(c.DefaultQuery("brandId", "0"), 10, 64)
	productCategoryId, _ := strconv.ParseInt(c.DefaultQuery("productCategoryId", "0"), 10, 64)
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	sort, _ := strconv.Atoi(c.DefaultQuery("sort", "5"))
	page, err := esProductService.SearchByDetail(keyword, brandId, productCategoryId, pageNum, pageSize, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(page))
}

//  Recommend godoc
// @Summary 根据商品id推荐商品
// @Description  根据商品id推荐商品
// @Tags 搜索商品管理
// @ID v1/esProduct/Recommend
// @Accept  json
// @Produce  json
// @Param brandId path int true "brandId"
// @Param pageNum query int false "page number" defualt(0)
// @Param pageSize query int false "page size"  defualt(5)
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /recommend/{id} [get]
func (e esProductController) Recommend(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	page, err := esProductService.Recommend(id, pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))

		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(page))
}

//  SearchRelatedInfo godoc
// @Summary 获取搜索的相关品牌、分类及筛选属性
// @Description  获取搜索的相关品牌、分类及筛选属性
// @Tags 搜索商品管理
// @ID v1/esProduct/SearchRelatedInfo
// @Accept  json
// @Produce  json
// @Param keyword query string false "keyword" defualt("")
// @Success 200 {object} response.ResponseMsg{data=domain.EsProductRelatedInfo} "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /search/relate [get]
func (e esProductController) SearchRelatedInfo(c *gin.Context) {
	keyword := c.Query("keyword")
	info, err := esProductService.SearchRelatedInfo(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(info))
}
