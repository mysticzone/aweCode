package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Dialing: ", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello123", &reply)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("[reply:] ", reply)
}
