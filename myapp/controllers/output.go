package controllers

import "github.com/astaxie/beego"

type User struct {
	beego.Controller
}

type UserInfo struct {
	Name string
	Age  int
	Sex  string
}

func (this *User) JsonFunc() {
	user := &UserInfo{"renwoxing", 20, "man"}
	this.Data["json"] = user
	this.ServeJSON()
}

func (this *User) XmlFunc() {
	user := &UserInfo{"renwoxing", 20, "man"}
	this.Data["xml"] = user
	this.ServeXML()
}

func (this *User) YamlFunc() {
	user := &UserInfo{"renwoxing", 20, "man"}
	this.Data["yaml"] = user
	this.ServeYAML()
}

func (this *User) JsonpFunc() {
	user := &UserInfo{"renwoxing", 20, "man"}
	this.Data["jsonp"] = user
	this.ServeJSONP()
}
