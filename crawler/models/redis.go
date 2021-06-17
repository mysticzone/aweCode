package models

import (
	"fmt"
	"github.com/astaxie/goredis"
)

var (
	client goredis.Client
)

const (
	URL_QUEUE     = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

func ConnectRedis(addr string) {
	client.Addr = addr
}

func PutInQueue(url string) {
	err := client.Lpush(URL_QUEUE, []byte(url))
	if err != nil {
		fmt.Printf("PutInQueue Err:%v\n", err)
		panic(err)
	}
}

func PopFromQueue() string {
	url, err := client.Rpop(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return string(url)
}

func GetQueueLength() int {
	length, err := client.Llen(URL_QUEUE)
	if err != nil {
		return 0
	}
	return length
}

func AddToSet(url string) {
	ret, err := client.Sadd(URL_VISIT_SET, []byte(url))
	if err != nil {
		panic(err)
	}
	fmt.Printf("AddToSet: %v\n", ret)
}

func IsVisit(url string) bool {
	bIsVisit, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil {
		return false
	}
	return bIsVisit
}
