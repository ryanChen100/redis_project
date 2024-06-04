package enum

import "time"

var (
	Round     = 0
	StartTime time.Time
)

const (
	RoundSecond    = 60               // 每一局的時間
	DefaultBalance = 1000             // 玩家初始化金額
	UserMember     = "game"           // 儲存所有使用者的Balance   	Redis:`Sorted-Set`	SCORE -> USER
	BetThisRound   = "bet_this_round" // 儲存目前局的下注狀況		 	Redis:`Sorted-Set`  SCORE -> USER
)

const (
	ChatChannel = "chatroom"
)
