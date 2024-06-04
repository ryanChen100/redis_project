package controller

import (
	"errors"
	"fmt"
	"net/http"
	"ryan/redis/project/enum"
	"ryan/redis/project/model"
	gameserver "ryan/redis/project/pkg/game_server"
	"ryan/redis/project/pkg/redis_tool"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GetUserBalance(c *gin.Context) {
	var (
		user model.User
		rc   = redis_tool.GetRedisClient()
		ctx  = redis_tool.GetRedisContext()
	)

	user.Id = c.Param("user")
	fmt.Println(user)
	balance, err := rc.ZScore(ctx, enum.UserMember, user.Id).Result()
	if err == redis.Nil { //查無使用者，註冊新帳號
		balance = enum.DefaultBalance

		rc.ZAdd(ctx, enum.UserMember, &redis.Z{
			Score:  float64(balance),
			Member: user.Id,
		}).Result()
	}
	user.Balance = int(balance)
	wrapResponse(c, user, nil)

}

func Bet(c *gin.Context) {
	var (
		user      model.User
		rc        = redis_tool.GetRedisClient()
		ctx       = redis_tool.GetRedisContext()
		amountStr = c.Param("amount")
	)

	user.Id = c.Param("user")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		wrapResponse(c, nil, errors.New("下注金額有誤"))
		return
	}

	balance, err := rc.ZScore(ctx, enum.UserMember, user.Id).Result()
	if err == redis.Nil {
		wrapResponse(c, nil, errors.New("查無此用戶，請先註冊"))
		return
	}
	user.Balance = int(balance)
	if amount <= 0 {
		wrapResponse(c, nil, errors.New("下注金額需為正整數"))
		return
	}
	if amount > user.Balance {
		wrapResponse(c, nil, errors.New("餘額不足"))
		return
	}

	user.Balance -= amount
	rc.ZIncrBy(ctx, enum.UserMember, float64(-amount), user.Id)
	rc.ZIncrBy(ctx, enum.BetThisRound, float64(amount), user.Id)

	wrapResponse(c, user, nil)
}

func GetCurrentPrize(c *gin.Context) {
	wrapResponse(c, gameserver.GetCurrentPrize(), nil)
}

func GetUserBets(c *gin.Context) {
	UserBets := gameserver.GetUserBets()
	if len(UserBets) == 0 {
		wrapResponse(c, nil, errors.New("目前沒有任何記錄"))
		return
	}
	wrapResponse(c, UserBets, nil)
}

func wrapResponse(c *gin.Context, data interface{}, err error) {
	type ret struct {
		Status string      `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}

	d := ret{
		Status: "ok",
		Msg:    "",
		Data:   []struct{}{},
	}

	if data != nil {
		d.Data = data
	}

	if err != nil {
		d.Status = "failed"
		d.Msg = err.Error()
	}

	c.JSON(http.StatusOK, d)
}
