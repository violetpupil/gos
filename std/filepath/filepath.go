package filepath

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

var (
	// 拼接文件路径，使用os特定分隔符
	Join = filepath.Join
)

// Walk 遍历以root为根的文件树，打印树上的每个文件和目录信息，包括根
// 按词汇顺序遍历
func Walk(root string) error {
	err := filepath.Walk(root, WalkFunc)
	return err
}

// WalkFunc 遍历处理函数，打印信息
func WalkFunc(path string, info fs.FileInfo, err error) error {
	fmt.Printf("Path=%s IsDir=%v Error=%v\n", path, info.IsDir(), err)
	return nil
}

// Filenames 获取指定目录下所有文件路径
func Filenames(path string) ([]string, error) {
	// 根据路径模式，返回文件列表
	// *代表任意字符串
	return filepath.Glob(filepath.Join(path, "*"))
}
