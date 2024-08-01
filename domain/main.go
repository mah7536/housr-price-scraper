package domain

import "time"

type TimeStamp time.Time

type Item struct {
	ChatId int64 // 要通知的Id
	*House
}

type House struct {
	HouseId     int     `json:"houseid"`
	Kind        int     `json:"kind"`
	KindName    string  `json:"kind_name"`
	ShapeName   string  `json:"shape_name"`
	RegionName  string  `json:"region_name"`
	SectionName string  `json:"section_name"`
	Title       string  `json:"title"`
	Room        string  `json:"room"`
	Floor       string  `json:"floor"`
	MainArea    float32 `json:"mainarea"`
	Area        float32 `json:"area"`
	// PhotoNum     string `json:"photoNum"`
	RefreshTime   interface{} `json:"refreshtime"`
	PhotoUrl      string      `json:"photo_url"`
	NickName      string      `json:"nick_name"`
	HouseType     int         `json:"housetype"`
	IsNew         int         `json:"isnew"`
	PostTime      *TimeStamp  `json:"posttime"`
	HouseAge      int         `json:"houseage"`
	Address       string      `json:"address"`
	UnitPrice     interface{} `json:"unitprice"`
	Price         int         `json:"price"`
	IsDownPrice   int         `json:"is_down_price"`
	IsHurryPrice  int         `json:"is_hurry_price"`
	CommunityLink string      `json:"community_link"`
	Url           string      `json:"url"`
}

type Data struct {
	HouseList []*House `json:"house_list"`
	Total     string   `json:"total"`
}

type Res struct {
	Status    int    `json:"status"`
	Timestamp string `json:"timesamp"`
	Data      *Data  `json:"data"`
}
