package telegram

import (
	"fmt"
	"os"
	"time"

	"scrape/config"
	"scrape/delivery/telegram/lib"
	"scrape/usecase"
	"strings"

	"188.166.240.198/GAIUS/lib/errorCode"

	"188.166.240.198/GAIUS/lib/errorhandler"
	"188.166.240.198/GAIUS/lib/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	TelegramSys *TelegramServer
)

type command struct {
	tgcommand tgbotapi.BotCommand
	fn        func(update *tgbotapi.Update) (int, []tgbotapi.MessageConfig, error)
}

type callbackCommand struct {
	tgcommand tgbotapi.BotCommand
	fn        func(message *tgbotapi.Update, req []byte) (int, tgbotapi.MessageConfig, error)
}

type TelegramServer struct {
	// usecase
	usecase *usecase.Usecase

	// internal
	tgChannel    tgbotapi.UpdatesChannel
	mainBot      *tgbotapi.BotAPI
	commandList  map[string]*command
	sendChan     chan tgbotapi.Chattable
	callbackChan chan tgbotapi.CallbackConfig
	deletemsChan chan tgbotapi.DeleteMessageConfig
}

func NewTelegramServer(usecase *usecase.Usecase) *TelegramServer {
	bot, err := tgbotapi.NewBotAPI(config.Telegram_token)
	if err != nil {
		logger.Error(err)
		os.Exit(2)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	channel := bot.GetUpdatesChan(u)

	TelegramSys = &TelegramServer{
		usecase: usecase,

		tgChannel:    channel,
		mainBot:      bot,
		commandList:  make(map[string]*command),
		sendChan:     make(chan tgbotapi.Chattable),
		callbackChan: make(chan tgbotapi.CallbackConfig),
		deletemsChan: make(chan tgbotapi.DeleteMessageConfig),
	}

	// 新增command
	TelegramSys.AddCommandList("start", "for start command", TelegramSys.Start)
	TelegramSys.AddCommandList("isalive", "check sysOk", IsAlive)
	TelegramSys.AddCommandList("h", "get house info", TelegramSys.GetHouseInfo)

	// 設定command list
	_, _, err = TelegramSys.SetTgCommandList()
	if err != nil {
		logger.Error(err)
		os.Exit(2)
	}

	return TelegramSys
}

func (server *TelegramServer) RunJob() {
	for {
		select {

		case ms := <-server.sendChan:

			_, err := server.mainBot.Send(ms)
			if err != nil {
				logger.Error(err)
				continue
			}
			time.Sleep(500 * time.Millisecond)
		case newItem := <-server.usecase.NewItemChan:
			code, content, err := lib.FormatMs(config.Member, newItem)
			if code != errorCode.Success {
				logger.Error(err)
				continue
			}
			_, err = server.mainBot.Send(content)
			if err != nil {
				logger.Error(err)
				continue
			}
		}
	}
}

// 將可用的指令 統一整理
func (server *TelegramServer) AddCommandList(act string, des string, fn func(update *tgbotapi.Update) (int, []tgbotapi.MessageConfig, error)) {
	server.commandList[act] = &command{
		tgcommand: tgbotapi.BotCommand{
			Command:     act,
			Description: des,
		},
		fn: fn,
	}
}

// 將callback指令 統一整理
// func (server *TelegramServer) AddCallBackList(act string, des string, fn func(message *tgbotapi.Update, req []byte) (int, tgbotapi.MessageConfig, error)) {
// 	server.callbackList[act] = &callbackCommand{
// 		tgcommand: tgbotapi.BotCommand{
// 			Command:     act,
// 			Description: des,
// 		},
// 		fn: fn,
// 	}
// }

// 設定機器人的menu list
func (server *TelegramServer) SetTgCommandList() (code int, data interface{}, err error) {
	tgCommandList := []tgbotapi.BotCommand{}
	for _, command := range server.commandList {
		if command.tgcommand.Command != "start" {
			tgCommandList = append(tgCommandList, command.tgcommand)
		}
	}
	// err = server.mainBot.SetMyCommands(tgCommandList)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	return
}

// 傳送訊息
func (server *TelegramServer) SendMs(ms tgbotapi.Chattable) {
	server.sendChan <- ms
}

// run telegram server
func (server *TelegramServer) RunServer() {
	for message := range server.tgChannel {
		if message.Message != nil {

			// 此區 判斷是否能執行command
			fmt.Println(message.Message.Chat.ID)
			code, _, _ := lib.CheckChatID(message.Message.Chat.ID)
			if code != errorCode.Success {
				response := errorhandler.NewResponse(code)
				response.SetExtra(fmt.Sprintf("有未知的帳號 開通機器人, 請注意 該人的名稱為%s %s 語系為%s Id為%d", message.Message.From.LastName, message.Message.From.FirstName, message.Message.From.LanguageCode, message.Message.From.ID))
				logger.Error(response)
				server.sendChan <- lib.AlertMessage(message.Message.Chat.ID, "Hi")
				continue
			}
			// 僅針對 是command的message
			if message.Message.IsCommand() {
				command, exist := server.commandList[strings.ToLower(message.Message.Command())]
				if exist {
					code, res, err := command.fn(&message)
					if code != errorCode.Success {
						response := errorhandler.NewResponse(code)
						response.SetExtra(err)
						logger.Error(response)
						server.sendChan <- (lib.WarnMessage(message.CallbackQuery.Message.Chat.ID, response.Message))
						continue
					}
					for _, eachRes := range res {
						server.sendChan <- eachRes
					}
				} else {
					server.sendChan <- (lib.WarnMessage(message.Message.Chat.ID, "不在喔喔喔喔"))
				}
				continue
			}

			if message.Message.LeftChatMember != nil {
				continue
			}

			server.sendChan <- lib.AlertMessage(message.Message.Chat.ID, "Hi")
			continue
		}

	}
}
