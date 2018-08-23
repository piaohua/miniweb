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

	//mongodb init
	dbHost := beego.AppConfig.String("mdb.host")
	dbPort := beego.AppConfig.String("mdb.port")
	dbUser := beego.AppConfig.String("mdb.user")
	dbPassword := beego.AppConfig.String("mdb.password")
	dbName := beego.AppConfig.String("mdb.name")
	models.InitMgo(dbHost, dbPort, dbUser, dbPassword, dbName)

	beego.Run()
}
