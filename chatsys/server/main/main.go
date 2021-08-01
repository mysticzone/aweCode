package main

import (
	"chatsys/common"
	"chatsys/server/model"
	"chatsys/server/process"
	"chatsys/server/utils"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

// clientProcessor
// 有点 client 控制器
type ClientProcessor struct {
	conn net.Conn
	buf  [8192]byte
}

// 根据消息类型，做不同的处理
func (cp *ClientProcessor) ServerProcessMsg(msg *common.Message) (err error) {

	switch msg.Type {
	case common.LoginMesType:
		up := &process.UserProcessor{
			Conn: cp.conn,
		}
		err = up.ServerProcessLogin(msg)
	case common.RegisterMesType:
		up := &process.UserProcessor{
			Conn: cp.conn,
		}
		err = up.ServerProcessRegister(msg)
	default:
		err = errors.New("Message type invalid.")
	}
	return
}

//
func (cp *ClientProcessor) Process() (err error) {
	for {
		var msg common.Message

		// 通过 Transfer 实例 返回消息包给客户端
		tf := &utils.Transfer{
			Conn: cp.conn,
		}

		// 从客户端读取消息
		msg, err := tf.ServerReadPkg()
		if err != nil {
			fmt.Println("Server read package failed. err=", err)
			return err
		}

		// 封装好的消息交给 ServerProcessMsg 处理
		err = cp.ServerProcessMsg(&msg)
		if err != nil {
			fmt.Println("Process msg failed. err=", err)
			return err
		} else {
			//如果我们处理没有任何错误，作为测试，就先退出这个goroutine
			//即退出这个for循环，否则会继续执行 readPackage
			//但是，客户端没有发消息，这就服务器会报 server read package err= read header failed
			// 如果希望服务器端一直读取客户端信息，下面的return去掉即可
			return err
		}

	}
}

// Redis
var pool *redis.Pool

func initRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     idleConn,
		MaxActive:   maxConn,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	return
}

// 从连接池中取一个链接
func GetConn() redis.Conn {
	return pool.Get()
}

// 关闭一个从连接池中取到的链接
func PutConn(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		return
	}
}

func initUserDao() {
	// redis 连接池 pool 是一个全局变量，在此可直接使用
	model.MyUserDao = model.NewUserMgr(pool)
}

//func process(conn net.Conn) {
//	cp := &ClientProcessor{}
//}

func execute(conn net.Conn) {

	// 创建client，有点控制器
	cp := &ClientProcessor{
		conn: conn,
	}

	// 客户端进行处理，将不同的请求分配给不同的处理器
	err := cp.Process()
	if err != nil {
		fmt.Println("Client process failed. err=", err)
		return
	}
}

func main() {

	// 服务端连接Redis
	initRedis("192.168.146.54:6379", 16, 1024, time.Second*300)

	// 服务端初始化UserDao, 用于操作Redis
	initUserDao()

	var addr = "0.0.0.0:8989"
	fmt.Println("服务器端8989端口监听,等待客户端连接....")
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Listen failed. err=", err)
		return
	}

	for {
		// 等待连接
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept failed. err=", err)
			return
		}

		go execute(conn)
	}
}
