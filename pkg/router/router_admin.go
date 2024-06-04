package router

import (
	"ryan/redis/project/pkg/router/controller"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) {
	lottery := r.Group("/lottery")
	lottery.POST("/bet/:user", controller.GetUserBalance) // 玩家註冊（不須密碼，填入帳號即可）`user`區分大小寫
	lottery.POST("/bet/:user/:amount", controller.Bet)     // 玩家對目前的局面進行下注，`amount`金額
	lottery.GET("/prize", controller.GetCurrentPrize)     // 此局目前的獎金池
	lottery.GET("/bets", controller.GetUserBets)          // 此局所有玩家目前的下注
}
