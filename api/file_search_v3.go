package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type FileSearchResultV3 struct {
	Items      []FileListItemV3 `json:"items"`
	NextMarker string           `json:"next_marker"`
}

type FileSearchItem struct {
	DriveID            string    `json:"drive_id"`
	DomainID           string    `json:"domain_id"`
	FileID             string    `json:"file_id"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	ContentType        string    `json:"content_type"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	FileExtension      string    `json:"file_extension"`
	MimeType           string    `json:"mime_type"`
	MimeExtension      string    `json:"mime_extension"`
	Hidden             bool      `json:"hidden"`
	Size               int       `json:"size"`
	Starred            bool      `json:"starred"`
	Status             string    `json:"status"`
	UserMeta           string    `json:"user_meta"`
	Labels             []string  `json:"labels"`
	UploadID           string    `json:"upload_id"`
	ParentFileID       string    `json:"parent_file_id"`
	Crc64Hash          string    `json:"crc64_hash"`
	ContentHash        string    `json:"content_hash"`
	ContentHashName    string    `json:"content_hash_name"`
	DownloadURL        string    `json:"download_url"`
	URL                string    `json:"url"`
	Thumbnail          string    `json:"thumbnail"`
	Category           string    `json:"category"`
	EncryptMode        string    `json:"encrypt_mode"`
	VideoMediaMetadata struct {
		Time      time.Time `json:"time"`
		Width     int       `json:"width"`
		Height    int       `json:"height"`
		ImageTags []struct {
			Confidence float64 `json:"confidence"`
			Name       string  `json:"name"`
			TagLevel   int     `json:"tag_level"`
			ParentName string  `json:"parent_name,omitempty"`
		} `json:"image_tags"`
		VideoMediaVideoStream []interface{} `json:"video_media_video_stream"`
		VideoMediaAudioStream []interface{} `json:"video_media_audio_stream"`
		Duration              string        `json:"duration"`
	} `json:"video_media_metadata,omitempty"`
	VideoPreviewMetadata struct {
		Bitrate     string `json:"bitrate"`
		Duration    string `json:"duration"`
		AudioFormat string `json:"audio_format"`
		VideoFormat string `json:"video_format"`
		FrameRate   string `json:"frame_rate"`
		Height      int    `json:"height"`
		Width       int    `json:"width"`
	} `json:"video_preview_metadata,omitempty"`
	PunishFlag         int `json:"punish_flag"`
	ImageMediaMetadata struct {
		Time        time.Time `json:"time"`
		Width       int       `json:"width"`
		Height      int       `json:"height"`
		Location    string    `json:"location"`
		Country     string    `json:"country"`
		Province    string    `json:"province"`
		City        string    `json:"city"`
		District    string    `json:"district"`
		Township    string    `json:"township"`
		AddressLine string    `json:"address_line"`
		ImageTags   []struct {
			Confidence float64 `json:"confidence"`
			Name       string  `json:"name"`
			TagLevel   int     `json:"tag_level"`
			ParentName string  `json:"parent_name,omitempty"`
		} `json:"image_tags"`
		Exif         string `json:"exif"`
		ImageQuality struct {
			OverallScore float64 `json:"overall_score"`
		} `json:"image_quality"`
		CroppingSuggestion []struct {
			AspectRatio      string  `json:"aspect_ratio"`
			Score            float64 `json:"score"`
			CroppingBoundary struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Top    int `json:"top"`
				Left   int `json:"left"`
			} `json:"cropping_boundary"`
		} `json:"cropping_suggestion"`
	} `json:"image_media_metadata,omitempty"`
}

func FileSearchV3() (result FileSearchResultV3, err error) {
	uri := "https://api.aliyundrive.com/adrive/v3/file/search"
	contentType := "application/json;charset=UTF-8"
	params := map[string]interface{}{
		"drive_id":                "9680003",
		"query":                   "type = \"file\"",
		"image_thumbnail_process": "image/resize,w_400/format,jpeg",
		"image_url_process":       "image/resize,w_1920/format,jpeg",
		"video_thumbnail_process": "video/snapshot,t_0,f_jpg,ar_auto,w_1000",
		"limit":                   100,
		"order_by":                "image_time DESC,created_at DESC",
	}
label:
	body, _ := json.Marshal(params)
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body))
	req.Header.Set("authorization", Authorization)
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	var items FileSearchResultV3
	err = json.Unmarshal(body, &items)
	if err == nil {
		result.Items = append(result.Items, items.Items...)
	}
	if items.NextMarker != "" {
		params["marker"] = items.NextMarker
		goto label
	}
	return
}
