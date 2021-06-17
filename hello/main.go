package main

import (
	_ "src/hello/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

