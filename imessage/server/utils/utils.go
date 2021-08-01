package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"imessage/common/message"
	"net"
)

// 将这些方法关联到结构体
type Transfer struct {
	Conn net.Conn
	// 传输时使用的缓存
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据....")
	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		//err = errors.New("conn.Read Header fail")
		return
	}
	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	// 根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail. err: ", err)
		//err = errors.New("conn.Read Body fail")
		return
	}

	// 把 pkgLen 反序列化成 message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail. err: ", err)
		//err = errors.New("json.Unmarshal fail")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度
	// 获取data的长度，先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))

	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	_, err = this.Conn.Write(this.Buf[:])
	//_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(byte) err: ", err)
		return
	}
	// 发送 data 本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}
	return
}
