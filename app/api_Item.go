package app

import (
	"GRUPO-WELEARN/API-PROYECTO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ItemsIndex(c *gin.Context) {
	s := models.Item{Title: "User", Notes: ""}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hola Welcome " + s.Title,
	})
}
