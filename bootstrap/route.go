package bootstrap

import (
	"net/http"
	"strings"

	"goapi/app/middlewares"
	"goapi/routes"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	registerGlobalMiddleWare(router)

	registerApiRoutes(router)

	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func registerApiRoutes(router *gin.Engine) {
	apiRoutes := routes.ApiRoute{}
	apiRoutes.RegisterRoutes(router)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
