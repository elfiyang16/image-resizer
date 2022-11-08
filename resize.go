package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/h2non/bimg"
)

const (
	Quality     = 60 // Default is 75/100
	Compression = 5  // Default is 6/10

)

func Resize(writer *aws.WriteAtBuffer, objectKey string) ([]byte, error) {
	converted, err := bimg.NewImage(writer.Bytes()).Convert(bimg.JPEG) // get jpeg -> bytes -> convert jpeg
	if err != nil {
		return nil, fmt.Errorf("failed to convert img %s, %v", objectKey, err)
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{
		Quality:     Quality,
		Compression: Compression,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to process img %s, %v", objectKey, err)
	}

	//err = bimg.Write(fmt.Sprintf(LocalDir+"/%s", objectKey), processed)
	//if err != nil {
	//	return fmt.Errorf("failed to write processed img to file %s, %v", objectKey, err)
	//}

	fmt.Println("Vips Memory Alloc", bimg.VipsMemory())

	return processed, nil
}
