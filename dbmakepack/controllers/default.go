package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"dbmakepack/models"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/widuu/goini"
	"strings"
	"os/exec"
)

type MainController struct {
	beego.Controller
}

type ProductController struct {
	beego.Controller
}

type TrynoController struct{
	beego.Controller
}

type PartnerController struct{
	beego.Controller
}

type MakepackInfoController struct{
	beego.Controller
}

type ResultController struct{
	beego.Controller
}

type StopController struct{
	beego.Controller
}

var JsonOutput = map[string]interface{}{
	"errorcode" : 0,
	"msg" : "",
	"data" : map[string]interface{}{},
	"version" : "1.0",
}

func init(){
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *ProductController) Get(){
	o := orm.NewOrm()
	var maps []orm.Params
	var productlist []interface{}
	num,_:=o.Raw("SELECT * FROM db_productlist").Values(&maps)
	for _,term := range maps{
		fmt.Println(term["id"], ":", term["name"], num)
		productlist = append(productlist, term["name"])
	}
	JsonOutput["data"] = productlist
	this.Data["json"] = JsonOutput     
    this.ServeJSON()
    return
}

func (this *TrynoController) Get(){
	o := orm.NewOrm()
	qs := o.QueryTable("db_trynolist")
	var trynolist []*models.DbTrynolist
	var itemlist []interface{}
	num,_ := qs.Filter("status", 1).All(&trynolist)
	for _,term:= range trynolist{
		fmt.Println(term.Id,term.Tryno, num)
		itemlist = append(itemlist, term.Tryno)
	}
	fmt.Println(trynolist)
	JsonOutput["data"] = itemlist
	this.Data["json"] = JsonOutput    
	this.ServeJSON()
	return
}

func (this *PartnerController) Get(){
	o := orm.NewOrm()
	qs := o.QueryTable("db_partnerlist")
	var partnerlist []*models.DbPartnerlist
	var itemlist []interface{}
	num,_ := qs.Filter("status", 1).All(&partnerlist)
	// for i := 0; i <num; i++{
	// 	itemlist = append(itemlist, partnerlist[i].Partner)
	// }
	for _,term:= range partnerlist{
		fmt.Println(term.Id,term.Partner, num)
		itemlist = append(itemlist, term.Partner)
	}
	fmt.Println(partnerlist)
	JsonOutput["data"] = itemlist
	this.Data["json"] = JsonOutput    
	this.ServeJSON()
	return
}

func (this *MakepackInfoController) Get(){
	o := orm.NewOrm()
	qs := o.QueryTable("db_packageinfo")
	var packageinfo []*models.DbPackageinfo
	var itemlist []interface{}
	num,_ := qs.All(&packageinfo)
	fmt.Println(num)
	for _,term := range packageinfo{
		var item []interface{}
		item = append(item, term.Taskid, term.Makepackagetime,term.Product, term.Isnewitem, term.Itemname, term.Tryno, term.Packagetype,
			term.Packagemodel, term.Tid1,term.Tid2,term.Tod1,term.Tod2,term.Fixuplive,term.Islokmp,term.Specialfile,term.Localname,term.User,term.Result, term.Installxml, term.Packetxml) 
		itemlist = append(itemlist, item)
	}
	JsonOutput["data"] = itemlist
	this.Data["json"] = JsonOutput    
	this.ServeJSON()
	return
}

func (this *ResultController) Get(){
	var result int
	o := orm.NewOrm()
	taskid, ret := this.GetInt("taskid")
	fmt.Println(ret)
	fmt.Println(taskid)
	pi := models.DbPackageinfo{Taskid : taskid}
	err := o.Read(&pi)
	if err == orm.ErrNoRows{
		fmt.Println("查询不到")
	}else if err == orm.ErrMissPK{
		fmt.Println("找不到主键")
	} else{
		result = pi.Result
	}
	// JsonOutput["data"] = result
	JsonOutput["result"] = result
	this.Data["json"] = JsonOutput    
	this.ServeJSON()
	return
}

func (this *StopController) Post(){
	conf := goini.SetConfig("stconfig.ini")
	text := conf.GetValue("taskkill", "processlist")
	fmt.Println(text)
	var processlist = strings.Split(text, "|")
	for _,term := range processlist{
		// fmt.Println(_)
		fmt.Println(term)
		c := exec.Command("taskkill.exe", "/f", "/im", term)
		err := c.Start()
		fmt.Println(err)
		if err == nil{
			fmt.Println("杀进程失败")
		}
	}
	this.Data["json"] = JsonOutput  
	this.ServeJSON()
	return
}