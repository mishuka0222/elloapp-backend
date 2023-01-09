package minio_util

import (
	"log"

	"github.com/minio/minio-go"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

// endpoint := "127.0.0.1:9000"
// accessKeyID := "Q3AM3UQ867SPQQA43P2F"
// secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
// useSSL := true

func MustNewMinioClient(c *MinioConfig) *minio.Core {
	core, err := minio.NewCore(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		log.Fatal(err)
	}
	return core
}
