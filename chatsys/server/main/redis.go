package main

//import (
//	"github.com/garyburd/redigo/redis"
//	"time"
//)
//
//var pool *redis.Pool
//
//func initRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration) {
//	pool = &redis.Pool{
//		MaxIdle:     idleConn,
//		MaxActive:   maxConn,
//		IdleTimeout: idleTimeout,
//		Dial: func() (redis.Conn, error) {
//			return redis.Dial("tcp", addr)
//		},
//	}
//	return
//}
//
//// 从连接池中取一个链接
//func GetConn() redis.Conn {
//	return pool.Get()
//}
//
//// 关闭一个从连接池中取到的链接
//func PutConn(conn redis.Conn) {
//	err := conn.Close()
//	if err != nil {
//		return
//	}
//}
