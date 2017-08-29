package routers

import (
	"github.com/astaxie/beego"
	"github.com/migege/anthill-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/status", &controllers.StatusController{})
	beego.Router("/status/stream", &controllers.StatusController{}, "get:Stream")
}
