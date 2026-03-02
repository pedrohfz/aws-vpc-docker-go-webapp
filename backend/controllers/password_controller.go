package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/services"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func GeneratePassword(c *gin.Context) {
	pass, err := services.GeneratePassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"password": pass,
	})
}
