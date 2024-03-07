package auth

import (
	"errors"

	"goapi/app/constants"
	"goapi/app/models"
	"goapi/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 从 gin.context 中获取当前登录用户
func User(c *gin.Context) models.User {
	userModel, ok := c.MustGet(constants.AuthUser).(models.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return models.User{}
	}
	// db is now a *DB value
	return userModel
}

// 从 gin.context 中获取当前登录用户 ID
func UserID(c *gin.Context) uint64 {
	if value, exists := c.Get(constants.AuthUserID); exists {
		userID, _ := value.(uint64)
		return userID
	}
	return 0
}
