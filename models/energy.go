package models

import (
	"time"

	"miniweb/pb"
)

//CheckEnergy 精力恢复
func CheckEnergy(user *User) (msg *pb.SPushProp) {
	if user == nil {
		return
	}
	now := time.Now().Unix()
	if user.EnergyTime == 0 {
		user.EnergyTime = now
		return
	}
	n := (now - user.EnergyTime) / 60
	if n <= 0 {
		return
	}
	user.EnergyTime += 60 * n
	msg = AddEnergyMsg(user, n)
	return
}

//AddEnergyMsg add energy msg
func AddEnergyMsg(user *User, num int64) (msg *pb.SPushProp) {
	user.AddEnergy(num)
	msg = &pb.SPushProp{
		//Type: pb.LOG_TYPE0,
		Num: num,
		PropInfo: &pb.PropData{
			Type: pb.PROP_TYPE3,
			Num:  user.Energy,
		},
	}
	return
}
