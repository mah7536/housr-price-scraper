package main

import (
	"os"
	"os/signal"
	"scrape/delivery/telegram"
	"scrape/respository/buyscraper"
	"scrape/respository/cache"
	"scrape/respository/scraper"
	"scrape/usecase"
	"syscall"

	"188.166.240.198/GAIUS/lib/errorCode"
	"188.166.240.198/GAIUS/lib/logger"
)

func main() {
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
