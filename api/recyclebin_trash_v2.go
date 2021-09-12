package api

import (
	"bytes"
	"encoding/json"
	"errors"
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
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}
	return err
}

func RecycleBinTrashBatchV2(driveId string, fileId []string) (BatchResultV2, error) {
	requests := make([]interface{}, 0)
	for _, id := range fileId {
		requests = append(requests, map[string]interface{}{
			"body": map[string]string{
				"drive_id": driveId,
				"file_id":  id,
			},
			"headers": map[string]string{
				"Content-Type": "application/json",
			},
			"id":     id,
			"method": "POST",
			"url":    "/recyclebin/trash",
		})
	}
	return BatchV2("file", requests)
}
