package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/services"
)

// Health() - Camada HTTP
// Responsável por receber a requisição e registrar o Health Check da aplicação.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// GeneratePassword() - Camada HTTP
// Responsável por receber a requisição, chamar a regra de negócio e devolver resposta ao cliente.
func GeneratePassword(c *gin.Context) {
	password, err := services.GeneratePassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"password": password,
	})
}
