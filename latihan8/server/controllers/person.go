package controllers

import (
	"net/http"
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeople godoc
// @Summary Get all people
// @Decription Get list people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// CreatePeople godoc
// @Summary Create people
// @Decription Create people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [post]
func (p *PersonController) CreatePeople(ctx *gin.Context) {

	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "CREATE_PEOPLE_SUCCESS",
		Payload: p,
	})

}

// UpdatePeople godoc
// @Summary Update people
// @Decription Update people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [put]
func (p *PersonController) UpdatePeople(ctx *gin.Context) {

	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PEOPLE_SUCCESS",
		Payload: p,
	})

}

// DeletePeople godoc
// @Summary Delete people
// @Decription Delete people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.DeletePeopleSwagger
// @Router /people/:id [delete]
func (p *PersonController) DeletePeople(ctx *gin.Context) {

	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE_PEOPLE_SUCCESS",
	})

}
