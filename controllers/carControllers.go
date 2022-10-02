package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID	string	`json:"car_id"`
	Brand	string	`json:"brand"`
	Model	string	`json:"model"`
	Price	int		`json:"price"`
}

var CarData	= []Car{

}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarData) + 1)
	CarData = append(CarData, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car" : newCar,
	})
}


func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarData {
		if carID == car.CarID {
			condition = true
			CarData[i] =  updateCar
			CarData[i].CarID = carID
			break 
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not Found!",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("Car with id %v has been successfully updated", carID),
	})
}


func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car 

	for i, car := range CarData {
		if carID == car.CarID {
			condition = true
			carData = CarData[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Cr with ID %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"car":carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarData {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Cr with ID %v not found", carID),
		})
		return
	}

	copy(CarData[carIndex:], CarData[carIndex + 1:])
	CarData[len(CarData) - 1] = Car{}
	CarData = CarData[:len(CarData) - 1]

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("Car with id %v has been successfully deleted", carID),
	})
}