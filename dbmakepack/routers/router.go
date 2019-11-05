package routers

import (
	"dbmakepack/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dbpackage/getproductlist", &controllers.ProductController{})
	beego.Router("/dbpackage/gettrynolist", &controllers.TrynoController{})
	beego.Router("/dbpackage/getpartnerlist", &controllers.PartnerController{})
	beego.Router("/dbpackage/getallmakepacketinfo", &controllers.MakepackInfoController{})
	beego.Router("/dbpackage/getresultbytaskid", &controllers.ResultController{})
	beego.Router("/dbpackage/stopmakepackage", &controllers.StopController{})
	beego.Router("/dbpackage/makepackage", &controllers.MakePackController{})
}
