/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-10-02 11:26:22
 * Filename      : node.go
 * Description   : node actor
 * *******************************************************/

package controllers

import (
	"time"

	"miniweb/models"
	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/astaxie/beego"
	"github.com/gogo/protobuf/proto"
)

var (
	//NodePid manages pid
	NodePid *actor.PID
)

//NodeActor 管理服务
type NodeActor struct {
	Name string
	//nodes name - pid
	nodes map[string]*actor.PID
	//rooms roomid - room TODO 优化为独立pid
	rooms map[string]*models.Room
	//roles userid - role
	roles map[string]*models.RoomRole
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (n *NodeActor) Receive(ctx actor.Context) {
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
		n.Handler(msg, ctx)
	default:
		beego.Error("unknown message ", msg)
	}
}

func newNodeActor() actor.Actor {
	n := new(NodeActor)
	n.Name = beego.AppConfig.String("node.name")
	n.nodes = make(map[string]*actor.PID)
	n.rooms = make(map[string]*models.Room)
	n.roles = make(map[string]*models.RoomRole)
	n.stopCh = make(chan struct{})
	return n
}

//NewRemote 启动服务
func NewRemote() {
	bind := beego.AppConfig.String("node.bind")
	kind := beego.AppConfig.String("node.kind")
	remote.Start(bind)
	nodeProps := actor.FromProducer(newNodeActor)
	remote.Register(kind, nodeProps)
	var err error
	NodePid, err = actor.SpawnNamed(nodeProps, kind)
	if err != nil {
		beego.Critical("NodePid err: ", err)
	}
	beego.Info("NodePid: ", NodePid.String())
	NodePid.Tell(new(pb.ServeStart))
}

//StopNode 停止服务
func StopNode() {
	beego.Info("stop NodePid: ", NodePid.String())
	NodePid.Tell(new(pb.ServeStop))
}

//CloseNode 关闭服务消息
func CloseNode() {
	beego.Info("close NodePid: ", NodePid.String())
	timeout := 5 * time.Second
	msg := new(pb.ServeClose)
	res, err := NodePid.RequestFuture(msg, timeout).Result()
	if err != nil {
		beego.Error("close node error: ", err)
	}
	if response, ok := res.(*pb.ServeClosed); ok {
		beego.Info("close node response: ", response)
	} else {
		beego.Error("close node res: ", res)
	}
}
