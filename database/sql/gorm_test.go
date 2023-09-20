package sql

import (
	"fmt"
	"testing"

	"github.com/Ho-J/base/config"
	"github.com/Ho-J/base/logs"
	"github.com/Ho-J/base/viper_c"
)

func TestConnDB(t *testing.T) {
	viper_c.SetViper("../../config/default.yaml")
	c := config.Config{}

	viper_c.LoadConfig(&c)
	fmt.Printf("\n\nTestConnDB: %v", c)

	logs.InitLogger(logs.DefaultLogConfigs)

	InitDB(c.Mysqls)

	res := []map[string]interface{}{}
	if err := GetDB().Raw("select * from users limit ?", 10).Scan(&res).Error; err != nil {
		fmt.Printf("err: %s", err)
		return
	}
}
