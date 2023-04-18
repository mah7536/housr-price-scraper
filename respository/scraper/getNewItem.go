package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"scrape/domain"
	"time"

	"188.166.240.198/GAIUS/lib/errorCode"
	"188.166.240.198/GAIUS/lib/logger"
)

func (s *Scraper) GetNewItem() (code int, data []*domain.House, err error) {
	newReq, err := http.NewRequest(http.MethodGet, Item591, nil)
	if err != nil {
		code = errorCode.Error
		return
	}
	q := url.Values{}
	q.Add("type", "2")
	q.Add("shType", "list")
	q.Add("section", "117")
	q.Add("regionid", "8")
	q.Add("shape", "3,4")
	q.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	newReq.URL.RawQuery = q.Encode()
	newReq.Header.Set("X-CSRF-TOKEN", s.XCSRFToken)

	res, err := s.Client.Do(newReq)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		logger.Error(res.Status)
		return
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return
	}

	d := &domain.Res{}
	err = json.Unmarshal(bodyBytes, d)
	if err != nil {
		logger.Error(err)
		return
	}

	for i, _ := range d.Data.HouseList {
		d.Data.HouseList[i].Url = fmt.Sprintf("https://sale.591.com.tw/home/house/detail/2/%d.html", d.Data.HouseList[i].HouseId)
	}
	data = d.Data.HouseList
	return
}
