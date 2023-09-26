package viper_c

import (
	"fmt"

	"github.com/Ho-J/base/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type viperC struct {
	configPath string
	viper      *viper.Viper
	config     *config.Config
}

func (vc *viperC) setViper() {
	vc.viper = viper.New()
	vc.viper.SetConfigFile(vc.configPath)
	vc.viper.SetConfigType("yaml")
	err := vc.viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}

	vc.viper.WatchConfig()
}

func (vc *viperC) LoadConfig() error {
	vc.setViper()

	if err := vc.viper.Unmarshal(&vc.config); err != nil {
		return err
	}

	vc.viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vc.viper.Unmarshal(vc.config); err != nil {
			fmt.Println(err)
		}
	})

	return nil
}

func (vc *viperC) SetConfigPath(path string) {
	vc.configPath = path
}

func (vc viperC) GetConfig() *config.Config {
	return vc.config
}

func NewViperC(path string) viperC {
	return viperC{configPath: path}
}
