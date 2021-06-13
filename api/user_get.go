package api

import "github.com/openset/aliyundrive/app"

func fetchUser() (user User) {
	app.Client.Request("/user/get").JSONUnmarshal(&user)
	return
}

type User struct {
	DomainID                    string   `json:"domain_id"`
	UserID                      string   `json:"user_id"`
	Avatar                      string   `json:"avatar"`
	CreatedAt                   int64    `json:"created_at"`
	UpdatedAt                   int64    `json:"updated_at"`
	Email                       string   `json:"email"`
	NickName                    string   `json:"nick_name"`
	Phone                       string   `json:"phone"`
	Role                        string   `json:"role"`
	Status                      string   `json:"status"`
	UserName                    string   `json:"user_name"`
	Description                 string   `json:"description"`
	DefaultDriveID              string   `json:"default_drive_id"`
	UserData                    UserData `json:"user_data"`
	DenyChangePasswordBySelf    bool     `json:"deny_change_password_by_self"`
	NeedChangePasswordNextLogin bool     `json:"need_change_password_next_login"`
}

type BackUpConfig interface{}

type UserData struct {
	BackUpConfig BackUpConfig `json:"back_up_config"`
	Share        string       `json:"share"`
}
