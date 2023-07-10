package controllers

import (
	"goapi/app/services"
	"goapi/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service services.UserService
}

func (ctrl *UserController) Index(c *gin.Context) {
	response.Data(c, ctrl.Service.GetList())
}
