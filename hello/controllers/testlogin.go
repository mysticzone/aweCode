package controllers

import beego "github.com/beego/beego/v2/server/web"

type TestLoginController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func (t *TestLoginController) Login() {
	//id := t.GetString("id")
	//t.Ctx.WriteString(id)
	name := t.Ctx.GetCookie("name")
	password := t.Ctx.GetCookie("password")

	_ = t.SetSession("name", name)
	_ = t.SetSession("password", password)

	if name != "" {
		t.Ctx.WriteString("Username: " + name + " Password: " + password)
	} else {

		t.Ctx.WriteString(
			`<html>
	<form action="http://127.0.0.1:8082/test_login" method="post">
		<input type="text" name="Username" />
		<input type="password" name="Password" />
		<input type="submit" value="submit" />
		</form>
				</html>`)
	}

}

func (t *TestLoginController) Post() {

	u := User{}
	if err := t.ParseForm(&u); err != nil {
		//
	}
	t.Ctx.SetCookie("name", u.Username, 100, "/")
	t.Ctx.SetCookie("password", u.Password, 100, "/")

	t.Ctx.WriteString("Username: " + u.Username + " Password: " + u.Password)
}
