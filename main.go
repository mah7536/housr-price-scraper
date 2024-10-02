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
	code, s, err := scraper.NewScraper("591")
	if code != errorCode.Success {
		logger.Error(err)
		os.Exit(2)
	}

	code, buyS, err := buyscraper.NewBuyScraper("buy")
	if code != errorCode.Success {
		logger.Error(err)
		os.Exit(2)
	}
	cache := cache.NewCache("591", "buy")

	usecase := usecase.NewUsecase(cache, s, buyS)

	usecase.Run()

	tg := telegram.NewTelegramServer(usecase)
	go tg.RunServer()
	go tg.RunJob()
	logger.Info("server is running.....")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	<-shutdown

	logger.Info("stop server .....")

}
