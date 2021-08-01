package process

import "fmt"

// 在线用户管理
type ClientMgr struct {
	// 使用切片来保存所有登录用户信息
	onlineUsers map[int]*UserProcessor
}

var (
	clientMgr *ClientMgr
)

// 初始化
func init() {
	clientMgr = &ClientMgr{
		// 切片初始化
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

// 一个 ClientProcessor 实例，就对应一个登录用户
func (cm *ClientMgr) AddClient(userId int, client *UserProcessor) {
	cm.onlineUsers[userId] = client
}

// 根据 userId 获取用户
func (cm *ClientMgr) GetClient(userId int) (client *UserProcessor) {
	client, ok := cm.onlineUsers[userId]
	if !ok {
		_ = fmt.Errorf("userId=%d的用户，不存在", userId)
		return
	}
	return
}

// 获取所有的用户
func (cm *ClientMgr) GetAllUsers() map[int]*UserProcessor {
	return cm.onlineUsers
}

// 当一个用户离线后，就从onlineUsers切片中删除
func (cm *ClientMgr) DeleteClient(userId int) {
	delete(cm.onlineUsers, userId)
}
