package main

//
//import (
//	"chatsys/common"
//	"chatsys/server/process"
//	"chatsys/server/utils"
//	"errors"
//	"fmt"
//	"net"
//)
//
//// 有点 client 控制器
//type ClientProcessor struct {
//	conn net.Conn
//	buf  [8192]byte
//}
//
//// 根据消息类型，做不同的处理
//func (cp *ClientProcessor) ServerProcessMsg(msg *common.Message) (err error) {
//
//	switch msg.Type {
//	case common.LoginMesType:
//		up := &process.UserProcessor{
//			Conn: cp.conn,
//		}
//		err = up.ServerProcessLogin(msg)
//	case common.RegisterMesType:
//		up := &process.UserProcessor{
//			Conn: cp.conn,
//		}
//		err = up.ServerProcessRegister(msg)
//	default:
//		err = errors.New("Message type invalid.")
//	}
//	return
//}
//
////
//func (cp *ClientProcessor) Process() (err error) {
//	for {
//		var msg common.Message
//
//		// 通过 Transfer 实例 返回消息包给客户端
//		tf := &utils.Transfer{
//			Conn: cp.conn,
//		}
//
//		// 从客户端读取消息
//		msg, err := tf.ServerReadPkg()
//		if err != nil {
//			fmt.Println("Server read package failed. err=", err)
//			return
//		}
//
//		// 封装好的消息交给 ServerProcessMsg 处理
//		err = cp.ServerProcessMsg(&msg)
//		if err != nil {
//			fmt.Println("Process msg failed. err=", err)
//			return
//		} else {
//			//如果我们处理没有任何错误，作为测试，就先退出这个goroutine
//			//即退出这个for循环，否则会继续执行 readPackage
//			//但是，客户端没有发消息，这就服务器会报 server read package err= read header failed
//			// 如果希望服务器端一直读取客户端信息，下面的return去掉即可
//			return err
//		}
//
//	}
//}
