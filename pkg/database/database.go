package database

import (
	"fmt"
	"goapi/pkg/app"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {

	var err error
	app.DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	app.SQL_DB, err = app.DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = app.DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	return deleteMySQLTables()
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	tables := []string{}

	// 读取所有数据表
	err := app.DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	app.DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err := app.DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	// 开启 MySQL 外键检测
	app.DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
