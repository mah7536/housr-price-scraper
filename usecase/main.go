package usecase

import (
	"fmt"
	"scrape/domain"
	"scrape/respository/cache"
	"time"

	"188.166.240.198/GAIUS/lib/errorCode"
	"188.166.240.198/GAIUS/lib/logger"
)

type Usecase struct {
	cache       *cache.Cache
	scrapers    []domain.Scraper
	NewItemChan chan *domain.House
}

func NewUsecase(cache *cache.Cache, scrapers ...domain.Scraper) *Usecase {
	return &Usecase{
		cache:       cache,
		scrapers:    scrapers,
		NewItemChan: make(chan *domain.House),
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
						u.NewItemChan <- eachResult
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
