package sql

import (
	"fmt"
	"testing"

	"github.com/Ho-J/base/logs"
	"github.com/Ho-J/base/viper_c"
)

func TestConnDB(t *testing.T) {
	viperC := viper_c.NewViperC("../../config/default.yaml")
	viperC.LoadConfig()
	c := viperC.GetConfig()

	fmt.Printf("\n\nTestConnDB: %v", c)

	logs.InitLogger(c.Log)

	InitDB(c.Mysqls)

	res := []map[string]interface{}{}
	if err := GetDB().Raw("select * from users limit ?", 10).Scan(&res).Error; err != nil {
		fmt.Printf("err: %s", err)
		return
	}
}
