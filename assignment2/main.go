package main

import (
	"assignment2/config"
	"assignment2/controllers"
	_ "assignment2/docs"
	"assignment2/repos"
	"assignment2/route"
	"assignment2/services"

	"github.com/gin-gonic/gin"
)

// @title Orders API
// @description Sample API Spec for Orders
// @version v2.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:8081
// @contact.name zayyan
// @contact.email zayyanramadhan@gmail.com

func main() {
	db, err := config.ConnectMysqlGORM()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	itemRepo := repos.NewItemRepository(db)
	orderRepo := repos.NewOrderRepository(db)

	orderService := services.NewOrderService(orderRepo, itemRepo)

	orderController := controllers.NewOrderController(orderService)

	app := route.NewRouter(router, orderController)
	app.SetupRouter(":8081")

}
