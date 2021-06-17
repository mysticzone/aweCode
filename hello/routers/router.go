package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
	"src/hello/controllers"
)

func init() {
	beego.Router("/home", &controllers.MainController{}, "get:Test")
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/test_input", &controllers.TestInputController{})
	beego.Router("/test_login", &controllers.TestLoginController{}, "get:Login;post:Post")

	beego.Router("/test_module", &controllers.TestModuleController{}, "get:Get;post:Post")
	beego.Router("/test_httplib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
	beego.Router("/test_context", &controllers.TestContextController{}, "get:Get;post:Post")
	beego.Get("/", func(ctx *context2.Context) {
		_ = ctx.Output.Body([]byte("hello world"))
	})

}
