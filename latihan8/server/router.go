package server

import (
	"sesi8-latihan/server/controllers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	people *controllers.PersonController
}

func NewRouter(people *controllers.PersonController) *Router {
	return &Router{people: people}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	router.GET("/people", r.people.GetPeople)
	router.POST("/people", r.people.CreatePeople)
	router.DELETE("/people/:id", r.people.DeletePeople)
	router.PUT("/people", r.people.CreatePeople)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(port)
}
