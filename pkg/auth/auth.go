package auth

import (
	"errors"

	"goapi/app/models"
	"goapi/app/services"
	"goapi/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) models.User {
	userModel, ok := c.MustGet("current_user").(models.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return models.User{}
	}
	// db is now a *DB value
	return userModel
}

// 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}

// 账号密码登录
func Attempt(username string, password string) (models.User, error) {
	s := new(services.UserService)
	userModel := s.GetByUsername(username)
	if userModel.ID == 0 {
		return models.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return models.User{}, errors.New("密码错误")
	}

	return userModel, nil
}
