package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	DbDef  DbDef        `mapstructure:"db"`
}

type ServerConfig struct {
	ServerHost string `mapstructure:"host"`
	ServerPort uint16 `mapstructure:"port"`
	Timeout    int    `mapstructure:"timeout"`
}

type DbDef struct {
	DbType string `mapstructure:"db-type"`
}

func LoadConfig(configPath string) Config {
	var cfg Config

	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return cfg
}
