package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", controllers.Health)
	r.GET("/generate", controllers.GeneratePassword)

	if err := r.Run("0.0.0.0:25000"); err != nil {
		panic(err)
	}
}
