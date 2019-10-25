package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"dbmakepack/models"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
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

type DbTrynolist struct{
	Id				int `orm:"column(id);pk"`
	Tryno			int
	Description		string
	Addtime 		time.Time
	Status			int
	Product 		string
}


var JsonOutput = map[string]interface{}{
	"errorcode" : 0,
	"msg" : "",
	"data" : map[string]interface{}{},
	"version" : "1.0",
}

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "testgo:testgo123@tcp(127.0.0.1:3306)/dbmakepack?charset=utf8", maxIdle, maxConn)
	fmt.Println("models is inited!")
	orm.RegisterModel(new(DbTrynolist))
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func printSlice(x []interface{}){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }



func (this *ProductController) Get(){
	//var o orm.Ormer
	// fmt.Println("test")
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// //最大空闲链接
	// maxIdle := 30
	// maxConn := 30
	// //orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)
	// _ = orm.RegisterDataBase("default", "mysql", "testgo:testgo123@tcp(127.0.0.1:3306)/dbmakepack?charset=utf8", maxIdle, maxConn)
	o := orm.NewOrm()
	var maps []orm.Params
	var productlist []interface{}
	num,_:=o.Raw("SELECT * FROM db_productlist").Values(&maps)
	for _,term := range maps{
		fmt.Println(term["id"], ":", term["name"], num)
		productlist = append(productlist, term["name"])
	}
	printSlice(productlist)
	JsonOutput["data"] = productlist
	this.Data["json"] = JsonOutput                     // json对象
    this.ServeJSON()
    return
}

func (this *TrynoController) Get(){
	// maxIdle := 30
	// maxConn := 30
	// orm.RegisterDataBase("default", "mysql", "testgo:testgo123@tcp(127.0.0.1:3306)/dbmakepack?charset=utf8", maxIdle, maxConn)
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("db_trynolist")
	var trynolist []*DbTrynolist
	var itemlist []interface{}
	num,_ := qs.Filter("status", 1).All(&trynolist)
	for _,term:= range trynolist{
		fmt.Println(term.Id,term.Tryno, num)
		itemlist = append(itemlist, term.Tryno)
	}
	fmt.Println(trynolist)
	JsonOutput["data"] = itemlist
	this.Data["json"] = JsonOutput                     // json对象
	this.ServeJSON()
	return
}