package utils

import (
	"chatsys/common"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

// 1. 向服务器端发送 2. 接收消息包读取
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (tf *Transfer) ClientReadPackage() (msg common.Message, err error) {
	// var buf [8192]byte
	n, err := tf.Conn.Read(tf.Buf[0:4])
	if n != 4 {
		err = errors.New("Read header failed.")
		fmt.Println("Sleep 30 seconds if reading failed, again....")
		time.Sleep(time.Second * 30)
		return
	}

	fmt.Println("Read package: ", tf.Buf[0:4][0:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	fmt.Println("Receive len: %d", pkgLen)

	n, err = tf.Conn.Read(tf.Buf[0:pkgLen])
	if n != int(pkgLen) {
		err = errors.New("Read body failed.")
		return
	}

	// fmt.Println("Receive data: %s\n", string[buf[0:pkgLen]])
	err = json.Unmarshal(tf.Buf[0:pkgLen], &msg)
	if err != nil {
		fmt.Println("Unmarshal failed. err: ", err)
	}
	return
}
