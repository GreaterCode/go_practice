package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type TagValidationController struct {
	beego.Controller
}

type Worker struct {
	User     string `valid:"Required;MinSize(6)"` //不为空且最小长度为6
	Pwd      string `valid:"Required;MinSize(6)"` //不为空且最小长度为6
	RealName string `valid:"Required"`            //不为空
	Age      string `valid:"Numeric"`             //为数字字符串
	IdCard   string `valid:"Required;Length(18)"` //不为空且长度为18
	Email    string `valid:"Required;Email"`      //不为空且格式为Email
	Tel      string `valid:"Required;Mobile"`     //不为空且格式为mobile

}

func (this *TagValidationController) Get() {
	this.TplName = "validation.html"
}

func (this *TagValidationController) Post() {
	this.TplName = "validation.html"
	var worker Worker
	// 解析数据到结构体
	if err := this.ParseForm(&worker); err != nil {
		this.Data["PaseFormErr"] = "数据解析失败"
	} else {
		valid := validation.Validation{}

		// 自定义消息提示
		var MessageTmpls = map[string]string{
			"Required": "不能为空",
			"MinSize":  "最短长度为 %d",
			"Length":   "长度必须为 %d",
			"Numeric":  "必须是有效的数字",
			"Email":    "无效的电子邮件地址",
			"Mobile":   "无效的手机号码",
		}
		validation.SetDefaultMessage(MessageTmpls)

		// 验证是否struct tag是否正确
		validResult, err := valid.Valid(&worker)
		if err != nil {
			this.Data["ParseFormerErr"] = err
		}
		if !validResult {
			// 验证没通过
			for _, err := range valid.Errors {
				data := "Verify" + err.Field
				this.Data[data] = err.Message
			}
		}

	}
}
