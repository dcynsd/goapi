package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (ctrl *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"content": "Welcome to GOAPI!",
	})
}
