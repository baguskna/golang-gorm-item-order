package controller

import (
	"golang-gorm-item-order/model/domain"
	"golang-gorm-item-order/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *orderController {
	return &orderController{orderService}
}

func (ic *orderController) GetOrders(c *gin.Context) {
	orders, err := ic.orderService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (ic *orderController) PostOrder(c *gin.Context) {
	var orderRequest domain.Order

	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	order, err := ic.orderService.Create(orderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func (ic *orderController) Updateorder(c *gin.Context) {
	var orderReq domain.Order

	err := c.ShouldBindJSON(&orderReq)
	if err != nil {
		// errMessages := []string{}
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		// 	errMessages = append(errMessages, errMessage)
		// }

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	order, err := ic.orderService.Update(id, orderReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func (ic *orderController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	order, err := ic.orderService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func (ic *orderController) GetOneOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	b, err := ic.orderService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})
}
