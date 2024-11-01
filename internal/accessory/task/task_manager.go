package tasks

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

type TaskManager struct {
	cron *cron.Cron
	mu   sync.Mutex
	jobs map[cron.EntryID]string
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		cron: cron.New(cron.WithSeconds()),
		jobs: make(map[cron.EntryID]string),
	}
}

func (tm *TaskManager) AddTask(schedule string, name string, task func()) (cron.EntryID, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	id, err := tm.cron.AddFunc(schedule, task)
	if err != nil {
		return 0, fmt.Errorf("failed to add task: %w", err)
	}

	tm.jobs[id] = name
	return id, nil
}

func (tm *TaskManager) Start() {
	tm.cron.Start()
}
