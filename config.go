package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/umtkas/image-resizer-lambda/configs"
)

func New3Downloader(configuration configs.Configuration) (*s3manager.Downloader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(configuration.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init s3 session, %v", err)
	}

	return s3manager.NewDownloader(sess), nil
}

func NewS3Uploader(configuration configs.Configuration) (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(configuration.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init s3 session, %v", err)
	}

	return s3manager.NewUploader(sess), nil
}
