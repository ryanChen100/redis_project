package chatserver

import (
	"log"
	"net/http"
	"ryan/redis/project/pkg/redis_tool"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(hub *Hub, redisClient *redis_tool.RedisClient, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump(redisClient)
}
