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
	case *pb.LoginedElse:
		ws.logon(arg, ctx)
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
	case *pb.CShareInfo:
		ws.shareInfo(arg)
	case *pb.CInviteInfo:
		ws.inviteInfo(arg)
	case *pb.CShare:
		ws.share(arg)
	case *pb.CInvite:
		ws.invite(arg)
	case *pb.ChangeCurrency:
		ws.change(arg)
	case *pb.LoginElse:
		ws.loginElse(arg, ctx)
	case *pb.Invite:
		models.SetInvite(arg.GetUserid(), ws.user)
	case proto.Message:
		//响应
		ws.Send(arg)
	default:
		beego.Error("unknown message ", arg)
	}
}

//微信登录验证
func (ws *WSConn) wxlogin(arg *pb.CWxLogin, ctx actor.Context) {
	if ws.online {
		beego.Error("wxlogin already")
		s2c := new(pb.SWxLogin)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	userInfo, err := models.VerifyUserInfo(arg, ws.session)
	beego.Debug("wxlogin userInfo: ", userInfo)
	if err != nil {
		beego.Error("wxlogin err: ", err)
		s2c := new(pb.SWxLogin)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	ws.userInfo = userInfo
	//检测别处登录
	ws.loginElseMsg(pb.WXLOGIN, ctx)
}

//检测别处登录
func (ws *WSConn) loginElseMsg(Type pb.LoginType, ctx actor.Context) {
	msg := &pb.LoginElse{
		WSPid:   ws.pid,
		Session: ws.session,
		Type:    Type,
	}
	MSPid.Request(msg, ctx.Self())
}

//别处登录
func (ws *WSConn) loginElse(arg *pb.LoginElse, ctx actor.Context) {
	if arg.GetSession() != ws.session {
		beego.Error("loginElse session error: ", arg)
		return
	}
	//下线消息
	msg := &pb.SLoginOut{
		Type: pb.OUT_TYPE1,
	}
	ws.Send(msg)
	//断开连接
	ws.stop(ctx) //TODO 优化
	//响应消息
	rsp := new(pb.LoginedElse)
	rsp.Session = arg.GetSession()
	rsp.Type = arg.GetType()
	arg.WSPid.Tell(rsp)
}

//验证通过后正式登录
func (ws *WSConn) logon(arg *pb.LoginedElse, ctx actor.Context) {
	user, err := models.LoginUserInfo(ws.userInfo, ws.session, arg.GetType())
	switch arg.GetType() {
	case pb.WXLOGIN:
		s2c := new(pb.SWxLogin)
		if err != nil {
			beego.Error("wxlogin err: ", err)
			s2c.Error = pb.LoginFailed
			ws.Send(s2c)
			return
		}
		s2c.Userid = user.ID
		ws.Send(s2c)
	case pb.CODELOGIN:
		s2c := new(pb.SLogin)
		if err != nil {
			beego.Error("code login err: ", err)
			s2c.Error = pb.LoginFailed
			ws.Send(s2c)
			return
		}
		s2c.Userid = user.ID
		ws.Send(s2c)
	default:
		ws.Close()
		return
	}
	ws.userInfo = nil
	ws.user = user
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
	//初始化
	ws.shareInit()
	ws.inviteInit()
	ws.user.LoginCount++
}

//普通登录验证
func (ws *WSConn) login(arg *pb.CLogin, ctx actor.Context) {
	if models.RunMode() {
		beego.Error("code login runmode")
		s2c := new(pb.SLogin)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	if ws.online {
		beego.Error("code login already")
		s2c := new(pb.SLogin)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	userInfo, err := models.VerifyUserLogin(arg, ws.session)
	beego.Info("code login userInfo: ", userInfo)
	if err != nil {
		beego.Error("code login err: ", err)
		s2c := new(pb.SLogin)
		s2c.Error = pb.LoginFailed
		ws.Send(s2c)
		return
	}
	ws.userInfo = userInfo
	//检测别处登录
	ws.loginElseMsg(pb.CODELOGIN, ctx)
}
