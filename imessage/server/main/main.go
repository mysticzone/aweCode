package main

import (
	"fmt"
	"net"
)



//// 编写 ServiceProcessLogin 函数。专门处理登录相关信息
//func ServiceProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//	// 从 mes 中取出 Data, 反序列化成 LoginMes
//	var LoginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &LoginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal error. err: ", err)
//		return
//	}
//	// 声明一个 resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	// 声明一个 LoginResMes, 并完成赋值
//	var loginResMes message.LoginResMes
//
//	// 如果用户的 UserID=100 && UserPwd=123456, 认为合法
//	if LoginMes.UserId == "100" && LoginMes.UserPwd == "123456" {
//		fmt.Println("认证成功")
//		loginResMes.Code = 200
//		//LoginResMes.Error = nil
//	} else {
//		// 不合法
//		loginResMes.Code = 500 // 500 表示用户不存在
//		loginResMes.Error = "用户不存在，请注册再使用"
//	}
//
//	// 将 LoginResMes 序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal error. err: ", err)
//		return
//	}
//
//	// 将 data 赋值给 resMes 中的 Data
//	resMes.Data = string(data)
//
//	// 将 resMes 序列化，准备发送
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal error. err: ", err)
//		return
//	}
//
//	// 发送 data，封装到 writePkg 函数中
//	err = client.WritePkg(conn, []byte(data))
//	return
//}

//// 编写 ServerProcessMes 函数。根据客户端发送消息的种类决定调用哪一个函数来处理
//func ServerProcessMes(conn net.Conn,  mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		// 处理登录的逻辑
//		err = ServiceProcessLogin(conn, mes)
//	case message.RegisterMesType:
//		// 处理注册逻辑
//	default:
//		fmt.Println("消息类型不存在，无法处理....")
//	}
//	return
//}

// 处理客户端的通讯
func process(conn net.Conn) {
	// 延迟关闭
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	//// 循环客户端读取信息
	//for {
	//	// 使用readpkg封装读取消息的操作，返回Message,Err
	//	mes, err := client.ReadPkg(conn)
	//	if err != nil {
	//		if err == io.EOF {
	//			fmt.Println("client exit; server exit")
	//			return
	//		} else {
	//			fmt.Println("readPkg fail. err: ", err)
	//			return
	//		}
	//
	//	}
	//	fmt.Println(mes)
	//	err = ServerProcessMes(conn, &mes)
	//	if err != nil {
	//		fmt.Println("ServerProcessMes(conn, &mes) error. err: ", err)
	//		return
	//	}
	//}
	// 总控 创建
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端链接错误=err: ", err)
		return
	}
}

func main() {
	fmt.Println("服务在监听8989端口....")
	listen, err := net.Listen("tcp", "0.0.0.0:8989")
	if err != nil {
		fmt.Println("监听无效。。。。")
		panic(err)
	}
	// 一旦监听成功，等待客户端来连接服务器
	for {
		fmt.Println("等待连接....")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error: ", err)
		}
		// 一旦连接成功，和客户端保持通讯
		go process(accept)
	}
}
