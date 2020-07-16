package websocket

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
	"vue-admin-go/utils"
	"vue-admin-go/vo"
)

type WebsocketService struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}
var connectMap = make(map[*websocket.Conn]bool)

func (this *WebsocketService) Connect() {
	log := utils.GetLogger()

	connect, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		panic("客户端建立连接错误 " + err.Error())
	}

	log.Info("客户端连接建立 " + connect.RemoteAddr().String())

	connectMap[connect] = true

	go sendMessage(connect)
}

func sendMessage(connect *websocket.Conn) {

	for {
		time.Sleep(time.Second * 1)

		lineResponse := vo.LineResponse{
			utils.RandomInt(200, 500),
			utils.RandomInt(200, 500),
			utils.RandomInt(200, 500),
			time.Now().Unix() * 1000,
		}

		message, _ := json.Marshal(lineResponse)

		log.Info(" 发送到客户端的消息：", string(message))

		err := connect.WriteMessage(1, message)
		if err != nil {
			log.Info("客户端连接关闭")
			connect.Close()
			delete(connectMap, connect)
			break
		}
		log.Info("发送消息成功")
	}

}
