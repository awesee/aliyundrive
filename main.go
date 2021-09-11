package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/openset/aliyundrive/api"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

var home, _ = os.UserHomeDir()

var allFilesPath = filepath.Join(home, "all_files.json")

func main() {
	AllFiles()
	DeleteDuplicateFile()
}

func AllFiles() {
	allFiles := make(map[string][]api.FileListItemV3)
	err := Walk("", "root", func(filePath string, item api.FileListItemV3) {
		item.FullName = path.Join(filePath, item.Name)
		allFiles[item.ContentHash] = append(allFiles[item.ContentHash], item)
		n := len(allFiles[item.ContentHash])
		fmt.Println(n, item.FullName)
	})
	fmt.Println(err)
	data, _ := json.MarshalIndent(allFiles, "", "\t")
	_ = os.WriteFile(allFilesPath, data, os.ModePerm)
}

func Walk(filePath, root string, fn func(string, api.FileListItemV3)) error {
	if root != "root" && !strings.HasPrefix(filePath, "来自分享") {
		return nil
	}
	result, err := api.FileListV3(root)
	if err != nil {
		return err
	}
	for _, item := range result.Items {
		if item.Type == "folder" {
			filePath := path.Join(filePath, item.Name)
			err = Walk(filePath, item.FileID, fn)
			if err != nil {
				return err
			}
		} else {
			fn(filePath, item)
		}
	}
	return nil
}

func DeleteDuplicateFile() {
	allFiles := make(map[string][]api.FileListItemV3)
	data, err := os.ReadFile(allFilesPath)
	check(err)
	err = json.Unmarshal(data, &allFiles)
	check(err)
	fileIds := make([]string, 0)
	for _, items := range allFiles {
		if len(items) < 2 || items[0].Size < 0*MB {
			continue
		}
		sort.Slice(items, func(i, j int) bool {
			return items[i].Name > items[j].Name
		})
		for _, item := range items[1:] {
			fileIds = append(fileIds, item.FileID)
			fmt.Println(items[0].FullName, item.FullName, err)
		}
	}
	result, _ := api.RecycleBinTrashBatchV2(fileIds)
	data, err = json.MarshalIndent(result, "", "\t")
	fmt.Printf("%s\n%v\n", data, err)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
