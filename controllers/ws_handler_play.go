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
    ws.Send(s2c)
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
    //key := models.PropKey(int32(arg.GetPtype()))
    //if val, ok := ws.user.Prop[key]; ok {
    //    if val.
    //}
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
}

//game start handler
func (ws *WSConn) gameStart(arg *pb.CStart) {
    s2c := new(pb.SStart)
    ws.Send(s2c)
}
