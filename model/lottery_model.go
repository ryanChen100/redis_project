package model

type LotterySetting struct {
	Port          string `mapstructure:"port"`
	RedisIp       string `mapstructure:"redis_ip"`
	RedisPort     string `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}

type User struct {
	Id      string `json:"Id"`      // 玩家 ID
	Balance int    `json:"balance"` // 玩家餘額
}

type UserBet struct {
	Id     string `json:"Id"`
	Round  int    `json:"round"`  // 局數
	Amount int    `json:"amount"` // 下注金額
}
