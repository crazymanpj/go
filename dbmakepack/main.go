package main

import (
	_ "dbmakepack/routers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	fmt.Println("test")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//最大空闲链接
	maxIdle := 30
	maxConn := 30
	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)
	_ = orm.RegisterDataBase("default", "mysql", "testgo:testgo123@tcp(127.0.0.1:3306)/dbmakepack?charset=utf8", maxIdle, maxConn)
	fmt.Println("models is inited!")
	o := orm.NewOrm()
	var maps []orm.Params
	num,_:=o.Raw("SELECT * FROM db_productlist").Values(&maps)
	for _,term := range maps{
		fmt.Println(term["id"], ":", term["name"], num)
	}
	orm.DefaultTimeLoc = time.UTC
	beego.Run()
	fmt.Println("test")
}

