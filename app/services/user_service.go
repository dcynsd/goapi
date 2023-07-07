package services

import (
	"goapi/app/models"
	"goapi/pkg/database"
)

type UserService struct{}

func (s *UserService) GetList() (users []models.User) {
	database.DB.Select("id,name,avatar,created_at").Find(&users)
	return
}

func (s *UserService) GetByID(id uint64) (userModel models.User) {
	database.DB.Where("id", id).First(&userModel)
	return
}

func (s *UserService) GetByUsername(username string) (userModel models.User) {
	database.DB.Where("username = ?", username).First(&userModel)
	return
}
