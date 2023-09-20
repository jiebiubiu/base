package minio_c

import (
	"log"

	"github.com/Ho-J/base/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	PublicBuck  = "pub"
	PrivateBuck = "private"
)

var MinioClient *minio.Client

func NewMinioClient(minioC config.Minio) *minio.Client {
	useSSL := false
	// Initialize minio client object.
	client, err := minio.New(minioC.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioC.AccessKeyID, minioC.SecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func InitMinio(minioC config.Minio) {
	MinioClient = NewMinioClient(minioC)
}
