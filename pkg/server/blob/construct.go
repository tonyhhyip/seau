package blob

import (
	"github.com/minio/minio-go"
	"github.com/tonyhhyip/seau/pkg/server/config"
)

func NewFromConfig(config *config.Config) (m *MinioManager, err error) {
	blob := config.Blob
	client, err := minio.New(blob.Endpoint, blob.AccessKey, blob.SecretKey, blob.Secure)
	if err != nil {
		return
	}

	return NewFromClient(client, blob.BucketName), nil
}

func New(endpoint, accessKey, secretKey, bucketName string, secure bool) (m *MinioManager, err error) {
	client, err := minio.New(endpoint, accessKey, secretKey, secure)
	if err != nil {
		return
	}

	m = NewFromClient(client, bucketName)

	return
}

func NewFromClient(client *minio.Client, bucketName string) *MinioManager {
	return &MinioManager{
		Client:     client,
		BucketName: bucketName,
	}
}
