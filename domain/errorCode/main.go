package errorCode

import "scrape/domain/errorhandler"

const (
	// 一般常用訊息
	undefined       = -999
	Success         = 0 // 成功
	Error           = 1 // 發生錯誤
	EncodeJsonError = 2
	DBNoData        = 3
	DecodeJsonError = 4
)

func init() {
	Errorcodes := map[int]string{
		// 尚未定義到的錯誤訊息
		undefined: "InternalError", // 發生錯誤

		// 一般常用訊息
		Success: "",              // 成功
		Error:   "InternalError", // 發生錯誤
	}

	for code, message := range Errorcodes {
		errorhandler.Errorcodes[code] = message
	}
}
