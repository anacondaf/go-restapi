package awsService

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

var S3Client *s3.Client

func NewS3Client() {

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	S3Client = s3.NewFromConfig(cfg)
}

func PutObject(cfg *s3.PutObjectInput) {
	_, err := S3Client.PutObject(context.TODO(), cfg)

	if err != nil {
		log.Panic(err)
	}
}
