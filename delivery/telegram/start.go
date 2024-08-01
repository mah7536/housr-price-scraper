package telegram

import (
	"fmt"
	"scrape/config"
	"scrape/delivery/telegram/lib"

	"scrape/domain/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	SomeOneTalkToBot = "有人與Gaius機器人開通, 請注意 該人的名稱為%s %s 語系為%s Id為%d"
)

func (server *TelegramServer) Start(message *tgbotapi.Update) (code int, res []tgbotapi.MessageConfig, err error) {
	res = []tgbotapi.MessageConfig{lib.NewResponseMs(message.Message.Chat.ID, fmt.Sprintf("Hi! %s, welcome to Gaius", message.Message.From.UserName))}
	if !message.Message.From.IsBot {
		messageFromCheckBot := fmt.Sprintf(SomeOneTalkToBot, message.Message.From.LastName, message.Message.From.FirstName, message.Message.From.LanguageCode, message.Message.From.ID)
		logger.Info(messageFromCheckBot)
		server.SendMs(lib.NewCommonMessage(config.Arther, lib.TypeInfo, "New Talker To Bot", messageFromCheckBot))
	} else {
		messageFromCheckBot := fmt.Sprintf(SomeOneTalkToBot+"(bot)", message.Message.From.LastName, message.Message.From.FirstName, message.Message.From.LanguageCode, message.Message.From.ID)
		logger.Info(messageFromCheckBot)
		server.SendMs(lib.NewCommonMessage(config.Arther, lib.TypeInfo, "New Talker To Bot", messageFromCheckBot))
	}

	return
}
