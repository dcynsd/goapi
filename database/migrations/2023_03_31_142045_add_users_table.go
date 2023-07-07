package migrations

import (
	"database/sql"

	"goapi/app/models"
	"goapi/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		Username string `gorm:"column:username;type:varchar(32);not null;unique;"`
		Password string `gorm:"column:password;type:varchar(255);not null;"`
		Name     string `gorm:"column:name;type:varchar(50);not null;"`
		Avatar   string `gorm:"column:avatar;type:varchar(255);default:null;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2023_03_31_142045_add_users_table", up, down)
}
