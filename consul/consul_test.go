package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	capi "github.com/hashicorp/consul/api"
	"github.com/jiebiubiu/base/config"
	"github.com/spf13/viper"
)

func TestConsulLoadConfig(t *testing.T) {
	// consulC := NewConsul("consul://127.0.0.1:18500/", "")

	// if err := consulC.LoadConfig(); err != nil {
	// 	fmt.Sprintf("err: %s", err)
	// 	return
	// }
	// def := capi.DefaultConfig()
	client, err := capi.NewClient(&capi.Config{
		Address:   "127.0.0.1:18500",
		Scheme:    "http",
		Transport: &http.Transport{},
	})
	if err != nil {
		panic(err)
	}

	kv, _, err := client.KV().Get("snoopy_config", nil)
	if err != nil {
		log.Fatalln("consul获取配置失败:", err)
	}

	fmt.Printf("---------------: %s", string(kv.Value))

	defaultConfig := viper.New()
	defaultConfig.SetConfigType("json")
	err = defaultConfig.ReadConfig(bytes.NewBuffer(kv.Value))
	if err != nil {
		log.Fatalln("Viper解析配置失败:", err)
	}

	var c = config.Config{}
	if err := defaultConfig.Unmarshal(&c); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ttttttttt_____: ", c)
}

func TestXxx(t *testing.T) {
	var c = config.Config{}
	bs, _ := json.Marshal(&c)
	fmt.Println(string(bs))
}

func TestLoadConfig(t *testing.T) {
	consul := NewConsul("127.0.0.1:18500", "snoopy_config")

	if err := consul.LoadConfig(); err != nil {
		fmt.Printf("err: %v", err)
	}
	c := consul.GetConfig()
	fmt.Printf("config: %+v", c)
}
