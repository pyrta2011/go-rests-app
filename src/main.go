package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-app/controllers"
	"go-rest-app/models"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack)
	route.DELETE("/tracks/:id", controllers.DeleteTrack)

	err := route.Run()
	if err != nil {
		return 
	}
}
