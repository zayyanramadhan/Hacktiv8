package router

import (
	"webserver/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/cars", controller.CreateCar)
	router.GET("/cars/:carID", controller.GetCarByID)
	router.PUT("/cars/:carID", controller.UpdateCar)
	router.DELETE("/cars/:carID", controller.DeleteCar)

	return router
}
