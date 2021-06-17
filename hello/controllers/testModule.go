package controllers

import (
	"fmt"
	orm "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type UserInfo struct {
	Id       int64
	Username string
	Password string
}

type TestModuleController struct {
	beego.Controller
}

func (c *TestModuleController) Get() {
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.146.154:3306)/my_db?charset=utf8")
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	user := UserInfo{Username: "tom", Password: "123"}
	id, err := o.Insert(&user)
	c.Ctx.WriteString(fmt.Sprintf("id: %d\terr:%v", id, err))
}

func (c *TestModuleController) Post() {

}
