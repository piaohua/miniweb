package controllers

import (
	"miniweb/models"
	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/astaxie/beego"
)

//Handler 消息处理
func (a *NodeActor) Handler(msg interface{}, ctx actor.Context) {
	switch arg := msg.(type) {
	case *pb.ServeStop:
		//关闭服务
		a.stop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	case *pb.ServeClose:
		beego.Info("node ServeClose ", msg)
		a.handlerStop(ctx)
		//响应
		rsp := new(pb.ServeClosed)
		ctx.Respond(rsp)
	case *pb.ServeStart:
		beego.Info("node ServeStart ", msg)
		a.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.Logout:
		delete(a.roles, arg.GetUserid())
	case *pb.CFightCreate:
		beego.Info("CFightCreate ", arg)
		a.fightCreate(arg, ctx)
	case *pb.CFightChangeSet:
		beego.Info("CFightChangeSet ", arg)
		a.fightChange(arg, ctx)
	case *pb.CFightEnter:
		beego.Info("CFightEnter ", arg)
		a.fightEnter(arg, ctx)
	case *pb.CFightMatchExit:
		beego.Info("CFightMatchExit ", arg)
		a.fightExit(arg, ctx)
	case *pb.CFightStart:
		beego.Info("CFightStart ", arg)
		a.fightStart(arg, ctx)
	case *pb.CFightingScore:
		beego.Info("CFightingScore ", arg)
		a.fightScore(arg, ctx)
	case *pb.CFightMatch:
		beego.Info("CFightMatch ", arg)
		a.fightMatch(arg, ctx)
	case *pb.CFightingCancelGird:
		beego.Info("CFightingCancelGird ", arg)
		a.fightGird(arg, ctx)
	case *pb.CFightingEnd:
		beego.Info("CFightingEnd ", arg)
		a.fightEnd(arg, ctx)
	default:
		beego.Error("unknown message ", msg)
	}
}

func (a *NodeActor) handlerStop(ctx actor.Context) {
	beego.Info("handlerStop")
	//TODO clean mailbox
	for k, v := range a.nodes {
		beego.Info("node ", k, v)
	}
	for k, v := range a.rooms {
		ok := v.Upsert()
		beego.Info("room ", k, v, ok)
	}
}

func (a *NodeActor) stop(ctx actor.Context) {
	select {
	case <-a.stopCh:
		return
	default:
		//停止发送消息
		close(a.stopCh)
	}
	beego.Info("a stop: ", ctx.Self().String())
	ctx.Self().Stop()
}

func (a *NodeActor) start(ctx actor.Context) {
	beego.Info("node start: ", ctx.Self().String())
	//TODO 连接其它节点
	//Pid = actor.NewPID(bind, kind)
	//name := beego.AppConfig.String("node.remoteName")
	//bind := beego.AppConfig.String("node.remoteBind")
	//kind := beego.AppConfig.String("node.remoteKind")
	//TODO 加载房间
	list := models.GetRoomList(a.Name)
	for k, v := range list {
		a.rooms[v.ID] = &list[k]
	}
}

// fight handler

// create ...
func (a *NodeActor) fightCreate(arg *pb.CFightCreate, ctx actor.Context) {
	rsp := new(pb.SFightCreate)
	room, err := models.CreateRoom(arg, a.Name)
	if err != nil {
		beego.Error("fight create error ", err)
		rsp.Error = pb.FightCreateFailed
		ctx.Respond(rsp)
		return
	}
	rsp.RoomInfo = models.RoomData(room)
	ctx.Respond(rsp)
	//
	a.rooms[room.ID] = room
	a.roles[arg.Userid] = &models.RoomRole{
		Roomid: room.ID,
		Pid:    ctx.Sender(),
	}
}

