package controllers

import (
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
}

func (ws *WSConn) stop(ctx actor.Context) {
	beego.Info("ws stop: ", ctx.Self().String())
	//断开连接
	ws.Close()
	//表示已经断开
	ws.online = false
	ctx.Self().Stop()
}
