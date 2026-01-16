package corn

import (
	"fmt"
	"sync"
	"x_admin/app/service/cornService"
	"x_admin/core"
	"x_admin/util"

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
func (tm *CronManager) AddTask(taskID, CronExpr string, task cornService.Task) error {

	cmd := task.TaskFunc

	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// 如果任务已存在，先移除
	if id, exists := tm.taskIDs[taskID]; exists {
		tm.cron.Remove(id)
	}

	// 添加新任务
	id, err := tm.cron.AddFunc(CronExpr, func() {
		// 不加锁
		if !task.Lock {
			cmd()
			return
		}
		// 加锁
		lockKey := fmt.Sprintf("lock:%s", taskID)
		lock := util.NewRedisLock(lockKey, task.LockTTL) // 锁自动过期 10s

		if !lock.Lock() {
			core.Logger.Infof("任务加锁失败:%s: %s", task.TaskCode, task.TaskDesc)
			return
		}
		defer func() {
			core.Logger.Infof("任务解锁:%s: %s", task.TaskCode, task.TaskDesc)
			// 解锁失败时，记录日志
			if err := lock.Unlock(); err != nil {
				core.Logger.Error("任务解锁失败:%s: %s, err: %v", task.TaskCode, task.TaskDesc, err)
			}
		}()
		cmd()
	})
	if err != nil {
		return fmt.Errorf("添加任务失败: %w", err)
	}
	tm.taskIDs[taskID] = id
	fmt.Printf("任务 '%s' 已添加/更新\n", taskID)
	return nil
}

// 批量添加任务，先移除所有任务
func (tm *CronManager) AddTasksBeforeRemoveAll(tasks []cornService.RunTask) error {
	// 移除所有任务
	tm.RemoveAllTask()
	for _, task := range tasks {
		if task.TaskId == "" || task.CronExpr == "" || task.Task == nil {
			return fmt.Errorf("任务ID、Cron表达式或任务函数不能为空")
		}
		if err := tm.AddTask(task.TaskId, task.CronExpr, *task.Task); err != nil {
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
