package bootstrap

import (
	"fmt"
	"time"

	"goapi/pkg/config"
	"goapi/pkg/database"
	"goapi/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() {

	var dbConfig gorm.Dialector
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Config.DBUsername,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBDatabase,
		config.Config.DBCharset,
	)
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})
	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())
	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(100)
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(300) * time.Second)
}
