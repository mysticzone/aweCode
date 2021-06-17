package main

import (
	"fmt"
	"go_dev/day7/example1/balance"
	"hash/crc32"
	// "hash/crc64"
	"math/rand"
)

func init() {
	balance.RegisterBalance("hash", &HashBalance{})
}

type HashBalance struct {
	key string
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int31())
	if len(key) > 0 {
		defKey = key[0]
	}
	lens := len(insts)
	if lens == 0 {
		fmt.Errorf("No found instances")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
