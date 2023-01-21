package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getWorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	return wd
}

func LoadEnvVars() {
	err := godotenv.Load(fmt.Sprintf("%v/.env", getWorkDir()))

	if err != nil {
		log.Panic("Error when loading .env file")
	}
}

func GetEnv(key string) (string, error) {
	env := os.Getenv(key)

	if len(env) == 0 {
		return "", errors.New(fmt.Sprintf("Cannot load env value of key %v", key))
	}

	return env, nil
}
