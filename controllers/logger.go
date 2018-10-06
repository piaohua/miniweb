package controllers

import (
	"time"

	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/mailbox"
	"github.com/AsynkronIT/protoactor-go/router"
	"github.com/astaxie/beego"
	"github.com/gogo/protobuf/proto"
)

var (
	// LoggerPid logger pid
	LoggerPid *actor.PID
)

const maxConcurrency = 5

//LoggerActor 日志记录服务
type LoggerActor struct {
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *LoggerActor) Receive(ctx actor.Context) {
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
	case proto.Message:
		a.Handler(msg, ctx)
	default:
		beego.Error("unknown message ", msg)
	}
}

func newLoggerActor() actor.Actor {
	a := new(LoggerActor)
	return a
}

//newLoggerProps 启动
func newLoggerProps() *actor.Props {
	return router.NewRoundRobinPool(maxConcurrency).
		WithProducer(newLoggerActor).
		WithMailbox(mailbox.Unbounded())
}

// init ...
func initLogger() *actor.PID {
	props := newLoggerProps()
	return actor.Spawn(props) //启动一个进程
}

//NewLogger 启动服务
func NewLogger() {
	LoggerPid = initLogger()
	beego.Info("LoggerPid: ", LoggerPid.String())
}

//StopLogger 停止服务
func StopLogger() {
	beego.Info("stop LoggerPid: ", LoggerPid.String())
	LoggerPid.Tell(new(pb.ServeStop))
}

//CloseLogger 关闭服务消息
func CloseLogger() {
	beego.Info("close LoggerPid: ", LoggerPid.String())
	timeout := 5 * time.Second
	msg := new(pb.ServeClose)
	res, err := LoggerPid.RequestFuture(msg, timeout).Result()
	if err != nil {
		beego.Error("close Logger error: ", err)
	}
	if response, ok := res.(*pb.ServeClosed); ok {
		beego.Info("close Logger response: ", response)
	} else {
		beego.Error("close Logger res: ", res)
	}
}
