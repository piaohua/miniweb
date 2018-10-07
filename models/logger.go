/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-10-02 23:36:57
 * Filename      : logger.go
 * Description   : logger
 * *******************************************************/

package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// LogRoom log room
type LogRoom struct {
	ID        string     `bson:"_id" json:"id"`                //unique ID roomid
	Userid    string     `bson:"userid" json:"userid"`         //创建者
	Code      string     `bson:"code" json:"code"`             //
	Node      string     `bson:"node" json:"node"`             //
	Type      int32      `bson:"type" json:"type"`             //type
	Match     int32      `bson:"match" json:"match"`           //
	UserProp  int32      `bson:"user_prop" json:"user_prop"`   //
	Number    int        `bson:"number" json:"number"`         //
	User      []RoomUser `bson:"user" json:"user"`             //
	RoomCtime time.Time  `bson:"room_ctime" json:"room_ctime"` //room创建时间
	Ctime     time.Time  `bson:"ctime" json:"ctime"`           //record创建时间(结束时创建)
}

//Save 写入数据库
func (t *LogRoom) Save() bool {
	t.Ctime = bson.Now()
	return Insert(LogRooms, t)
}

// LogRoomUser log room user
type LogRoomUser struct {
	ID     string    `bson:"_id" json:"id"`        //unique ID
	Roomid string    `bson:"roomid" json:"roomid"` //unique == LogRoom.ID
	Userid string    `bson:"userid" json:"userid"` //userid
	Score  int32     `bson:"score" json:"score"`   //score
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
}

//Save 写入数据库
func (t *LogRoomUser) Save() bool {
	t.Ctime = bson.Now()
	return Insert(LogRoomUsers, t)
}

//LogRoomRecord ...
func LogRoomRecord(room *Room) {
	logRoomid := GenUniqueID()
	logRoom := &LogRoom{
		ID:        logRoomid,
		Userid:    room.Userid,
		Code:      room.ID,
		Node:      room.Node,
		Type:      room.Type,
		Match:     room.Match,
		UserProp:  room.UserProp,
		Number:    room.Number,
		User:      room.User,
		RoomCtime: room.Ctime,
	}
	logRoom.Save()
	for _, v := range room.User {
		logRoomUser := &LogRoomUser{
			Roomid: logRoomid,
			Userid: v.Userid,
			Score:  v.Score,
		}
		logRoomUser.Save()
	}
}
