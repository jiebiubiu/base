package minio_c

import (
	"context"
	"fmt"
	"testing"

	"github.com/Ho-J/base/viper_c"
)

func TestMinio(t *testing.T) {
	viperC := viper_c.NewViperC("../config/default.yaml")
	viperC.LoadConfig()
	c := viperC.GetConfig()

	fmt.Printf("%+v", c)

	InitMinio(c.Minio)

	bks, err := MinioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	for _, bk := range bks {
		fmt.Println(bk)
	}
}
