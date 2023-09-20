package minio_c

import (
	"context"
	"fmt"
	"testing"

	"github.com/Ho-J/base/config"
	"github.com/Ho-J/base/viper_c"
)

func TestMinio(t *testing.T) {
	viper_c.SetViper("../config/default.yaml")
	minioC := config.Minio{}
	viper_c.LoadMinio(&minioC)
	InitMinio(minioC)

	bks, err := MinioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	for _, bk := range bks {
		fmt.Println(bk)
	}
}
