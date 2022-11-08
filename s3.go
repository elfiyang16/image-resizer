package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
)

func DownloadFile(config *Config, writer io.WriterAt, objectKey string) error {
	_, err := config.Downloader.Download(writer,
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
	_, err := config.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(objectKey),
		Body:   input,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file %s, %v", objectKey, err)
	}
	return nil
}
