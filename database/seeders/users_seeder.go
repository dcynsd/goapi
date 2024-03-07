package seeders

import (
	"fmt"

	"goapi/app/models"
	"goapi/database/factories"
	"goapi/pkg/app"
	"goapi/pkg/console"
	"goapi/pkg/logger"
	"goapi/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	// 添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		// 创建 10 个用户对象
		users := factories.MakeUsers(10)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		var userModel models.User
		app.DB.Where("id", 1).First(&userModel)
		userModel.Name = "Administrator"
		userModel.Username = "admin"
		userModel.Password = "123456"
		userModel.Save()

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
