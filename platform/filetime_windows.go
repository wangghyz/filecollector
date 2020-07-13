// Package platform 说明
package platform

import (
	"os"
	"syscall"
)

// GetFileCreateTime 获得文件创建时间
func GetFileCreateTime(fi os.FileInfo) syscall.Timespec {
	filetime := fi.Sys().(*syscall.Win32FileAttributeData).CreationTime
	return syscall.Timespec{
		Nsec: filetime.Nanoseconds(),
	}
}
