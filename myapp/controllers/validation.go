package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ValidationController struct {
	beego.Controller
}

type Info struct {
	User     string
	Pwd      string
	RealName string
	Age      string
	IdCard   string
	Email    string
	Tel      string
}

func (this *ValidationController) Get() {
	this.TplName = "validation.html"
}

func (this *ValidationController) Post() {
	this.TplName = "validation.html"
	var info Info

	// 解析数据
	if err := this.ParseForm(&info); err != nil {
		this.Data["ParseFormErr"] = "数据解析到结构体错误"
	} else {
		valid := validation.Validation{}
		/*
			// 验证用户名
			userresult := valid.Required(info.User, "User")
			if !userresult.Ok {
				this.Data["VerifyUser"] = "用户名不能为空"
			} else {
				userresult = valid.MinSize(info.User, 6, "User")
				if !userresult.Ok {
					this.Data["VerifyUser"] = "用户名最小长度为6位"
				}
			}

			// 验证密码
			pwdresult := valid.Required(info.Pwd, "Pwd")
			if !pwdresult.Ok {
				this.Data["VerifyPwd"] = "密码不能为空"
			} else {
				pwdresult = valid.MinSize(info.Pwd, 6, "Pwd")
				if !pwdresult.Ok {
					this.Data["VerifyPwd"] = "密码最小长度为6位"
				}
			}

			// 验证密码是否一致
			cfmpwd := this.GetString("CfmPwd")
			if cfmpwd != "" {
				if cfmpwd != info.Pwd {
					this.Data["VerifyCfm"] = "两次密码不一致"
				}
			} else {
				this.Data["VerifyCfm"] = "确认密码不能为空"
			}

			// 验证姓名不能为空
			realnameresult := valid.Required(info.RealName, "RealName")
			if !realnameresult.Ok {
				this.Data["VerifyRealName"] = "真实姓名不能为空"
			}

			// 验证年龄是否为数字
			ageresult := valid.Numeric(info.Age, "Age")
			if !ageresult.Ok {
				this.Data["VerifyAge"] = "年龄只能为数字"
			}

			// 验证身份证是否合法
			idcardresult := valid.Required(info.IdCard, "IdCard")
			if !idcardresult.Ok {
				this.Data["VerifyIdCard"] = "身份证不能为空"
			} else {
				idcardresult = valid.Length(info.IdCard, 18, "IdCard")
				if !idcardresult.Ok {
					this.Data["VerifyIdCard"] = "身份证错误"
				}
			}

			// 验证邮箱
			emailresult := valid.Required(info.Email, "Email")
			if !emailresult.Ok {
				this.Data["VerifyEmail"] = "邮箱不能为空"
			} else {
				emailresult = valid.Email(info.Email,"Email")
				if !emailresult.Ok {
					this.Data["VerifyEmail"] = "邮箱格式错误"
				}
			}

			// 验证手机
			telresult := valid.Required(info.Tel, "Tel")
			if !telresult.Ok {
				this.Data["VerifyTel"] = "手机不能为空"
			} else {
				telresult = valid.Mobile(info.Tel, "Tel")
				if !telresult.Ok {
					this.Data["VerifyTel"] = "手机格式错误"
				}
			}*/

		valid.Required(info.User, "User").Message("用户名不能为空")
		valid.MinSize(info.User, 6, "User").Message("用户名最短为6位")
		valid.Required(info.Pwd, "Pwd").Message("密码不能为空")
		valid.MinSize(info.User, 6, "Pwd").Message("密码最短为6位")
		valid.Required(info.RealName, "RealName").Message("真实姓名不能为空")
		valid.Required(info.Age, "Age").Message("年龄不能为空")
		valid.Numeric(info.Age, "Age").Message("年龄只能为数字")
		valid.Required(info.IdCard, "IdCard").Message("身份证不能为空")
		valid.Length(info.IdCard, 18, "IdCard").Message("身份证格式非法")
		valid.Required(info.Email, "Email").Message("邮箱不能为空")
		valid.Email(info.Email, "Email").Message("邮箱格式非法")
		valid.Required(info.Tel, "Tel").Message("手机不能为空")
		valid.Mobile(info.Tel, "Tel").Message("手机格式非法")

		if valid.HasErrors() {
			for _, err := range valid.Errors {
				data := "Verify" + err.Key
				this.Data[data] = err.Message
			}
		}
	}
}
