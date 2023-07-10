package controllers

import (
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/auth"
	"goapi/pkg/jwt"
	"goapi/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type AuthController struct{}

func (ctrl *AuthController) Store(c *gin.Context) {
	request := requests.AuthRequest{}
	if ok := requests.Validate(c, &request, requests.AuthSave); !ok {
		return
	}

	userModel, err := auth.Attempt(request.Username, request.Password)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	token := jwt.NewJWT().IssueToken(userModel.ID, userModel.Name)

	response.Data(c, gin.H{
		"access_token": token,
	})
}

func (ctrl *AuthController) Me(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	var me models.Me
	mapstructure.Decode(userModel, &me)

	response.Data(c, me)
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
