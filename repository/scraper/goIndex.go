package scraper

import (
	"net/http"

	"scrape/domain/errorCode"
	"scrape/domain/logger"

	"github.com/PuerkitoBio/goquery"
)

type SaleIndex struct {
	Status int `json:"status"`
	Data   struct {
		DeviceId string `json:"device_id"`
	} `json:"data"`
}

func (s *Scraper) GoIndex() (code int, newReq *http.Request, err error) {
	newReq, err = http.NewRequest(http.MethodGet, Www591, nil)
	if err != nil {
		code = errorCode.Error
		return
	}
	return
}

func (s *Scraper) GoSaleIndex() (code int, newReq *http.Request, err error) {
	newReq, err = http.NewRequest(http.MethodGet, Sale591, nil)
	if err != nil {
		code = errorCode.Error
		return
	}

	newReq.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	res, err := s.Client.Do(newReq)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logger.Error(err)
		return
	}

	sr, e := doc.Find("meta[name='csrf-token']").Attr("content")
	if !e {
		logger.Error("not found")
		return
	}

	s.XCSRFToken = sr
	s.DeviceId = ""
	return
}
