package controllers

import (
	"goapi/app/services"
	"goapi/pkg/api"
	"goapi/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	api.Api
	Service services.UserService
}

func (ctrl *UserController) Index(c *gin.Context) {
	ctrl.MakeContext(c).MakeService(&ctrl.Service.BaseService)
	response.Data(c, ctrl.Service.All())
}
