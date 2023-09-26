package consul

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/Ho-J/base/config"
	capi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// 参考 https://blog.csdn.net/weixin_44615512/article/details/125391556 加载配置
// 没有做监听

type cosulClient struct {
	configKey    string
	consulClient *capi.Client
	config       *config.Config
}

func NewConsul(ecpoint, configKey string) *cosulClient {
	client, err := capi.NewClient(&capi.Config{
		Address:   ecpoint,
		Scheme:    "http",
		Transport: &http.Transport{},
	})
	if err != nil {
		panic(fmt.Sprintf("NewConsul| ConsulClient|%s", err))
	}

	return &cosulClient{consulClient: client, configKey: configKey, config: &config.Config{}}
}

func (c *cosulClient) LoadConfig() error {
	kv, _, err := c.consulClient.KV().Get(c.configKey, nil)
	if err != nil {
		return fmt.Errorf("LoadConfig| consul获取配置失败: %v", err)
	}

	defaultConfig := viper.New()
	defaultConfig.SetConfigType("json")
	err = defaultConfig.ReadConfig(bytes.NewBuffer(kv.Value))
	if err != nil {
		return fmt.Errorf("LoadConfig| viper解析配置失败: %v", err)
	}

	if err := defaultConfig.Unmarshal(&c.config); err != nil {
		return fmt.Errorf("unmarshal| viper解析配置失败: %v", err)
	}

	return nil
}

func (c *cosulClient) GetConfig() *config.Config {
	return c.config
}
