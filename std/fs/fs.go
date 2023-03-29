package fs

import (
	"io/fs"
	"time"
)

// FileInfo 文件信息接口
type FileInfo = fs.FileInfo

// FileInfoS 文件信息结构体
type FileInfoS struct {
	Name    string      // 文件名
	Size    int64       // 文件字节数
	Mode    fs.FileMode // 文件模式
	ModTime time.Time   // 更新时间
	IsDir   bool        // 是否为目录
}

// FileInfoPath 文件信息和绝对路径
type FileInfoPath struct {
	fs.FileInfo
	Path string
}
