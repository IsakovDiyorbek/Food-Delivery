package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() (*minio.Client, error) {
    endpoint := "127.0.0.1:9000"
    accessKeyID := "pen0FH8fIYhbCcT0zCGK"
    secretAccessKey := "q73ArcpEWn3bypibpFu0wcW7rsC6kyHQ8u9yOVUD"
	useSSL := false
	
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}	

	return minioClient, nil
}
