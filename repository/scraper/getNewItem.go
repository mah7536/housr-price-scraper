package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"scrape/domain"
	"strings"
	"time"

	"scrape/domain/errorCode"
	"scrape/domain/logger"
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
	q.Add("shape", "4,3")
	q.Add("category", "1")
	q.Add("order", "posttime_desc")
	q.Add("totalRows", "579")
	q.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	newReq.URL.RawQuery = q.Encode()
	newReq.Header.Set("X-CSRF-TOKEN", s.XCSRFToken)
	newReq.Header.Set("Device", "pc")
	newReq.Header.Set("Cookie",
		strings.Join(
			[]string{
				"is_new_index=1",
				"is_new_index_redirect=1",
				"user_index_role=2",
				"user_browse_recent=a%3A3%3A%7Bi%3A0%3Ba%3A2%3A%7Bs%3A4%3A%22type%22%3Bi%3A2%3Bs%3A7%3A%22post_id%22%3Bi%3A15320140%3B%7Di%3A1%3Ba%3A2%3A%7Bs%3A4%3A%22type%22%3Bi%3A2%3Bs%3A7%3A%22post_id%22%3Bi%3A15415273%3B%7Di%3A2%3Ba%3A2%3A%7Bs%3A4%3A%22type%22%3Bi%3A2%3Bs%3A7%3A%22post_id%22%3Bi%3A15515269%3B%7D%7D",
				"webp=1",
				fmt.Sprintf("PHPSESSID=%s", s.XCSRFToken),
				fmt.Sprintf("T591_TOKEN=%s", s.XCSRFToken),
				"_gcl_au=1.1.1107047101.1727790781",
				"_gid=GA1.3.387915728.1727790781",
				"tw591__privacy_agree=0",
				"timeDifference=-1",
				"__one_id__=01J945P4TW1Q9YSEJ947NJ0SZW",
				"_clck=1kgipf0%7C2%7Cfpn%7C0%7C1735",
				"urlJumpIp=8",
				"index_keyword_search_analysis=%7B%22role%22%3A%222%22%2C%22type%22%3A1%2C%22keyword%22%3A%22%22%2C%22selectKeyword%22%3A%22%22%2C%22menu%22%3A%22%22%2C%22hasHistory%22%3A0%2C%22hasPrompt%22%3A0%2C%22history%22%3A0%7D",
				"_fbp=fb.2.1727790864705.39049446182990457",
				"_clsk=1qn9t55%7C1727791098825%7C32%7C0%7Ct.clarity.ms%2Fcollect",
				"_gat=1",
				"last_search_type=2",
				"_ga=GA1.1.552015766.1686475181",
				"_ga_HDSPSZ773Q=GS1.1.1727790781.4.1.1727791124.29.0.0",
				"_ga_H07366Z19P=GS1.3.1727790781.2.1.1727791124.59.0.0",
			},
			";",
		))
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