// change ...
func (a *NodeActor) fightChange(arg *pb.CFightChangeSet, ctx actor.Context) {
	rsp := new(pb.SFightUser)
	room := a.rooms[arg.GetRoomid()]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight change error ", arg)
		//rsp.Error = pb.FightRoomNotExist
		//ctx.Respond(rsp)
		return
	}
	room.Match = int32(arg.GetMatch())
	room.UserProp = int32(arg.GetUserProp())
	if !models.ChangeRoom(arg) {
		beego.Error("fight change failed ", arg)
		//rsp.Error = pb.FightChangeFailed
		//ctx.Respond(rsp)
		return
	}
	rsp.RoomInfo = models.RoomData(room)
	//ctx.Respond(rsp)
	a.broadcast(room.ID, rsp)
}

//TODO broadcast to nodes
func (a *NodeActor) broadcast(roomid string, msg interface{}) {
	if v, ok := a.rooms[roomid]; ok {
		for _, v := range v.User {
			a.send2userid(v.Userid, msg)
		}
	}
}

//TODO broadcast to nodes
func (a *NodeActor) send2userid(userid string, msg interface{}) {
	if v, ok := a.roles[userid]; ok {
		v.Pid.Tell(msg)
	}
}

// exit ...
func (a *NodeActor) fightExit(arg *pb.CFightMatchExit, ctx actor.Context) {
	rsp := new(pb.SFightMatchExit)
	roomid := a.userid2roomid(arg.GetUserid())
	room := a.rooms[roomid]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight exit error ", arg)
		rsp.Error = pb.FightNotInRoom
		ctx.Respond(rsp)
		return
	}
	for k, v := range room.User {
		if v.Userid == arg.GetUserid() {
			room.User = append(room.User[:k], room.User[k+1:]...)
			break
		}
	}
	userInfo, err := models.GetRoomUser(arg.GetUserid())
	if err != nil {
		beego.Error("fight exit error ", err)
	}
	rsp.UserInfo = userInfo
	room.Status = 0
	room.Upsert()
	a.rooms[roomid] = room
	a.broadcast(roomid, rsp)
}

// enter ...
func (a *NodeActor) fightEnter(arg *pb.CFightEnter, ctx actor.Context) {
	rsp := new(pb.SFightEnter)
	room := a.rooms[arg.GetRoomid()]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight enter error ", arg)
		rsp.Error = pb.FightRoomNotExist
		ctx.Respond(rsp)
		return
	}
	if a.roomFull(arg.GetRoomid()) {
		rsp.Error = pb.FightRoomFull
		ctx.Respond(rsp)
		return
	}
	//
	a.roles[arg.Userid] = &models.RoomRole{
		Roomid: room.ID,
		Pid:    ctx.Sender(),
	}
	userInfo := models.RoomUser{
		Userid: arg.GetUserid(),
	}
	room.User = append(room.User, userInfo)
	room.Number++
	a.rooms[room.ID] = room
	rsp.RoomInfo = models.RoomData(room)
	ctx.Respond(rsp)
	for _, v := range room.User {
		if v.Userid == arg.GetUserid() {
			continue
		}
		msg := new(pb.SFightUser)
		msg.RoomInfo = rsp.RoomInfo
		a.send2userid(v.Userid, msg)
	}
}

func (a *NodeActor) roomFull(roomid string) bool {
	if v, ok := a.rooms[roomid]; ok {
		switch v.Type {
		case int32(pb.FIGHT_TYPE0):
			return len(v.User) == 2
		case int32(pb.FIGHT_TYPE1):
			return len(v.User) == 2
		case int32(pb.FIGHT_TYPE2):
			return len(v.User) == 4
		case int32(pb.FIGHT_TYPE3):
			return len(v.User) == 4
		}
	}
	return false
}

func (a *NodeActor) userid2roomid(userid string) (roomid string) {
	if v, ok := a.roles[userid]; ok {
		roomid = v.Roomid
	}
	return
}

// start ...
func (a *NodeActor) fightStart(arg *pb.CFightStart, ctx actor.Context) {
	rsp := new(pb.SFightStart)
	roomid := a.userid2roomid(arg.GetUserid())
	room := a.rooms[roomid]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight start error ", arg)
		rsp.Error = pb.FightNotInRoom
		ctx.Respond(rsp)
		return
	}
	if !a.roomFull(roomid) {
		rsp.Error = pb.FightStartFailed
		ctx.Respond(rsp)
		return
	}
	room.Status = 1
	room.Upsert()
	a.rooms[roomid] = room
	a.broadcast(roomid, rsp)
}

