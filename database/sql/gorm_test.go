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
	mysqls := config.Mysqls{}

	viper_c.LoadMysql(&mysqls)
	fmt.Printf(" TestConnDB: %v", mysqls)

	logs.InitLogger(logs.DefaultLogConfigs)
	InitDB(mysqls)

	res := []map[string]interface{}{}
	if err := GetDb().Raw("select * from users limit ?", 10).Scan(&res).Error; err != nil {
		fmt.Printf("err: %s", err)
		return
	}
}
