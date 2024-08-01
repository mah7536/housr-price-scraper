package buyscraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"scrape/domain"

	"scrape/domain/errorCode"

	"scrape/domain/logger"
)

func (s *BuyScraper) GetNewItem() (code int, data []*domain.House, err error) {
	baseUrl, err := url.Parse(fmt.Sprintf(IndexUrl, "台中市", "大雅區", "透天-別墅"))
	if err != nil {
		code = errorCode.Error
		return
	}

	newReq, err := http.NewRequest(http.MethodGet, baseUrl.String(), nil)
	if err != nil {
		code = errorCode.Error
		return
	}

	newReq.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	res, err := s.Client.Do(newReq)
	if err != nil {
		code = errorCode.Error
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		code = errorCode.Error
		logger.Error(res.Status)
		return
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		code = errorCode.Error
		logger.Error(err)
		return
	}

	d := &domain.BuyRes{}
	err = json.Unmarshal(bodyBytes, d)
	if err != nil {
		code = errorCode.Error
		logger.Error(err)
		return
	}

	if d.Code != http.StatusOK {
		code = errorCode.Error
		logger.Error(err)
		return
	}

	data = []*domain.House{}
	for _, eachHouse := range d.WebCaseGrouping {
		newData := &domain.House{
			Title:         eachHouse.CaseName,
			HouseId:       eachHouse.Sid,
			Price:         int(eachHouse.TotalPrice),
			Area:          eachHouse.LandPin,
			MainArea:      eachHouse.BuildPin,
			HouseAge:      int(eachHouse.BuildAge),
			PostTime:      (*domain.TimeStamp)(eachHouse.NewKeyInDate),
			PhotoUrl:      eachHouse.MPicUrl,
			CommunityLink: eachHouse.CaseUrl,
			RefreshTime:   eachHouse.NewKeyInDaysName,
			Url:           eachHouse.CaseUrl,
		}

		if eachHouse.NewKeyInDaysName == "今天新上架" {
			newData.IsNew = 1
		}

		if eachHouse.OrigPrice != nil {
			newData.IsDownPrice = 1
		}
		data = append(data, newData)
	}
	return
}
