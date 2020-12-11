package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (login *LoginController) Get() {
	login.TplName = "register.tpl"
}

func (login *LoginController) Post() {
	user := login.GetString("user")
	pwd := login.GetString("pwd")
	email := login.GetString("email")
	login.Ctx.WriteString("user=" + user + " password=" + pwd + " email=" + email + "\n")
}
