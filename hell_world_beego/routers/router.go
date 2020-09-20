package routers

import (
	"github.com/astaxie/beego"
	"hell_world_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//新增路由
	beego.Router("/hello", &controllers.HelloController{})
}
