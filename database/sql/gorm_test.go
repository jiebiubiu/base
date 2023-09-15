package sql

import (
	"fmt"
	"testing"

	"github.com/Ho-J/base/logs"
)

func TestConnDB(t *testing.T) {
	logs.InitLogger(logs.DefaultLogConfigs)
	InitDB("root:123456@tcp(127.0.0.1:3306)/snoopy?charset=utf8mb4&parseTime=True&loc=Local")

	res := []map[string]interface{}{}
	if err := DB.Raw("select * from user limit ?", 10).Scan(&res).Error; err != nil {
		fmt.Printf("err: %s", err)
		return
	}
}
