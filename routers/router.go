package routers

import (
	"github.com/astaxie/beego"
	"github.com/migege/anthill-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ant/:id([0-9]+)", &controllers.AntController{}, "get:ById")
	beego.Router("/ant/stream", &controllers.AntController{}, "get:Stream")
	beego.Router("/ant/cmd", &controllers.AntController{}, "post:DoCmd")
}
