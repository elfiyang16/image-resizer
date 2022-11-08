package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"
)

// DownloadFile downloads file from s3 bucket
func DownloadFile(config *Config, objectKey string) (string, error) {
	downloader, err := NewS3Downloader(config)
	if err != nil {
		return "", fmt.Errorf("failed to init s3 downloader, %v", err)
	}
	tmpFilePath := LocalDir + "/" + filepath.Base(objectKey)
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file, %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("failed to close file", err)
		}
	}()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(config.Bucket),
			Key:    aws.String(objectKey),
		})

	if err != nil {
		return "", fmt.Errorf("failed to download item %q, %v", objectKey, err)
	}
	// Logger
	fmt.Println("succesful downloaded file:", file.Name(), numBytes, "bytes")

	return tmpFilePath, nil
}

// Resizer

//func UploadFiles(config *Config, files []string) {
//	for _, f := range files {
//		uploadFile(config, f)
//	}
//}

// uploads resized image to bucket
func uploadFile(config *Config, fileName string) error {
	uploader, err := NewS3Uploader(config)
	if err != nil {
		return fmt.Errorf("failed to init s3 uploader, %v", err)
	}
	//tmpFilePath := LocalDir + "/" + fileName
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file %v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("failed to close file", err)
		}
	}()

	objectKey := strings.Split(fileName, "/")[1] // /tmp/xxx -> [tmp, xxx]
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file %s, %v", fileName, err)
	}
	return nil
}
