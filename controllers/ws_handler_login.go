/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-27 21:36:19
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
		//响应
		if ws.online {
			ws.handlerLogined(arg, ctx)
			//ws.Send(arg)
		}
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
	switch arg := msg.(type) {
	case *pb.CUserData:
		beego.Debug("CUserData ", arg)
		s2c := new(pb.SUserData)
		s2c.Data = &pb.UserData{
			Userid:    ws.user.ID,
			NickName:  ws.user.NickName,
			AvatarUrl: ws.user.AvatarUrl,
			Gender:    ws.user.Gender,
		}
		ws.Send(s2c)
	case *pb.CGameData:
		beego.Debug("CGameData ", arg)
		s2c := new(pb.SGameData)
		s2c.NextInfo = arg.GetGameInfo()
		s2c.NextInfo.Gate++
		ws.Send(s2c)
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
		s2c.Error = pb.LoginFaild
		ws.Send(s2c)
		return
	}
	if user == nil {
		s2c.Error = pb.LoginFaild
		ws.Send(s2c)
		return
	}
	ws.user = user
	s2c.Userid = user.ID
	ws.Send(s2c)
	//成功后处理
	ws.logined(user.ID, false, ctx)
}

//登录成功处理
func (ws *WSConn) logined(userid string, isRegist bool,
	ctx actor.Context) {
	//登录成功消息
	//msg := new(pb.LoginSuccess)
	//msg.IsRegist = isRegist
	//msg.Ip = ws.GetIPAddr()
	//msg.Userid = userid
	//msg.WsPid = ctx.Self()
	//pid已经切换为rsPid
	//ws.pid.Tell(msg)
	//登录成功
	ws.online = true
	//成功
	ctx.SetReceiveTimeout(0) //login Successfully, timeout off
	beego.Info("login success: ", userid)
}

//普通登录验证
func (ws *WSConn) login(arg *pb.CLogin, ctx actor.Context) {
	s2c := new(pb.SLogin)
	if models.RunMode() {
		beego.Error("login runmode")
		s2c.Error = pb.LoginFaild
		ws.Send(s2c)
		return
	}
	user, err := models.VerifyUserLogin(arg, ws.session)
	beego.Info("login user: ", user)
	if err != nil {
		beego.Error("login err: ", err)
		s2c.Error = pb.LoginFaild
		ws.Send(s2c)
		return
	}
	if user == nil {
		s2c.Error = pb.LoginFaild
		ws.Send(s2c)
		return
	}
	ws.user = user
	s2c.Userid = user.ID
	ws.Send(s2c)
	//成功后处理
	ws.logined(user.ID, false, ctx)
}
