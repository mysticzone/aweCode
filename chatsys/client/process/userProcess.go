package process

import (
	"chatsys/client/utils"
	"chatsys/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const address = "localhost:8989"

// 处理用户相关结构体
type UserProcess struct {
	// Conn net.Conn
	// Buf [8192]byte
}

// New user register
func (up *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	// Connect redis
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Dialing failed. err=", err)
		return
	}

	//
	var msg common.Message
	// 指定消息的种类
	msg.Type = common.RegisterMesType

	var registerMes common.RegisterMes
	registerMes.User.Id = userId
	registerMes.User.Name = userName
	registerMes.User.Pwd = userPwd

	// 序列化用于传输
	data, err := json.Marshal(registerMes)
	if err != nil {
		return
	}

	// 注册内容，放到msg中的Data字段
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	// 注册内容的长度
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 先发送内容的长度，便于服务端检测收到的数据是否完整
	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("Write data failed. err=", err)
		return
	}

	// 发送消息给服务器
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	// Transfer 读服务器返回的消息包
	tf := &utils.Transfer{
		Conn: conn,
	}
	msg, err = tf.ClientReadPackage()
	if err != nil {
		fmt.Println("ReadPackage failed. err=", err)
		return
	}

	var registerResMes common.RegisterResMes
	// 返回的消息(Data)内容反序列化读取
	err = json.Unmarshal([]byte(msg.Data), &registerResMes)
	if err != nil {
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("Register success....")
		os.Exit(0)
	}
	if registerResMes.Code == 100 {
		fmt.Println("Regsiter userId exist....")
		os.Exit(0)
	}
	return
}

// Old user login
func (up *UserProcess) Login(userId int, userPwd string) (err error) {

	// Connect redis
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Dialing error. err=", err.Error())
		return
	}

	//
	var msg common.Message
	// 指定消息的种类
	msg.Type = common.LoginMesType

	var loginMes common.LoginMes
	loginMes.Id = userId
	loginMes.Pwd = userPwd

	// 序列化用于传输
	data, err := json.Marshal(loginMes)
	if err != nil {
		return
	}

	// 登录内容，放到msg中的Data字段
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	// 登录内容的长度
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 先发送内容的长度，便于服务端检测收到的数据是否完整
	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("Write data failed. err=", err)
		return
	}

	// 发送消息给服务器
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	// Transfer 读服务器返回的消息包
	tf := &utils.Transfer{
		Conn: conn,
	}

	msg, err = tf.ClientReadPackage()
	if err != nil {
		fmt.Println("ReadPackage failed. err=", err)
		return
	}

	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(msg.Data), &loginResMes)
	if err != nil {
		return
	}
	if loginResMes.Code == 500 {
		fmt.Println("User not exist. Please register first....")
		os.Exit(0)
	}
	if loginResMes.Code == 200 {
		// 登录成功后，获取当前在线用户列表
		fmt.Println("========用户在线列表如下========")
		for _, v := range loginResMes.User {
			// 在线用户列表，排除自己
			if v == userId {
				continue
			}
			fmt.Println("在线用户ID:\t", v)

			// 客户端记录在线用户，使用切片保存用户列表信息
			user := &common.User{
				Id: v,
			}
			onlineUserMap[user.Id] = user
		}
		fmt.Println("\n\n")

		// 登录成功，需要实时读取服务端发送的消息并处理，专门启动一个 goroutine
		go ProcessServerMessage(conn)

		// 如果登录成功，显示主菜单（循环显示）
		for {
			ShowMenu(conn)
		}
	}
	return
}
