package controllers

import (
	"perfshow/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// "strconv"
)

// Operations about Users
type CreateTaskinfoController struct {
	beego.Controller
}


// Operations about Users
type GetTaskinfoController struct {
	beego.Controller
}

type DelTaskController struct{
	beego.Controller
}

type Mtaskinfo struct{
	Id				int64
	Phone		    string
	Name			string
	Duration 		int
	Tester			string
	Time 			time.Time
	Interval		int
}

// func Get_JsonOutput()(map[string]interface{}){
// 	var json_output = map[string]interface{}{
// 		"errorcode": 0,
// 		"msg" : "ok",
// 	}
// 	return json_output
// }

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *CreateTaskinfoController) Post() {
	logs.Debug("CreateTaskinfoController")
	var taskinfo models.Taskinfo
	json_output := Get_JsonOutput()
	// name := this.Ctx.Request.PostFormValue("name")
	name,_ := this.GetInt("id")
	// id := this.Input().Get("id")
	intid := this.Ctx.Request.PostFormValue("id")
	logs.Debug(intid)
	// name := this.Ctx.Request.PostForm["name"][0]
	fmt.Println("isenable")
	fmt.Println(name)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &taskinfo)
	fmt.Println(taskinfo.Id)
	taskinfo.Time = time.Now()
	logs.Debug(taskinfo)
	if err != nil{
		this.Data["json"] = "error"
	}
	// if err := this.ParseForm(&gameinfo); err != nil{
	// 	fmt.Println("param error")
	// }
	// fmt.Println(taskinfo)
	ret,merr :=models.AddTaskinfo(taskinfo)
	fmt.Println(ret)
	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = COMMON_ERROR
		this.Data["json"] = json_output
	}else{
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}

func (this *GetTaskinfoController) Get() {
	logs.Debug("gettaskinfo")
	json_output := Get_JsonOutput()
	data,merr := models.GetAllTaskinfo()
	var taskinfo []Mtaskinfo
	qb,_ := orm.NewQueryBuilder("mysql")
	gamename := this.GetString("gamename")
	phone := this.GetString("phone")
	logs.Debug(gamename)
	if gamename != ""{
		logs.Debug("query gamename")
		logs.Debug(gamename)
		qb.Select("taskinfo.id", "gameinfo.name", "taskinfo.phone", "taskinfo.frequency", "taskinfo.tester", "taskinfo.time", "taskinfo.interval").
			From("taskinfo").
			InnerJoin("gameinfo").On("gameinfo.id = taskinfo.gameid").
			And("gameinfo.name = ?")
		sql := qb.String()
		logs.Debug(sql)
		o := orm.NewOrm()
		o.Raw(sql, gamename).QueryRows(&taskinfo)
		logs.Debug(taskinfo)
	}
	if gamename != "" && phone != ""{
		// logs.Debug(qb.Subquery("taskinfo.phone = ?", phone))
		qb.And("taskinfo.phone = ?")
		sql := qb.String()
		logs.Debug(sql)
		o := orm.NewOrm()
		o.Raw(sql, gamename, phone).QueryRows(&taskinfo)
		logs.Debug(taskinfo)
	}
	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}else{
		json_output["data"] = data
		this.Data["json"] = json_output
	}

	this.ServeJSON()
}

type Userinfo_del struct{
	Taskid			int64
	Username		string
	Session			string
	Gameid			int
	Perfdataid		int
}

func (this * DelTaskController) Post(){
	logs.Debug("DelTaskController")
	json_output := Get_JsonOutput()
	var userinfo_del Userinfo_del 
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userinfo_del)
	logs.Debug(userinfo_del.Taskid)
	logs.Debug(userinfo_del.Username)
	logs.Debug(userinfo_del.Session)
	if userinfo_del.Username == "" || userinfo_del.Session == ""{
		json_output["errorcode"] = AUTH_ERROR
		json_output["msg"] = "auth error"
		this.Data["json"] = json_output
	}
	logs.Debug(userinfo_del.Taskid)
	o := orm.NewOrm()
	num,err := o.Delete(&models.Taskinfo{Id:int64(userinfo_del.Taskid)})
	if err == nil && num!=0{
		logs.Debug(num)
		this.Data["json"] = json_output
	}else{
		logs.Debug(err)
		json_output["errorcode"] = COMMON_ERROR
		json_output["msg"] = "delete error!"
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}