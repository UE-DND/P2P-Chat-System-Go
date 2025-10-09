// src\main\java\com\uednd\p2pchat\util\DirectoryUtils.java equivalent
package util

import "os"

// 目录创建工具类
func CreateDirectoryIfNotExists(path string) bool {
	info, err := os.Stat(path)  // 获取目录信息

	if err == nil {
		// 如果路径存在，检查是否为目录
		return info.IsDir()
	}

	if os.IsNotExist(err) {
		// 如果路径不存在，尝试创建目录
		return os.MkdirAll(path, 0755) == nil
	}

	return false  // 创建失败
}
