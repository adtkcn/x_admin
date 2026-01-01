package corn

import (
	"fmt"
	"sync"
	"x_admin/app/service/cornService"

	"github.com/robfig/cron/v3"
)

func NewCronManager() *CronManager {
	return &CronManager{
		cron:    cron.New(cron.WithSeconds()), // 启用秒级别的调度
		taskIDs: make(map[string]cron.EntryID),
	}
}

// CronManager 任务管理器
type CronManager struct {
	cron    *cron.Cron
	taskIDs map[string]cron.EntryID // 通过字符串ID映射到cron的内部任务ID
	mutex   sync.RWMutex
}

// NewTaskManager 创建一个新的任务管理器

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
func (tm *CronManager) RemoveAllTask() {
	// 移除所有任务
	for _, EntryID := range tm.taskIDs {
		tm.cron.Remove(EntryID)
	}
	tm.taskIDs = make(map[string]cron.EntryID)
	fmt.Printf("所有任务已移除\n")
}

// AddTask 添加、更新任务
func (tm *CronManager) AddTask(taskID, CronExpr string, cmd func()) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// 如果任务已存在，先移除
	if id, exists := tm.taskIDs[taskID]; exists {
		tm.cron.Remove(id)
	}

	// 添加新任务
	id, err := tm.cron.AddFunc(CronExpr, cmd)
	if err != nil {
		return fmt.Errorf("添加任务失败: %w", err)
	}
	tm.taskIDs[taskID] = id
	fmt.Printf("任务 '%s' 已添加/更新\n", taskID)
	return nil
}

// 批量添加任务，先移除所有任务
func (tm *CronManager) AddTasksBeforeRemoveAll(tasks []cornService.RunTask) error {
	// // 获取已存在的任务ID
	// existingTaskIDs := []string{}
	// for id := range tm.taskIDs {
	// 	existingTaskIDs = append(existingTaskIDs, id)
	// }
	// // 获取需要删除idkey
	// // 从任务列表中移除已存在的任务ID
	// for _, id := range existingTaskIDs {
	// 	for _, task := range tasks {
	// 		if task.TaskId == id {
	// 			tm.RemoveTask(id)
	// 		}
	// 	}
	// }

	// 移除所有任务
	tm.RemoveAllTask()
	for _, task := range tasks {
		if task.TaskId == "" || task.CronExpr == "" || task.Task == nil {
			return fmt.Errorf("任务ID、Cron表达式或任务函数不能为空")
		}
		if err := tm.AddTask(task.TaskId, task.CronExpr, task.Task.TaskFunc); err != nil {
			return err
		}
	}
	return nil
}

// Start 启动任务调度器
func (tm *CronManager) Start() {
	tm.cron.Start()
}

// Stop 停止任务调度器
func (tm *CronManager) Stop() {
	tm.cron.Stop()
}
