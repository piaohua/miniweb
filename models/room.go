/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-10-02 17:37:31
 * Filename      : room.go
 * Description   : room
 * *******************************************************/

package models

import (
	"errors"
	"time"

	"miniweb/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

//Room room info TODO type区分阵容
type Room struct {
	ID       string     `bson:"_id" json:"id"`              //unique ID
	Userid   string     `bson:"userid" json:"userid"`       //
	Node     string     `bson:"node" json:"node"`           //
	Type     int32      `bson:"type" json:"type"`           //type
	Match    int32      `bson:"match" json:"match"`         //
	UserProp int32      `bson:"user_prop" json:"user_prop"` //
	Status   int32      `bson:"status" json:"status"`       //0:准备,1游戏中
	Number   int        `bson:"number" json:"number"`       //
	User     []RoomUser `bson:"user" json:"user"`           //
	Ctime    time.Time  `bson:"ctime" json:"ctime"`         //创建时间
}

//RoomUser room user info
type RoomUser struct {
	Userid string `bson:"userid" json:"userid"` //userid
	Score  int32  `bson:"score" json:"score"`   //score
}

//RoomRole room role info
type RoomRole struct {
	Roomid string     `bson:"-" json:"-"`
	Pid    *actor.PID `bson:"-" json:"-"`
}

//Save 写入数据库
func (t *Room) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Rooms, t)
}

//Upsert ...
func (t *Room) Upsert() bool {
	return Upsert(Rooms, bson.M{"_id": t.ID}, t)
}

//Get 加载
func (t *Room) Get() {
	Get(Rooms, t.ID, t)
}

//GetRoomList ...
func GetRoomList(node string) []Room {
	var list []Room
	ListByQ(Rooms, bson.M{"node": node}, &list)
	return list
}

//GetByRoomID  通过id获取
func GetByRoomID(id string) (t *Room) {
	t = new(Room)
	GetByQ(Rooms, bson.M{"_id": id}, t)
	return
}

//HasCode code是否存在
func HasCode(code string) bool {
	return Has(Rooms, bson.M{"_id": code})
}

//DeleteByRoomID 删除记录
func DeleteByRoomID(id string) bool {
	return Delete(Rooms, bson.M{"_id": id})
}

//GetRoomData  通过roomid获取房间信息
func GetRoomData(roomid string) *pb.RoomData {
	t := new(Room)
	t.ID = roomid
	t.Get()
	return RoomData(t)
}

//RoomData room info
func RoomData(t *Room) *pb.RoomData {
	return &pb.RoomData{
		Userid:   t.Userid,
		Roomid:   t.ID,
		Type:     pb.FightType(t.Type),
		Match:    pb.AllowType(t.Match),
		UserProp: pb.AllowType(t.UserProp),
		UserInfo: getRoomUsers(t.User),
	}
}

// getRoomUsers 获取房间成员信息
func getRoomUsers(ls []RoomUser) []*pb.RoomUser {
	ids := make([]string, 0)
	ms := make(map[string]int32)
	for _, v := range ls {
		ids = append(ids, v.Userid)
		ms[v.Userid] = v.Score
	}
	list, err := getRoomUserInfo(ids)
	if err != nil {
		beego.Error("getRoomUsers failed ", err)
	}
	for k, v := range list {
		v.Score = ms[v.Userid]
		list[k] = v
	}
	return list
}

// getRoomUserInfo room user info
func getRoomUserInfo(ids []string) ([]*pb.RoomUser, error) {
	var list []bson.M
	selector := make(bson.M, 4)
	selector["nick_name"] = true
	selector["avatar_url"] = true
	selector["gender"] = true
	selector["_id"] = true
	q := bson.M{"_id": bson.M{"$in": ids}}
	err := Users.
		Find(q).Select(selector).
		All(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none record")
	}
	infos := make([]*pb.RoomUser, 0)
	for _, v := range list {
		info := roomUserInfo(v)
		if info == nil {
			continue
		}
		infos = append(infos, info)
	}
	return infos, nil
}

func roomUserInfo(v bson.M) *pb.RoomUser {
	info := new(pb.RoomUser)
	if val, ok := v["_id"]; ok {
		info.Userid = val.(string)
	}
	if val, ok := v["nick_name"]; ok {
		info.NickName = val.(string)
	}
	if val, ok := v["avatar_url"]; ok {
		info.AvatarUrl = val.(string)
	}
	if val, ok := v["gender"]; ok {
		info.Gender = val.(int32)
	}
	if info.Userid == "" {
		return nil
	}
	return info
}

// GetRoomUser room user info
func GetRoomUser(id string) (*pb.RoomUser, error) {
	var user bson.M
	selector := make(bson.M, 4)
	selector["nick_name"] = true
	selector["avatar_url"] = true
	selector["gender"] = true
	selector["_id"] = true
	q := bson.M{"_id": id}
	err := Users.
		Find(q).Select(selector).
		One(&user)
	if err != nil {
		return nil, err
	}
	info := roomUserInfo(user)
	if info == nil {
		return nil, errors.New("none record")
	}
	return info, nil
}

// CreateRoom create room
func CreateRoom(arg *pb.CFightCreate, node string) (room *Room, err error) {
	code, err := getNewCode()
	if err != nil {
		return
	}
	room = &Room{
		ID:       code,
		Userid:   arg.GetUserid(),
		Node:     node,
		Type:     int32(arg.GetType()),
		Match:    int32(arg.GetMatch()),
		UserProp: int32(arg.GetUserProp()),
		Status:   0,
		Number:   1,
	}
	roomUser := RoomUser{
		Userid: arg.GetUserid(),
		Score:  0,
	}
	room.User = append(room.User, roomUser)
	if !room.Save() {
		err = errors.New("room save failed")
		return
	}
	return
}

// ChangeRoom update room
func ChangeRoom(arg *pb.CFightChangeSet) bool {
	return Update(Rooms, bson.M{"_id": arg.Roomid, "userid": arg.Userid},
		bson.M{"$set": bson.M{"match": int32(arg.Match),
			"user_prop": int32(arg.UserProp)}})
}
