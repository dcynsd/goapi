package app

import (
	"database/sql"
	"time"

	"goapi/pkg/config"

	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	SQL_DB *sql.DB
)

func IsLocal() bool {
	return config.Config.AppEnv == "local"
}

func IsProduction() bool {
	return config.Config.AppEnv == "production"
}

func IsTesting() bool {
	return config.Config.AppEnv == "testing"
}

// 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.Config.Timezone)
	return time.Now().In(chinaTimezone)
}
