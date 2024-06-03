package main

import (
	"ryan/redis/project/config"
	gameserver "ryan/redis/project/pkg/game_server"
	"ryan/redis/project/pkg/redis_tool"
	"ryan/redis/project/pkg/router"
)

func main() {
	config.Config()
	redis_tool.RedisInit()
	gameserver.GameServer()
	router.Start()

}
