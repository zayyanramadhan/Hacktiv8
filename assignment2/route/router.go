package route

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	router *gin.Engine
	order  *controllers.OrderController
}

func NewRouter(r *gin.Engine, order *controllers.OrderController) *router {
	return &router{
		router: r,
		order:  order,
	}
}

func (r *router) SetupRouter(port string) {
	r.router.GET("/orders", r.order.GetAll)
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.PUT("/orders/:id", r.order.UpdateOrder)
	r.router.DELETE("/orders/:id", r.order.DeleteOrder)
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.router.Run(port)
}
