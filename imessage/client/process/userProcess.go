package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"imessage/client/utils"
	"imessage/common/message"
	"net"
)

type UserProcess struct {
	// 暂时不需要字段
}

// 关联用户登录的方法
// 登录函数
func (this *UserProcess) Login(userId string, userPwd string) (err error) {
	// 开始定义协议
	//fmt.Printf("userId=%v\tuserPwd=%v\n", userId, userPwd)
	//return nil
	conn, err := net.Dial("tcp", "0.0.0.0:8989")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	// 通过conn发送消息给server
	var mes message.Message
	mes.Type = message.LoginMesType

	// 创建LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	// 序列化后的值赋给mes.Data
	mes.Data = string(data)

	// mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	// 就是要发送的消息
	//num, err := conn.Write(data)
	//if err != nil {
	//	fmt.Println("conn.Write err: ", err)
	//	return
	//}
	//fmt.Println("Send byte: ", num)
	// 获取data的长度，先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}
	fmt.Println("客户端发送消息长度：", len(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err: ", err)
		return
	}

	//// 休眠20秒
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20秒")
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg error. err: ", err)
		return
	}
	// 将 mes 的Data 反序列化成 LoginResMes
	var LoginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &LoginResMes)
	//if err != nil {
	//	fmt.Println("json.Unmarshal error. err: ", err)
	//	return
	//}
	if LoginResMes.Code == 200 {
		fmt.Println("登录成功1")
	} else if LoginResMes.Code == 500 {
		fmt.Println("登录失败： ", LoginResMes.Error)
	}
	return nil
}
