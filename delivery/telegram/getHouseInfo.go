package telegram

import (
	"scrape/delivery/telegram/lib"
	"scrape/domain"
	"sort"
	"time"

	"scrape/domain/errorCode"
	"scrape/domain/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HomeSlice []*domain.House

func (a HomeSlice) Len() int {
	return len(a)
}
func (a HomeSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a HomeSlice) Less(i, j int) bool {
	return time.Time(*a[i].PostTime).After(time.Time(*a[j].PostTime))
}

func (server *TelegramServer) GetHouseInfo(message *tgbotapi.Update) (code int, res []tgbotapi.MessageConfig, err error) {
	code, data, err := server.usecase.GetAllData()
	if code != errorCode.Success {
		code = errorCode.DBNoData
		return
	}

	list := []*domain.House{}
	for _, i := range data {
		for _, eachData := range i {
			list = append(list, eachData)
		}
	}

	sort.Sort(HomeSlice(list))

	res = []tgbotapi.MessageConfig{}

	for _, each := range list {

		code, ms, err := lib.FormatMs(message.Message.Chat.ID, each)
		if code != errorCode.Success {
			logger.Error(err)
			continue
		}

		res = append(res, ms)

	}

	return
}
