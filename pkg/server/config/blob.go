package config

import "os"

type Blob struct {
	Endpoint   string
	Secure     bool
	AccessKey  string
	SecretKey  string
	BucketName string
}

func blobFromEnv() *Blob {
	endpoint := os.Getenv("S3_ENDPOINT")
	if endpoint == "" {
		// Use S3 for default
		endpoint = "s3.amazonaws.com"
	}

	return &Blob{
		Endpoint:   endpoint,
		Secure:     os.Getenv("S3_INSECURE") != "true",
		AccessKey:  os.Getenv("S3_ACCESS_KEY_ID"),
		SecretKey:  os.Getenv("S3_SECRET_ACCESS_KEY"),
		BucketName: os.Getenv("S3_BUCKET_NAME"),
	}
}
