package controllers

import (
	"github.com/astaxie/beego"
	"vue-admin-go/models"
	"vue-admin-go/repository"
)

type SkuController struct {
	beego.Controller
}

func (this *SkuController) SkuPage() {

	var pageNo, _ = this.GetInt("pageNo")
	var pageSize, _ = this.GetInt("pageSize")
	name := this.GetString("name")

	if pageSize == 0{
		panic("参数错误")
	}

	var skus []models.Sku

	sqlParam := make(map[string]interface{})
	if name != "" {
		sqlParam["name"] = name
	}

	pageResponse := repository.PageQuery(pageNo, pageSize, sqlParam, "t_sku", &skus)

	this.Data["json"] = pageResponse
	this.ServeJSON()
}
