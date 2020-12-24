package app

import (
	"GRUPO-WELEARN/API-PROYECTO/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TrabajoIndex(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	lis := []models.Trabajo{}
	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func TrabajoCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	var d models.Trabajo
	//d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func TrabajoGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Trabajo
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}

func TrabajoUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Trabajo
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}

func TrabajoDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Trabajo

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
