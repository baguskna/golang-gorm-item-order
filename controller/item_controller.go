package controller

import (
	"golang-gorm-item-order/model/domain"
	"golang-gorm-item-order/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type itemController struct {
	itemService service.ItemService
}

func NewItemController(itemService service.ItemService) *itemController {
	return &itemController{itemService}
}

func (ic *itemController) GetItems(c *gin.Context) {
	items, err := ic.itemService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (ic *itemController) PostItem(c *gin.Context) {
	var itemRequest domain.Items

	err := c.ShouldBindJSON(&itemRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	item, err := ic.itemService.Create(itemRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func (ic *itemController) UpdateItem(c *gin.Context) {
	var itemReq domain.Items

	err := c.ShouldBindJSON(&itemReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	item, err := ic.itemService.Update(id, itemReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func (ic *itemController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	item, err := ic.itemService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func (ic *itemController) BooksHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	b, err := ic.itemService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})
}
