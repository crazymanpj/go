package models

import(
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Product struct{
	id            int
	name          string
	description   string
	addtime       time.Time
	status        int
}

type DbPartnerlist struct{
	Id				int `orm:"column(id);pk"`
	Partner			string
	Description		string
	Addtime 		time.Time
	Status			int
	Product 		string
}

type DbTrynolist struct{
	Id				int `orm:"column(id);pk"`
	Tryno			int
	Description		string
	Addtime 		time.Time
	Status			int
	Product 		string
}

type DbPackageinfo struct{
	Taskid				int `orm:"column(taskid);pk"`
	Makepackagetime		time.Time
	Product 			string
	Isnewitem			int
	Itemname			string
	Tryno				string
	Packagetype			string
	Packagemodel		string
	Tid1				string
	Tid2				string
	Tod1				string
	Tod2				string
	Fixuplive			int
	Islokmp				int
	Specialfile			string
	Localname			string
	User				string
	Result				int
	Installxml			string
	Packetxml			string
}


func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "testgo:testgo123@tcp(127.0.0.1:3306)/dbmakepack?charset=utf8", maxIdle, maxConn)
	orm.RegisterModel(new(DbTrynolist))
	orm.RegisterModel(new(DbPartnerlist), new(DbPackageinfo))
}