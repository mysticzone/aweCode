package common

// 常量表示用户状态
const (
	UserStatusOnline = 1
UserStatusOffline = iota
)

// 用户信息结构体
type User struct {
	Id int `json:"userId"`
	Name string `json:"userName"`
	Pwd string `json:"userPwd"`
	Status int `json:"status"`
	//后续可扩展的字段
	// Nick      string `json:"nick"`
	// Sex       string `json:"sex"`
	// Header    string `json:"header"`
	// LastLogin string `json:"last_login"`
}