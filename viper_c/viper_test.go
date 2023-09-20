package viper_c

import (
	"fmt"
	"testing"
	"time"

	"github.com/Ho-J/base/config"
)

func TestLoadConfig(t *testing.T) {
	SetViper("../config/default.yaml")

	mysqls := config.Mysqls{}
	if err := LoadMysql(&mysqls); err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	for {
		fmt.Println(mysqls)
		time.Sleep(2 * time.Second)
	}
}
