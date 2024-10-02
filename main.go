package main

import (
	"os"
	"os/signal"
	"scrape/config"
	"scrape/delivery/telegram"
	"scrape/repository/buyscraper"
	"scrape/repository/cache"
	"scrape/repository/scraper"
	"scrape/usecase"
	"strconv"
	"strings"
	"syscall"

	"scrape/domain/errorCode"
	"scrape/domain/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.Telegram_token = os.Getenv("Telegram_token")
	var err error
	config.Member, err = strconv.ParseInt(os.Getenv("Member"), 10, 64)
	if err != nil {
		logger.Error("Member 需要 整數")
		os.Exit(2)
	}

	if strings.TrimSpace(config.Telegram_token) == "" || config.Member == 0 {
		logger.Error("請先設置telegram token 跟 要通知的member")
		os.Exit(2)
	}

	logger.Info("start server .....")
	setting_daya := &scraper.DetailSetting{
		Type:     "2",
		ShType:   "list",
		Section:  "117",
		Regionid: "8",
		Shape:    "3,4",
	}
	code, s_taya, err := scraper.NewScraper("591-大雅", setting_daya)
	if code != errorCode.Success {
		logger.Error(err)
		os.Exit(2)
	}

	setting_taichung := &scraper.DetailSetting{
		Type:     "2",
		ShType:   "list",
		Section:  "104",
		Regionid: "8",
		Shape:    "1,2",
		Price:    "750_1000",
	}
	code, s_taichung, err := scraper.NewScraper("591-台中", setting_taichung)
	if code != errorCode.Success {
		logger.Error(err)
		os.Exit(2)
	}

	code, buyS, err := buyscraper.NewBuyScraper("buy")
	if code != errorCode.Success {
		logger.Error(err)
		os.Exit(2)
	}
	cache := cache.NewCache(s_taya, buyS, s_taichung)

	receivers := []*usecase.Receiver{
		{
			ChatId:       config.Member,
			ReceiverType: []string{s_taya.GetSourceName(), buyS.GetSourceName()},
		},
	}
	usecase := usecase.NewUsecase(cache, receivers, s_taya, buyS, s_taichung)

	tg := telegram.NewTelegramServer(usecase)
	go tg.RunServer()
	go tg.RunJob()

	usecase.Run()

	logger.Info("server is running.....")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	<-shutdown

	logger.Info("stop server .....")

}
