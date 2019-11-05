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
	"time"
	"encoding/json"
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

type MakePackController struct{
	beego.Controller
}

type db_packageinfo struct{
	taskid				int
}

var JsonOutput = map[string]interface{}{
	"errorcode" : 0,
	"msg" : "",
	// "data" : map[string]interface{}{},
	"version" : "1.0",
}

func init(){
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
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
	JsonOutput["productlist"] = productlist
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
	product := this.GetString("product")
	fmt.Println(product)
	var partnerlist []*models.DbPartnerlist
	var itemlist []interface{}
	num,_ := qs.Filter("status", 1).Filter("product", product).All(&partnerlist)
	// for i := 0; i <num; i++{
	// 	itemlist = append(itemlist, partnerlist[i].Partner)
	// }
	for _,term:= range partnerlist{
		fmt.Println(term.Id,term.Partner, num)
		itemlist = append(itemlist, term.Partner)
	}
	fmt.Println(partnerlist)
	JsonOutput["partnerlist"] = itemlist
	this.Data["json"] = JsonOutput    
	this.ServeJSON()
	return
}

func (this *MakepackInfoController) Get(){
	o := orm.NewOrm()
	// qs := o.QueryTable("db_packageinfo")
	var packageinfo []*models.DbPackageinfo
	m_packageinfo := new(models.DbPackageinfo)
	num,_:= o.QueryTable(m_packageinfo).OrderBy("makepackagetime").All(&packageinfo)
	fmt.Println(num)
	var itemlist []interface{}
	// num,_ := qs.All(&packageinfo)
	// 也可以直接使用对象作为表名
	// user := new(User)
	// qs = o.QueryTable(user) // 返回 QuerySeter
	// fmt.Println(num)
	for _,term := range packageinfo{
		// var item map[string]interface{}
		item := map[string]interface{}{"taskid" :term.Taskid, "makepackagetime" :term.Makepackagetime, "product" : term.Product, "isnewitem" : term.Isnewitem,
	"Itemname" : term.Itemname, "tryno" : term.Tryno, "packagetype":term.Packagetype, "packagemodel":term.Packagemodel, "tid1":term.Tid1, "tid2": term.Tid2,
	"tod1":term.Tod1, "tod2":term.Tod2,"fixuplive":term.Fixuplive,"islokmp":term.Islokmp,"specialfile":term.Specialfile,"localname":term.Localname,"user":term.User,
	"result":term.Result, "installxml":term.Installxml, "packetxml":term.Packetxml}
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
		var process string
		process = term
		c := exec.Command("taskkill.exe", "/f", "/im", process)
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

func (this *MakePackController) Post(){
	// product := this.GetString("product")
	// isnewitem,err := this.GetInt("isnewitem")
	// itemname := this.GetString("itemname")
	// tryno := this.GetString("tryno")
	// packagetype := this.GetString("packagetype")
	// packetmodel := this.GetString("packetmodel")
	// tid1 := this.GetString("tid1")
	// tid2 := this.GetString("tid2")
	// tod1 := this.GetString("tod1")
	// tod2  := this.GetString("tod2")
	// fixuplive,err := this.GetInt("fixuplive")
	// islokmp,err := this.GetInt("islokmp")
	// specialfile := this.GetString("specialfile")
	// localname := this.GetString("localname")
	// user := this.GetString("user")
	// result,err := this.GetInt("result")
	// installxml := this.GetString("installxml")
	// packetxml := this.GetString("packetxml")
	// mtime := time.Now().Format("2006-01-02 15:04:05")

	o := orm.NewOrm()
	var packageinfo models.DbPackageinfo
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &packageinfo)
	if err != nil{
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(packageinfo.Product)
	packageinfo.Makepackagetime = time.Now()
	// packageinfo.Makepackagetime = time.Now()
	// packageinfo.Product = product
	// packageinfo.Isnewitem = isnewitem
	// packageinfo.Itemname = itemname
	// packageinfo.Tryno = tryno
	// packageinfo.Packagetype = packagetype
	// packageinfo.Packagemodel = packetmodel
	// packageinfo.Tid1 = tid1
	// packageinfo.Tid2 = tid2
	// packageinfo.Tod1 = tod1
	// packageinfo.Tod2 = tod2
	// packageinfo.Fixuplive = fixuplive
	// packageinfo.Islokmp = islokmp
	// packageinfo.Specialfile = specialfile
	// packageinfo.Localname = localname
	// packageinfo.User = user
	// packageinfo.Result = result
	// packageinfo.Installxml = installxml
	// packageinfo.Packetxml = packetxml

	id, err := o.Insert(&packageinfo)
	if err == nil{
		fmt.Println(id)
	}else{
		fmt.Println("插入失败")
		fmt.Println(err)
	}

	this.Data["json"] = JsonOutput  
	JsonOutput["msg"] = "ok"
	filename := "D:\\svn\\DubaTest\\Tools\\makedata\\makedata.exe"
	argvs := fmt.Sprintf("-product=%s -isnewitem=%d -itemname=%s -tryno=%s -packettype=%s -packetmodel=%s -tid1=%s -tid2=%s -tod1=%s -tod2=%s -fixuplive=%d -islokmp=%d -specialfile=%s -localname=%s -taskid=%d", 
	packageinfo.Product, packageinfo.Isnewitem, packageinfo.Itemname, packageinfo.Tryno, packageinfo.Packagetype, packageinfo.Packagemodel, packageinfo.Tid1, packageinfo.Tid2, packageinfo.Tod2, packageinfo.Tod2, 
	packageinfo.Fixuplive, packageinfo.Islokmp, packageinfo.Specialfile, packageinfo.Localname, packageinfo.Taskid)
	fmt.Println(filename)
	fmt.Println(argvs)
	c := exec.Command(filename, argvs)
	err = c.Start()
	fmt.Println(err)
	if err == nil{
		fmt.Println(err)
		fmt.Println("成功")
	}else{
		fmt.Println(err)
		fmt.Println("失败")
	}
	this.ServeJSON()
	return
}