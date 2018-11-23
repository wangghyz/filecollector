package tool

import (
	"errors"
	"os"
)

// IsFolder 是否为文件夹
func IsFolder(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, errors.New("文件或目录错误！")
	}
	return fi.IsDir(), nil
}
