package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/openset/aliyundrive/api"
)

func main() {
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
