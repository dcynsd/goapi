package routes

import (
	"goapi/app/controllers"

	"github.com/gin-gonic/gin"
)

type WebRoute struct{}

func (webRoute *WebRoute) RegisterRoutes(r *gin.Engine) {

	home := new(controllers.HomeController)

	r.Any("/", home.Index)
}
