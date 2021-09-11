package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/openset/aliyundrive/api"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	//AllFiles()
	//DeleteDuplicateFile()
}

func AllFiles() {
	allFiles := make(map[string][]api.FileListItemV3)
	err := Walk("root", func(item api.FileListItemV3) {
		allFiles[item.ContentHash] = append(allFiles[item.ContentHash], item)
		n := len(allFiles[item.ContentHash])
		fmt.Println(item.Name, n)
		if n >= 2 {
			fmt.Println(allFiles[item.ContentHash])
			time.Sleep(5 * time.Second)
		}
	})
	fmt.Println(err)
	data, _ := json.Marshal(allFiles)
	_ = os.WriteFile("all_files.json", data, os.ModePerm)
}

func Walk(root string, fn func(api.FileListItemV3)) error {
	result, err := api.FileListV3(root)
	if err != nil {
		return err
	}
	for _, item := range result.Items {
		if item.Type == "folder" {
			err = Walk(item.FileID, fn)
			if err != nil {
				return err
			}
		} else {
			fn(item)
		}
	}
	return nil
}

func DeleteDuplicateFile() {
	allFiles := make(map[string][]api.FileListItemV3)
	data, err := os.ReadFile("all_files.json")
	check(err)
	err = json.Unmarshal(data, &allFiles)
	check(err)
	for _, items := range allFiles {
		if len(items) < 2 || items[0].Size < GB {
			continue
		}
		sort.Slice(items, func(i, j int) bool {
			return items[i].Name < items[j].Name
		})
		for _, item := range items[1:] {
			err := api.RecycleBinTrashV2(item.FileID)
			fmt.Println(item.Name, err)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
