package routers

import (
	"github.com/astaxie/beego"
	"vue-admin-go/controllers"
	"vue-admin-go/websocket"
)

func init() {
	beego.Router("/order-sku/order/page", &controllers.OrderController{}, "get:OrderPage")
	beego.Router("/order-sku/sku/page", &controllers.SkuController{}, "get:SkuPage")
	beego.Router("/wss", &websocket.WebsocketService{}, "get:Connect")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
}
