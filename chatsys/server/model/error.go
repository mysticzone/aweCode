package model

import "errors"

// 自定义登录，注册错误
var (
	ErrUserNotExsit  = errors.New("User not exist")
	ErrInvalidPasswd = errors.New("UserPwd or UserName not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("User exist")
)
