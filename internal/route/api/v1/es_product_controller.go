package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/internal/service"
	"net/http"
	"strconv"
)

var esProductService = service.EsProductService

func ImportAllList(c *gin.Context) {
	count, err := esProductService.ImportAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(count))

}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	err = esProductService.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))

}

func DeleteBatch(c *gin.Context) {
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

func Create(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product, err := esProductService.Create(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(product))
}
func SearchSimple(c *gin.Context) {
	keyword := c.Query("keyword")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))

	page, err := esProductService.SearchByKeyword(keyword, pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(page))

}
func SearchDetail(c *gin.Context) {
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

func Recommend(c *gin.Context) {
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

func SearchRelatedInfo(c *gin.Context) {
	keyword := c.Query("keyword")
	info, err := esProductService.SearchRelatedInfo(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(info))
}
