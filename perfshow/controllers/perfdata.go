package controllers

import (
	"perfshow/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Operations about Users
type CreatePerfdataController struct {
	beego.Controller
}

type GetPerfdataController struct{
	beego.Controller
}

type CreatePerfdatashowController struct{
	beego.Controller
}

type QueryPerfdataController struct{
	beego.Controller
}

type DelPerfdataController struct{
	beego.Controller
}

type QueryPerfdatashowController struct{
	beego.Controller
}

type GetPerfdatashowController struct{
	beego.Controller
}

type UpdateStarttimeController struct{
	beego.Controller
}

type GetPhoneinfoController struct{
	beego.Controller
}


type Querydata struct{
	Gamename		string
	Platform		string
	Gamever			string
	Clientver		string
	Phone			string
}

type Perfdataq struct{
	Id				int `orm:"column(id);pk;auto"`
	Taskid			int64
	Name			string
	Gameid			int
	Phone			string
	FpsMax			float32
	FpsAvr			float32
	FpsMin			float32
	CpuMax			float32
	CpuAvr			float32
	CpuMin			float32
	MemMax			float32
	MemAvr			float32
	MemMin			float32
	NetMax		    float32
	NetAvr			float32
	NetMin			float32
	BatteryAvr 		float32
	TempAvr			float32
	Duration 		float32
	Starttime 		float32
}

type Perfdatashowq struct{
	Id				int `orm:"column(id);pk;auto"`
	Taskid			int64
	Name			string
	Gameid			int
	Phone			string
	FpsMax			float32
	FpsAvr			float32
	FpsMin			float32
	CpuMax			float32
	CpuAvr			float32
	CpuMin			float32
	MemMax			float32
	MemAvr			float32
	MemMin			float32
	NetMax		    float32
	NetAvr			float32
	NetMin			float32
	BatteryAvr 		float32
	TempAvr			float32
	Starttime 		float32
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *CreatePerfdataController) Post() {
	fmt.Println("CreatePerfdataController")
	json_output := Get_JsonOutput()
	var perfdata models.Perfdata
	// name := this.Ctx.Request.PostFormValue("name")
	// name,_ := this.GetBool("isenable")
	// name := this.Ctx.Request.PostForm["name"][0]
	// fmt.Println("isenable")
	// fmt.Println(name)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &perfdata)
	if err != nil{
		this.Data["json"] = "error"
	}
	logs.Debug(perfdata)
	fmt.Println(perfdata.Taskid)
	if perfdata.Taskid == 0{
		json_output["msg"] = "param error"
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
		this.ServeJSON()
		return
	}

	fmt.Println(perfdata)
	ret,merr :=models.AddPerfData(perfdata)
	fmt.Println(ret)
	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}else{
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}

func (this * GetPerfdataController) Get(){
	logs.Debug("GetPerfdataController")
	// var perfdatas []Perfdata
	// qb,_ := orm.NewQueryBuilder("mysql")
	var perfdatas []models.Perfdata
	// qb, _ := orm.NewQueryBuilder("mysql")
	taskid,_:= this.GetInt("taskid")

	logs.Debug(taskid)
	logs.Debug(taskid)
	if taskid != 0{
		logs.Debug("query taskid")
		logs.Debug(taskid)
		o := orm.NewOrm()
		qs := o.QueryTable("perfdata")
		perfdata := new(models.Perfdata)
		qs.Filter("taskid", taskid).All(&perfdatas)
		logs.Debug(perfdatas)
		logs.Debug(qs)
		logs.Debug(perfdata)
		// qb.Select("perf.taskid", "")
	}else{
		logs.Debug("gamename null")
	}
	json_output := Get_JsonOutput()
	ret, merr := models.GetAllPerfdata()
	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = COMMON_ERROR
		this.Data["json"] = json_output
	}else{
		json_output["data"] = ret
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}

func (this *CreatePerfdatashowController) Post() {
	logs.Debug("CreatPerfdatashowController")
	json_output := Get_JsonOutput()
	var perfdatashow models.Perfdatashow
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &perfdatashow)
	if err != nil{
		this.Data["json"] = "error"
	}
	logs.Debug(perfdatashow)
	fmt.Println(perfdatashow.Taskid)
	if perfdatashow.Gameid == 0 || perfdatashow.Phone == ""{
		json_output["msg"] = "param error"
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
		this.ServeJSON()
		return
	}

	fmt.Println(perfdatashow)
	ret,merr :=models.AddPerfDatashow(perfdatashow)
	fmt.Println(ret)
	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}else{
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}

