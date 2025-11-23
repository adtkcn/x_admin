package core

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

// CronManager 任务管理器
type CronManager struct {
	cron    *cron.Cron
	taskIDs map[string]cron.EntryID // 通过字符串ID映射到cron的内部任务ID
	mutex   sync.RWMutex
}

// NewTaskManager 创建一个新的任务管理器
func NewCronManager() *CronManager {
	return &CronManager{
		cron:    cron.New(cron.WithSeconds()), // 启用秒级别的调度
		taskIDs: make(map[string]cron.EntryID),
	}
}

// AddTask 添加任务
func (tm *CronManager) AddTask(taskID, spec string, cmd func()) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// 如果任务已存在，先移除
	if id, exists := tm.taskIDs[taskID]; exists {
		tm.cron.Remove(id)
	}

	// 添加新任务
	id, err := tm.cron.AddFunc(spec, cmd)
	if err != nil {
		return fmt.Errorf("添加任务失败: %w", err)
	}
	tm.taskIDs[taskID] = id
	fmt.Printf("任务 '%s' 已添加/更新\n", taskID)
	return nil
}

// RemoveTask 删除任务
func (tm *CronManager) RemoveTask(taskID string) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	if id, exists := tm.taskIDs[taskID]; exists {
		tm.cron.Remove(id)
		delete(tm.taskIDs, taskID)
		fmt.Printf("任务 '%s' 已移除\n", taskID)
	} else {
		fmt.Printf("任务 '%s' 不存在\n", taskID)
	}
}

// Start 启动任务调度器
func (tm *CronManager) Start() {
	tm.cron.Start()
}

// Stop 停止任务调度器
func (tm *CronManager) Stop() {
	tm.cron.Stop()
}
