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

	// Show
	beego.Router("/show/shop", &controllers.ShowController{}, "get:Shop")
	beego.Router("/show/prize", &controllers.ShowController{}, "get:Prize")
	beego.Router("/show/prop", &controllers.ShowController{}, "get:Prop")

	// Set
	beego.Router("/set/shop", &controllers.SetController{}, "post:Shop")
	beego.Router("/set/prize", &controllers.SetController{}, "post:Prize")
	beego.Router("/set/prop", &controllers.SetController{}, "post:Prop")
}
