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
	beego.Router("/code", &controllers.MainController{}, "options:Code")
	beego.Router("/code", &controllers.MainController{}, "get:Code")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/login", &controllers.WebSocketController{}, "get:Login")

	// Show
	beego.Router("/show/shop", &controllers.ShowController{}, "get:Shop")
	beego.Router("/show/prize", &controllers.ShowController{}, "get:Prize")
	beego.Router("/show/prop", &controllers.ShowController{}, "get:Prop")
	beego.Router("/show/gate", &controllers.ShowController{}, "get:Gate")
	beego.Router("/show/share", &controllers.ShowController{}, "get:Share")
	beego.Router("/show/invite", &controllers.ShowController{}, "get:Invite")

	// Set
	beego.Router("/set/shop", &controllers.SetController{}, "post:Shop")
	beego.Router("/set/prize", &controllers.SetController{}, "post:Prize")
	beego.Router("/set/prop", &controllers.SetController{}, "post:Prop")
	beego.Router("/set/coin", &controllers.SetController{}, "post:Coin")
	beego.Router("/set/diamond", &controllers.SetController{}, "post:Diamond")
	beego.Router("/set/close", &controllers.SetController{}, "post:Close")
	beego.Router("/set/gate", &controllers.SetController{}, "post:Gate")
	beego.Router("/set/share", &controllers.SetController{}, "post:Share")
	beego.Router("/set/invite", &controllers.SetController{}, "post:Invite")
}
