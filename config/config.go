package config

import (
	"time"
)

var (
	Telegram_token          = ""
	Member            int64 = 0
	AllowChatId             = []int64{Member}
	TaiwanLocation, _       = time.LoadLocation("Asia/Taipei")
	Power                   = true
)
