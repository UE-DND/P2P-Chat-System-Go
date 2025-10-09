// src\main\java\com\uednd\p2pchat\util\NetworkUtils.java equivalent
package util

import (
	"fmt"
	"net"
	"os"
)

// 网络基础工具
type PortStatus int

const (
	AVAILABLE 	   PortStatus = 1
	IN_USE    	   PortStatus = 2
	INVALID_RANGE  PortStatus = 3
)

// 获取端口状态
func (ps PortStatus) String() string {
    switch ps {
    case AVAILABLE:
        return "AVAILABLE"
    case IN_USE:
        return "IN_USE"
    case INVALID_RANGE:
        return "INVALID_RANGE"
    default:
        return fmt.Sprintf("UNKNOWN-STATUS: %d", int(ps))
    }
}

// 获取主机名
func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "localhost"
	}
	return hostname
}

// 获取本机IP地址
func GetLocalIpAddress() string {
	// 获取所有接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "127.0.0.1"
	}

	// 遍历所有接口
	// _表示不使用该变量，因为我们不需要关心数组的索引。
	// go不允许声明了但未使用的变量
	for _, iface := range interfaces {
		// 如果是环回接口和未启用的接口，跳过
		if iface.Flags & net.FlagLoopback != 0 || iface.Flags & net.FlagUp == 0 {
			continue
		}

		// 获取接口的所有地址
		addresses, err := iface.Addrs()
		if err != nil {
			continue
		}

		// 遍历所有地址
		for _, address := range addresses {
			// 类型断言
			// if 变量名, 布尔值 := 接口变量.(具体类型); 布尔值
			// 如果类型断言成功，布尔值为true，执行if块内语句
			if ipnet, success := address.(*net.IPNet) ; success {
				// 如果是IPv4且不为环回，返回地址
				if ipnet.IP.To4() != nil && !ipnet.IP.IsLoopback() {
					return ipnet.IP.String()
				}
			}
		}
	}

	return "127.0.0.1"
}

func CheckPort(port int) PortStatus {
	// 检查端口范围是否越界
	if port < 1 || port > 65535 {
		return INVALID_RANGE
	}

	// 尝试监听端口
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return IN_USE
	}

	listener.Close()
	return AVAILABLE
}