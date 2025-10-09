// src\main\java\com\uednd\p2pchat\model\FileInfo.java equivalent
package model

import "fmt"

// 文件信息类
type FileInfo struct {
	// 文件名
	fileName string

	// 文件大小（字节）
	fileSize int64

	// 文件二进制数据
	fileData []byte

	// 发送者
	sender string

	// 接收者
	receiver string
}

func NewFileInfo(fileName string, fileSize int64, fileData []byte, sender string, receiver string) *FileInfo {
	return &FileInfo{
		fileName: fileName,
		fileSize: fileSize,
		fileData: fileData,
		sender:   sender,
		receiver: receiver,
	}
}

func (f *FileInfo) GetFileName() string {
	return f.fileName
}

func (f *FileInfo) GetFileSize() int64 {
	return f.fileSize
}

func (f *FileInfo) GetFileData() []byte {
	return f.fileData
}

func (f *FileInfo) GetSender() string {
	return f.sender
}

func (f *FileInfo) GetReceiver() string {
	return f.receiver
}

func (f *FileInfo) SetFileName(fileName string) {
	f.fileName = fileName
}

func (f *FileInfo) SetFileSize(fileSize int64) {
	f.fileSize = fileSize
}

func (f *FileInfo) SetFileData(fileData []byte) {
	f.fileData = fileData
}

func (f *FileInfo) SetSender(sender string) {
	f.sender = sender
}

func (f *FileInfo) SetReceiver(receiver string) {
	f.receiver = receiver
}

func (f *FileInfo) String() string {
	return "文件: " + f.fileName + " (大小: " + f.formatFileSize(f.fileSize) + ")"
}

// 格式化文件大小
func (f *FileInfo) formatFileSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
	}
}
