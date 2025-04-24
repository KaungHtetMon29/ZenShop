package pkg

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Config struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	session         *session.Session
}

func NewS3Config() *S3Config {
	config := &S3Config{
		Region:          "ap-southeast-1",
		AccessKeyID:     "AKIAYHMTXAPOS75MZASH",
		SecretAccessKey: "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",
		BucketName:      "zenshopkmd",
	}

	var err error
	config.session, err = session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyID,
			config.SecretAccessKey,
			""),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
	}

	return config
}

func (awsS3 *S3Config) S3ImageUpload(fileBytes []byte, fileName string) (string, error) {
	bucketName := awsS3.BucketName
	// Create a unique key using filename and timestamp
	fileNameWithoutExt := strings.Split(fileName, ".")[0]
	fileExt := fileName
	if len(strings.Split(fileName, ".")) > 1 {
		fileExt = "." + strings.Split(fileName, ".")[1]
	} else {
		fileExt = ""
	}
	s3Key := fmt.Sprintf("%s_%d%s", fileNameWithoutExt, time.Now().Unix(), fileExt)

	// Upload the file to S3

	svc := s3.New(awsS3.session)
	// Determine content type based on file extension
	contentType := "application/octet-stream" // default
	if fileExt != "" {
		switch strings.ToLower(fileExt) {
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		case ".gif":
			contentType = "image/gif"
		case ".webp":
			contentType = "image/webp"
		case ".pdf":
			contentType = "application/pdf"
		}
	}

	uploadInput := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3Key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	}
	fmt.Println("upload file")
	_, err := svc.PutObject(uploadInput)
	fmt.Println("put object")
	if err != nil {
		return "", err
	}
	s3URL := fmt.Sprintf("https://%s.s3.ap-southeast-1.amazonaws.com/%s", bucketName, s3Key)
	fmt.Println(s3URL)
	return s3URL, nil
}

func (awsS3 *S3Config) S3ImageDelete(s3Key string) error {
	bucketName := awsS3.BucketName

	svc := s3.New(awsS3.session)

	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(s3Key),
	}

	_, err := svc.DeleteObject(deleteInput)
	if err != nil {
		return err
	}
	return nil
}
