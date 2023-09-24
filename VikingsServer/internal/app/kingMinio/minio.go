package kingMinio

import (
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

const (
	BucketName = "vikings-server"
)

func NewMinioClient(logger *logrus.Logger) *minio.Client {
	endpoint := "localhost:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio124"
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)

	if err != nil {
		logger.Fatalf("error: %s", err)
	}
	location := "us-east-1"

	err = minioClient.MakeBucket(BucketName, location)
	if err != nil {
		exists, err2 := minioClient.BucketExists(BucketName)
		if err2 == nil && exists {
			logger.Infof("We already own %s", BucketName)
		} else {
			logger.Fatalf("error: %s", err2)
		}
	} else {
		logger.Infof("Successfully created %s\n", BucketName)
	}

	return minioClient
}
