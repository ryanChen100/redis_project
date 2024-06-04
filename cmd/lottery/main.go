package main

import (
	"ryan/redis/project/config"
	gameserver "ryan/redis/project/pkg/game_server"
	redistool "ryan/redis/project/pkg/redis_tool"
	"ryan/redis/project/pkg/router"
)

func main() {
	config.LotteryConfig()
	redistool.RedisInit()
	gameserver.GameServer()
	router.Start()
}