func (this *QueryPerfdataController) Post(){
	logs.Debug("QueryPerfdataController")
	json_output := Get_JsonOutput()
	var querydata Querydata
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &querydata)
	// var gamename = querydata.gamename
	logs.Debug(querydata.Gamename)
	logs.Debug(querydata.Platform)
	logs.Debug(querydata.Clientver)
	logs.Debug(querydata.Gamever)
	logs.Debug(querydata.Phone)
	if err != nil{
		json_output["msg"] = "param error"
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}
	o := orm.NewOrm()
	// var maps []orm.Params
	var perfdata []Perfdataq
	// var ret string
	qb, _ := orm.NewQueryBuilder("mysql")
	var sql string
	
	qb.Select("p.id, g.name, p.taskid, p.gameid, p.phone, p.fps_max, p.fps_avr, p.fps_min, p.cpu_max, p.cpu_avr, p.cpu_min, p.mem_max, p.mem_avr, p.mem_min, p.net_max, p.net_avr, p.net_min, p.battery_avr, p.temp_avr, p.duration, p.starttime").
	From("perfdata as p").
	InnerJoin("gameinfo as g").On("p.gameid=g.id")

	if querydata.Gamename != ""{
		qb.And("g.name = ?")
	}

	if querydata.Platform != ""{
		qb.And("g.platform = ?")
	}

	if querydata.Clientver != ""{
		qb.And("g.Clientver = ?")
	}

	if querydata.Gamever != ""{
		qb.And("g.gamever = ?")
	}

	if querydata.Phone != ""{
		qb.And("p.phone = ?")
	}

	sql = qb.String()
	logs.Debug(sql)
	var num int64
	var merr error
	if querydata.Gamename == ""{
		if querydata.Phone == ""{
			logs.Debug("query all")
			num, merr =o.Raw(sql).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Phone).QueryRows(&perfdata)
		}
	}else if querydata.Platform == ""{
		if querydata.Phone == ""{
			num, merr =o.Raw(sql, querydata.Gamename).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Phone).QueryRows(&perfdata)
		}
	}else if querydata.Clientver == ""{
		if querydata.Phone == ""{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Phone).QueryRows(&perfdata)
		}
	}else if querydata.Gamever == ""{
		if querydata.Phone ==""{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver, querydata.Phone).QueryRows(&perfdata)
		}
	}else{
		if querydata.Phone == ""{
			logs.Debug(o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Gamever))
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver, querydata.Gamever, querydata.Phone).QueryRows(&perfdata)
		}
	}
	logs.Debug(num)
	logs.Debug(merr)
	if num == 0{
		json_output["data"] = ""
	}else{
		json_output["data"] = perfdata
	}
	logs.Debug(perfdata)
	this.Data["json"] = json_output
	this.ServeJSON()
}

func (this * DelPerfdataController) Post(){
	logs.Debug("DelPerfdataController")
	json_output := Get_JsonOutput()
	var userinfo_del Userinfo_del 
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userinfo_del)
	logs.Debug(userinfo_del.Perfdataid)
	logs.Debug(userinfo_del.Username)
	logs.Debug(userinfo_del.Session)
	if userinfo_del.Username == "" || userinfo_del.Session == ""{
		json_output["errorcode"] = AUTH_ERROR
		json_output["msg"] = "auth error"
		this.Data["json"] = json_output
	}
	logs.Debug(userinfo_del.Perfdataid)
	o := orm.NewOrm()
	num,err := o.Delete(&models.Perfdata{Id:userinfo_del.Perfdataid})
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

