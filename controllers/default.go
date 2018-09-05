package controllers

import (
	"strings"

	"miniweb/models"
	"miniweb/pb"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type MainController struct {
	baseController      // Embed to use methods that are implemented in baseController.
	token          bool // have token
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//
	json := make(map[string]interface{}, 0)
	json["errcode"] = int(pb.Failed)
	json["errmsg"] = "failed"
	c.jsonResult(json)
}

// Code method handles POST requests for AppController.
func (this *MainController) Code() {
	jsonData := &models.SessionResult{}
	defer this.jsonResult(jsonData)

	// Get form value.
	jscode := this.GetString("js_code")

	// Check valid.
	if len(jscode) == 0 {
		jsonData.WxErr.ErrCode = int(pb.Failed)
		jsonData.WxErr.ErrMsg = "parameter error"
		return
	}
	beego.Info("get session by jscode: " + jscode)

	if models.RunMode() {
		jsonData.WxErr.ErrCode = int(pb.Failed)
		jsonData.WxErr.ErrMsg = "mode closed"
		return
	}

	if !this.isPost() {
		jsonData.WxErr.ErrCode = int(pb.Failed)
		jsonData.WxErr.ErrMsg = "method error"
		return
	}

	if !this.token {
		jsonData.WxErr.ErrCode = int(pb.Failed)
		jsonData.WxErr.ErrMsg = "token error"
		return
	}

	//test TODO 控制频率
	ip := this.getClientIp()
	session, err := models.GetSessionByCode(jscode, ip)
	if err != nil {
		jsonData.WxErr.ErrCode = int(pb.Failed)
		jsonData.WxErr.ErrMsg = err.Error()
		return
	}

	wsaddr := beego.AppConfig.String("ws.addr")
	jsonData.Session = session
	jsonData.WsAddr = wsaddr + session
}

// Prepare implemented Prepare() method for baseController.
func (this *MainController) Prepare() {
	token := this.Ctx.Request.Header.Get("token")
	setToken := beego.AppConfig.String("set.token")
	beego.Debug("token: ", token, ", setToken: ", setToken)
	if token != "" && token == setToken {
		this.token = true
	}
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
	// Set config
	this.Ctx.Output.Header("X-Powered-By", "miniweb/"+beego.AppConfig.String("version"))
	this.Ctx.Output.Header("X-Author-By", "piaohua")
	// Set header
	if origin := this.Ctx.Request.Header.Get("Origin"); origin != "" {
		this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
		this.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		this.Ctx.Output.Header("Access-Control-Allow-Headers", "*")
	}
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

//获取用户IP地址
func (this *baseController) getClientIp() string {
	if p := this.Ctx.Input.Proxy(); len(p) > 0 {
		return p[0]
	}
	return this.Ctx.Input.IP()
}

// 输出json
func (this *baseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// 是否POST提交
func (this *baseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}
