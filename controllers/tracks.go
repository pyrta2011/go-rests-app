package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest-app/models"
	"net/http"
)

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title string `json:"title" binding:"required"`
}

func GetAllTracks(ctx *gin.Context)  {
	var tracks []models.Track
	models.DB.Find(&tracks)

	ctx.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

func CreateTrack(ctx *gin.Context)  {
	var input CreateTrackInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}
	models.DB.Create(&track)

	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
func DeleteTrack(ctx *gin.Context) {
	var track models.Track
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&track)

	ctx.JSON(http.StatusOK, gin.H{"track": true})
}
