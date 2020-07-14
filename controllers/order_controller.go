package controllers

import (
	"github.com/astaxie/beego"
	"vue-admin-go/models"
	"vue-admin-go/repository"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) OrderPage() {

	var pageNo, _ = this.GetInt("pageNo")
	var pageSize, _ = this.GetInt("pageSize")
	orderNo := this.GetString("orderNo")
	status := this.GetString("status")

	var orders []models.Order

	sqlParam := make(map[string]interface{})
	if orderNo != "" {
		sqlParam["order_no"] = orderNo
	}
	if status != "" {
		sqlParam["status"] = status
	}

	pageResponse := repository.PageQuery(pageNo, pageSize, sqlParam, "t_order", &orders)

	this.Data["json"] = pageResponse
	this.ServeJSON()
}
