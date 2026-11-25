package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  int
	DatabaseDSN string
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetDefault("server.port", 8080)
	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) && !os.IsNotExist(err) {
			return nil, err
		}
	}
	return &Config{
		ServerPort:  viper.GetInt("server.port"),
		DatabaseDSN: viper.GetString("database.dsn"),
	}, nil
}
