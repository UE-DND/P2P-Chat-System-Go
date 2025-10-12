// src\main\java\com\uednd\p2pchat\core\BackgroundService.java equivalent
package core

import (
	"context"
	"fmt"
	"sync"
)

// 封装后台任务的抽象类，简化线程管理
type BackgroundService struct {
	context    context.Context             // 管理后台任务
	cancel     context.CancelFunc          // 停止后台任务
	waitgroup  sync.WaitGroup              // 确保任务完全退出
	running    bool                        // 表示服务是否正在运行
	runningMux sync.RWMutex                // 读写锁
	taskFn     func(context.Context) error // 返回错误信息
	errHandler func(error)                 // 打印错误信息
}

// 创建后台任务
func NewBackgroundService(taskFn func(context.Context) error, errHandler func(error)) (*BackgroundService, error) {
	if errHandler == nil {
		errHandler = func(err error) {
			fmt.Printf("创建 %T 后台任务失败: %v\n", taskFn, err)
		}
	}
	return &BackgroundService{taskFn: taskFn, errHandler: errHandler}, nil
}

// 启动后台任务
func (bs *BackgroundService) Start() {
	bs.runningMux.Lock() // 写上锁，避免竟态

	// 如果正在运行，直接退出启动
	if bs.running {
		bs.runningMux.Unlock()
		return
	}

	bs.context, bs.cancel = context.WithCancel(context.Background()) // 创建上下文
	bs.running = true
	bs.runningMux.Unlock()
	bs.waitgroup.Add(1) // 确保停止后台时可等待此 goroutine 完成

	go func() {
		defer bs.waitgroup.Done() // 退出时标记为完成

		// 循环执行
		for {
			select {
			case <-bs.context.Done():
				return // 后台被取消，退出
			default:
				if err := bs.taskFn(bs.context); err != nil { // 调用用户定义的任务函数，传入上下文
					bs.Stop()
					return
				}
			}
		}
	}()
}

// 停止后台任务
func (bs *BackgroundService) Stop() {
	bs.runningMux.Lock()

	// 如果未运行，直接退出
	if !bs.running {
		bs.runningMux.Unlock()
		return
	}

	bs.running = false
	bs.runningMux.Unlock()

	// 执行退出
	bs.cancel()
	bs.waitgroup.Wait()
}

func (bs *BackgroundService) IsRunning() bool {
	defer bs.runningMux.RUnlock()

	// 获取并返回读锁状态
	bs.runningMux.RLock()
	return bs.running
}
