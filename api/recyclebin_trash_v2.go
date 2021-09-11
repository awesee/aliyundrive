package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RecycleBinTrashV2(fileId string) error {
	uri := "https://api.aliyundrive.com/v2/recyclebin/trash"
	contentType := "application/json;charset=UTF-8"
	params := map[string]interface{}{
		"drive_id": "83410",
		"file_id":  fileId,
	}
	body, _ := json.Marshal(params)
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body))
	req.Header.Set("authorization", Authorization)
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp.Status)
	return err
}