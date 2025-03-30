package config

import "github.com/spf13/viper"

type Config struct {
	DenyIPList     []string `mapstructure:"denyIPList"`
	DenyHTTPBody   []string `mapstructure:"denyHTTPBody"`
	DenyHTTPHeader []string `mapstructure:"denyHTTPHeader"`
	Backend        []string `mapstructure:"backend"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
