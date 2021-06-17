package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yanghao@luojilab.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Ctx.WriteString("Method Post")
}

func (c *MainController) Test() {
	c.Data["Website"] = "Test.me"
	c.Data["Email"] = "Test@luojilab.com"
	c.TplName = "index.tpl"
}


