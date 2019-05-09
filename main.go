package main

import (
	_ "myweb/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"myweb/models"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}
func main() {
	beego.Run()
}

