package process

import (
	"chatsys/common"
	"chatsys/server/model"
	"chatsys/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 用户控制器，用于处理与用户相关。每一个登录用户对应一个 UserProcessor 实例
type UserProcessor struct {
	Conn   net.Conn
	userId int
	Buf    [8192]byte
}

func (up *UserProcessor) NotifyOtherUserOnline(userId int) {

	// 获取当前所有在线用户
	users := clientMgr.GetAllUsers()
	for id, client := range users {
		// 从在线用户中过滤掉自己
		if id == userId {
			continue
		}
		// 每个 client 对应一个 *UserProcessor 实例，每个 *UserProcessor 维护自己的 conn
		// 每取出一个当前在线用户，就通知这个在线用户，userId（上线了）
		client.NotifyUserOnline(userId)
	}
}

func (up *UserProcessor) NotifyUserOnline(userId int) {

	var respMsg common.Message
	// 类型为用户上线响应类型
	respMsg.Type = common.UserStatusNotifyMesType

	var userStatusNotifyMes common.UserStatusNotifyMes
	userStatusNotifyMes.UserId = userId
	userStatusNotifyMes.Status = common.UserStatusOnline

	data, err := json.Marshal(userStatusNotifyMes)
	if err != nil {
		fmt.Println("Marshal failed. err=", err)
		return
	}

	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("Marshal failed. err=", err)
		return
	}

	// 返回消息包给客户端
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.ServerWritePkg(data)
	if err != nil {
		fmt.Println("Send msg failed. err=", err)
		return
	}
	return
}

func (up *UserProcessor) ServerProcessLogin(msg *common.Message) (err error) {

	// 对msg.Data进行反序列化，得到用户名和密码
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(msg.Data), &loginMes)
	if err != nil {
		fmt.Println("unmarshal failed, err=", err)
		return
	}

	// 构建登录响应消息包
	var resMsg common.Message
	resMsg.Type = common.LoginResMesType

	var loginResMes common.LoginResMes

	// Redis中验证用户是否存在。用户ID和密码都正确，则登录成功；否则登录失败
	// 返回登录响应消息包
	var user *common.User
	user, err = model.MyUserDao.Login(loginMes.Id, loginMes.Pwd)
	if err != nil {
		// 登录失败，添加到失败错误信息，返回不同的Code码值
		fmt.Printf("Login failed. err=%v\n", err)
		loginResMes.Code = 500
	} else {
		// 登录成功，输出这个user信息
		fmt.Printf("%v login sucess.\n", user)
		loginResMes.Code = 200
	}

	// 登录成功后，将当前登录用户添加到在线用户列表中(切片)
	clientMgr.AddClient(loginMes.Id, up)
	up.userId = loginMes.Id

	// 当用户[LoginMes.Id]登录成功之后，给其他用户发送该用户上线提醒
	up.NotifyOtherUserOnline(loginMes.Id)

	// 登录成功后，需要将当前在线用户(ID)添加到返回包User的切片中 LoginResMes.User
	userMap := clientMgr.GetAllUsers()
	for userId, _ := range userMap {
		loginResMes.User = append(loginResMes.User, userId)
	}

	// 序列化登录返回包，用于传输
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("Marshal failed. err=", err)
		return
	}

	// 序列化后返回的消息 data 放入到 resMes 的 Data
	resMsg.Data = string(data)
	// 序列化消息结构体，用于传输
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("Marshal failed. err=", err)
		return
	}
	fmt.Println("返回pkg包=%s", string(data))

	// 通过 Transfer 实例返回消息包给客户端
	tf := utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.ServerWritePkg(data)
	if err != nil {
		fmt.Println("Send msg failed. err=", err)
		return
	}
	return
}

// 处理用户注册
func (up *UserProcessor) ServerProcessRegister(msg *common.Message) (err error) {

	var registerMes common.RegisterMes
	// msg.Data 反序列化，获取用户名和密码
	err = json.Unmarshal([]byte(msg.Data), &registerMes)
	if err != nil {
		fmt.Println("Unmarshal failed. err=", err)
		return
	}

	// 构建注册响应消息包
	var resMsg common.Message
	// 消息类型为注册响应类型
	resMsg.Type = common.RegisterResMesType

	var registerResMes common.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		fmt.Printf("Register userId exist. err=", err)
		registerResMes.Code = 100
	} else {
		fmt.Println("Register success....")
		registerResMes.Code = 200
	}

	// 返回消息序列化，用于传输
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("Marshal failed. err=", err)
		return
	}

	//将序列化后的返回消息data，放到resMsg的Data字段
	resMsg.Data = string(data)

	//再对resMsg序列化化便于传输
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("Marshal failed. err=", err)
		return
	}
	//因为要返回消息包给客户端，因此创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.ServerWritePkg(data)
	if err != nil {
		fmt.Println("Send msg failed. err=", err)
		return
	}
	return
}
