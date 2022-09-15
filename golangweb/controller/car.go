package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID int    `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var cars = []Car{}

func CreateCar(ctx *gin.Context) {
	var req Car
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(cars) < 1 {
		req.CarID = len(cars) + 1
	} else {
		req.CarID = cars[len(cars)-1].CarID + 1
	}

	cars = append(cars, req)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"car":     req,
	})
}

func GetCarByID(ctx *gin.Context) {

	var carData Car

	carID, isExist := ctx.Params.Get("carID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "carId",
			"error_message": fmt.Sprintf("carId not exist"),
		})
		return
	}

	condition := false
	for i, car := range cars {
		carid, _ := strconv.Atoi(carID)
		if carid == car.CarID {
			carData = cars[i]
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("carID %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})

}

func UpdateCar(ctx *gin.Context) {
	var updatecarData Car

	carID, isExist := ctx.Params.Get("carID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "carId",
			"error_message": fmt.Sprintf("carId not exist"),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&updatecarData); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	condition := false
	for i, car := range cars {
		carid, _ := strconv.Atoi(carID)
		if carid == car.CarID {
			updatecarData.CarID = carid
			cars[i] = updatecarData
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("carID %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car":     updatecarData,
		"message": fmt.Sprintf("carID %v success updated", carID),
	})
}

func DeleteCar(ctx *gin.Context) {
	var carIndex int
	var lastcarData Car

	carID, isExist := ctx.Params.Get("carID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "carId",
			"error_message": fmt.Sprintf("carId not exist"),
		})
		return
	}

	condition := false
	for i, car := range cars {
		carid, _ := strconv.Atoi(carID)
		if carid == car.CarID {
			carIndex = i
			lastcarData = cars[i]
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("carID %v not found", carID),
		})
		return
	}

	copy(cars[carIndex:], cars[carIndex+1:])
	cars[len(cars)-1] = Car{}
	cars = cars[:len(cars)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"car":     lastcarData,
		"message": fmt.Sprintf("carID %v success deleted", carID),
	})
}
