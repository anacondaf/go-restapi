package wd

import (
	"log"
	"os"
)

func GetWorkDirectory() string {
	currentPath, err := os.Getwd()

	if err != nil {
		log.Panic(err)
	}

	return currentPath
}
