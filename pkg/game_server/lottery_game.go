package gameserver

import (
	"fmt"
	"log"
	"math/rand"
	"ryan/redis/project/enum"
	"ryan/redis/project/model"
	"ryan/redis/project/pkg/redis_tool"
	"time"
)

func GameServer() {
	rand.Seed(time.Now().UTC().UnixNano())
	var (
		ticker = time.NewTicker(enum.RoundSecond * time.Second) // 每過RoundSecond秒，執行一次以下迴圈
		rc     = redis_tool.GetRedisClient()
		ctx    = redis_tool.GetRedisContext()
	)
	go func() {
		for {
			enum.Round++
			enum.StartTime = time.Now()
			log.Println(enum.StartTime.Format("2006-01-02 15:04:05"), "\t round", enum.Round, "start")
			<-ticker.C

			var prizePool = GetCurrentPrize()
			var userBets = GetUserBets()
			if len(userBets) == 0 {
				log.Println("Round", enum.Round, "沒有任何玩家下注")
				continue
			}

			// 抽獎選贏家
			winNum := rand.Intn(prizePool + 1)
			var winner string
			for _, userBet := range userBets {
				winNum -= userBet.Amount
				if winNum <= 0 {
					winner = userBet.Id
					break
				}
			}
			log.Println("獎金池:", prizePool, "\t 得主:", winner)

			// 發獎金給得主
			rc.ZIncrBy(ctx, enum.UserMember, float64(prizePool), winner)

			// 刪除現有Table
			rc.Del(ctx, enum.BetThisRound)
		}
	}()
}

func GetCurrentPrize() (prizePool int) {
	var (
		rc  = redis_tool.GetRedisClient()
		ctx = redis_tool.GetRedisContext()
	)

	bets, _ := rc.ZRangeWithScores(ctx, enum.BetThisRound, 0, -1).Result()
	for _, bet := range bets {
		var userBet model.UserBet
		userBet.Amount = int(bet.Score)
		prizePool += userBet.Amount
	}
	return
}

func GetUserBets() (userBets []model.UserBet) {
	var (
		rc  = redis_tool.GetRedisClient()
		ctx = redis_tool.GetRedisContext()
	)
	bets, _ := rc.ZRangeWithScores(ctx, enum.BetThisRound, 0, -1).Result()
	for _, bet := range bets {
		var userBet model.UserBet
		userBet.Id = fmt.Sprintf("%s", bet.Member)
		userBet.Amount = int(bet.Score)
		userBet.Round = enum.Round
		userBets = append(userBets, userBet)
	}
	return
}
