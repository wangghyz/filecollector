// Package platform 说明
package platform

import (
	"os"
	"syscall"
)

// GetFileCreateTime 获得文件创建时间
func GetFileCreateTime(fi os.FileInfo) syscall.Timespec {
	return fi.Sys().(*syscall.Stat_t).Mtim
}
