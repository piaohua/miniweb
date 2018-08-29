/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-29 17:07:48
 * Filename      : ws_hander_play.go
 * Description   : play handler
 * *******************************************************/

package controllers

import (
	"miniweb/models"
	"miniweb/pb"

	"github.com/astaxie/beego"
)

//get user data info
func (ws *WSConn) getUserData(arg *pb.CUserData) {
	s2c := new(pb.SUserData)
	s2c.UserInfo = &pb.UserData{
		Userid:    ws.user.ID,
		NickName:  ws.user.NickName,
		AvatarUrl: ws.user.AvatarUrl,
		Gender:    ws.user.Gender,
		Diamond:   ws.user.Diamond,
		Coin:      ws.user.Coin,
		Energy:    ws.user.Energy,
	}
	ws.Send(s2c)
}

//get currency info
func (ws *WSConn) getCurrency() {
	s2c := new(pb.SGetCurrency)
	s2c.Coin = ws.user.Coin
	s2c.Diamond = ws.user.Diamond
	s2c.Energy = ws.user.Energy
	ws.Send(s2c)
}

//get gate data info
func (ws *WSConn) getGateData() {
	s2c := new(pb.SGateData)
	for _, v := range ws.user.Gate {
		gate := &pb.GateData{
			Type:   pb.GateType(v.Type),
			Gateid: v.Gateid,
			Num:    v.Num,
			Star:   v.Star,
		}
		s2c.GateInfo = append(s2c.GateInfo, gate)
	}
	ws.Send(s2c)
}

//get prop data info
func (ws *WSConn) getPropData() {
	s2c := new(pb.SPropData)
	for _, v := range ws.user.Prop {
		prop := &pb.PropData{
			Type: pb.PropType(v.Type),
			Num:  v.Num,
			Attr: v.Attr,
		}
		prop.Name = models.GetPropName(v.Type)
		s2c.PropInfo = append(s2c.PropInfo, prop)
	}
	ws.Send(s2c)
}

//get shop data info
func (ws *WSConn) getShopData() {
	s2c := new(pb.SShop)
	list := models.GetShops()
	for _, v := range list {
		shop := &pb.Shop{
			Id:     v.Id,
			Status: pb.ShopStatus(v.Status),
			Type:   pb.PropType(v.Ptype),
			Way:    pb.PayWay(v.Payway),
			Number: v.Number,
			Price:  v.Price,
			Name:   v.Name,
			Info:   v.Info,
		}
		s2c.List = append(s2c.List, shop)
	}
	ws.Send(s2c)
}

//buy handler
func (ws *WSConn) buy(arg *pb.CBuy) {
	s2c := new(pb.SBuy)
	//shop := models.GetShop(arg.GetId())
	ws.Send(s2c)
}

//over handler
func (ws *WSConn) overData(arg *pb.COverData) {
	s2c := new(pb.SOverData)
	Type := int32(arg.GetType())
	gateID := arg.GetGateid()
	key := models.GateKey(Type, gateID)
	if val, ok := ws.user.Gate[key]; ok {
		//send prize
		coin, energy := overPrize(arg.GetStar(), val.Num == 0)
		if coin > 0 {
			msg1 := models.AddCoinMsg(ws.user, coin)
			ws.Send(msg1)
		}
		if energy > 0 {
			msg2 := models.AddEnergyMsg(ws.user, energy)
			ws.Send(msg2)
		}
		//response
		//update gateid
		models.AddGate(ws.user, Type, gateID, arg.GetStar())
		//add new gateid
		nextid := gateID + 1
		models.AddNewGate(ws.user, Type, nextid)
		nextKey := models.GateKey(Type, nextid)
		if nextVal, ok := ws.user.Gate[nextKey]; ok {
			s2c.NextGate = &pb.GateData{
				Type:   arg.GetType(),
				Gateid: gateID + 1,
				Num:    nextVal.Num,
				Star:   nextVal.Star,
			}
		}
		ws.Send(s2c)
		return
	}
	s2c.Error = pb.GateUnreachable
	ws.Send(s2c)
	beego.Error("overData error ", arg)
}

//over prize
func overPrize(star int32, first bool) (coin, energy int64) {
	if first {
		switch star {
		case 1:
			coin, energy = 100, 1
		case 2:
			coin, energy = 300, 3
		case 3:
			coin, energy = 500, 5
		}
		return
	}
	switch star {
	case 1:
		coin, energy = 30, 0
	case 2:
		coin, energy = 70, 0
	case 3:
		coin, energy = 100, 0
	}
	return
}

//card handler
func (ws *WSConn) card(arg *pb.CCard) {
	s2c := new(pb.SCard)
	ws.Send(s2c)
}

//login prize handler
func (ws *WSConn) loginPrize(arg *pb.CLoginPrize) {
	s2c := new(pb.SLoginPrize)
	ws.Send(s2c)
}

//use prop handler
func (ws *WSConn) useProp(arg *pb.CUseProp) {
	s2c := new(pb.SUseProp)
	key := models.PropKey(int32(arg.GetPtype()))
	if _, ok := ws.user.Prop[key]; !ok {
		s2c.Error = pb.PropNotEnough
		ws.Send(s2c)
		beego.Error("useProp error ", arg)
		return
	}
	var msg interface{}
	switch arg.GetPtype() {
	case pb.PROP_TYPE10:
		msg = models.AddEnergyMsg(ws.user, 1)
	case pb.PROP_TYPE11:
		msg = models.AddEnergyMsg(ws.user, 5)
	case pb.PROP_TYPE12:
		msg = models.AddEnergyMsg(ws.user, 30)
	}
	ws.Send(s2c)
	if msg != nil {
		ws.Send(msg)
	}
	msg2 := models.AddPropMsg(ws.user, key, -1, arg.GetPtype())
	if msg2 != nil {
		ws.Send(msg2)
	}
}

//game start handler
func (ws *WSConn) gameStart(arg *pb.CStart) {
	s2c := new(pb.SStart)
	//检测关卡
	key := models.GateKey(int32(arg.GetType()), arg.GetGateid())
	if val, ok := ws.user.Gate[key]; ok {
		s2c.GateInfo = &pb.GateData{
			Type:   arg.GetType(),
			Gateid: arg.GetGateid(),
			Num:    val.Num,
			Star:   val.Star,
		}
		ws.Send(s2c)
		return
	}
	s2c.Error = pb.GateUnreachable
	ws.Send(s2c)
	beego.Error("gameStart error ", arg)
}
