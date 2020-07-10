package routers

import (
	"github.com/astaxie/beego"
	"vue-admin-go/controllers"
)

func init() {
	beego.Router("/api/order/page", &controllers.OrderController{}, "get:OrderPage")

	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)


	//ns := beego.NewNamespace("v1",
	//		beego.NSNamespace("/api/order/page",
	//			beego.NSInclude(
	//				&controllers.OrderController{},
	//			),
	//		),
	//	)
	//beego.AddNamespace(ns)
}
