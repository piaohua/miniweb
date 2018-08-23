package routers

import (
	"miniweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/code", &controllers.MainController{}, "post:Code")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/login", &controllers.WebSocketController{}, "get:Login")
}
