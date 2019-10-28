package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"dbmakepack/models"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
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
	o.Using("default")
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
	o.Using("default")
	qs := o.QueryTable("db_partnerlist")
	var partnerlist []*models.DbPartnerlist
	var itemlist []interface{}
	num,_ := qs.Filter("status", 1).All(&partnerlist)
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