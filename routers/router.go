package routers

import (
	"github.com/astaxie/beego"
	"github.com/migege/anthill-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/status/:id([0-9]+)", &controllers.StatusController{}, "get:ById")
	beego.Router("/status/stream", &controllers.StatusController{}, "get:Stream")
}
