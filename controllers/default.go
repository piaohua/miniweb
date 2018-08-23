package controllers

import (
	"strings"

	"miniweb/models"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type MainController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//
	json := make(map[string]interface{}, 0)
	json["errcode"] = 10001
	json["errmsg"] = "failed"
	c.Data["json"] = json
	c.ServeJSON()
}

// Code method handles POST requests for AppController.
func (this *MainController) Code() {
	// Get form value.
	jscode := this.GetString("js_code")

	// Check valid.
	if len(jscode) == 0 {
		this.Redirect("/", 302)
		return
	}
	beego.Info("join jscode: " + jscode)

	wsaddr := beego.AppConfig.String("ws.addr")
	this.Data["json"] = &models.SessionResult{
		Session: "3rd_session",
		WsAddr:  wsaddr + "3rd_session",
	}
	this.ServeJSON()
}

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

// baseController represents base router for all other app routers.
// It implemented some methods for the same implementation;
// thus, it will be embedded into other routers.
type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (this *baseController) Prepare() {
	// Reset language option.
	this.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	// 2. Default language is English.
	if len(this.Lang) == 0 {
		this.Lang = "en-US"
	}

	// Set template level language option.
	this.Data["Lang"] = this.Lang
}
