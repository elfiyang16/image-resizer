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
	Region     string
	Bucket     string
	Downloader *s3manager.Downloader
	Uploader   *s3manager.Uploader
}

func NewConfig() (*Config, error) {
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init s3 session, %v", err)
	}

	downloader := s3manager.NewDownloader(sess)
	uploader := s3manager.NewUploader(sess)
	return &Config{
		Region:     region,
		Bucket:     bucket,
		Downloader: downloader,
		Uploader:   uploader,
	}, nil
}
