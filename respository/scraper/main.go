package scraper

import (
	"net/http"
	"net/http/cookiejar"
	"os"
	"scrape/domain"
	"time"

	"scrape/domain/errorCode"

	"scrape/domain/logger"
)

const (
	Index591 = "https://591.com.tw/"
	Sale591  = "https://sale.591.com.tw"

	Www591    = "https://www.591.com.tw"
	Union591  = "https://union.591.com.tw"
	Double591 = "https://doubleclick.net"

	Item591 = "https://sale.591.com.tw/home/search/list"
)

type DetailSetting struct {
	Type     string
	ShType   string
	Section  string
	Regionid string
	Shape    string
	Price    string
}
type Scraper struct {
	Client        *http.Client
	XCSRFToken    string
	Name          string
	DetailSetting *DetailSetting
}

func NewScraper(name string, detailSetting *DetailSetting) (code int, s domain.Scraper, err error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	jar, err := cookiejar.New(nil)

	if err != nil {
		logger.Error(err)
		os.Exit(2)
	}

	tmp := &Scraper{
		Name: name,
		Client: &http.Client{
			Jar:       jar,
			Timeout:   30 * time.Second,
			Transport: tr,
		},
		XCSRFToken:    "",
		DetailSetting: detailSetting,
	}

	code, _, err = tmp.GoSaleIndex()

	if code != errorCode.Success {
		logger.Error(err)
		return
	}

	s = tmp
	return
}
