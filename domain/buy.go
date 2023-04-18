package domain

import "time"

type WebCase struct {
	Sid         int     `json:"sid"`
	DetailUrl   string  `json:"detailUrl"`
	GroupID     string  `json:"groupID"`
	CaseName    string  `json:"caseName"`
	CaseFrom    string  `json:"caseFrom"`
	SimpAddress string  `json:"simpAddress"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	Road        string  `json:"road"`
	Rm          float32 `json:"rm"`
	LivingRm    float32 `json:"livingRm"`
	BathRm      float32 `json:"bathRm"`
	// SpaceRm           `json:"spaceRm"`
	FromFloor   string   `json:"fromFloor"`
	ToFloor     string   `json:"toFloor" `
	UpFloor     int      `json:"upFloor"`
	TotalPrice  float32  `json:"totalPrice"`
	OrigPrice   *float32 `json:"origPrice"`
	BuildPin    float32  `json:"buildPin"`
	LandPin     float32  `json:"landPin"`
	CaseUrl     string   `json:"caseUrl"`
	UnitPrice   float32  `json:"unitPrice"`
	BudName     string   `json:"budName"`
	BuildAge    float32  `json:"buildAge"`
	GroupCount  int      `json:"groupCount"`
	Lat         float32  `json:"lat"`
	Lng         float32  `json:"lng"`
	PurPoseName string   `json:"purPoseName"`
	KWsLabel    string   `json:"kWsLabel"`
	// CommunityId       `json:"communityId"`
	ParkingYN string `json:"parkingYN"`
	// LastTime          `json:"lastTime"`
	NewKeyInDate     *time.Time `json:"newKeyInDate"`
	NewKeyInDaysName string     `json:"newKeyInDaysName"`
	// DownRatio         `json:"downRatio"`
	// PinPriceRatio     `json:"pinPriceRatio"`
	// PinPriceDiff      `json:"pinPriceDiff"`
	// WeightsNumber     `json:"weightsNumber"`
	// UpdateTime        `json:"updateTime"`
	// LastTotalPrice    `json:"lastTotalPrice"`
	MainPin      float32 `json:"mainPin"`
	CaseTypeName string  `json:"caseTypeName"`
	CaseTypePic  string  `json:"CaseTypePic"`
	MPicUrl      string  `json:"mPicUrl"`
	// CommunityTag      `json:"communityTag"`
	// CaseFromList      `json:"caseFromList"`
	// CommunityKwsLabel `json:"communityKwsLabel"`
	// ImageFileList     `json:"imageFileList"`
	// Tags              `json:"tags"`
}

type BuyRes struct {
	WebCaseGrouping []*WebCase `json:"webCaseGrouping"`
	Code            int        `json:"code"`
}
