package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//func index(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Form: ", r.Form)
//	fmt.Println("Path: ", r.URL.Path)
//
//	_, _ = fmt.Fprintf(w, "It works!")
//}

//func main() {
//	http.HandleFunc("/", index)
//
//	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello: " + request
	return nil
}

func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("ListenTCP error: ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept error: ", err)
	}

	rpc.ServeConn(conn)
}
