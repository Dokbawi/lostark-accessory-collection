package commonPkg

type AuctionResponse struct {
	TotalCount int           `json:"TotalCount"`
	PageSize   int           `json:"PageSize"`
	Items      []interface{} `json:"Items"`
}

type CustomAuctionItem struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Grade        string  `json:"grade"`
	AuctionPrice int     `json:"auction_price"`
	BuyPrice     int     `json:"buy_price"`
	TradeLeft    int     `json:"trade_left"`
	Quality      float64 `json:"quality"`
}

type AuctionRequest struct {
	ItemGradeQuality int         `json:"item_grade_quality"`
	ItemTier         int         `json:"item_tier"`
	ItemGrade        string      `json:"item_grade"`
	CategoryCode     int         `json:"category_code"`
	EtcOptions       []EtcOption `json:"etc_options"`
	PageNo           int         `json:"page_no"`
	Sort             string      `json:"sort"`
	SortCondition    string      `json:"sort_condition"`
}

type EtcOption struct {
	FirstOption  int `json:"first_option"`
	SecondOption int `json:"second_option"`
	MinValue     int `json:"min_value"`
	MaxValue     int `json:"max_value"`
}

var BASE_URL = "https://developer-lostark.game.onstove.com"
