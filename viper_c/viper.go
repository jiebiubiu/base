package viper_c

import (
	"fmt"

	"github.com/Ho-J/base/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var vC *viper.Viper

func SetViper(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	vC = v
}

func LoadConfig(config *config.Config) error {
	if err := vC.Unmarshal(config); err != nil {
		fmt.Println(err)
		return err
	}

	vC.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vC.Unmarshal(config); err != nil {
			fmt.Println(err)
		}
	})

	fmt.Printf("%+v", config)

	return nil
}

// func GetMysqlC() []config.Mysql {
// 	return
// }

// func LoadMysql(mysqlC *config.Mysqls) error {
// 	if err := vC.Unmarshal(mysqlC); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	vC.OnConfigChange(func(e fsnotify.Event) {
// 		fmt.Println("config file changed:", e.Name)
// 		if err := vC.Unmarshal(mysqlC); err != nil {
// 			fmt.Println(err)
// 		}
// 	})

// 	return nil
// }

// func LoadLog(cof *config.Log) error {
// 	if err := vC.Unmarshal(cof); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	vC.OnConfigChange(func(e fsnotify.Event) {
// 		fmt.Println("config file changed:", e.Name)
// 		if err := vC.Unmarshal(cof); err != nil {
// 			fmt.Println(err)
// 		}
// 	})

// 	return nil
// }

// func LoadMinio(cof config.Minio) error {
// 	if err := vC.Unmarshal(&cof); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println(cof)
// 	vC.OnConfigChange(func(e fsnotify.Event) {
// 		fmt.Println("config file changed:", e.Name)
// 		if err := vC.Unmarshal(&cof); err != nil {
// 			fmt.Println(err)
// 		}
// 	})

// 	return nil
// }

// func LoadJaeger(cof *config.Jaeger) error {
// 	if err := vC.Unmarshal(cof); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	vC.OnConfigChange(func(e fsnotify.Event) {
// 		fmt.Println("config file changed:", e.Name)
// 		if err := vC.Unmarshal(cof); err != nil {
// 			fmt.Println(err)
// 		}
// 	})

// 	return nil
// }
