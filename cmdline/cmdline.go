// Package cmdline 命令行操作面板
package cmdline

import (
	"filecollector/service"
	"filecollector/tool"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var (
	// 批次大小
	batchSize = 10
	// 元文件夹
	sourceFolder string
	// 目标文件夹
	targetFolder string
)

// 操作菜单
const constMenu string = `
--------------------------------------------------------------
请选择操作:
	0: 批次大小 %v
	1: 设置源目录 %v
	2: 设置目标目录 %v
	3: 执行整理程序
	9: 退出系统
--------------------------------------------------------------
`

// ShowOperationMenu 输出控制菜单
func ShowOperationMenu() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	for {
		clearCmdline()

		printMenu()

		selectedMenu, _ := scanUserInput()
		switch selectedMenu {
		case "0":
			fmt.Print("请输入批次大小：")
			s, err := scanUserInput()
			if err != nil {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}

			size, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}
			if size <= 0 {
				fmt.Println("批次大小不能小于1！")
				continue
			}
			batchSize = int(size)
		case "1":
			fmt.Print("请输入元目录：")
			s, err := scanUserInput()
			if err != nil {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}
			rst, err := tool.IsFolder(s)
			if err != nil {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}
			if rst {
				sourceFolder = s
			} else {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}
		case "2":
			fmt.Print("请输入目标目录：")
			t, err := scanUserInput()
			if err != nil {
				fmt.Println("err")
				fmt.Scanln()
				continue
			}
			rst, err := tool.IsFolder(t)
			if err != nil {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}
			if rst {
				targetFolder = t
			} else {
				fmt.Println(err)
				fmt.Scanln()
				continue
			}
		case "3":
			err := service.Collect(sourceFolder, targetFolder, batchSize)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Scanln()
		case "9":
			os.Exit(0)
		default:
			fmt.Println("输入的选项不存在！")
			fmt.Scanln()
			continue
		}
	}
}

// ParseFlag 设置命令行参数，输出Help帮助
func ParseFlag() {
	flag.IntVar(&batchSize, "bs", 10, "批次大小(一个批次处理的文件大小)")
	flag.StringVar(&sourceFolder, "s", "", "元文件夹")
	flag.StringVar(&targetFolder, "t", "", "目标文件夹")
	flag.Parse()
}

// printMenu 打印菜单
func printMenu() {
	fmt.Printf(constMenu, batchSize, sourceFolder, targetFolder)
}

// scanUserInput 获取用户输入
func scanUserInput() (string, error) {
	fmt.Print("> ")

	var inputStr string
	_, err := fmt.Scanln(&inputStr)

	if err != nil {
		return "", err
	}

	return inputStr, nil
}

// clearCmdline 清屏
func clearCmdline() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
