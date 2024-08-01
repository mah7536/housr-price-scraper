package config

import (
	"time"
)

var (
	Telegram_token          = "5374196105:AAF35ntUrMMUkJZzC2aPwNOfv0omCeJbxk0"
	Arther            int64 = -4066657420
	MenLun            int64 = 1769823110
	RDMember                = []int64{Arther}
	AllowChatId             = []int64{Arther}
	TaiwanLocation, _       = time.LoadLocation("Asia/Taipei")
	Power                   = true
)
