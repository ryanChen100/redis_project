package main

import (
	"log"
	"net/http"
	"ryan/redis/project/config"
	"ryan/redis/project/enum"
	chatserver "ryan/redis/project/pkg/chat_server"
	redistool "ryan/redis/project/pkg/redis_tool"
)

func main() {
	config.ChatConfig()
	redistool.RedisInit()
	redisClient := redistool.GetRedis()
	hub := chatserver.NewHub()

	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chatserver.ServeWs(hub, redisClient, w, r)
	})

	go func() {
		pubsub := redisClient.Subscribe(enum.ChatChannel)
		for msg := range pubsub.Channel() {
			hub.Broadcast <- []byte(msg.Payload)
		}
	}()

	log.Println("Server started on :", config.GetChatSetting().Port)
	err := http.ListenAndServe(":"+config.GetChatSetting().Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
