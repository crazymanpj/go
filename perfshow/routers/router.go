// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"perfshow/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/perfshow/creategameinfo", &controllers.CreateGameinfoController{})
	beego.Router("/perfshow/createtaskinfo", &controllers.CreateTaskinfoController{})
	beego.Router("/perfshow/gettaskinfo", &controllers.GetTaskinfoController{})
	beego.Router("/perfshow/getgameinfo", &controllers.GetGameinfoController{})
	beego.Router("/perfshow/createperfdata", &controllers.CreatePerfdataController{})
	beego.Router("/perfshow/getperfdata", &controllers.GetPerfdataController{})
	beego.Router("/perfshow/deltask", &controllers.DelTaskController{})
	beego.Router("/perfshow/createperfdatashow", &controllers.CreatePerfdatashowController{})
	beego.Router("/perfshow/queryperfdata", &controllers.QueryPerfdataController{})
	beego.Router("/perfshow/delgame", &controllers.DelGameController{})
	beego.Router("/perfshow/querygameinfo", &controllers.QueryGameinfoController{})
	beego.Router("/perfshow/delperfdata", &controllers.DelPerfdataController{})
	beego.Router("/perfshow/queryperfdatashow", &controllers.QueryPerfdatashowController{})
	beego.Router("/perfshow/getperfdatashow", &controllers.GetPerfdatashowController{})
	beego.Router("/perfshow/updatestarttime", &controllers.UpdateStarttimeController{})
	beego.Router("/perfshow/getphoneinfo", &controllers.GetPhoneinfoController{})
}
