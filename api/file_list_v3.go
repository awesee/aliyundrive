package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type FileListResultV3 struct {
	Items      []FileListItemV3 `json:"items"`
	NextMarker string           `json:"next_marker"`
}
type FileListItemV3 struct {
	DriveID         string    `json:"drive_id"`
	DomainID        string    `json:"domain_id"`
	FileID          string    `json:"file_id"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Hidden          bool      `json:"hidden"`
	Starred         bool      `json:"starred"`
	Status          string    `json:"status"`
	UserMeta        string    `json:"user_meta,omitempty"`
	ParentFileID    string    `json:"parent_file_id"`
	EncryptMode     string    `json:"encrypt_mode"`
	ContentType     string    `json:"content_type,omitempty"`
	FileExtension   string    `json:"file_extension,omitempty"`
	MimeType        string    `json:"mime_type,omitempty"`
	MimeExtension   string    `json:"mime_extension,omitempty"`
	Size            int       `json:"size,omitempty"`
	UploadID        string    `json:"upload_id,omitempty"`
	Crc64Hash       string    `json:"crc64_hash,omitempty"`
	ContentHash     string    `json:"content_hash,omitempty"`
	ContentHashName string    `json:"content_hash_name,omitempty"`
	DownloadURL     string    `json:"download_url,omitempty"`
	URL             string    `json:"url,omitempty"`
	Category        string    `json:"category,omitempty"`
	PunishFlag      int       `json:"punish_flag,omitempty"`
}

func FileListV3(parentFileId string) (result FileListResultV3, err error) {
	uri := "https://api.aliyundrive.com/adrive/v3/file/list"
	contentType := "application/json;charset=UTF-8"
	params := map[string]interface{}{
		"drive_id":                "83410",
		"parent_file_id":          parentFileId,
		"limit":                   100,
		"all":                     false,
		"url_expire_sec":          1600,
		"image_thumbnail_process": "image/resize,w_400/format,jpeg",
		"image_url_process":       "image/resize,w_1920/format,jpeg",
		"video_thumbnail_process": "video/snapshot,t_0,f_jpg,ar_auto,w_300",
		"fields":                  "*",
		"order_by":                "name",
		"order_direction":         "ASC",
	}
	body, _ := json.Marshal(params)
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body))
	req.Header.Set("authorization", Authorization)
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	//dump, err := httputil.DumpResponse(resp, true)
	//fmt.Printf("%s\n%s\n", dump, err)
	if err != nil {
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	return
}
