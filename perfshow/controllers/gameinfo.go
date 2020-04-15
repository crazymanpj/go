package controllers

import (
	"perfshow/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// "net/http"
	// "strings"
	// "io/ioutil"
)

// Operations about Users
type CreateGameinfoController struct {
	beego.Controller
}

// Operations about Users
type GetGameinfoController struct {
	beego.Controller
}

type DelGameController struct{
	beego.Controller
}

type QueryGameinfoController struct{
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *CreateGameinfoController) Post() {
	var gameinfo models.Gameinfo
	json_output := Get_JsonOutput()
	// name := this.Ctx.Request.PostFormValue("name")
	// name,_ := this.GetBool("isenable")
	// name := this.Ctx.Request.PostForm["name"][0]
	// fmt.Println("isenable")
	// fmt.Println(name)
	json.Unmarshal(this.Ctx.Input.RequestBody, &gameinfo)
	// if err := this.ParseForm(&gameinfo); err != nil{
	// 	fmt.Println("param error")
	// }
	fmt.Println(gameinfo)
	models.AddGameinfo(gameinfo)

	json_output["msg"] = "ok"
	this.Data["json"] = json_output
	this.ServeJSON()
}

func (this *GetGameinfoController) Get() {
	data,merr := models.GetAllGameinfo()
	json_output := Get_JsonOutput()
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



func (this * DelGameController) Post(){
	logs.Debug("DelGameController")
	json_output := Get_JsonOutput()
	var userinfo_del Userinfo_del 
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userinfo_del)
	logs.Debug(userinfo_del.Gameid)
	logs.Debug(userinfo_del.Username)
	logs.Debug(userinfo_del.Session)
	// logs.Debug(LOGIN_SERVER_TEST + "/user/certify_session/")
	// resp, err := http.Post(LOGIN_SERVER_TEST + "/user/certify_session/",
	// 			"application/json",
	// 			strings.NewReader("username=pangjian"))

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	fmt.Println(err)
	if userinfo_del.Username == "" || userinfo_del.Session == ""{
		json_output["errorcode"] = AUTH_ERROR
		json_output["msg"] = "auth error"
		this.Data["json"] = json_output
	}
	logs.Debug(userinfo_del.Gameid)
	o := orm.NewOrm()
	num,err := o.Delete(&models.Gameinfo{Id:userinfo_del.Gameid})
	if err == nil && num!=0{
		logs.Debug(num)
		this.Data["json"] = json_output
	}else{
		logs.Debug(err)
		json_output["errorcode"] = COMMON_ERROR
		json_output["msg"] = "delete error!"
		this.Data["json"] = json_output
	}
	this.Data["json"] = json_output
	this.ServeJSON()
}

func (this * QueryGameinfoController) Post(){
	logs.Debug("QueryGameinfoController")
	json_output := Get_JsonOutput()
	var querydata Querydata
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &querydata)
	// var gamename = querydata.gamename
	logs.Debug(querydata.Gamename)
	logs.Debug(querydata.Platform)
	logs.Debug(querydata.Clientver)
	logs.Debug(querydata.Gamever)
	
	var namelist []string
	var platformlist []string
	var clientverlist []string
	var gameverlist []string
	if err != nil{
		json_output["msg"] = "param error"
		json_output["errorcode"] = -1
		this.Data["json"] = json_output
	}

	o := orm.NewOrm()
	// qb, _ := orm.NewQueryBuilder("mysql")
	// var sql string

	if true{
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("distinct name").
		From("gameinfo")
		sql := qb.String()
		logs.Debug(sql)
		o.Raw(sql).QueryRows(&namelist)
	}

	if querydata.Gamename != ""{
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("distinct platform").
		From("gameinfo").
		Where("name = ?")
		sql := qb.String()
		logs.Debug(sql)
		o.Raw(sql, querydata.Gamename).QueryRows(&platformlist)	
	}

	if querydata.Platform != "" && querydata.Gamename != ""{
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("distinct clientver").
		From("gameinfo").
		Where("name = ?").
		And("platform = ?")
		sql := qb.String()
		logs.Debug(sql)
		o.Raw(sql, querydata.Gamename, querydata.Platform).QueryRows(&clientverlist)	
	}

	if querydata.Clientver != "" && querydata.Gamename != "" && querydata.Platform != ""{
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("distinct gamever").
		From("gameinfo").
		Where("name = ?").
		And("platform = ?").
		And("clientver = ?")
		sql := qb.String()
		logs.Debug(sql)
		o.Raw(sql, querydata.Gamename, querydata.Platform, querydata.Clientver).QueryRows(&gameverlist)
	}



	logs.Debug(namelist)
	logs.Debug(platformlist)
	logs.Debug(clientverlist)
	logs.Debug(gameverlist)
	json_output["namelist"] = namelist
	json_output["platformlist"] = platformlist
	json_output["clientverlist"] = clientverlist
	json_output["gameverlist"] = gameverlist
	this.Data["json"] = json_output
	this.ServeJSON()
}

// // @Title GetAll
// // @Description get all Users
// // @Success 200 {object} models.User
// // @router / [get]
// func (u *UserController) GetAll() {
// 	users := models.GetAllUsers()
// 	u.Data["json"] = users
// 	u.ServeJSON()
// }
