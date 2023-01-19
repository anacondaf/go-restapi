package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func init() {
	var env map[string]string
	env, err := godotenv.Read(".env")

	if err != nil {
		log.Error("Error loading .env file")
	}

	log.Info(env["S3_BUCKET"])
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello World")

		if err != nil {
			return err
		}

		return nil
	})

	app.Post("/avatar", func(c *fiber.Ctx) error {
		file, err := c.FormFile("avatar")
		if err != nil {
			return err
		}

		// Get Buffer from file
		buffer, err := file.Open()

		currentPath, err := os.Getwd()
		if err != nil {
			return err
		}

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"error":   1,
				"message": err.Error(),
			})
		}

		fileData, err := io.ReadAll(buffer)
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%v/images/%v", currentPath, file.Filename), fileData, 0666)
		if err != nil {
			return err
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "OK",
		})
	})

	app.Listen(":3000")
}
