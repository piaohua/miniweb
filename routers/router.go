/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-27 21:26:59
 * Filename      : router.go
 * Description   : 路由
 * *******************************************************/

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
