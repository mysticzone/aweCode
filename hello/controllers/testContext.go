package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type TestContextController struct {
	beego.Controller
}

func (t *TestContextController) Get() {
	ip := t.Ctx.Input.IP()
	port := t.Ctx.Input.Port()
	t.Ctx.WriteString(ip + ":" + strconv.Itoa(port))
	name := t.Ctx.Input.Query("name")
	t.Ctx.WriteString("name: " + name)

	m1 := make(map[string]float64)
	m1["zhangsan"] = 99.1
	m1["lisi"] = 90.2
	_ = t.Ctx.Output.JSON(m1, false, false)
}
