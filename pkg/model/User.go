// src\main\java\com\uednd\p2pchat\model\User.java equivalent
package model

import "fmt"

// 在之前的编译器项目中，通过 C 结构体来模拟实现面向对象。Go 和 C 差不多，这里也用结构体实现。

// 用户类，表示聊天用户
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
	return &User{
		username:  username,
		ipAddress: ipAddress,
		port:      port,
	}
	// return &User{username, ipAddress, port} 为位置参数，适用于不需要初始化的情况下的简短写法
	// 这里统一为命名参数写法
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) GetIpAddress() string {
    return u.ipAddress
}

func (u *User) SetIpAddress(ipAddress string) {
    u.ipAddress = ipAddress
}

func (u *User) GetPort() int {
    return u.port
}

func (u *User) SetPort(port int) {
    u.port = port
}

func (u *User) String() string {
    return fmt.Sprintf("User{username='%s', ipAddress='%s', port=%d}", u.username, u.ipAddress, u.port)
}