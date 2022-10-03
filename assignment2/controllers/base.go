package controllers

import (
	"assignment2/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, response *views.Response) {
	ctx.JSON(response.Status, response)
}
