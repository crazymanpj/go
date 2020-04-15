package models

import(
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
	"reflect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// "errors"
)

type Gameinfo struct{
	Id            int `orm:"column(id);pk;auto"`
	Name          string
	Gametype  	  string
	Platform      string
	Clientver     string
	Gamever		  string
	Isenable	  int
}

type Taskinfo struct{
	Id				int64 `orm:"pk"`
	Gameid			int
	Taskname		string
	Phone		    string
	Tester			string
	Time 			time.Time
	Interval		int
	Frequency 		int
}

type Perfdata struct{
	Id				int `orm:"column(id);pk;auto"`
	Taskid			int64
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

type Perfdatashow struct{
	Id				int `orm:"column(id);pk;auto"`
	Taskid			int64
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

type Starttimeinfo struct{
	Id				int `orm:"pk"`
	Gameid			int
	starttime		float32
	Phone		    string
}

func RegisterDB(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	USER, PASS, MYSQL_URLS, MYSQLDB := beego.AppConfig.String("mysqluser"),beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"), beego.AppConfig.String("mysqldb")
	beego.AppConfig.String("mysqluser")
	beego.AppConfig.String("mysqlpass")
	beego.AppConfig.String("mysqlurls")
	beego.AppConfig.String("mysqldb")
	mysqlurl := USER + ":" + PASS + "@tcp(" + MYSQL_URLS + ":3306)/" + MYSQLDB + "?charset=utf8&loc=Local"
	fmt.Println(mysqlurl)
	orm.RegisterDataBase("default", "mysql", mysqlurl, maxIdle, maxConn)
	orm.RegisterModel(new(Gameinfo), new(Taskinfo), new(Perfdata), new(Perfdatashow))
}

//gameinfo
func AddGameinfo(g Gameinfo) int64 {
	o := orm.NewOrm()
	id,err := o.Insert(&g)
	if err == nil{
		fmt.Println(id)
	}else{
		fmt.Println("插入失败")
		fmt.Println(err)
		return -1
	}
	return id
}

func AddTaskinfo(t Taskinfo) (int64, error) {
	var merr error
	o := orm.NewOrm()
	id,err := o.Insert(&t)
	if err == nil{
		fmt.Println(id)
		return id,nil
	}else{
		fmt.Println("插入失败")
		fmt.Println(err)
		merr =  err
		return 1,merr
	}

}

//taskinfo
func GetAllTaskinfo()(string, error){
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select a.id, a.gameid, a.taskname, b.name, phone, tester, frequency, `interval`, time, platform from taskinfo a LEFT JOIN gameinfo b ON a.gameid=b.id").Values(&maps)
	if err == nil && num > 0{
		result,_ := json.Marshal(maps)
		ret :=string(result)
		return ret,nil
	}else{
		return "", err
	}
}


func GetAllGameinfo()(string, error){
	fmt.Println("getallgameinfo")
	o := orm.NewOrm()
	var maps []orm.Params
	num,_ := o.Raw("SELECT * from gameinfo").Values(&maps)
	fmt.Println(num)
	// fmt.Println(maps)
	result,_ := json.Marshal(maps)
	ret :=string(result)
	// fmt.Println(ret)
	fmt.Println(reflect.TypeOf(ret))
	return ret,nil
}

//perfdata
func AddPerfData(p Perfdata) (int64, error) {
	var merr error
	o := orm.NewOrm()
	id,err := o.Insert(&p)
	if err == nil{
		return id,nil
	}else{
		fmt.Println("插入失败")
		fmt.Println(err)
		merr =  err
		return 1,merr
	}

}

func AddPerfDatashow(p Perfdatashow) (int64, error) {
	var merr error
	o := orm.NewOrm()
	id,err := o.Insert(&p)
	if err == nil{
		return id,nil
	}else{
		fmt.Println("插入失败")
		fmt.Println(err)
		merr =  err
		return 1,merr
	}

}

func GetAllPerfdata()(string,error){
	logs.Debug("GetAllPerfdata")
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("select * from perfdata").Values(&maps)
	result,merr := json.Marshal(maps)
	if merr == nil{
		ret :=string(result)
		return ret,nil
	}else{
		return "",nil
	}
}

func GetAllPerfdatashow()(string,error){
	logs.Debug("GetAllPerfdatashow")
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("select * from perfdatashow").Values(&maps)
	result,merr := json.Marshal(maps)
	if merr == nil{
		ret :=string(result)
		return ret,nil
	}else{
		return "",nil
	}
}

func UpdateStarttime(taskid int64, starttime float32)(error){
	logs.Debug("UpdateStarttime")
	o := orm.NewOrm()
	// var perfdata Perfdata
	logs.Debug(taskid)
	logs.Debug(starttime)
	num,err:=o.Raw("UPDATE perfdata SET starttime = ? where taskid = ?", starttime, taskid).Exec()
	logs.Debug(num)
	if err == nil{
		return nil
	}else{
		fmt.Println("插入失败")
		fmt.Println(err)
		return err
	}

}