package main

import (
	"Assignment-02/config"
	"Assignment-02/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DbInit()
	inDB   := &controllers.InDB{DB : db}

	r := gin.Default()
	
	r.POST("/order", inDB.PostData)
	r.PUT("/order/:orderId", inDB.UpdateData)
	r.GET("/orders", inDB.GetAllData)
	r.DELETE("/order/:id", inDB.DeleteData)

	r.Run(":3000")
}