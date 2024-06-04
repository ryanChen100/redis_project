package config

import (
	"fmt"
	"ryan/redis/project/model"

	"github.com/spf13/viper"
)

var setting model.Setting

// Config 讀取config.yaml
func Config() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %s\n", err)
		return
	}

	if err := viper.Unmarshal(&setting); err != nil {
		fmt.Printf("无法映射配置到结构体: %s\n", err)
		return
	}

	fmt.Println(setting)
}

func GetSetting() model.Setting {
	return setting
}
