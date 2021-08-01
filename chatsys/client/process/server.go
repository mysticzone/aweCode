package process

import (
	"chatsys/client/utils"
	"chatsys/common"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ProcessServerMessage(conn net.Conn) {
	// 实例化 Transfer 不停的读取信息...
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		msg, err := tf.ClientReadPackage()
		if err != nil {
			fmt.Println("Read failed. err=", err)
			fmt.Println("msg=", msg, "err=", err)
			// os.Exit(0)
		}

		var userStatus common.UserStatusNotifyMes

		err = json.Unmarshal([]byte(msg.Data), &userStatus)
		if err != nil {
			fmt.Println("Unmarshal failed. err=", err)
			return
		}

		switch msg.Type {
		case common.UserStatusNotifyMesType:
			// updateUserStatus(userStatus)
		}
	}
}

// 用户登录成功，显示的菜单
func ShowMenu(conn net.Conn) {
	fmt.Println("1. 显示在线用户列表")
	fmt.Println("2. 发送信息")
	fmt.Println("3. 信息列表")
	fmt.Println("4. 退出系统")

	//
	var key int
	func(key *int) {
		_, err := fmt.Scanf("%d\n", key)
		if err != nil {
			fmt.Println("Scanf failed. err=", err)
			return
		}
	}(&key)
	fmt.Println("Scanf key: ", key)

	switch key {
	case 1:

		outputUserOnline()
	case 2:
	//enterTalk()
	case 3:
	// listUnReadMsg()
	case 4:
		os.Exit(0)
	}
}
