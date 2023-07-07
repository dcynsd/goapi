package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"goapi/pkg/console"
	"goapi/pkg/helpers"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/joho/godotenv"
)

var Config *AppConfig
var once sync.Once

type AppConfig struct {
	AppName  string `config:"app_name"`
	AppEnv   string `config:"app_env"`
	AppDebug bool   `config:"app_debug"`
	AppPort  string `config:"app_port"`
	AppKey   string `config:"app_key"`
	AppUrl   string `config:"app_url"`

	Timezone string `config:"timezone"`

	DBHost     string `config:"db_host"`
	DBPort     string `config:"db_port"`
	DBDatabase string `config:"db_database"`
	DBUsername string `config:"db_username"`
	DBPassword string `config:"db_password"`
	DBCharset  string `config:"db_charset"`

	RedisHost     string `config:"redis_host"`
	RedisPort     string `config:"redis_port"`
	RedisDatabase int    `config:"redis_database"`
	RedisUsername string `config:"redis_username"`
	RedisPassword string `config:"redis_password"`

	LogLevel     string `config:"log_level"`
	LogType      string `config:"log_type"`
	LogFilename  string `config:"log_filename"`
	LogMaxSize   int    `config:"log_max_size"`
	LogMaxBackup int    `config:"log_max_backup"`
	LogMaxAge    int    `config:"log_max_age"`
	LogCompress  bool   `config:"log_compress"`

	JwtExpireTime      int64 `config:"jwt_expire_time"`
	JwtDebugExpireTime int64 `config:"jwt_debug_expire_time"`
	JwtMaxRefreshTime  int64 `config:"jwt_max_refresh_time"`

	PerPage         int    `config:"per_page"`
	UrlQueryPage    string `config:"url_query_page"`
	UrlQuerySort    string `config:"url_query_sort"`
	UrlQueryOrder   string `config:"url_query_order"`
	UrlQueryPerPage string `config:"url_query_per_page"`
}

func SetupConfig(commandKey string) {
	once.Do(func() {
		loadConfig(commandKey)
	})
}

func loadConfig(commandKey string) {
	err := godotenv.Load(".env")
	console.ExitIf(err)

	for _, env := range os.Environ() {
		key := env[:strings.Index(env, "=")]
		value := os.Getenv(key)
		os.Setenv(key, value)
	}

	config.WithOptions(config.ParseEnv)
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})

	config.AddDriver(yaml.Driver)

	configFiles, err := getAllFilesInDirectory("config")
	console.ExitIf(err)

	err = config.LoadFiles(configFiles...)
	console.ExitIf(err)

	Config = &AppConfig{}
	err = config.Decode(&Config)
	console.ExitIf(err)

	if commandKey != "goapi key" && helpers.Empty(Config.AppKey) {
		console.Exit("Please execute the command `go run main.go key` to generate app_key!")
	}
}

func getAllFilesInDirectory(directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
