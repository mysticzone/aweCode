package utils

import (
	"chatsys/common"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//这个结构体，完成对客户端发送和接收消息包的读取
type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

// 读取客户端发送的消息包，封装到Message中
func (tf *Transfer) ServerReadPkg() (msg common.Message, err error) {
	n, err := tf.Conn.Read(tf.Buf[0:4])
	if n != 4 {
		err = errors.New("Read header failed.")
		return
	}
	fmt.Println("Read package:", tf.Buf[0:4], string(tf.Buf[0:4]))

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	fmt.Printf("Receive len:%d\n", pkgLen)

	// 读取客户端发送消息的长度
	n, err = tf.Conn.Read(tf.Buf[0:pkgLen])
	if n != int(pkgLen) {
		err = errors.New("Read body failed.")
		return
	}

	fmt.Printf("Receive data: %d\n", string(tf.Buf[0:pkgLen]))
	err = json.Unmarshal(tf.Buf[0:pkgLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}

// 发送消息包给客户端
func (tf *Transfer) ServerWritePkg(data []byte) (err error) {

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(tf.Buf[0:4], pkgLen)
	n, err := tf.Conn.Write(tf.Buf[0:4])
	if err != nil {
		fmt.Println("Write data length failed. err=", err)
		return
	}

	n, err = tf.Conn.Write(data)
	if err != nil {
		fmt.Println("Write data failed. err=", err)
		return
	}

	if n != int(pkgLen) {
		fmt.Println("Write data missing.")
		err = errors.New("Write data missing.")
		return
	}
	return
}
