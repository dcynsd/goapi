package bootstrap

import (
	"fmt"

	"goapi/pkg/config"
	"goapi/pkg/redis"
)

func SetupRedis() {

	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.Config.RedisHost, config.Config.RedisPort),
		config.Config.RedisUsername,
		config.Config.RedisPassword,
		config.Config.RedisDatabase,
	)
}
