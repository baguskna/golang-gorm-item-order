package main

import (
	"golang-gorm-item-order/app"
	"golang-gorm-item-order/controller"
	"golang-gorm-item-order/repository"
	"golang-gorm-item-order/service"

	"github.com/gin-gonic/gin"
)

func main() {
	app.StartDB()
	db := app.GetDB()

	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	r := gin.Default()
	r.GET("/orders", orderController.GetOrders)
	r.POST("/orders", orderController.PostOrder)
	r.GET("/orders/:id", orderController.GetOneOrder)
	r.DELETE("/orders/:id", orderController.Delete)
	r.GET("/items", itemController.GetItems)
	r.POST("/items", itemController.PostItem)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run()
}
