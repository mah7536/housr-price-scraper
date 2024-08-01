package usecase

import (
	"fmt"
	"scrape/domain"
	"scrape/respository/cache"
	"time"

	"scrape/domain/errorCode"

	"scrape/domain/logger"
)

type Usecase struct {
	cache       *cache.Cache
	scrapers    []domain.Scraper
	receiver    []*Receiver
	NewItemChan chan *domain.Item
}

type Receiver struct {
	ChatId       int64
	ReceiverType []string
}

func NewUsecase(cache *cache.Cache, receiver []*Receiver, scrapers ...domain.Scraper) *Usecase {
	return &Usecase{
		cache:       cache,
		scrapers:    scrapers,
		NewItemChan: make(chan *domain.Item),
		receiver:    receiver,
	}
}

func (u *Usecase) Run() {
	ticker := time.NewTicker(30 * time.Minute)
	for index, eachScraper := range u.scrapers {
		go func(eachScraper domain.Scraper, index int) {
			for {
				logger.Info(fmt.Sprintf("更新 時間為:%s", time.Now().Format(time.RFC3339)))
				code, result, err := eachScraper.GetNewItem()
				if code != errorCode.Success {
					logger.Error(err)
					continue
				}

				for _, eachResult := range result {
					code, _, err := u.cache.IsDataExist(eachScraper.GetSourceName(), eachResult)
					if code == errorCode.DBNoData {
						fmt.Println(fmt.Sprintf("Source:%s id:%d p:%v area: %f age: %d isNew: %b  is down price: %d post time:%s   refresh times:%s", eachScraper.GetSourceName(), eachResult.HouseId, eachResult.Price, eachResult.MainArea, eachResult.HouseAge, eachResult.IsNew, eachResult.IsDownPrice, eachResult.PostTime, eachResult.RefreshTime))
						code, _, err = u.cache.Add(eachScraper.GetSourceName(), eachResult)
						if code != errorCode.Success {
							logger.Error(err)
							continue
						}

						for _, eachReceiver := range u.receiver {
							for _, eachType := range eachReceiver.ReceiverType {
								if eachScraper.GetSourceName() != eachType {
									continue
								}
								u.NewItemChan <- &domain.Item{
									ChatId: eachReceiver.ChatId,
									House:  eachResult,
								}
							}
						}

					}
					if code != errorCode.Success {
						logger.Error(err)
						continue
					}

				}
				select {
				case <-ticker.C:

				}

			}
		}(eachScraper, index)
	}

}

func (u *Usecase) GetAllData() (code int, data map[string]map[int]*domain.House, err error) {
	code, data, err = u.cache.GetAll()
	if code != errorCode.Success {
		logger.Error(err)
		return
	}

	return
}
