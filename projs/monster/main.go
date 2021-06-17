package monster

import (
	"fmt"
	"testing"
)

//测试用例,测试 Store 方法
func TestStore(t *testing.T) {

	//先创建一个Monster 实例
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火.",
	}
	res := monster.Store()
	if !res {
		fmt.Printf("monster.Store() 错误，希望为=%v 实际为=%v", true, res)
	}
	fmt.Println("monster.Store() 测试成功!")
}

func main() {
}
