package pkg

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
	"strings"
	c "textgopher/config"
)

type AwsClient struct {
	Client *session.Session
	Err    error
}

func GetSession() *AwsClient {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			c.Configure().AWSAccessKeyId,
			c.Configure().AWSSecretAccessKey,
			"",
		),
	})
	if err != nil {
		log.Fatal("Error creating session: %s", err)
	}
	client := &AwsClient{
		Client: sess,
		Err:    err,
	}
	return client
}

func (client *AwsClient) CreateBucket(bucketName string) error {
	svc := s3.New(client.Client)
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func (client *AwsClient) ListBuckets() ([]*s3.Bucket, error) {
	svc := s3.New(client.Client)
	result, err := svc.ListBuckets(nil)
	return result.Buckets, err
}

func (client *AwsClient) DeleteBucket(bucketName string) error {
	svc := s3.New(client.Client)
	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	return err
}

func (client *AwsClient) UploadFile(bucketName string, filePath string) error {
	var fileName string
	svc := s3.New(client.Client)

	file, err := os.Open(filePath)
	fileName = file.Name()[strings.LastIndex(file.Name(), "/")+1:]

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   aws.ReadSeekCloser(file),
	})
	return err
}
