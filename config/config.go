package config

import "github.com/spf13/viper"

type Config struct {
	DenyIPList     []string `mapstructure:"denyIPList"`
	DenyHTTPBody   []string `mapstructure:"denyHTTPBody"`
	DenyHTTPHeader []string `mapstructure:"denyHTTPHeader"`
	Backend        []string `mapstructure:"backend"`
}

func NewConfig() *Config {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
