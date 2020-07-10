package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {

	beego.LoadAppConfig("ini", "..\\conf\\app.conf")
	sqlConn := beego.AppConfig.String("sqlconn")

	fmt.Print(sqlConn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlConn, 30)

	orm.RegisterModel(new(Order))
	orm.RegisterModel(new(Sku))
	orm.Debug = true
}
