package models

import (
	"goapi/pkg/database"
	"goapi/pkg/hash"

	"gorm.io/gorm"
)

type User struct {
	BaseModel

	Username string `json:"username,omitempty"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar,omitempty"`

	CommonTimestampsField
}

type Me struct {
	BaseModel

	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}

func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
