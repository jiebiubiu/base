package viper_c

import (
	"fmt"
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	viperC := NewViperC("../config/default.yaml")
	if err := viperC.LoadConfig(); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	c := viperC.GetConfig()

	// fmt.Printf("%+v\n", c)
	for {
		fmt.Println(c)
		time.Sleep(2 * time.Second)
	}
}
