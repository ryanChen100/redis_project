package router

import (
	"ryan/redis/project/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	InitAdminRouter(r)
	r.Run(":" + config.GetLotterySetting().Port)
}
