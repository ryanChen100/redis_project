package config

import (
	"fmt"
	"ryan/redis/project/model"

	"github.com/spf13/viper"
)

var lotterySetting *model.LotterySetting
var chatSetting *model.ChatSetting

// Config 讀取config.yaml
func LotteryConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("lottery_config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %s\n", err)
		return
	}

	if err := viper.Unmarshal(&lotterySetting); err != nil {
		fmt.Printf("无法映射配置到结构体: %s\n", err)
		return
	}

	fmt.Println("lotterySetting", lotterySetting)
}

func GetLotterySetting() *model.LotterySetting {
	return lotterySetting
}

// Config 讀取config.yaml
func ChatConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("chat_config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %s\n", err)
		return
	}

	if err := viper.Unmarshal(&chatSetting); err != nil {
		fmt.Printf("无法映射配置到结构体: %s\n", err)
		return
	}

	fmt.Println("chatSetting", chatSetting)
}

func GetChatSetting() *model.ChatSetting {
	return chatSetting
}
