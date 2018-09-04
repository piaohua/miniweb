/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-29 17:07:48
 * Filename      : ws_hander_play.go
 * Description   : play handler
 * *******************************************************/

package controllers

import (
	"time"

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
	//gate count
	list := models.GetGateCount()
	for _, v := range list {
		count := &pb.GateCount{
			Type: pb.GateType(v.Type),
			Num:  v.Number,
		}
		s2c.Counts = append(s2c.Counts, count)
	}
	ws.Send(s2c)
}

//get prop data info
func (ws *WSConn) getPropData() {
	s2c := new(pb.SPropData)
	for _, v := range ws.user.Prop {
		prop := &pb.PropData{
			Type: pb.PropType(v.Type),
			Num:  int64(v.Num),
			Attr: v.Attr,
		}
		prop.Name = models.GetPropName(v.Type)
		s2c.PropInfo = append(s2c.PropInfo, prop)
	}
	ws.Send(s2c)
}

//get temp shop data info
func (ws *WSConn) getTempShopData(arg *pb.CTempShop) {
	s2c := new(pb.STempShop)
	Type := int32(arg.GetType())
	gateID := arg.GetGateid()
	key := models.GateKey(Type, gateID)
	s2c.Error = ws.isReachable(Type, key)
	if s2c.Error != pb.OK {
		ws.Send(s2c)
		beego.Error("get temp shop error ", arg)
		return
	}
	m := make(map[string]bool, 0)
	gate := models.GetGate(Type, gateID)
	if gate != nil {
		for _, v := range gate.TempShop {
			m[v] = true
		}
	}
	//默认 TODO 优化
	if len(m) == 0 {
		m["17"] = true
		m["18"] = true
		m["19"] = true
	}
	list := models.GetShops() //优化
	for _, v := range list {
		switch v.Status {
		case int32(pb.SHOP_STATUS4): //temp shop
		default:
			continue
		}
		if _, ok := m[v.ID]; !ok {
			continue
		}
		shop := &pb.Shop{
			Id:     v.ID,
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

//gate is reachable
func (ws *WSConn) isReachable(Type int32, key string) pb.ErrCode {
	switch Type {
	case int32(pb.GATE_TYPE1), //单人
		int32(pb.GATE_TYPE2): //副本
		if _, ok := ws.user.Gate[key]; !ok {
			return pb.GateUnreachable
		}
	default:
		return pb.GateUnreachable
	}
	return pb.OK
}

//get shop data info
func (ws *WSConn) getShopData() {
	s2c := new(pb.SShop)
	list := models.GetShops()
	for _, v := range list {
		switch v.Status {
		case int32(pb.SHOP_STATUS4): //temp shop
			continue
		}
		shop := &pb.Shop{
			Id:     v.ID,
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
	shop := models.GetShop(arg.GetId())
	if shop == nil {
		beego.Error("order failed ", arg.GetId())
		s2c.Error = pb.OrderFailed
		ws.Send(s2c)
		return
	}
	switch shop.Status {
	case int32(pb.SHOP_STATUS3):
		beego.Error("buy failed ", arg.GetId())
		s2c.Error = pb.OrderFailed
		ws.Send(s2c)
		return
	}
	switch shop.Payway {
	case int32(pb.PAY_WAY0):
		//TODO RMB
		beego.Error("buy failed ", arg.GetId())
		s2c.Error = pb.OrderFailed
		ws.Send(s2c)
		return
	case int32(pb.PAY_WAY1):
		if ws.user.Diamond < int64(shop.Price) {
			s2c.Error = pb.DiamondNotEnough
			s2c.Status = pb.BuyFailed
			ws.Send(s2c)
			return
		}
		msg1 := models.AddDiamondMsg(ws.user, -1*int64(shop.Price))
		ws.Send(msg1)
	case int32(pb.PAY_WAY2):
		if ws.user.Coin < int64(shop.Price) {
			s2c.Error = pb.CoinNotEnough
			s2c.Status = pb.BuyFailed
			ws.Send(s2c)
			return
		}
		msg1 := models.AddCoinMsg(ws.user, -1*int64(shop.Price))
		ws.Send(msg1)
	}
	s2c.Status = pb.BuySuccess
	ws.Send(s2c)
	//发货
	key := models.PropKey(int32(shop.Ptype))
	msg2 := models.AddPropMsg(ws.user, key, int64(shop.Number), pb.PropType(shop.Ptype))
	ws.Send(msg2)
	//奖励发放
	ws.sendShopPrize(shop.Prize)
	//TODO 下单购买日志
}

//buy temp prop handler
func (ws *WSConn) buyTemp(ids []string) (err pb.ErrCode) {
	if len(ids) == 0 {
		return
	}
	var coin, diamond int64
	l := []struct {
		Ptype  int32
		Number uint32
	}{}
	for _, id := range ids {
		shop := models.GetShop(id)
		switch shop.Payway {
		case int32(pb.PAY_WAY1):
			diamond += int64(shop.Price)
		case int32(pb.PAY_WAY2):
			coin += int64(shop.Price)
		default:
			err = pb.OrderFailed
			return
		}
		l = append(l, struct {
			Ptype  int32
			Number uint32
		}{Ptype: shop.Ptype, Number: shop.Number})
	}
	if ws.user.Coin < int64(coin) {
		err = pb.CoinNotEnough
		return
	}
	if ws.user.Diamond < int64(diamond) {
		err = pb.DiamondNotEnough
		return
	}
	if diamond > 0 {
		msg1 := models.AddDiamondMsg(ws.user, -1*int64(diamond))
		ws.Send(msg1)
	}
	if coin > 0 {
		msg2 := models.AddCoinMsg(ws.user, -1*int64(coin))
		ws.Send(msg2)
	}
	//奖励发放
	for _, v := range l {
		key := models.PropKey(int32(v.Ptype))
		msg2 := models.AddTempPropMsg(ws.user, key, int64(v.Number), pb.PropType(v.Ptype))
		ws.Send(msg2)
	}
	return
}

//change change currency
func (ws *WSConn) change(arg *pb.ChangeCurrency) {
	if arg.Diamond != 0 {
		msg1 := models.AddDiamondMsg(ws.user, int64(arg.Diamond))
		ws.Send(msg1)
	}
	if arg.Coin != 0 {
		msg2 := models.AddCoinMsg(ws.user, int64(arg.Coin))
		ws.Send(msg2)
	}
}

//奖励发放
func (ws *WSConn) sendShopPrize(list []models.ShopPrizeProp) {
	for _, v := range list {
		key := models.PropKey(int32(v.Type))
		msg := models.AddPropMsg(ws.user, key, int64(v.Number), pb.PropType(v.Type))
		ws.Send(msg)
	}
}

//over handler
func (ws *WSConn) overData(arg *pb.COverData) {
	s2c := new(pb.SOverData)
	Type := int32(arg.GetType())
	gateID := arg.GetGateid()
	key := models.GateKey(Type, gateID)
	if val, ok := ws.user.Gate[key]; ok {
		//send prize
		coin, energy := overPrize(arg.GetGateid(), arg.GetStar(), val.Num == 0)
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
		//gate info
		s2c.GateInfo = &pb.GateData{
			Type:   arg.GetType(),
			Gateid: gateID,
			Num:    ws.user.Gate[key].Num,
			Star:   ws.user.Gate[key].Star,
		}
		ws.Send(s2c)
		ws.tempClean()
		return
	}
	s2c.Error = pb.GateUnreachable
	ws.Send(s2c)
	beego.Error("overData error ", arg)
}

//temp prop clean
func (ws *WSConn) tempClean() {
	ws.user.TempProp = make(map[string]models.TempPropInfo)
}

//over prize
func overPrize(gate, star int32, first bool) (coin, energy int64) {
	switch star {
	case 1:
	case 2:
		energy = 2
	case 3:
		energy = 5
	}
	if first {
		coin = int64((10 * (gate - 1)) + 500 + (100 * star * star))
	} else {
		coin = int64(50 + (30 * star))
	}
	return
}

//card handler
func (ws *WSConn) card(arg *pb.CCard) {
	//test
	msg1 := models.AddCoinMsg(ws.user, 30000)
	ws.Send(msg1)
	msg2 := models.AddDiamondMsg(ws.user, 1000)
	ws.Send(msg2)
	s2c := new(pb.SCard)
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
	switch arg.GetPtype() {
	case pb.PROP_TYPE9,
		pb.PROP_TYPE10,
		pb.PROP_TYPE11:
		if ws.user.Energy >= 30 {
			s2c.Error = pb.EnergyEnough
			ws.Send(s2c)
			return
		}
	}
	var msg interface{}
	switch arg.GetPtype() {
	case pb.PROP_TYPE9:
		msg = models.AddEnergyMsg(ws.user, 1)
	case pb.PROP_TYPE10:
		msg = models.AddEnergyMsg(ws.user, 5)
	case pb.PROP_TYPE11:
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
	ws.tempClean()
	s2c := new(pb.SStart)
	if ws.user.Energy < 5 {
		s2c.Error = pb.EnergyNotEnough
		ws.Send(s2c)
		return
	}
	Type := int32(arg.GetType())
	gateID := arg.GetGateid()
	//检测关卡
	key := models.GateKey(Type, gateID)
	if val, ok := ws.user.Gate[key]; ok {
		//购买临时道具
		err := ws.buyTemp(arg.GetIds())
		if err != pb.OK {
			s2c.Error = err
			ws.Send(s2c)
			return
		}
		s2c.GateInfo = &pb.GateData{
			Type:   arg.GetType(),
			Gateid: gateID,
			Num:    val.Num,
			Star:   val.Star,
		}
		//data 配置数据
		gate := models.GetGate(Type, gateID)
		if gate != nil {
			s2c.GateInfo.Data = gate.Data
		}
		ws.Send(s2c)
		msg := models.AddEnergyMsg(ws.user, -5)
		ws.Send(msg)
		return
	}
	s2c.Error = pb.GateUnreachable
	ws.Send(s2c)
	beego.Error("gameStart error ", arg)
}

//更新连续登录奖励
func (ws *WSConn) loginPrizeInit() {
	setLoginPrize(ws.user)
	ws.user.LoginTime = time.Now().Local()
}

//连续登录奖励处理
func (ws *WSConn) loginPrize(arg *pb.CLoginPrize) {
	msg := new(pb.SLoginPrize)
	msg.Type = arg.Type
	switch arg.Type {
	case pb.LoginPrizeSelect:
		msg.List = loginPrizeInfo(ws.user)
	case pb.LoginPrizeDraw:
		l, errCode := getLoginPrize(arg.Day, ws.user)
		if errCode == pb.OK {
			//奖励发放
			ws.sendLoginPrize(l)
			msg.List = loginPrizeInfo(ws.user)
		} else {
			msg.Error = errCode
		}
	}
	ws.Send(msg)
}

//奖励发放
func (ws *WSConn) sendLoginPrize(list []models.LoginPrizeProp) {
	for _, v := range list {
		key := models.PropKey(int32(v.Type))
		msg := models.AddPropMsg(ws.user, key, int64(v.Number), pb.PropType(v.Type))
		ws.Send(msg)
	}
}

//setLoginPrize 连续登录处理
func setLoginPrize(user *models.User) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	yesterDay := today.AddDate(0, 0, -1)
	if user.LoginTime.Before(yesterDay) {
		//隔天登录重置
		user.LoginTimes = (1 << 0)
		user.LoginPrize = 0
		return
	}
	//是否昨天登录过
	if user.LoginTime.Before(today) {
		//全部领取完重置
		if user.LoginTimes == 127 && user.LoginPrize == 127 {
			user.LoginTimes = (1 << 0)
			user.LoginPrize = 0
			return
		}
		//新的一天
		var i uint32
		for i = 0; i < 7; i++ {
			if (user.LoginTimes & (1 << i)) == 0 {
				user.LoginTimes |= (1 << i)
				break
			}
		}
	}
}

//getLoginPrize 领取连续登录奖励
func getLoginPrize(day uint32, user *models.User) (l []models.LoginPrizeProp, err pb.ErrCode) {
	if (user.LoginPrize & (1 << day)) != 0 {
		beego.Error("getLoginPrize error ", day, user.LoginPrize)
		err = pb.AlreadyAward
		return
	}
	if (user.LoginTimes & (1 << day)) == 0 {
		beego.Error("getLoginPrize failed ", day, user.LoginTimes)
		err = pb.AwardFailed
		return
	}
	prize := models.GetLoginPrize(day)
	if prize == nil {
		beego.Error("getLoginPrize failed ", day, user.LoginTimes)
		err = pb.AwardFailed
		return
	}
	user.LoginPrize |= (1 << day)
	return prize.Prize, pb.OK
}

//loginPrizeInfo 获取连续登录信息
func loginPrizeInfo(user *models.User) (msg []*pb.LoginPrize) {
	list := models.GetLoginPrizes()
	for _, v := range list {
		msg2 := new(pb.LoginPrize)
		msg2.Day = v.Day
		for _, val := range v.Prize {
			msg3 := &pb.LoginPrizeProp{
				Type:   pb.PropType(val.Type),
				Number: val.Number,
			}
			msg3.Name = models.GetPropName(val.Type)
			msg2.Prize = append(msg2.Prize, msg3)
		}
		if (user.LoginPrize & (1 << v.Day)) != 0 {
			msg2.Status = pb.LoginPrizeGot
		} else if (user.LoginTimes & (1 << v.Day)) != 0 {
			msg2.Status = pb.LoginPrizeDone
		}
		msg = append(msg, msg2)
	}
	return
}