func (this * QueryPerfdatashowController) Post(){
	logs.Debug("QueryPerfdatashowController")
	json_output := Get_JsonOutput()
	var querydata Querydata
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &querydata)
	logs.Debug(querydata.Gamename)
	logs.Debug(querydata.Platform)
	logs.Debug(querydata.Clientver)
	logs.Debug(querydata.Gamever)
	logs.Debug(querydata.Phone)
	if err != nil{
		json_output["msg"] = "param error"
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}

	o := orm.NewOrm()
	var perfdata []Perfdatashowq
	qb, _ := orm.NewQueryBuilder("mysql")
	var sql string

	qb.Select("p.id, g.name, p.taskid, p.gameid, p.phone, p.fps_max, p.fps_avr, p.fps_min, p.cpu_max, p.cpu_avr, p.cpu_min, p.mem_max, p.mem_avr, p.mem_min, p.net_max, p.net_avr, p.net_min, p.battery_avr, p.temp_avr, p.starttime").
	From("perfdatashow as p").
	InnerJoin("gameinfo as g").On("p.gameid=g.id")

	if querydata.Gamename != ""{
		qb.And("g.name = ?")
	}

	if querydata.Platform != ""{
		qb.And("g.platform = ?")
	}

	if querydata.Clientver != ""{
		qb.And("g.Clientver = ?")
	}

	if querydata.Gamever != ""{
		qb.And("g.gamever = ?")
	}

	if querydata.Phone != ""{
		qb.And("p.phone = ?")
	}

	sql = qb.String()
	logs.Debug(sql)
	var num int64
	var merr error
	if querydata.Gamename == ""{
		if querydata.Phone == ""{
			logs.Debug("query all")
			num, merr =o.Raw(sql).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Phone).QueryRows(&perfdata)
		}
	}else if querydata.Platform == ""{
		if querydata.Phone == ""{
			num, merr =o.Raw(sql, querydata.Gamename).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Phone).QueryRows(&perfdata)
		}
	}else if querydata.Clientver == ""{
		if querydata.Phone == ""{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Phone).QueryRows(&perfdata)
		}
	}else if querydata.Gamever == ""{
		if querydata.Phone == ""{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver, querydata.Phone).QueryRows(&perfdata)
		}
	}else{
		if querydata.Phone == ""{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver, querydata.Gamever).QueryRows(&perfdata)
		}else{
			num, merr =o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver, querydata.Gamever, querydata.Phone).QueryRows(&perfdata)
		}
	}
	logs.Debug(num)
	logs.Debug(merr)
	if num == 0{
		json_output["data"] = ""
	}else{
		json_output["data"] = perfdata
	}
	logs.Debug(perfdata)
	this.Data["json"] = json_output
	this.ServeJSON()
}

func (this * GetPerfdatashowController) Get(){
	logs.Debug("GetPerfdatashowController")
	// var perfdatas []models.Perfdatashow
	json_output := Get_JsonOutput()
	ret, merr := models.GetAllPerfdatashow()
	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = COMMON_ERROR
		this.Data["json"] = json_output
	}else{
		json_output["data"] = ret
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}

type Updatestdata struct{
	Taskid			int64
	Starttime 		float32
}

func (this * UpdateStarttimeController) Post(){
	logs.Debug("UpdateStarttimeController")
	json_output := Get_JsonOutput()
	var updatestdata Updatestdata
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &updatestdata)
	logs.Debug(updatestdata.Taskid)
	logs.Debug(updatestdata.Starttime)
	merr := models.UpdateStarttime(updatestdata.Taskid, updatestdata.Starttime)
	if err != nil{
		json_output["msg"] = "param error"
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}

	if merr != nil{
		json_output["msg"] = merr
		json_output["errorcode"] = COMMON_ERROR
		this.Data["json"] = json_output
	}else{
		this.Data["json"] = json_output
	}
	this.ServeJSON()
}

func (this * GetPhoneinfoController) Get(){
	logs.Debug("GetPhoneinfoController")
	var phonelist []string
	json_output := Get_JsonOutput()
	this.Data["json"] = json_output
	o := orm.NewOrm()
	qb,_ := orm.NewQueryBuilder("mysql")
	var sql string
	qb.Select("distinct perfdata.phone").
	From("perfdata")

	sql = qb.String()
	logs.Debug(sql)
	num,merr:=o.Raw(sql).QueryRows(&phonelist)
	logs.Debug(merr)
	if num == 0 && merr == nil{
		json_output["data"] = ""
	}else{
		json_output["data"] = phonelist
	}
	this.ServeJSON()
}