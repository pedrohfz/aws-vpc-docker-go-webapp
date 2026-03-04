package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"web/controllers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	url := os.Getenv("BACKEND_URL")
	if url == "" {
		url = "http://localhost:25000"
	}

	page := controllers.NewPageController(url)
	api := controllers.NewAPIController(url)

	r.GET("/", page.Index)
	r.GET("/api/generate", api.Generate)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
