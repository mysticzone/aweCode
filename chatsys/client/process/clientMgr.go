package process

import (
	"chatsys/common"
	"fmt"
)

// 用于维护线上用户列表
var onlineUserMap map[int]*common.User = make(map[int]*common.User, 10)
var UserId int

// 输出当前在线用户信息列表
func outputUserOnline() {
	fmt.Println("========在线用户列表如下======== ")
	for id, _ := range onlineUserMap {
		// 在线列表，排除自己
		if id == UserId {
			continue
		}
		fmt.Println("在线用户ID: \t", id)
	}
}

// 更新客户端 当前在线用户列表 (服务端会返回新登录用户信息包，更新返回这个用户的状态即可)
// 1. 新登录的用户，添加到 onelineUserMap
// 2. 登录用户下线，更新
func updateUserStatus(userStatusNotifyMes common.UserStatusNotifyMes) {

	// 上线用户是否已经在线上列表
	user, ok := onlineUserMap[userStatusNotifyMes.UserId]
	if !ok {
		// 上线用户不在用户线上列表，将其添加
		user = &common.User{}
		user.Id = userStatusNotifyMes.UserId
	}

	// 更新状态 1.在线，2. 离线
	user.Status = userStatusNotifyMes.Status

	// 更新onelineUserMap
	onlineUserMap[user.Id] = user

	// 输出更新后的在线用户在线列表
	outputUserOnline()
}
