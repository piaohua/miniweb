package main

import (
	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/astaxie/beego"
)

//Handler 消息处理
func (a *LoggerActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.ServeStop:
		//关闭服务
		a.handlerStop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	case *pb.ServeClose:
		beego.Debug("ms ServeClose ", arg)
		//响应
		rsp := new(pb.ServeClosed)
		ctx.Respond(rsp)
	default:
		beego.Error("unknown message ", msg)
	}
}

func (a *LoggerActor) handlerStop(ctx actor.Context) {
	beego.Info("handlerStop")
	//TODO clean mailbox
}
