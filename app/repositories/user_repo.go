package repositories

import (
	"goapi/app/models"
	"goapi/pkg/database"
)

type UserRepo struct{}

func (r *UserRepo) GetList() (users []models.User) {
	database.DB.Select("id,name,avatar,created_at,updated_at").Find(&users)
	return
}

func (r *UserRepo) GetByID(id uint64) (userModel models.User) {
	database.DB.Where("id", id).First(&userModel)
	return
}

func (r *UserRepo) GetByUsername(username string) (userModel models.User) {
	database.DB.Where("username = ?", username).First(&userModel)
	return
}
