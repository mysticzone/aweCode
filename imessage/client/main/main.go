package main

import (
	"fmt"
	"imessage/client"
	"os"
)

// 表示用户ID
var userId string

// 表示用户密码
var userPwd string

func main() {
	// 接收用户选择
	var key int
	// 判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("=====================欢迎登陆多人聊天系统=====================")
		fmt.Println("\t\t1 登陆聊天系统")
		fmt.Println("\t\t2 注册用户")
		fmt.Println("\t\t3 退出系统")
		fmt.Println("请选择(1-3): ")

		_, _ = fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false

		case 2:
			fmt.Println("注册用户")
			loop = false

		case 3:
			fmt.Println("退出系统")
			loop = false
			os.Exit(1)
		default:
			fmt.Println("无效输入")
		}

	}

	if key == 1 {
		// 用户登录
		fmt.Println("请输入用户ID：")
		_, _ = fmt.Scanf("%s\n", &userId)
		fmt.Println("请输入密码：")
		_, _ = fmt.Scanf("%s\n", &userPwd)
		_ = client.Login(userId, userPwd)
		//if err != nil {
		//	fmt.Println("登录失败....")
		//	panic(err)
		//} else {
		//	fmt.Println("登录成功2")
		//}
	} else if key == 2 {
		fmt.Println("用户注册逻辑处理....")
	}
}
