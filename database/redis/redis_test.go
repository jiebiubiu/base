package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Ho-J/base/viper_c"
)

func TestRedisClient(t *testing.T) {
	ic := viper_c.NewViperC("../../config/default.yaml")
	ic.LoadConfig()
	c := ic.GetConfig()

	InitRedis(c.RedisConns)

	sta := GetDB().Set(context.Background(), "hello", "world!", 3*time.Second)
	if sta.Err() != nil {
		fmt.Printf("err: %v", sta.Err())
		return
	}

	scmd := GetDB().Get(context.Background(), "hello")
	if scmd.Err() != nil {
		fmt.Printf("err: %v", sta.Err())
		return
	}
	fmt.Println(scmd.Val())
}
