package domain

import "time"

type TimeStamp time.Time

const (
	TimeLayout = "2006-01-02 15:04:05"
)

type House struct {
	// Type        string  `josn:"type"`
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
	Total     int      `json:"total"`
}

type Res struct {
	Status    int    `json:"status"`
	Timestamp string `json:"timesamp"`
	Data      *Data  `json:"data"`
}
