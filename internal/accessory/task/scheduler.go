package task

import (
	"context"
	"fmt"
	"log"
	"time"

	common "lostark/internal/common"
	commonPkg "lostark/pkg/common"
	"lostark/pkg/models"

	"github.com/hibiken/asynq"
)

const TypeAccessoryPriceUpdate = "price:update"

type TaskHandler struct {
	accessoryRepo *models.AccessoryRepository
}

func NewTaskHandler(accessoryRepo *models.AccessoryRepository) *TaskHandler {
	return &TaskHandler{
		accessoryRepo: accessoryRepo,
	}
}

func (h *TaskHandler) HandlePriceUpdateTask(ctx context.Context, _ *asynq.Task) error {
	log.Printf("Starting accessory price update at %v", time.Now())

	// rowsAffected, err := h.accessoryRepo.UpdatePrices(ctx)
	// if err != nil {
	// 	return err
	// }

	itemsChan := make(chan []commonPkg.CustomAuctionItem, 10)

	go func() {
		common.GetAllAuctionsByLostarkAPI("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IktYMk40TkRDSTJ5NTA5NWpjTWk5TllqY2lyZyIsImtpZCI6IktYMk40TkRDSTJ5NTA5NWpjTWk5TllqY2lyZyJ9", itemsChan)
	}()

	items := <-itemsChan
	for _, item := range items {
		fmt.Printf("아이템 이름: %s, ", item.Grade)
	}

	// log.Printf("Successfully updated prices for %d accessories", rowsAffected)
	return nil
}

func (h *TaskHandler) RegisterHandlers(mux *asynq.ServeMux) {
	mux.HandleFunc(TypeAccessoryPriceUpdate, h.HandlePriceUpdateTask)
}

type Scheduler struct {
	scheduler *asynq.Scheduler
}

func NewScheduler(redisAddr string) (*Scheduler, error) {
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: redisAddr},
		&asynq.SchedulerOpts{},
	)

	return &Scheduler{
		scheduler: scheduler,
	}, nil
}

func (s *Scheduler) RegisterPriceUpdateTask() error {
	task := asynq.NewTask(TypeAccessoryPriceUpdate, nil)
	entryID, err := s.scheduler.Register("0 * * * *", task)
	if err != nil {
		return err
	}

	log.Printf("Registered price update task with entry ID: %s", entryID)
	return nil
}

func (s *Scheduler) Start() error {
	return s.scheduler.Start()
}

func (s *Scheduler) Shutdown() {
	s.scheduler.Shutdown()
}
