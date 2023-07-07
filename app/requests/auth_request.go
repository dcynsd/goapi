package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AuthRequest struct {
	Username string `valid:"username" form:"username"`
	Password string `valid:"password" form:"password"`
}

func AuthSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"username": []string{"required", "min:1", "max:32"},
		"password": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"username": []string{
			"required:账号不能为空",
			"min:账号长度至少 1 位",
			"max:账号长度不能超过 32 位",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码长度至少 6 位",
		},
	}

	return InternalValidate(data, rules, messages)
}
