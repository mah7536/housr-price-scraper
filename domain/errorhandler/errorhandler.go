package errorhandler

import (
	"encoding/json"
)

var Errorcodes map[int]string

type Response struct {
	Code    int         `json:"code"`              // 錯誤碼: 0代表成功 其餘代表有狀況
	Message string      `json:"message,omitempty"` // 錯誤訊息
	Data    interface{} `json:"data,omitempty"`    // 資料
	Extra   interface{} `json:"extra,omitempty"`   // 額外訊息
}

const (
	// 尚未定義到的錯誤訊息
	Undefined = -999
)

func init() {
	Errorcodes = map[int]string{
		// 尚未定義到的錯誤訊息
		Undefined: "系統內部錯誤",
	}
}

func NewResponse(code int) *Response {
	if _, ok := Errorcodes[code]; ok {
		return &Response{
			Code:    code,
			Message: Errorcodes[code],
		}
	}
	return &Response{
		Code:    Undefined,
		Message: Errorcodes[Undefined],
	}
}

func (res *Response) SetExtra(extra interface{}) *Response {
	switch extra.(type) {
	case error:
		res.Extra = extra.(error).Error()
	default:
		res.Extra = extra
	}

	return res
}

func (res *Response) SetData(data interface{}) *Response {
	res.Data = data
	return res
}

func (res *Response) String() string {
	data, _ := json.Marshal(res)
	return string(data)
}
