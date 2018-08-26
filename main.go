package main

import (
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

	beego.Run()
}
