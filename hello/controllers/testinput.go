package controllers

import beego "github.com/beego/beego/v2/server/web"

type TestInputController struct {
	beego.Controller
}

//type User struct {
//	Username string
//	Password string
//}

func (t *TestInputController) Get() {
	//id := t.GetString("id")
	//t.Ctx.WriteString(id)

	name := t.GetSession("name")
	password := t.GetSession("password")
	if name != "" {
		t.Ctx.WriteString("Username: " + name.(string) + " Password: " + password.(string))

	} else {
		t.Ctx.WriteString(
			`<html>
	<form action="http://127.0.0.1:8082/test_input" method="post">
		<input type="text" name="Username" />
		<input type="password" name="Password" />
		<input type="submit" value="submit" />
		</form>
				</html>`)
	}

}

func (t *TestInputController) Post() {
	u := User{}
	if err := t.ParseForm(&u); err != nil {
		//
	}
	t.Ctx.WriteString("Username: " + u.Username + " Password: " + u.Password)
}
