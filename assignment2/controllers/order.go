package controllers

import (
	"assignment2/models"
	"assignment2/services"
	"assignment2/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

// CreateOrderAndListItems godoc
// @Summary Create Order and create list of items with same order id
// @Decription Create Order and create list of items with same order id
// @Tags order
// @Accept json
// @Produce json
// @Param order body models.CreateOrderRequest true "order info"
// @Success 200 {object} views.SuccesCreateOrder
// @Failure 400 {object} views.Failed
// @Router /orders [POST]
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, views.M_BAD_REQUEST)
		return
	}

	response := c.service.CreateOrder(&req)

	WriteJsonResponse(ctx, response)
}

// UpdateOrderAndListItems godoc
// @Summary Update Order By Order id and update list of items with same order id
// @Decription Update Order By Order id and delete all list of items with same order id
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body models.UpdateOrderRequest true "order info"
// @Success 200 {object} views.SuccessUpdateOrder
// @Failure 400 {object} views.Failed
// @Router /orders/{id} [PUT]
func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var req models.UpdateOrderRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": views.M_BAD_REQUEST, "error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": views.M_BAD_REQUEST, "error": err.Error()})
		return
	}

	response := c.service.UpdateOrder(&req, id)

	WriteJsonResponse(ctx, response)
}

// GetAllOrderWithItems godoc
// @Summary Get all order with list of items
// @Decription Get all order with list of items
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllOrderWithItemsSuccess
// @Router /orders [get]
func (c *OrderController) GetAll(ctx *gin.Context) {
	response := c.service.FindAllOrder()
	WriteJsonResponse(ctx, response)
}

// DeleteOrderAndListOfItem godoc
// @Summary Delete Order By Order id and delete all list of items with same order id
// @Decription Delete Order By Order id and delete all list of items with same order id
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} views.SuccessDeleteOrder
// @Failure 400 {object} views.Failed
// @Router /orders/{id} [DELETE]
func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": views.M_BAD_REQUEST, "error": err.Error()})
		return
	}

	response := c.service.DeleteOrder(id)

	WriteJsonResponse(ctx, response)
}
