/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-29 00:28:48
 * Filename      : ws_hander_login.go
 * Description   : login handler
 * *******************************************************/

package controllers

import (
	"time"

	"miniweb/models"
	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/astaxie/beego"
	"github.com/gogo/protobuf/proto"
)

//玩家数据请求处理
func (ws *WSConn) handlerLogin(msg interface{}, ctx actor.Context) {
	switch arg := msg.(type) {
	case *pb.CPing:
		ws.Send(&pb.SPing{Time: time.Now().Unix()})
	case *pb.CWxLogin:
		beego.Debug("CWxLogin ", arg)
		ws.wxlogin(arg, ctx)
	case *pb.CLogin:
		beego.Debug("CLogin ", arg)
		ws.login(arg, ctx)
	case proto.Message:
		ws.handlerLogined(arg, ctx)
	default:
		beego.Error("unknown message ", arg)
	}
}

//玩家登录后数据请求处理
func (ws *WSConn) handlerLogined(msg interface{}, ctx actor.Context) {
	if !ws.online {
		return
	}
	if ws.user == nil {
		return
	}
	beego.Debug("userid ", ws.user.ID, " msg ", msg)
	switch arg := msg.(type) {
	case *pb.CUserData:
		ws.getUserData(arg)
	case *pb.CGateData:
		ws.getGateData()
	case *pb.CPropData:
		ws.getPropData()
	case *pb.CGetCurrency:
		ws.getCurrency()
	case *pb.CShop:
		ws.getShopData()
	case *pb.CTempShop:
		ws.getTempShopData(arg)
	case *pb.CBuy:
		ws.buy(arg)
	case *pb.COverData:
		ws.overData(arg)
	case *pb.CCard:
		ws.card(arg)
	case *pb.CLoginPrize:
		ws.loginPrize(arg)
	case *pb.CUseProp:
		ws.useProp(arg)
	case *pb.CStart:
		ws.gameStart(arg)
	case *pb.ChangeCurrency:
		ws.change(arg)
	case proto.Message:
		//响应
		ws.Send(arg)
	default:
		beego.Error("unknown message ", arg)
	}
}

//微信登录验证
func (ws *WSConn) wxlogin(arg *pb.CWxLogin, ctx actor.Context) {
	s2c := new(pb.SWxLogin)
	user, err := models.VerifyUserInfo(arg, ws.session)
	beego.Info("wxlogin user: ", user)
	if err != nil {
		beego.Error("wxlogin err: ", err)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	if user == nil {
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	ws.user = user
	s2c.Userid = user.ID
	ws.Send(s2c)
	//成功后处理
	ws.logined(user.ID, ctx)
}

//登录成功处理
func (ws *WSConn) logined(userid string, ctx actor.Context) {
	ctx.SetReceiveTimeout(0) //login Successfully, timeout off
	//登录成功
	ws.online = true
	beego.Info("login success: ", userid)
	ws.user.LoginIP = ws.GetIPAddr()
	//ws.user.LoginTime = time.Now()
	//初始化
	models.PropInit(ws.user)
	models.GateInit(ws.user)
	//精力恢复
	msg := models.CheckEnergy(ws.user)
	if msg != nil {
		ws.Send(msg)
	}
	//更新连续登录
	ws.loginPrizeInit()
	//管理消息
	msg2 := &pb.Login{
		Userid:  ws.user.ID,
		Session: ws.session,
		WSPid:   ws.pid,
	}
	MSPid.Request(msg2, ctx.Self())
}

//普通登录验证
func (ws *WSConn) login(arg *pb.CLogin, ctx actor.Context) {
	s2c := new(pb.SLogin)
	if models.RunMode() {
		beego.Error("login runmode")
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	user, err := models.VerifyUserLogin(arg, ws.session)
	beego.Info("login user: ", user)
	if err != nil {
		beego.Error("login err: ", err)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	if user == nil {
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	ws.user = user
	s2c.Userid = user.ID
	ws.Send(s2c)
	//成功后处理
	ws.logined(user.ID, ctx)
}
