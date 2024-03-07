package api

import (
	"goapi/app/models"
	"goapi/app/services"
	"goapi/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Context *gin.Context
	Error   *models.Error
	UserID  uint64
}

func (a *Api) MakeContext(c *gin.Context) *Api {
	a.Context = c
	a.UserID = auth.UserID(c)
	a.makeError()
	return a
}

func (a *Api) MakeService(s *services.BaseService) *Api {
	s.Error = a.Error
	s.UserID = a.UserID
	return a
}

func (a *Api) makeError() *Api {
	a.Error = &models.Error{
		Err:        nil,
		StatusCode: 500,
	}
	return a
}
