package controllers

import (
	"github.com/astaxie/beego"
	orm2 "github.com/astaxie/beego/orm"
	"vue-admin-go/models"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) OrderPage() {

	//pageNo := this.GetString("pageNo")
	//pageSize := this.GetString("pageSize")
	//orderNo := this.GetString("orderNo")
	//status := this.GetString("status")
	//
	//if pageNo == nil {
	//	pageNo := 1
	//}
	//if pageSize == nil{
	//	pageSize := 10
	//}

	orm := orm2.NewOrm()
	orderTabel := new(models.Order)
	var orders []models.Order
	orm.QueryTable(orderTabel).All(&orders)

	this.Data["json"] = orders
	this.ServeJSON()
}
