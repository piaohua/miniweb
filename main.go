/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-27 21:26:59
 * Filename      : main.go
 * Description   : 主文件
 * *******************************************************/

package main

import (
	"miniweb/controllers"
	"miniweb/models"
	_ "miniweb/routers"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	APP_VER = "0.1.1.0021"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	//log file config
	beego.AppConfig.Set("version", APP_VER)
	if models.RunMode() {
		beego.SetLevel(beego.LevelInformational)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log_file")+`"}`)
		beego.BeeLogger.DelLogger("console")
	} else {
		beego.SetLevel(beego.LevelDebug)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log_file")+`"}`)
	}

	//mongodb init
	dbHost := beego.AppConfig.String("mdb.host")
	dbPort := beego.AppConfig.String("mdb.port")
	dbUser := beego.AppConfig.String("mdb.user")
	dbPassword := beego.AppConfig.String("mdb.password")
	dbName := beego.AppConfig.String("mdb.name")
	models.InitMgo(dbHost, dbPort, dbUser, dbPassword, dbName)

	//var init
	models.GenIDChInit()
	models.SessionChInit()
	//models.HealthCheck()
	//models.RunStatics()
	//models.RunTask()
	models.GenCodeChInit()

	//init cache
	models.InitPropList()
	models.InitShopList()
	models.InitLoginPrizeList()
	models.InitGateList()
	models.InitShareList()
	models.InitInviteList()

	//init pid
	controllers.NewRemote()
	controllers.NewMS()
	controllers.NewLogger()

	beego.Run()
}
