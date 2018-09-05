/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-09-03 14:59:07
 * Filename      : ms_actor.go
 * Description   : manages actor
 * *******************************************************/

package controllers

import (
	"time"

	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/astaxie/beego"
	"github.com/gogo/protobuf/proto"
)

var (
	//MSPid manages pid
	MSPid *actor.PID
)

//MSActor 管理服务
type MSActor struct {
	//ws进程 userid - pid
	online map[string]*actor.PID
	//session - userid
	userids map[string]string
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (ms *MSActor) Receive(ctx actor.Context) {
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
		ms.Handler(msg, ctx)
	default:
		beego.Error("unknown message ", msg)
	}
}

//初始化
func (ms *MSActor) initMS() *actor.PID {
	props := actor.FromProducer(func() actor.Actor { return ms }) //实例
	return actor.Spawn(props)                                     //启动一个进程
}

//NewMS 启动服务
func NewMS() {
	ms := new(MSActor)
	ms.online = make(map[string]*actor.PID)
	ms.userids = make(map[string]string)
	ms.stopCh = make(chan struct{})
	MSPid = ms.initMS()
	beego.Info("MSPid: ", MSPid.String())
	MSPid.Tell(new(pb.ServeStart))
}

//StopMS 停止服务
func StopMS() {
	beego.Info("stop MSPid: ", MSPid.String())
	MSPid.Tell(new(pb.ServeStop))
}

//CloseMS 关闭服务消息
func CloseMS() {
	beego.Info("close MSPid: ", MSPid.String())
	timeout := 5 * time.Second
	msg := new(pb.ServeClose)
	res, err := MSPid.RequestFuture(msg, timeout).Result()
	if err != nil {
		beego.Error("close MS error: ", err)
	}
	if response, ok := res.(*pb.ServeClosed); ok {
		beego.Info("close MS response: ", response)
	} else {
		beego.Error("close MS res: ", res)
	}
}
