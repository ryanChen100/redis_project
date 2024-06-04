package chatserver

import (
	"log"
	"ryan/redis/project/enum"
	"ryan/redis/project/pkg/redis_tool"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump(redisClient *redis_tool.RedisClient) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		// Publish the message to Redis
		err = redisClient.PublishMessage(enum.ChatChannel, message)
		if err != nil {
			log.Println("Publish message error:", err)
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
