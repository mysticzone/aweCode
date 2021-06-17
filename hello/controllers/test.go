package controllers

import beego "github.com/beego/beego/v2/server/web"

type TestController struct {
	beego.Controller
}

func (t *TestController) Get() {
	t.Ctx.WriteString("<font style='color:red;'>Controller</font>")
}
