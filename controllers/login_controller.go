package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"vue-admin-go/vo"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	var loginRequest vo.LoginRequest
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &loginRequest)

	loginResponse := vo.LoginResponse{}
	if err != nil {
		loginResponse.Success = false
		loginResponse.Message = "用户名或密码错误"
	} else {
		loginResponse.Token = uuid.NewV4().String()
		loginResponse.Success = true
	}

	this.Data["json"] = loginResponse
	this.ServeJSON()
}
