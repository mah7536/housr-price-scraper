package buyscraper

import (
	"net/http"
	"net/http/cookiejar"
	"os"
	"scrape/domain"
	"time"

	"scrape/domain/logger"
)

const (
	IndexUrl = "https://buy.houseprice.tw/ws/list/%s_city/%s_zip"
)

type BuyScraper struct {
	Client     *http.Client
	XCSRFToken string
	Name       string
}

func NewBuyScraper(name string) (code int, s domain.Scraper, err error) {
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

	s = &BuyScraper{
		Name: name,
		Client: &http.Client{
			Jar:       jar,
			Timeout:   30 * time.Second,
			Transport: tr,
		},
		XCSRFToken: "",
	}

	return
}
