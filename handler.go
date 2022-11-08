package main

import (
	"bytes"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"log"
)

func Handler(ctx context.Context, s3Event events.S3Event) {
	config := NewConfig()
	var writer aws.WriteAtBuffer

	for _, record := range s3Event.Records {
		objectKey := record.S3.Object.Key
		if err := DownloadFile(config, &writer, objectKey); err != nil {
			log.Println("failed to download image with objectkey %s, %v", objectKey, err)
		}

		resized, err := Resize(&writer, objectKey)
		if err != nil {
			log.Println("failed to resize image with objectkey %s, %v", objectKey, err)
		}

		reader := bytes.NewReader(resized)
		if err := UploadFile(reader, config, objectKey); err != nil {
			log.Println("failed to upload image with objectkey %s, %v", objectKey, err)
		}
	}
}
