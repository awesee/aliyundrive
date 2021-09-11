package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BatchResultV2 struct {
	Responses []BatchResponsesV2 `json:"responses"`
}

type BatchResponsesV2 struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

func BatchV2(resource string, requests []interface{}) (result BatchResultV2, err error) {
	if len(requests) == 0 {
		return
	}
	uri := "https://api.aliyundrive.com/v2/batch"
	contentType := "application/json;charset=UTF-8"
	params := map[string]interface{}{
		"resource": resource,
		"requests": requests,
	}
	body, _ := json.Marshal(params)
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body))
	req.Header.Set("authorization", Authorization)
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	return
}
