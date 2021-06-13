package app

import (
	"encoding/json"
	"github.com/openset/aliyundrive/config"
	"time"
)

var Client = client{
	config: clientConfig{
		authTokenKey:   "token",
		shareTokenKey:  "shareToken",
		pdsEndpoint:    config.Global.PdsEndpoint,
		svEndpoint:     config.Global.SvEndpoint,
		memberEndpoint: config.Global.MemberEndpoint,
		version:        "v2",
	},
}

type client struct {
	config         clientConfig
	authTokenInfo  authTokenInfo
	shareTokenInfo shareTokenInfo
}

type clientConfig struct {
	authTokenKey   string
	shareTokenKey  string
	pdsEndpoint    string
	svEndpoint     string
	memberEndpoint string
	version        string
}

type authTokenInfo struct {
	AccessToken        string        `json:"access_token"`
	RefreshToken       string        `json:"refresh_token"`
	ExpiresIn          int           `json:"expires_in"`
	TokenType          string        `json:"token_type"`
	UserID             string        `json:"user_id"`
	UserName           string        `json:"user_name"`
	Avatar             string        `json:"avatar"`
	NickName           string        `json:"nick_name"`
	DefaultDriveID     string        `json:"default_drive_id"`
	DefaultSboxDriveID string        `json:"default_sbox_drive_id"`
	Role               string        `json:"role"`
	Status             string        `json:"status"`
	ExpireTime         time.Time     `json:"expire_time"`
	State              string        `json:"state"`
	ExistLink          []interface{} `json:"exist_link"`
	NeedLink           bool          `json:"need_link"`
	PinSetup           bool          `json:"pin_setup"`
	IsFirstLogin       bool          `json:"is_first_login"`
	NeedRpVerify       bool          `json:"need_rp_verify"`
}

type shareTokenInfo struct {
	ShareToken string    `json:"share_token"`
	ExpireTime time.Time `json:"expire_time"`
	ExpiresIn  int       `json:"expires_in"`
}

func (c *client) getAuthTokenInfo() authTokenInfo {
	return c.authTokenInfo
}

func (c *client) isLogin() {

}

func (c *client) getHost(hostType string) string {
	switch hostType {
	case "member":
		return c.config.memberEndpoint
	default:
		return c.config.pdsEndpoint
	}
}

func (c *client) onRefreshTokenError() {

}

func (c *client) onRequestError() {

}

func (c *client) logOut() {

}

func (c *client) getShareToken() {

}

func (c *client) getClient() *client {
	return c
}

type request struct {
	url    string
	method string
}

func (r request) Method(method string) request {
	r.method = method
	return r
}

func (r request) JSONUnmarshal(v interface{}) error {
	var data []byte
	return json.Unmarshal(data, v)
}

func (c *client) Request(url string) request {
	return request{url: url}
}

func (c *client) refreshAuthToken() {

}
