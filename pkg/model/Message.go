// src\main\java\com\uednd\p2pchat\model\Message.java equivalent
package model

import "fmt"

// 消息类，表示聊天消息
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

func (msg *Message) GetSender() string {
	return msg.sender
}

func (msg *Message) GetReceiver() string {
	return msg.receiver
}

func (msg *Message) GetContent() string {
	return msg.content
}

func (msg *Message) GetType() string {
	return msg.msgType
}

func (msg *Message) GetFilePath() string {
	return msg.filePath
}

func (msg *Message) SetSender(sender string) {
	msg.sender = sender
}

func (msg *Message) SetReceiver(receiver string) {
	msg.receiver = receiver
}

func (msg *Message) SetContent(content string) {
	msg.content = content
}

func (msg *Message) SetType(msgType string) {
	msg.msgType = msgType
}

func (msg *Message) SetFilePath(filePath string) {
	msg.filePath = filePath
}

func (msg *Message) String() string {
	switch msg.msgType {
	case "FILE":
		return fmt.Sprintf("%s: [文件] %s", msg.sender, msg.content)
	case "SYSTEM":
		return fmt.Sprintf("[系统消息] %s", msg.content)
	default:
		return fmt.Sprintf("%s: %s", msg.sender, msg.content)
	}
}
