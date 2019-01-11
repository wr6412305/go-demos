package routers

import (
	"go-demos/beego-demo/helloprj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
