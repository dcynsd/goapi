package controllers

import (
	"goapi/app/requests"
	"goapi/app/services"
	"goapi/pkg/api"
	"goapi/pkg/auth"
	"goapi/pkg/jwt"
	"goapi/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	api.Api
	Service services.AuthService
}

func (ctrl *AuthController) Store(c *gin.Context) {
	request := requests.AuthRequest{}
	if ok := requests.Validate(c, &request, requests.AuthSave); !ok {
		return
	}

	ctrl.MakeContext(c).MakeService(&ctrl.Service.BaseService)

	userModel, err := ctrl.Service.Login(request)
	if err != nil {
		response.AbortWithStatus(c, ctrl.Error)
		return
	}

	token := jwt.NewJWT().IssueToken(userModel.ID, userModel.Name)

	response.Data(c, gin.H{
		"access_token": token,
	})
}

func (ctrl *AuthController) Me(c *gin.Context) {
	userModel := auth.User(c)

	response.Data(c, gin.H{
		"id":     userModel.ID,
		"name":   userModel.Name,
		"avatar": userModel.Avatar,
	})
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Unauthorized(c, "令牌刷新失败")
		return
	}

	response.Data(c, gin.H{
		"access_token": token,
	})
}
