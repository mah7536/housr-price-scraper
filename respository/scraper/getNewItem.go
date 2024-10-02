package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"scrape/domain"
	"time"

	"scrape/domain/logger"

	"scrape/domain/errorCode"
)

func (s *Scraper) GetNewItem() (code int, data []*domain.House, err error) {
	newReq, err := http.NewRequest(http.MethodGet, Item591, nil)
	if err != nil {
		code = errorCode.Error
		return
	}
	q := url.Values{}
	if s.DetailSetting.Type != "" {
		q.Add("type", s.DetailSetting.Type)
	}

	if s.DetailSetting.ShType != "" {
		q.Add("shType", s.DetailSetting.ShType)
	}

	if s.DetailSetting.Section != "" {
		q.Add("section", s.DetailSetting.Section)
	}
	if s.DetailSetting.Regionid != "" {
		q.Add("regionid", s.DetailSetting.Regionid)
	}
	if s.DetailSetting.Shape != "" {
		q.Add("shape", s.DetailSetting.Shape)
	}

	if s.DetailSetting.Price != "" {
		q.Add("price", s.DetailSetting.Price)
	}

	q.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	q.Add("order", "posttime_desc")
	newReq.URL.RawQuery = q.Encode()
	newReq.Header.Set("X-CSRF-TOKEN", s.XCSRFToken)
	newReq.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

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
