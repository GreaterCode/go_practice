package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}
type Person struct {
	User  string `form:"user"`
	Pwd   string `form:"pwd"`
	Email string `form:"email"`
}

func (login *LoginController) Get() {
	login.TplName = "register.tpl"
}

func (login *LoginController) Post() {
	//user := login.GetString("user")
	//pwd := login.GetString("pwd")
	//email := login.GetString("email")
	var p Person
	if err := login.ParseForm(&p); err != nil {
		login.Ctx.WriteString("parse info error!!")
	}
	login.Ctx.WriteString("user=" + p.User + " password=" + p.Pwd + " email=" + p.Email + "\n")
}
