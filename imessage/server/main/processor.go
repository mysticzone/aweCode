package main

import (
	"fmt"
	"imessage/common/message"
	processes "imessage/server/process"
	"imessage/server/utils"
	"io"
	"net"
)
//
type Processor struct {
Conn net.Conn
}

// 编写 ServerProcessMes 函数。根据客户端发送消息的种类决定调用哪一个函数来处理
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
		case message.LoginMesType:
			// 处理登录的逻辑
			up := &processes.UserProcess{
			Conn: this.Conn,
			}
			err = up.ServiceProcessLogin(mes)
			//err = ServiceProcessLogin(this.Conn, mes)
		case message.RegisterMesType:
			// 处理注册逻辑
		default:
			fmt.Println("消息类型不存在，无法处理....")
		}
	return
}

func (this *Processor) process2() (err error){
	// 循环客户端读取信息
	for {
		// 使用readpkg封装读取消息的操作，返回Message,Err
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit; server exit")
				return
			} else {
				fmt.Println("readPkg fail. err: ", err)
				return
			}

		}
		fmt.Println(mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("ServerProcessMes(conn, &mes) error. err: ", err)
			return
		}
	}
}

