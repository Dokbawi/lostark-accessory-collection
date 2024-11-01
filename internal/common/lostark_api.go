package common

import (
	"log"
	"lostark/config"
	accessoryPkg "lostark/pkg/accessory"
	commonPkg "lostark/pkg/common"
	"sync"

	"github.com/go-resty/resty/v2"
)

func GetAllAuctionsByLostarkAPI(itemsChan chan<- []commonPkg.CustomAuctionItem) {
	client := resty.New()
	config := accessoryPkg.NewConfig()

	params := commonPkg.AuctionRequest{}
	params.ItemTier = 4
	params.CategoryCode = 200010
	params.ItemGrade = "고대"
	params.ItemGradeQuality = 67
	params.EtcOptions = []commonPkg.EtcOption{
		{FirstOption: 7, SecondOption: config.PolishingEffect["추가 피해"], MinValue: 3, MaxValue: 3},
	}

	var wg sync.WaitGroup

	for page := 1; page <= 3; page++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			params.PageNo = pageNum
			resp := fetchAuctionPage(client, params)

			if len(resp.Items) == 0 {
				return
			}
			itemsChan <- convertToCustomItems(resp.Items)
		}(page)
	}

	wg.Wait()
	close(itemsChan)
}

func fetchAuctionPage(client *resty.Client, params commonPkg.AuctionRequest) commonPkg.AuctionResponse {
	etcOptionsMap := make([]map[string]interface{}, len(params.EtcOptions))
	for i, opt := range params.EtcOptions {
		etcOptionsMap[i] = map[string]interface{}{
			"FirstOption":  opt.FirstOption,
			"SecondOption": opt.SecondOption,
			"MinValue":     opt.MinValue,
			"MaxValue":     opt.MaxValue,
		}
	}

	requestBody := map[string]interface{}{
		"ItemTier":         params.ItemTier,
		"CategoryCode":     params.CategoryCode,
		"ItemGradeQuality": params.ItemGradeQuality,
		"ItemGrade":        params.ItemGrade,
		"EtcOptions":       etcOptionsMap,
		"PageNo":           params.PageNo,
		"Sort":             "BUY_PRICE",
		"SortCondition":    "ASC",
	}

	cfg := config.GetConfig()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetAuthToken(cfg.LostarkAPIKey).
		SetResult(&commonPkg.AuctionResponse{}).
		Post(commonPkg.BASE_URL + "/auctions/items")

	if err != nil {
		log.Printf("Error fetching page %d: %v", params.PageNo, err)
		return commonPkg.AuctionResponse{}
	}

	return *resp.Result().(*commonPkg.AuctionResponse)
}

func convertToCustomItems(items []interface{}) []commonPkg.CustomAuctionItem {
	var customItems []commonPkg.CustomAuctionItem
	for _, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			customItem := commonPkg.CustomAuctionItem{
				// ID:           itemMap["AuctionInfo"].(map[string]interface{})["AuctionId"].(string),
				// Name:         itemMap["Name"].(string),
				Grade: itemMap["Grade"].(string),
				// AuctionPrice: int(itemMap["AuctionInfo"].(map[string]interface{})["StartPrice"].(float64)),
				// BuyPrice:     int(itemMap["AuctionInfo"].(map[string]interface{})["BuyPrice"].(float64)),
				// TradeLeft:    int(itemMap["TradeLeft"].(float64)),
				// Quality:      itemMap["GradeQuality"].(float64),
			}
			customItems = append(customItems, customItem)
		}
	}
	return customItems
}
