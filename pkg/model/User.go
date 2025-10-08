// src\main\java\com\uednd\p2pchat\model\User.java equivalent
package model

import "fmt"

// 在之前的编译器项目中，通过 C 结构体来模拟实现面向对象。Go 和 C 差不多，这里也用结构体实现。

/**
 * 用户类，表示聊天用户
 *
 * @version 1.0.0
 * @since 2025-10-08
 */
type User struct {
	// 用户名
	username string

	// IP地址
	ipAddress string

	// 端口号
	port int
}

// 新建用户
func NewUser(username, ipAddress string, port int) *User {
	return &User {username, ipAddress, port}
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) SetUsername(username string) string {
	user.username = username
}

func (user *User) GetIpAddress() string {
    return user.ipAddress
}

func (user *User) SetIpAddress(ipAddress string) {
    user.ipAddress = ipAddress
}

func (user *User) GetPort() int {
    return user.port
}

func (user *User) SetPort(port int) {
    user.port = port
}

func (user *User) String() string {
    return fmt.Sprintf("User{username='%s', ipAddress='%s', port=%d}", user.username, user.ipAddress, user.port)
}