package processes

import (
	"encoding/json"
	"fmt"
	"imessage/common/message"
	"imessage/server/utils"
	"net"
)

type UserProcess struct {
Conn net.Conn
}

// 编写 ServiceProcessLogin 函数。专门处理登录相关信息
func (this *UserProcess) ServiceProcessLogin(mes *message.Message) (err error) {
	// 从 mes 中取出 Data, 反序列化成 LoginMes
	var LoginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &LoginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error. err: ", err)
		return
	}
	// 声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 声明一个 LoginResMes, 并完成赋值
	var loginResMes message.LoginResMes

	// 如果用户的 UserID=100 && UserPwd=123456, 认为合法
	if LoginMes.UserId == "100" && LoginMes.UserPwd == "123456" {
		fmt.Println("认证成功")
		loginResMes.Code = 200
		//LoginResMes.Error = nil
	} else {
		// 不合法
		loginResMes.Code = 500 // 500 表示用户不存在
		loginResMes.Error = "用户不存在，请注册再使用"
	}

	// 将 LoginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal error. err: ", err)
		return
	}

	// 将 data 赋值给 resMes 中的 Data
	resMes.Data = string(data)

	// 将 resMes 序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal error. err: ", err)
		return
	}

	// 发送 data，封装到 writePkg 函数中
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg([]byte(data))
	return
}
