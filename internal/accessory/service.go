package accessory

import (
	"log"
	common "lostark/internal/common"
	commonPkg "lostark/pkg/common"
	pb "lostark/proto/accessory"
)

type AccessoryService struct {
	pb.UnimplementedAccessoryServiceServer
}

func NewService() *AccessoryService {
	return &AccessoryService{}
}

func (s *AccessoryService) GetAverageTradePrice(req *pb.GetAverageTradePriceRequest, stream pb.AccessoryService_GetAverageTradePriceServer) error {
	log.Println("req", req)

	progressChan := make(chan float32)
	itemsChan := make(chan []commonPkg.CustomAuctionItem, 10)

	go func() {
		common.GetAllAuctionsByLostarkAPI(itemsChan)
	}()

	var allItems []commonPkg.CustomAuctionItem
	for {
		select {
		case progress, ok := <-progressChan:
			if !ok {

				return stream.Send(&pb.GetAverageTradePriceResponseList{
					Items:    convertToProtoItems(allItems),
					Progress: 100,
				})
			}

			err := stream.Send(&pb.GetAverageTradePriceResponseList{Progress: progress})
			if err != nil {
				return err
			}
		case items, ok := <-itemsChan:
			if !ok {
				continue
			}
			allItems = append(allItems, items...)
		}
	}
}

func convertToProtoItems(items []commonPkg.CustomAuctionItem) []*pb.GetAverageTradePriceResponse {
	protoItems := make([]*pb.GetAverageTradePriceResponse, len(items))
	for i, item := range items {
		protoItems[i] = &pb.GetAverageTradePriceResponse{
			Id:           item.ID,
			Name:         item.Name,
			Grade:        item.Grade,
			AuctionPrice: int32(item.AuctionPrice),
			BuyPrice:     int32(item.BuyPrice),
			TradeLeft:    int32(item.TradeLeft),
			Quality:      float32(item.Quality),
		}
	}
	return protoItems
}
