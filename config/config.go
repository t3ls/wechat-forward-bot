package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	ForwardTargetUsername string `json:"forward_target_username"   mapstructure:"forward_target_username"   yaml:"forward_target_username"`
	keyword               string `json:"keyword,omitempty"   mapstructure:"keyword"   yaml:"keyword,omitempty"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

func GetForwardTargetUsername() string {
	username := getEnv("forward_target_username")

	if username != "" {
		return username
	}

	if config == nil {
		return ""
	}

	if username == "" {
		username = config.ForwardTargetUsername
	}
	return username
}

func GetWechatKeyword() string {
	keyword := getEnv("keyword")

	if keyword != "" {
		return keyword
	}

	if config == nil {
		return ""
	}

	if keyword == "" {
		keyword = config.keyword
	}
	return keyword
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = os.Getenv(strings.ToUpper(key))
	}

	if len(value) > 0 {
		return value
	}

	if config == nil {
		return ""
	}

	if len(value) > 0 {
		return value
	}

	if config.keyword != "" {
		value = config.keyword
	}
	return ""
}
