package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageController struct {
	url string
}

func NewPageController(url string) *PageController {
	return &PageController{url: url}
}

func (page *PageController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"url": page.url,
	})
}
