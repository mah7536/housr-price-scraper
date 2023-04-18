package scraper

import (
	"net/http"

	"188.166.240.198/GAIUS/lib/errorCode"
	"188.166.240.198/GAIUS/lib/logger"
	"github.com/PuerkitoBio/goquery"
)

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
	return
}
