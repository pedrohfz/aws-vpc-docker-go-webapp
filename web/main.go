package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// url := os.Getenv("BACKEND_URL")

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