// score ...
func (a *NodeActor) fightScore(arg *pb.CFightingScore, ctx actor.Context) {
	rsp := new(pb.SFightingScore)
	rsp.Userid = arg.GetUserid()
	rsp.Score = arg.GetScore()
	roomid := a.userid2roomid(arg.GetUserid())
	room := a.rooms[roomid]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight score error ", arg)
		rsp.Error = pb.FightNotInRoom
		ctx.Respond(rsp)
		return
	}
	if room.Status != 1 {
		rsp.Error = pb.FightNotStart
		ctx.Respond(rsp)
		return
	}
	for k, v := range room.User {
		if v.Userid == arg.GetUserid() {
			v.Score = arg.GetScore()
			room.User[k] = v
		}
	}
	room.Upsert()
	a.rooms[roomid] = room
	a.broadcast(roomid, rsp)
}

// match ...
func (a *NodeActor) fightMatch(arg *pb.CFightMatch, ctx actor.Context) {
	rsp := new(pb.SFightMatch)
	roomid := a.userid2roomid(arg.GetUserid())
	if _, ok := a.rooms[roomid]; ok {
		//TODO broadcast to nodes
		beego.Error("fight match error ", arg)
		rsp.Error = pb.FightMatchFailed
		ctx.Respond(rsp)
		return
	}
	var room *models.Room
	for k, v := range a.rooms {
		if v.Status == 1 {
			continue
		}
		if v.Type != int32(arg.GetType()) {
			continue
		}
		if a.roomFull(k) {
			continue
		}
		room = v
	}
	if room == nil {
		beego.Error("fight match failed ", arg)
		rsp.Error = pb.FightMatchFailed
		ctx.Respond(rsp)
		return
	}
	//
	a.roles[arg.Userid] = &models.RoomRole{
		Roomid: room.ID,
		Pid:    ctx.Sender(),
	}
	userInfo := models.RoomUser{
		Userid: arg.GetUserid(),
	}
	room.User = append(room.User, userInfo)
	room.Number++
	a.rooms[room.ID] = room
	rsp.RoomInfo = models.RoomData(room)
	ctx.Respond(rsp)
	for _, v := range room.User {
		if v.Userid == arg.GetUserid() {
			continue
		}
		msg := new(pb.SFightUser)
		msg.RoomInfo = rsp.RoomInfo
		a.send2userid(v.Userid, msg)
	}
}

// gird ...
func (a *NodeActor) fightGird(arg *pb.CFightingCancelGird, ctx actor.Context) {
	rsp := new(pb.SFightingCancelGird)
	rsp.Userid = arg.GetUserid()
	rsp.StartPosition = arg.GetStartPosition()
	rsp.EndPosition = arg.GetEndPosition()
	roomid := a.userid2roomid(arg.GetUserid())
	room := a.rooms[roomid]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight gird error ", arg)
		rsp.Error = pb.FightNotInRoom
		ctx.Respond(rsp)
		return
	}
	if room.Status != 1 {
		rsp.Error = pb.FightNotStart
		ctx.Respond(rsp)
		return
	}
	a.broadcast(roomid, rsp)
}

// end ...
func (a *NodeActor) fightEnd(arg *pb.CFightingEnd, ctx actor.Context) {
	rsp := new(pb.SFightingEnd)
	rsp.Userid = arg.GetUserid()
	roomid := a.userid2roomid(arg.GetUserid())
	room := a.rooms[roomid]
	if room == nil {
		//TODO broadcast to nodes
		beego.Error("fight gird error ", arg)
		rsp.Error = pb.FightNotInRoom
		ctx.Respond(rsp)
		return
	}
	if room.Status != 1 {
		rsp.Error = pb.FightNotStart
		ctx.Respond(rsp)
		return
	}
	a.broadcast(roomid, rsp)
	//
	models.LogRoomRecord(room)
	for _, v := range room.User {
		delete(a.roles, v.Userid)
	}
	delete(a.rooms, room.ID)
}
