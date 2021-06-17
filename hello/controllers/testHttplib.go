package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

type TestHttpLibController struct {
	beego.Controller
}

func (t *TestHttpLibController) Get() {
	req := httplib.Get("https://douban.com")
	//res, err := req.Response()
	//fmt.Println(err)
	//fmt.Println(res)
	res, err := req.String()
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	t.Ctx.WriteString(res)
	//t.Ctx.WriteString("<font style='color:red;'>Controller</font>")
}
