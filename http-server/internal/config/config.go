package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel   string     `yaml:"log_level" env:"LOG_LEVEL" env-default:"INFO"`
	HttpServer HttpServer `yaml:"http_server" env:",inline"`
}

type HttpServer struct {
	Host string `yaml:"host" env:"HTTP_SERVER_HOST"`
	Port int    `yaml:"port" env:"HTTP_SERVER_PORT"`
}

func LoadWithConfigFunc(cfg interface{}, f func() string) error {
	if f == nil {
		f = fetchConfigPath
	}

	configPath := f()
	if configPath != "" {
		if err := cleanenv.ReadConfig(configPath, cfg); err == nil {
			return nil
		}
	}

	return cleanenv.ReadEnv(cfg)
}

func Load(cfg *Config) error {
	return LoadWithConfigFunc(cfg, fetchConfigPath)
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config_path", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
