package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

const (
	LocalDir = "/tmp"
)

type Config struct {
	Region string
	Bucket string
}

func NewConfig() *Config {
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	return &Config{
		Region: region,
		Bucket: bucket,
	}
}

func NewS3Downloader(config *Config) (*s3manager.Downloader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String((*config).Region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init s3 session, %v", err)
	}

	return s3manager.NewDownloader(sess), nil
}

func NewS3Uploader(config *Config) (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String((*config).Region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init s3 session, %v", err)
	}

	return s3manager.NewUploader(sess), nil
}
