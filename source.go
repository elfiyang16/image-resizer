package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
)

func DownloadFile(config *Config, writer io.WriterAt, objectKey string) error {
	downloader, err := NewS3Downloader(config)
	if err != nil {
		return fmt.Errorf("failed to init s3 downloader, %v", err)
	}

	imgBytes, err := downloader.Download(writer,
		&s3.GetObjectInput{
			Bucket: aws.String(config.Bucket),
			Key:    aws.String(objectKey),
		})

	if err != nil {
		return fmt.Errorf("failed to download item %q, %v", objectKey, err)
	}
	return nil
}

func UploadFile(input io.Reader, config *Config, objectKey string) error {
	uploader, err := NewS3Uploader(config)
	if err != nil {
		return fmt.Errorf("failed to init s3 uploader, %v", err)
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(objectKey),
		Body:   input,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file %s, %v", objectKey, err)
	}
	return nil
}
