package main

import (
	"cmdline"
)

// init 程序初始化
func init() {
	// 输出命令行Help信息，设置命令行参数
	cmdline.ParseFlag()
}

// 程序main入口
func main() {
	// 启动命令行菜单
	cmdline.ShowOperationMenu()
}
