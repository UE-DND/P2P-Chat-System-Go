// src\main\java\com\uednd\p2pchat\repository\ChatRepository.java equivalent
package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/UE-DND/P2P-Chat-System-Go/pkg/model"
	"fmt"
)

// 数据库管理类，负责处理SQLite数据库操作
type ChatRepository struct {
	sql_path string  // 数据库文件夹路径
	db 		 *sql.DB // 数据库连接
}

func NewChatRepository(sql_path string) *ChatRepository {
	return &ChatRepository{
		sql_path: sql_path,
	}
}

// 初始化数据库连接
func (cr *ChatRepository) InitDatabase() error {
	db, err := sql.Open("sqlite3", cr.sql_path)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	cr.db = db  // 保存该连接

	// 更新数据库：创建用户表
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			ip_address TEXT,
			port INTEGER
		)`
	if _, err := cr.db.Exec(createUsersTable); err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

    // 更新数据库：创建消息记录表
    createMessagesTable := `
        CREATE TABLE IF NOT EXISTS messages (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            sender TEXT NOT NULL,
            receiver TEXT NOT NULL,
            content TEXT NOT NULL,
            type TEXT NOT NULL,
            file_path TEXT
        )`
	if _, err := cr.db.Exec(createMessagesTable); err != nil {
		return fmt.Errorf("failed to create messages table: %v", err)
	}

	return nil
}

// 保存用户信息
func (cr *ChatRepository) SaveUser(u *model.User) error {
	var userID int  // 作检查用，实际无作用
	query := "SELECT id FROM users WHERE username = ?"
	err := cr.db.QueryRow(query, u.GetUsername()).Scan(&userID)

	if err == sql.ErrNoRows {
		// 用户不存在，插入新用户
		insertQuery := `
			INSERT INTO users (username, ip_address, port)
			VALUES (?, ?, ?)
		`
		_, err := cr.db.Exec(insertQuery, u.GetUsername(), u.GetIpAddress(), u.GetPort())
		if err != nil {
			return fmt.Errorf("failed to insert user: %v", err)
		}
		return nil
	} else if err == nil {
		// 用户存在，更新用户信息
		updateQuery := `
			UPDATE users SET ip_address = ?, port = ?
			WHERE username = ?
		`

		// db.Exec() 用于修改操作
		_, err = cr.db.Exec(updateQuery, u.GetIpAddress(), u.GetPort(), u.GetUsername())
		if err != nil {
			return fmt.Errorf("failed to update user: %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("failed to query user: %v", err)
	}
}

// 保存消息记录
func (cr *ChatRepository) SaveMessage(msg *model.Message) error {
	insertQuery := `
		INSERT INTO messages (sender, receiver, content, type, file_path)
        VALUES (?, ?, ?, ?, ?)
	`

	// 这里使用占位符，修复源代码中可能存在的SQL注入问题
    _, err := cr.db.Exec(insertQuery,
        msg.GetSender(),
        msg.GetReceiver(),
        msg.GetContent(),
        msg.GetType(),
        msg.GetFilePath())

    if err != nil {
        return fmt.Errorf("failed to save message: %w", err)
    }

    return nil
}

// 获取与特定用户的聊天记录
func (cr *ChatRepository) GetChatHistory(user1, user2 string) ([]*model.Message, error) {
	getQuery := `
		SELECT sender, receiver, content, type, file_path
		FROM messages
		WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)
		ORDER BY id
	`

	// db.Query() 用于查询操作
	result, err := cr.db.Query(getQuery, user1, user2, user2, user1)
	if err != nil {
		return nil, fmt.Errorf("failed to query chat history: %v", err)
	}

	defer result.Close()  // defer：无论函数如何退出，都会执行

	var messages []*model.Message
	for result.Next() {
		message := model.NewMessage()
		var sender, receiver, content, msgType, filePath string

		err := result.Scan(&sender, &receiver, &content, &msgType, &filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %v", err)
		}

		message.SetSender(sender)
		message.SetReceiver(receiver)
		message.SetContent(content)
		message.SetType(msgType)
		message.SetFilePath(filePath)

		// 将当前消息加入查询结果
		messages = append(messages, message)
	}

	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return messages, nil
}

// 清除特定用户的聊天记录
func (cr *ChatRepository) ClearChatHistory(user1, user2 string) error {
	deleteQuery := `
		DELETE FROM messages
		WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)
	`

	_, err := cr.db.Exec(deleteQuery, user1, user2, user2, user1)
	if err != nil {
		return fmt.Errorf("failed to clear chat history: %v", err)
	}

	return nil
}

// 关闭数据库连接
func (cr *ChatRepository) CloseConnection() error {
	if cr.db != nil {
		err := cr.db.Close()
		cr.db = nil
		if err != nil {
			return fmt.Errorf("failed to close database connection: %v", err)
		}
	}
	return nil
}