// src\main\java\com\uednd\p2pchat\model\Message.java equivalent
package model

import "fmt"

/**
 * 消息类，表示聊天消息
 *
 * @version 1.0.0
 * @since 2025-10-09
 */
type Message struct {
	// 发送者
	sender string

	// 接收者
	receiver string

	// 消息内容
	content string

	// 消息类型（文本、文件、系统消息）
	msgType string

	// 文件路径（如果是文件消息）
	filePath string
}

// 空构造函数
func NewMessage() *Message {
	return &Message{}
}

// 文本消息构造函数
func NewTextMessage(sender, receiver, content string) *Message {
	return &Message{
		sender:   sender,
		receiver: receiver,
		content:  content,
		msgType:  "TEXT",
		filePath: "",
	}
}

// 文件消息构造函数
func NewFileMessage(sender, receiver, content, filePath string) *Message {
	return &Message{
		sender:   sender,
		receiver: receiver,
		content:  content,
		msgType:  "FILE",
		filePath: filePath,
	}
}

func (m *Message) GetSender() string {
	return m.sender
}

func (m *Message) GetReceiver() string {
	return m.receiver
}

func (m *Message) GetContent() string {
	return m.content
}

func (m *Message) GetType() string {
	return m.msgType
}

func (m *Message) GetFilePath() string {
	return m.filePath
}

func (m *Message) SetSender(sender string) {
	m.sender = sender
}

func (m *Message) SetReceiver(receiver string) {
	m.receiver = receiver
}

func (m *Message) SetContent(content string) {
	m.content = content
}

func (m *Message) SetType(msgType string) {
	m.msgType = msgType
}

func (m *Message) SetFilePath(filePath string) {
	m.filePath = filePath
}

func (m *Message) String() string {
	switch m.msgType {
	case "FILE":
		return fmt.Sprintf("%s: [文件] %s", m.sender, m.content)
	case "SYSTEM":
		return fmt.Sprintf("[系统消息] %s", m.content)
	default:
		return fmt.Sprintf("%s: %s", m.sender, m.content)
	}
}
