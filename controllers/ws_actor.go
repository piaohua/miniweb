/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-24 12:17:20
 * Filename      : ws_actor.go
 * Description   : ws actor
 * *******************************************************/

package controllers

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/astaxie/beego"
	"github.com/gogo/protobuf/proto"
)

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (ws *WSConn) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		beego.Notice("Starting, initialize actor here")
	case *actor.Stopping:
		beego.Notice("Stopping, actor is about to shut down")
	case *actor.Stopped:
		beego.Notice("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		beego.Notice("Restarting, actor is about to restart")
	case *actor.ReceiveTimeout:
		beego.Info("ReceiveTimeout pid: ", ctx.Self().String())
		//登录超时
		ws.Close()
	case proto.Message:
		ws.Handler(msg, ctx)
	default:
		beego.Error("unknown message ", msg)
	}
}

//初始化
func (ws *WSConn) initWs() *actor.PID {
	props := actor.FromProducer(func() actor.Actor { return ws }) //实例
	return actor.Spawn(props)                                     //启动一个进程
}
