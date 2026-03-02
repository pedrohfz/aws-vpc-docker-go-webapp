package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// main() - Ponto de entrada da aplicação.
// Responsável por subir o servidor HTTP e registrar as rotas da aplicação.
func main() {
	r := gin.Default()

	r.GET("/health", controllers.Health)
	r.GET("/generate", controllers.GeneratePassword)

	if err := r.Run("0.0.0.0:25000"); err != nil {
		panic(err)
	}
}
