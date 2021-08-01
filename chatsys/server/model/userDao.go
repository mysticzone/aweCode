package model

import (
	"chatsys/common"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

var MyUserDao *UserDao

// 工厂模式：获取一个 UserDao 变量（实例）
func NewUserMgr(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (ud *UserDao) getUser(conn redis.Conn, id int) (user *common.User, err error) {
	result, err := redis.String(conn.Do("HGet", "users", fmt.Sprintf("%d", id)))
	fmt.Printf("result=%v\n", result)

	// 如果没有对应的用户ID，返回 ErrUserNotExist
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExsit
		}
		return
	}

	// 声明 *User
	user = &common.User{}
	// 将 user 反序列化，返回序列化后的 user 实例
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		fmt.Printf("xxx~=%v\n", err)
		return
	}
	fmt.Printf("从数据库里得到 user=%v\n", user)
	return
}

// 完成登录操作，Redis验证。
// 1.登录信息有效，返回User实例
// 2.如果登录用户信息无效，返回nil和对应的自定义错误
func (ud *UserDao) Login(userId int, userPwd string) (user *common.User, err error) {
	// 得到一个链接
	conn := ud.pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close failed. err=", err)
			return
		}
	}(conn)

	// 通过ID来获取对应的用户信息
	user, err = ud.getUser(conn, userId)
	if err != nil {
		return
	}

	// 比较通过ID查询到的用户登录密码是否一致
	if user.Pwd != userPwd {
		err = ErrInvalidPasswd
		return
	}

	// user.Status = common.UserStatusOnline
	// user.LastLogin = fmt.Sprintf("%v", time.Now())
	return
}

// 在服务器端完成注册
func (ud *UserDao) Register(user *common.User) (err error) {
	// 获取链接
	conn := ud.pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close failed. err=", err)
			return
		}
	}(conn)

	_, err = ud.getUser(conn, user.Id)
	if err == nil { //如果获取到一个用户，说明该用户存在了
		err = ErrUserExist //用户存在了...
		return
	}

	//序列化这个user
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//存入数据库
	_, err = conn.Do("HSet", "users", fmt.Sprintf("%d", user.Id), string(data))
	if err != nil {
		return
	}
	return
}
