package bootstrap

import (
	"goapi/pkg/config"
	"goapi/pkg/logger"
)

func SetupLogger() {

	logger.InitLogger(
		config.Config.LogFilename,
		config.Config.LogMaxSize,
		config.Config.LogMaxBackup,
		config.Config.LogMaxAge,
		config.Config.LogCompress,
		config.Config.LogType,
		config.Config.LogLevel,
	)
}
