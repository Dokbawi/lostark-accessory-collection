package common

import (
	"log"
	accessoryPkg "lostark/pkg/accessory"
	commonPkg "lostark/pkg/common"
	"sync"

	"github.com/go-resty/resty/v2"
)

func GetAllAuctionsByLostarkAPI(token string, itemsChan chan<- []commonPkg.CustomAuctionItem) {
	defer close(itemsChan)

	client := resty.New()
	config := accessoryPkg.NewConfig()

	maxValue := config.GetMaxValue(config.PolishingEffect["공격력 %"])

	params := commonPkg.AuctionRequest{}
	params.ItemTier = 4
	params.CategoryCode = 200010
	params.ItemGrade = "고대"
	params.ItemGradeQuality = 67
	params.EtcOptions = []commonPkg.EtcOption{
		{FirstOption: 8, SecondOption: 1, MinValue: 13, MaxValue: 13},
		{FirstOption: 7, SecondOption: config.PolishingEffect["공격력 %"], MinValue: maxValue, MaxValue: maxValue},
	}

	var wg sync.WaitGroup

	for page := 1; page <= 10; page++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			resp := fetchAuctionPage(client, token, params)

			log.Println(resp.Items)
			itemsChan <- convertToCustomItems(resp.Items)
		}(page)
	}

	go func() {
		wg.Wait()
		close(itemsChan)
	}()
}

func fetchAuctionPage(client *resty.Client, token string, params commonPkg.AuctionRequest) commonPkg.AuctionResponse {
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

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetAuthToken(token).
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
