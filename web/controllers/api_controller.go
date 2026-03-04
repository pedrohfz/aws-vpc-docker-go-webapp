package controllers

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type APIController struct {
	url    string
	client *http.Client
}

func NewAPIController(url string) *APIController {
	return &APIController{
		url:    url,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (API *APIController) Generate(c *gin.Context) {
	resp, err := API.client.Get(API.url + "/generate")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "falha ao acessar o backend"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
