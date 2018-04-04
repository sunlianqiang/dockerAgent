package routers

import (
	"dockerAgent/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/v1/image", &controllers.ImageController{})
}
