package services

import (
	"goapi/app/models"
	"goapi/pkg/app"
)

type UserService struct {
	BaseService
}

func (s *UserService) All() []models.User {
	userModels := make([]models.User, 0)
	app.DB.Find(&userModels)
	return userModels
}
