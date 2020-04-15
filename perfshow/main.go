package main

import (
	_ "perfshow/routers"
	"github.com/astaxie/beego"
	"perfshow/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/logs"
	// "os/exec"
	// "strings"
	// "os"
	// "fmt"
)

// func getCurrentPath() string {
// 	fmt.Println("getCurrentPath")
// 	// s, err := exec.LookPath(os.Args[0])
// 	dir, err := os.Getwd()
// 	fmt.Println(err)
// 	if err == nil{
// 		fmt.Println(dir)
// 		// i := strings.LastIndex(dir, "\\")
// 		// path := string(s[0 : i+1])
// 		return dir
// 	}else{
// 		fmt.Println(err)
// 		return ""
// 	}

// }

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	models.RegisterDB()
	orm.RunCommand()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	logs.SetLogger("file", `{"filename":"logs/test.log"}`)
	logs.Debug("perfshow start")
	beego.Run()
}
