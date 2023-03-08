package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type DBConfig struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPass     string `mapstructure:"DB_PASS"`
}

type ServerConfig struct {
}

type Config struct {
	DB     DBConfig     `mapstructure:"DB"`
	SERVER ServerConfig `mapstructure:"SERVER"`
}

func LoadConfig() {
	viper.SetConfigName("app_settings")
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when loading config: %v", err)
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Error when unmarshalling config: %v", err)
	}

	fmt.Printf("%+v\n", c)
}
