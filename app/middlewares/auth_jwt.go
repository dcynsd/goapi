package middlewares

import (
	"fmt"

	"goapi/app/constants"
	"goapi/app/models"
	"goapi/pkg/app"
	"goapi/pkg/config"
	"goapi/pkg/jwt"
	"goapi/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.Config.AppName))
			return
		}

		var userModel models.User
		app.DB.Where("id", claims.UserID).First(&userModel)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set(constants.AuthUserID, userModel.ID)
		c.Set(constants.AuthUser, userModel)

		c.Next()
	}
}
