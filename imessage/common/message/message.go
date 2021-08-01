package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType     = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   string `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 状态码，500表示尚未注册，200表示登录成功
	Error string `json:"error"` // 返回错误信息
}
