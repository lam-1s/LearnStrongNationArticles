package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func dfs(path string) (ret uint) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(path)
		return 0xFFFFFFFF
	}

	for _, f := range files {
		if !f.IsDir() {
			if f.Size() > 0 {
				ret++
			}
			continue
		}
		subDir := dfs(path + "/" + f.Name())
		if subDir == 0 {
			fileName := path + "/" + f.Name()
			err = os.Remove(fileName)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
				continue
			}
			fmt.Printf("[ DIR ] Removed %s which is empty\n", fileName)
		}
		ret += subDir
	}

	return
}

func main() {
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		// Skip directories
		if err != nil || info.IsDir() {
			return err
		}

		if info.Size() < 128 {
			if info.Size() == 0 {
				fmt.Printf("[EMPTY] Removed file %s which is empty\n", path)
			} else {
				fmt.Printf("[ DEL ] Removed file %s which has size %d\n", path, info.Size())
			}
			err = os.Remove(path)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
				return err
			}
			return nil
		}

		thisFile, err := os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return err
		}
		defer thisFile.Close()
		data, _ := ioutil.ReadAll(thisFile)
		originalSize := len(data)
		dataStr := string(data)
		dataStr = strings.ReplaceAll(dataStr, "&nbsp;", "")
		dataStr = strings.ReplaceAll(dataStr, "\t", "")
		dataStr = strings.ReplaceAll(dataStr, "\v", "")
		for strings.Index(dataStr, "  ") != -1{
			dataStr = strings.ReplaceAll(dataStr, "  ", " ")
		}
		data = []byte(dataStr)
		trimmedSize := len(data)
		if trimmedSize == originalSize {
			return nil
		}

		if trimmedSize < 128 {
			thisFile.Close()
			if trimmedSize == 0 {
				fmt.Printf("[EMPTY] Removed file %s which is empty\n", path)
			} else {
				fmt.Printf("[ DEL ] Removed file %s which has size %d\n", path, trimmedSize)
			}
			err = os.Remove(path)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
				return err
			}
			return nil
		}

		fmt.Printf("[ TRI ] File %s, trimmed %d -> %d\n", info.Name(), originalSize, trimmedSize)
		thisFile.Truncate(0)
		thisFile.Write(data)
		return nil
	})
	dfs(".")
}
