package common

// 消息类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	UserStatusNotifyMesType = "UserStatusNotifyMes"
)

// 消息结构体：消息类型，消息内容
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 登录消息结构体
type LoginMes struct {
	Id   int    `json:"userId"`
	Name string `json:"userName"`
	Pwd  string `json:"userPwd"`
}

// 登录返回消息结构体：包含状态码(code)和错误信息(string)
type LoginResMes struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"` // 返回当前在线用户ID
	Error string `json:"error"`
}

// 用户注册消息结构体
type RegisterMes struct {
	User User `json:"user"`
}

// 用户注册返回消息结构体：包含状态码(code)和错误信息(string)
//这个RegisterResMes 和  LoginResMes ，其实是可以合并的
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// 用户状态通知消息结构体
type UserStatusNotifyMes struct {
	UserId int `json:"userId"`
	Status int `json:"userStatus"`
}
