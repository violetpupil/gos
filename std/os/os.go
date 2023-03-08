package os

import (
	"os"
	"time"
)

type FileInfo struct {
	Name    string      // 文件名
	Size    int64       // 文件字节数
	Mode    os.FileMode // 文件模式
	ModTime time.Time   // 更新时间
	IsDir   bool        // 是否为目录
}

// Stat 获取文件信息
func Stat(name string) (*FileInfo, error) {
	info, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	fi := &FileInfo{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
	}
	return fi, nil
}
