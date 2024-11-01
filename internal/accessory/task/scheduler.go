// tasks/accessory_price_task.go
package tasks

import (
	"fmt"
	"log"
	"lostark/internal/common"
	commonPkg "lostark/pkg/common"
	"time"
)

type AccessoryPriceTask struct {
	taskManager *TaskManager
}

func NewAccessoryPriceTask(tm *TaskManager) *AccessoryPriceTask {
	return &AccessoryPriceTask{
		taskManager: tm,
	}
}

func (t *AccessoryPriceTask) RunImmediately() {
	log.Println("Starting immediate price update...")
	t.UpdateAccessoryPrices()
}

func (t *AccessoryPriceTask) Schedule() error {
	_, err := t.taskManager.AddTask("0 0 * * * *", "accessory_price_update", t.UpdateAccessoryPrices)
	return err
}

func (t *AccessoryPriceTask) UpdateAccessoryPrices() {
	log.Printf("[AccessoryPriceTask] Started at %v", time.Now())

	itemsChan := make(chan []commonPkg.CustomAuctionItem, 10)

	go func() {
		common.GetAllAuctionsByLostarkAPI(itemsChan)
	}()

	items := <-itemsChan
	for _, item := range items {
		fmt.Printf("아이템 이름: %s, ", item.Grade)
	}

	log.Printf("[AccessoryPriceTask] Completed successfully")
}
