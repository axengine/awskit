package awskit

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func (ak *AWSKit) UploadFile(bucketName string, objectKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = ak.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	return err
}

func (ak *AWSKit) UploadBuf(bucketName string, objectKey string, buf []byte, size int64, contentType string) error {
	_, err := ak.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(objectKey),
		Body:               bytes.NewReader(buf),
		ContentLength:      size,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("attachment"),
	})
	return err
}

func (ak *AWSKit) DeleteOBJ(bucketName string, objectKey string) error {
	_, err := ak.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	return err
}
