package main

import (
	"vue-admin-go/models"
	_ "vue-admin-go/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
