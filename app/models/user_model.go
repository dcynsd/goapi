package models

import (
	"goapi/pkg/app"
	"goapi/pkg/hash"

	"gorm.io/gorm"
)

type User struct {
	BaseModel

	Username string `json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`

	*CommonTimestampsField
}

func (userModel *User) Save() (rowsAffected int64) {
	result := app.DB.Save(&userModel)
	return result.RowsAffected
}

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}

// ComparePassword 比对密码
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
