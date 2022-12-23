package main

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	var env map[string]string
	env, err := godotenv.Read(".env")

	if err != nil {
		log.Error("Error loading .env file")
	}

	fmt.Println(reflect.TypeOf(env))
	log.Info(env["S3_BUCKET"])
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		error := c.SendString("Hello World")

		if error != nil {
			return error
		}

		return nil
	})

	app.Listen(":3000")
}
