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
    }
    n := (now - user.EnergyTime) % 360
    if n <= 0 {
        return
    }
    user.EnergyTime += 360 * n
    msg = AddEnergyMsg(user, n)
    return
}

//AddEnergyMsg add energy msg
func AddEnergyMsg(user *User, n int64) (msg *pb.SPushProp) {
    user.AddEnergy(n)
    msg = &pb.SPushProp{
        //Type: pb.LOG_TYPE0,
        Ptype: pb.PROP_TYPE3,
        Num:   n,
    }
    return
}
