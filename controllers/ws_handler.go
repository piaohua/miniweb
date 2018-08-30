/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-23 13:10:13
 * Filename      : ws_hander.go
 * Description   : ws actor handler
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

//Handler 消息处理
func (ws *WSConn) Handler(msg interface{}, ctx actor.Context) {
	switch arg := msg.(type) {
	case *pb.ServeClose:
		beego.Debug("ws ServeClose ", arg)
		//断开连接
		ws.stop(ctx)
	case *pb.ServeStop:
		beego.Debug("ws ServeStop ", arg)
		//断开连接
		ws.stop(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStoped:
	case *pb.ServeStart:
		beego.Debug("ws ServeStart ", arg)
		ws.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStarted:
	case *pb.Tick:
		ws.ding(ctx)
	case proto.Message:
		//响应消息
		//ws.Send(msg)
		ws.handlerLogin(arg, ctx)
	default:
		beego.Error("unknown message ", arg)
	}
}

func (ws *WSConn) start(ctx actor.Context) {
	beego.Info("ws start: ", ctx.Self().String())
	ctx.SetReceiveTimeout(waitForLogin) //login timeout set
	//启动时钟
	go ws.ticker(ctx)
}

func (ws *WSConn) stop(ctx actor.Context) {
	beego.Info("ws stop: ", ctx.Self().String())
	//断开连接
	ws.Close()
	//表示已经断开
	ws.online = false
	//TODO 优化缓存
	if ws.user != nil {
		ws.user.Save()
	}
	ctx.Self().Stop()
}

//60秒同步一次
func (ws *WSConn) ding(ctx actor.Context) {
	ws.timer++
	if ws.timer != 60 {
		return
	}
	ws.timer = 0
	//TODO 优化
	if ws.user != nil {
		ws.user.Save()
	}
	msg := models.CheckEnergy(ws.user)
	if msg != nil {
		ws.Send(msg)
	}
	if !ws.online {
		return
	}
}

//时钟
func (ws *WSConn) ticker(ctx actor.Context) {
	tick := time.Tick(time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-ws.stopCh:
			beego.Info("ws ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-ws.stopCh:
			beego.Info("ws ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}
