/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-09-03 15:09:22
 * Filename      : ms_hander.go
 * Description   : ms actor handler
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
func (ms *MSActor) Handler(msg interface{}, ctx actor.Context) {
	switch arg := msg.(type) {
	case *pb.ServeClose:
		beego.Debug("ms ServeClose ", arg)
		//断开连接
		ms.stop(ctx)
	case *pb.ServeStop:
		beego.Debug("ms ServeStop ", arg)
		//断开连接
		ms.stop(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStoped:
	case *pb.ServeStart:
		beego.Debug("ms ServeStart ", arg)
		ms.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStarted:
	case *pb.Tick:
		ms.ding(ctx)
	case *pb.Login:
		ms.online[arg.GetUserid()] = arg.GetWSPid()
		ms.userids[arg.GetSession()] = arg.GetUserid()
		//响应
		//rsp := new(pb.Logouted)
		//rsp.Userid = arg.GetUserid()
		//ctx.Respond(rsp)
	case *pb.Logout:
		delete(ms.online, arg.GetUserid())
		delete(ms.userids, arg.GetSession())
		//响应
		//rsp := new(pb.Logouted)
		//rsp.Userid = arg.GetUserid()
		//ctx.Respond(rsp)
	case *pb.ChangeCurrency:
		if v, ok := ms.online[arg.GetUserid()]; ok {
			v.Tell(arg)
		} else {
			models.UpdateCurrency(arg)
		}
	case proto.Message:
		//TODO 响应消息
	default:
		beego.Error("unknown message ", arg)
	}
}

func (ms *MSActor) start(ctx actor.Context) {
	beego.Info("ms start: ", ctx.Self().String())
	//启动时钟
	go ms.ticker(ctx)
}

func (ms *MSActor) stop(ctx actor.Context) {
	beego.Info("ms stop: ", ctx.Self().String())
	beego.Info("ms userids: ", ms.userids)
	msg := new(pb.ServeStop)
	for k, v := range ms.online {
		beego.Info("ms online userid: ", k, ", pid: ", v.String())
		v.Tell(msg)
	}
	ctx.Self().Stop()
}

//1秒同步一次
func (ms *MSActor) ding(ctx actor.Context) {
}

//时钟
func (ms *MSActor) ticker(ctx actor.Context) {
	tick := time.Tick(time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-ms.stopCh:
			beego.Info("ms ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-ms.stopCh:
			beego.Info("ms ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}
