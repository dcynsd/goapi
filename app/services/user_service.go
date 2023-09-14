package services

import (
	"goapi/app/models"
	"goapi/app/repositories"
)

type UserService struct {
	Repo repositories.UserRepo
}

func (s *UserService) GetList() []models.User {
	return s.Repo.GetList()
}

func (s *UserService) GetByID(id uint64) models.User {
	return s.Repo.GetByID(id)
}

func (s *UserService) GetByUsername(username string) models.User {
	return s.Repo.GetByUsername(username)
}
