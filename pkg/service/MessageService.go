// src\main\java\com\uednd\p2pchat\service\MessageService.java equivalent
package service

import (
	"context"
	"errors"
	"time"

	"github.com/UE-DND/P2P-Chat-System-Go/pkg/core"
	"github.com/UE-DND/P2P-Chat-System-Go/pkg/repository"
)

// 消息处理器接口，定义了消息处理和错误处理方法
type MessageHandler interface {
	HandleMessage(message string)
	HandleError(errMsg string)
}

type MessageService struct {
	bgService        *core.BackgroundService   // 后台任务
	networkManager   *network.Manager          // 网络管理器
	chatRepo         repository.ChatRepository // 数据库接口
	localUsername    string                    // 本地用户名
	oppositeUsername string                    // 对方用户名
	handler          MessageHandler            // 消息处理器接口
	pollInterval     time.Duration             // 轮询间隔
}

// 创建消息服务后台
func NewMessageService(
	networkManager *network.Manager,
	chatRepo repository.ChatRepository,
	localUser string,
	oppositeUser string,
	handler MessageHandler,
) (*MessageService, error) {
	service := &MessageService{
		networkManager:   networkManager,
		chatRepo:         chatRepo,
		localUsername:    localUser,
		oppositeUsername: oppositeUser,
		handler:          handler,
		pollInterval:     100 * time.Millisecond,
	}

	// 定义后台任务（消息服务）
	taskFn := func(context context.Context) error {
		if networkManager.IsConnected() { // 检查连接状态
			msg, err := networkManager.ReceiveTextMessage(context) // 接收消息

			if err != nil {
				return err // 返回错误，触发停止
			}
			if msg == nil {
				return errors.New("对方已断开连接") // 处理断开
			}

			handler.HandleMessage(*msg) // 调用处理器处理消息

			messageRecord := model.NewMessage(localUser, oppositeUser, *msg) // 创建消息记录
			messageRecord.Type = model.MessageTypeText                       // 设置类型
			return chatRepo.SaveMessage(messageRecord)                       // 保存到数据库
		}

		select { // 未连接时轮询等待
		case <-context.Done():
			return context.Err() // 上下文取消时退出
		case <-time.After(service.pollInterval): // 等待间隔后继续
			return nil
		}
	}

	// 错误处理方式定义
	errHandler := func(err error) {
		handler.HandleError(err.Error())
	}

	bg, err := core.NewBackgroundService(taskFn, errHandler) // 调用顶层创建 BackgroundService
	if err != nil {
		return nil, err
	}
	service.bgService = bg
	return service, nil
}

func (msgbg *MessageService) Start() {
	msgbg.bgService.Start()
}

func (msgbg *MessageService) Stop() {
	msgbg.bgService.Stop()
}

func (msgbg *MessageService) IsRunning() bool {
	return msgbg.bgService.IsRunning()
}

func (msgbg *MessageService) SendTextMessage(context context.Context, content string) error {
	if !msgbg.networkManager.IsConnected() {
		return errors.New("未连接，无法发送消息")
	}

	if err := msgbg.networkManager.SendTextMessage(context, content); err != nil {
		return err
	}

	messageRecord := model.NewMessage(msgbg.localUsername, msgbg.oppositeUsername, content) // 创建记录
	messageRecord.Type = model.MessageTypeText                                              // 设置消息类型
	return msgbg.chatRepo.SaveMessage(messageRecord)                                        // 保存
}

func (msgbg *MessageService) GetChatHistory(context context.Context) ([]model.Message, error) {
	return msgbg.chatRepo.GetChatHistory(msgbg.localUsername, msgbg.oppositeUsername)
}

func (msgbg *MessageService) ClearChatHistory(context context.Context) error {
	return msgbg.chatRepo.ClearChatHistory(msgbg.localUsername, msgbg.oppositeUsername)
}
