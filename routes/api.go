package routes

import (
	"goapi/app/controllers/api"
	"goapi/app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiRoute struct{}

func (apiRoute *ApiRoute) RegisterRoutes(r *gin.Engine) {

	apiGroup := r.Group("/api")
	{
		apiGroup.Any("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "This is api route!",
			})
		})

		auth := new(api.AuthController)
		apiGroup.POST("/auth/login", auth.Store)

		apiGroup.Use(middlewares.AuthJWT())
		{
			apiGroup.GET("/auth/me", auth.Me)
			apiGroup.POST("/auth/refresh-token", auth.RefreshToken)

			users := new(api.UserController)
			apiGroup.GET("/users", users.Index)
		}

	}
}
