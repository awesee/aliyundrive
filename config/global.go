package config

var Global = GlobalConfig{
	PdsEndpoint:    "https://api.aliyundrive.com",
	MemberEndpoint: "https://member.aliyundrive.com",
	SvEndpoint:     "https://websv.aliyundrive.com",
	BaseURL:        "/drive",
	AuthEndpoint:   "https://auth.aliyundrive.com",
	ClientID:       "25dzX3vbYqktVxyX",
	RedirectURI:    "https://www.aliyundrive.com/sign/callback",
	EnableShare:    true,
	ShareFolder:    false,
	ShareVerify:    true,
	SharePromotion: false,
	MaxShare:       100,
	FileDownload:   true,
	FolderDownload: true,
	MultiDownload:  true,
	CdnHost:        "https://g.alicdn.com/aliyun-drive-fe/aliyun-drive/2.0.17-web/web",
	AliyunDriveEnv: "prod",
}

type GlobalConfig struct {
	PdsEndpoint    string `json:"pds_endpoint"`
	MemberEndpoint string `json:"member_endpoint"`
	SvEndpoint     string `json:"sv_endpoint"`
	BaseURL        string `json:"base_url"`
	AuthEndpoint   string `json:"auth_endpoint"`
	ClientID       string `json:"client_id"`
	RedirectURI    string `json:"redirect_uri"`
	EnableShare    bool   `json:"enable_share"`
	ShareFolder    bool   `json:"share_folder"`
	ShareVerify    bool   `json:"share_verify"`
	SharePromotion bool   `json:"share_promotion"`
	MaxShare       int    `json:"max_share"`
	FileDownload   bool   `json:"file_download"`
	FolderDownload bool   `json:"folder_download"`
	MultiDownload  bool   `json:"multi_download"`
	CdnHost        string `json:"cdnHost"`
	AliyunDriveEnv string `json:"aliyun_drive_env"`
}
