package services

import (
	"goapi/app/constants"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/app"
)

type AuthService struct {
	BaseService
}

// Login 用户登录
func (s *AuthService) Login(request requests.AuthRequest) (models.User, error) {
	var userModel models.User
	app.DB.Where("username = ?", request.Username).First(&userModel)
	if userModel.ID == 0 {
		return userModel, s.Error.SetCustomStatusCode(constants.Unauthorized).SetUnauthorizedError("账号不存在")
	}

	if !userModel.ComparePassword(request.Password) {
		return userModel, s.Error.SetUnauthorizedError("密码错误")
	}

	return userModel, nil
}
