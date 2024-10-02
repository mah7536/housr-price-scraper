package telegram

import (
	"scrape/delivery/telegram/lib"
	"time"

	"scrape/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TestJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}

func MenuList(message *tgbotapi.Update) (code int, res tgbotapi.MessageConfig, err error) {
	res = lib.NewResponseMs(message.Message.Chat.ID, "Gaius Bot Menu")
	str := "cmd"

	res.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Current time", time.Now().Format(domain.TimeLayout)),
			tgbotapi.NewInlineKeyboardButtonURL("Official Site", "https://www.gaiusauto.com/"),
			tgbotapi.NewInlineKeyboardButtonSwitch("switch_inline_query(1,3)", "cmd"),
			tgbotapi.InlineKeyboardButton{
				Text:                         "switch_inline_query_current_chat(1,4)",
				SwitchInlineQueryCurrentChat: &str,
			},
		),
	)
	return
}

func IsAlive(message *tgbotapi.Update) (code int, res []tgbotapi.MessageConfig, err error) {
	res = []tgbotapi.MessageConfig{lib.NewResponseMs(message.Message.Chat.ID, "👍")}
	return
}
