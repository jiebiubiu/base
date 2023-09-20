package viper_c

import (
	"fmt"
	"testing"
	"time"

	"github.com/Ho-J/base/config"
)

func TestLoadConfig(t *testing.T) {
	SetViper("../config/default.yaml")

	c := config.Config{}
	if err := LoadConfig(&c); err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	for {
		fmt.Println(c)
		time.Sleep(2 * time.Second)
	}
}
