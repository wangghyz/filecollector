package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"platform"
	"sync"
	"time"
	"tool"
)

// Collect 按时间整理文件
func Collect(source, target string, batchSize int) error {
	rst, err := tool.IsFolder(source)
	if err != nil {
		return err
	}
	if !rst {
		return errors.New("输入的元不是目录！")
	}
	rst2, err := tool.IsFolder(target)
	if err != nil {
		return err
	}
	if !rst2 {
		return errors.New("输入的目标不是目录！")
	}
	if source == target {
		return errors.New("目标目录不能和元目录相同！")
	}

	sourceFiles := getSourceFiles(source)
	cnt := len(sourceFiles)
	groupCnt := int(cnt / batchSize)

	var from, to int
	wg := new(sync.WaitGroup)
	chCnt := make(chan int, groupCnt+1)
	startTime := time.Now()

	for i := 0; i <= groupCnt; i++ {
		from = i * batchSize
		to = i*batchSize + batchSize
		if to > cnt {
			to = cnt
		}
		wg.Add(1)
		go copyFile(sourceFiles[from:to], target, wg, chCnt)
	}
	wg.Wait()

	endTime := time.Now()

	close(chCnt)
	sum := 0
	for c := range chCnt {
		sum += c
	}

	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf(
		"执行完成！\n总记录数：%v\t成功移动数：%v\n用时：%vs\n",
		cnt,
		sum,
		endTime.Sub(startTime).Seconds())
	fmt.Printf(
		"开始时间：%v\t结束时间：%v\n",
		startTime.Format("2006/01/02 15:04:05.999999999"),
		endTime.Format("2006/01/02 15:04:05.999999999"))
	fmt.Println("--------------------------------------------------------------------------------------")

	return nil
}

// getSourceFiles 获得元文件集合
func getSourceFiles(sourceFolder string) []string {
	files := make([]string, 0, 20)
	filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})
	return files
}

// copyFile 复制（移动）文件
func copyFile(sourceFiles []string, targetFolderBase string, wg *sync.WaitGroup, chCnt chan int) {
	cnt := 0
	defer wg.Done()

	for _, file := range sourceFiles {
		fi, err := os.Stat(file)
		if err != nil {
			log.Println(err)
			continue
		}

		mtim := platform.GetFileCreateTime(fi)
		t := time.Unix(mtim.Sec, mtim.Nsec)

		targetFile, err := createDateFolder(t, targetFolderBase)
		if err != nil {
			log.Println(err)
			continue
		}
		targetFile = targetFile + "/" + fi.Name()

		err2 := os.Rename(file, targetFile)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		fmt.Println(file + " > " + targetFile)
		cnt++
	}
	defer func(chCnt chan int, cnt int) {
		chCnt <- cnt
	}(chCnt, cnt)
}

func createDateFolder(mtime time.Time, targetFolder string) (string, error) {
	year := mtime.Format("2006")
	month := mtime.Format("01")

	// 目标目录：目标base/2006/2006-
	target := targetFolder + "/" + year + "/" + year + "-" + month
	_, err := os.Stat(target)

	if os.IsExist(err) {
		return target, nil
	}
	if os.IsNotExist(err) {
		err1 := os.MkdirAll(target, os.ModePerm)
		if err1 != nil {
			return "", err1
		}
	}
	return target, nil
}
