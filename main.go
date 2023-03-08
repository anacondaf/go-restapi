package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
	"io"
	"kai/config"
	"kai/helper/time"
	"kai/helper/wd"
	"kai/internal/aws"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

func init() {
	config.LoadEnvVars()
	awsService.NewS3Client()
	config.LoadConfig()
}

func main() {
	var wg sync.WaitGroup

	// Create a new engine
	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
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

		fmt.Println(reflect.TypeOf(file))

		if err != nil {
			return err
		}

		// Get Buffer from file
		buffer, err := file.Open()

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

		err = os.WriteFile(fmt.Sprintf("%v/images/%v-%v", wd.GetWorkDirectory(), timer.GetCurrentTime().Unix(), file.Filename), fileData, 0666)
		if err != nil {
			return err
		}

		wg.Add(1)

		// Upload to S3
		go func() {
			defer wg.Done()

			awsService.PutObject(&s3.PutObjectInput{
				Bucket: aws.String(config.GetEnv("S3_BUCKET")),
				Key:    aws.String(fmt.Sprintf("avatar/%v-%v", timer.GetCurrentTime().Unix(), file.Filename)),
				Body:   bytes.NewReader(fileData),
			})
		}()

		return c.Status(200).JSON(fiber.Map{
			"message": "OK",
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("hello_world", fiber.Map{
			"name": "1",
		})
	})

	app.Get("/avatar/metadata", func(c *fiber.Ctx) error {
		var url = c.Query("url")

		unixSec, err := strconv.ParseInt(url, 10, 64)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "success",
			"data":    time.Unix(unixSec, 0),
			"error":   false,
		})
	})

	app.Get("/avatars", func(c *fiber.Ctx) error {
		// Get the first page of results for ListObjectsV2 for a bucket
		output, err := awsService.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String(fmt.Sprintf(config.GetEnv("S3_BUCKET"))),
			Prefix: aws.String("avatar"),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Println("first page results:")

		//for _, object := range output.Contents {
		//	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
		//}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data": output.Contents,
		})
	})

	app.Listen(":3000")
}
