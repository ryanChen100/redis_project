package redis_tool

import (
	"context"
	"fmt"
	"log"
	"ryan/redis/project/config"
	"strings"

	"github.com/go-redis/redis/v8"
)

var (
	Redis *RedisClient
	ctx   = context.Background()
)

type RedisClient struct {
	Client *redis.Client
}

func RedisInit() {
	NewClient()
	cleanRedis()
}

func NewClient() { // 實體化redis.Client 並返回實體的位址

	Redis = &RedisClient{}
	redisOpt := &redis.Options{}
	if config.GetLotterySetting() != nil {
		fmt.Println("Lottery Redis")
		redisOpt = &redis.Options{
			Addr:     strings.Join([]string{config.GetLotterySetting().RedisIp, config.GetLotterySetting().RedisPort}, ":"),
			Password: config.GetLotterySetting().RedisPassword, // no password set
			DB:       config.GetLotterySetting().RedisDB,       // use default DB
		}
	} else {
		fmt.Println("Chat Redis")
		redisOpt = &redis.Options{
			Addr:     strings.Join([]string{config.GetChatSetting().RedisIp, config.GetChatSetting().RedisPort}, ":"),
			Password: config.GetChatSetting().RedisPassword, // no password set
			DB:       config.GetChatSetting().RedisDB,       // use default DB
		}
	}
	
	Redis.Client = redis.NewClient(redisOpt)
	if pong, err := Redis.Client.Ping(ctx).Result(); err != nil {
		log.Println("Redis connection failed", err)
	} else {
		log.Println("Redis connection success", pong)
	}
}

// 清空DB 0 Redis資料庫
func cleanRedis() {

	// 清空所有數據庫
	err := Redis.Client.FlushAll(ctx).Err()
	if err != nil {
		fmt.Println("清空數據庫失敗:", err)
		return
	}
	fmt.Println("所有數據庫已清空")

}

func GetRedisClient() *redis.Client {
	return Redis.Client
}

func GetRedisContext() context.Context {
	return ctx
}

func GetRedis() *RedisClient {
	return Redis
}

func (r *RedisClient) PublishMessage(channel string, message []byte) error {
	return r.Client.Publish(ctx, channel, message).Err()
}

func (r *RedisClient) Subscribe(channel string) *redis.PubSub {
	return r.Client.Subscribe(ctx, channel)
}
