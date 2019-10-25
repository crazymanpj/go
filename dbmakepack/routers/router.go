package routers

import (
	"dbmakepack/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dbpackage/getproductlist", &controllers.ProductController{})
	beego.Router("/dbpackage/gettrynolist", &controllers.TrynoController{})
}
