package lib

import (
	"encoding/json"
	"fmt"
	"scrape/domain"
	"strings"

	"scrape/domain/errorCode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	seperateLine = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, "=========")
	msTypeList   = map[string]string{
		TypeWarn:     "⚠️⚠️⚠️⚠️*Warn*⚠️⚠️⚠️⚠️",        //warn
		TypeDanger:   "❗️❗️❗️❗️❗️*Danger*❗️❗️❗️❗️❗️❗", //danger
		TypeInfo:     "👌👌👌👌👌*Info*👌👌👌👌👌",              //info
		TypeCommon:   "👍👍👍👍*Common*👍👍👍👍",              //common
		TypeUndefind: "❔❔❔❔*Undefined*❔❔❔❔",           //undefinded
	}
	StandardFormat = " %s \n " + seperateLine + " %s " + seperateLine + "\n `%s`"
)

const (
	TypeWarn     = "warn"
	TypeDanger   = "danger"
	TypeInfo     = "info"
	TypeCommon   = "common"
	TypeUndefind = "undefined"
)

type CallBackReq struct {
	Action string `json:"ac"`
	Req    string `json:"req"`
}

// passive message
func NewResponseMs(chatID int64, text string) (newMS tgbotapi.MessageConfig) {
	newMS = tgbotapi.NewMessage(chatID, text)

	return
}

func AlertMessage(chatID int64, text string) (newMS tgbotapi.MessageConfig) {
	text = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text)
	newMS = tgbotapi.NewMessage(chatID, fmt.Sprintf("`%s`", text))
	newMS.ParseMode = tgbotapi.ModeMarkdownV2
	return
}

func WarnMessage(chatID int64, text string) (newMS tgbotapi.MessageConfig) {
	text = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text)
	newMS = tgbotapi.NewMessage(chatID, fmt.Sprintf("*%s*", text))
	newMS.ParseMode = tgbotapi.ModeMarkdownV2
	return
}

func DangerMessage(chatID int64, text string) (newMS tgbotapi.MessageConfig) {
	text = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text)
	newMS = tgbotapi.NewMessage(chatID, fmt.Sprintf("__%s__", text))
	newMS.ParseMode = tgbotapi.ModeMarkdownV2
	return
}

func CheckChatID(id int64) (code int, data interface{}, err error) {
	// for _, userId := range config.AllowChatId {
	// 	if userId == id {
	// 		code = errorCode.Success
	// 		return
	// 	}
	// }
	// code = errorCode.TgNotFoundUser

	return
}

// active message
// 發送一般訊息
func NewCommonMessage(chatId int64, msType string, title string, text string) (newMS tgbotapi.MessageConfig) {
	text = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text)
	header, exist := msTypeList[strings.ToLower(msType)]
	if !exist {
		header = msTypeList[TypeUndefind]
	}
	newMS = tgbotapi.NewMessage(chatId, fmt.Sprintf(StandardFormat, header, title, text))
	newMS.ParseMode = tgbotapi.ModeMarkdownV2
	return
}

// 傳送事件 內容 及發生位置
func VenueMessage(chatId int64, msType string, title string, text string, latitude float64, longitude float64) (newVenue tgbotapi.VenueConfig) {
	text = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text)
	header, exist := msTypeList[strings.ToLower(msType)]
	if !exist {
		header = msTypeList[TypeUndefind]
	}
	newVenue = tgbotapi.NewVenue(chatId, header+"\n"+title, text, latitude, longitude)
	return
}

func JsonToString(jsonData interface{}) (code int, data string, err error) {
	byteData, err := json.Marshal(jsonData)
	if err != nil {
		code = errorCode.EncodeJsonError
		return
	}
	data = string(byteData)
	return
}

func StringToReq(reqStr string) (code int, req *CallBackReq, err error) {
	req = &CallBackReq{}
	err = json.Unmarshal([]byte(reqStr), req)
	if err != nil {
		code = errorCode.DecodeJsonError
		return
	}
	return
}

func SetCallBackReq(action string, req interface{}) (code int, data string, err error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		code = errorCode.EncodeJsonError
		return
	}
	byteData, jErr := json.Marshal(&CallBackReq{
		Action: action,
		Req:    string(reqData),
	})
	if jErr != nil {
		code = errorCode.EncodeJsonError
		return
	}
	data = string(byteData)
	return
}

func FormatMs(chatid int64, h *domain.House) (code int, data tgbotapi.MessageConfig, err error) {
	content := ""
	content += fmt.Sprintf("名稱:%s\n", h.Title)
	content += fmt.Sprintf("總價:%d\n", h.Price)
	content += fmt.Sprintf("主坪:%f\n", h.MainArea)
	content += fmt.Sprintf("建坪:%f\n", h.Area)
	content += fmt.Sprintf("屋齡:%d\n", h.HouseAge)
	content += fmt.Sprintf("單價:%s\n", h.UnitPrice)
	content += fmt.Sprintf("新po:%b\n", h.IsNew)
	content += fmt.Sprintf("是否降價:%b\n", h.IsDownPrice)
	content += fmt.Sprintf("刊登時間:%s\n", h.PostTime)
	content += fmt.Sprintf("連結: %s\n", h.Url)
	content += fmt.Sprintf("圖片:%s\n", h.PhotoUrl)
	content += "\n"
	data = NewResponseMs(chatid, content)

	return
}
