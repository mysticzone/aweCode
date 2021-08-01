package main

import (
	"chatsys/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {

	fmt.Println("-----------------欢迎登陆多人聊天系统---------------")
	fmt.Println("\t\t\t 1 登录聊天系统")
	fmt.Println("\t\t\t 2 注册用户")
	fmt.Println("\t\t\t 3 退出系统")
	fmt.Println("请选择(1-3):")
	fmt.Println("-----------------------------")

	// 接收用户输入的选择
	var key int
	func(key *int) {
		n, err := fmt.Scanf("%d\n", key)
		if err != nil {
			fmt.Println("Input key failed. err=", err)
		}
		fmt.Println("Input length. n=", n)
	}(&key)

	switch key {
	case 1:
		// 用户登录
		fmt.Println("========用户登录========")
		fmt.Println("请输入用户id:")
		_, _ = fmt.Scanf("%d\n", &userId)
		process.UserId = userId
		fmt.Println("请输入用户密码:")
		_, _ = fmt.Scanf("%s\n", &userPwd)
		fmt.Println("\n")

		// 通过 UserProcess 调用 Login 方法
		up := process.UserProcess{}
		err := up.Login(userId, userPwd)
		if err != nil {
			fmt.Println("用户登录失败。err=", err)
			return
		}
		fmt.Println("用户登录成功：", userId)
	case 2:
		// 用户注册
		fmt.Println("注册用户") //注册处理
		fmt.Println("请输入新用户id:")
		_, _ = fmt.Scanf("%d\n", &userId)

		fmt.Println("请输入新用户密码:")
		_, _ = fmt.Scanf("%s\n", &userPwd)

		fmt.Println("请输入新用户名字:")
		_, _ = fmt.Scanf("%s\n", &userName)

		fmt.Println("\n")

		// 通过 UserProcess 调用 Register 方法
		up := process.UserProcess{}
		err := up.Register(userId, userPwd, userName)
		if err != nil {
			fmt.Println("用户注册失败。err=", err)
			return
		}
		fmt.Println("用户注册成功：", userId)
		// 注册后退出系统
		os.Exit(0)
	case 3:
		fmt.Println("退出系统....")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入....")
	}
}
