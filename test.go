package main

import (
	"filecollector/platform"
	"fmt"
	"os"
	"time"
)

func main() {
	info, e := os.Stat(`D:\02.Workspaces\01.Myself\filecollector\platform\filetime_windows.go`)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	ct := platform.GetFileCreateTime(info)

	unix := time.Unix(ct.Sec, ct.Nsec)
	format := unix.Format("2006-01-02 15:04:05")
	fmt.Println(format)
}
