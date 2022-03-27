package service

import "github.com/aws/aws-sdk-go/service/s3"

type S3I interface {
	CreateBucket(bucketName string) error
	ListBuckets() ([]*s3.Bucket, error)
	DeleteBucket(bucketName string) error
	UploadFile(bucketName string, filePath string) error
}
