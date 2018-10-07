/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-10-03 17:12:08
 * Filename      : ws_hander_fight.go
 * Description   : fight handler
 * *******************************************************/

package controllers

import (
	"miniweb/pb"
)

// fight ...  TODO 优化控制
func (ws *WSConn) fight() {
	s2c := new(pb.SFight)
	s2c.Type = append(s2c.Type, pb.FIGHT_TYPE0,
		pb.FIGHT_TYPE1, pb.FIGHT_TYPE2, pb.FIGHT_TYPE3)
	ws.Send(s2c)
}

// match ...
func (ws *WSConn) fightMatch(arg *pb.CFightMatch) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// create ...
func (ws *WSConn) fightCreate(arg *pb.CFightCreate) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// enter ...
func (ws *WSConn) fightEnter(arg *pb.CFightEnter) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// exit ...
func (ws *WSConn) fightExit(arg *pb.CFightMatchExit) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// change ...
func (ws *WSConn) fightChange(arg *pb.CFightChangeSet) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// start ...
func (ws *WSConn) fightStart(arg *pb.CFightStart) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// gird ...
func (ws *WSConn) fightGird(arg *pb.CFightingCancelGird) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// score ...
func (ws *WSConn) fightScore(arg *pb.CFightingScore) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}

// end ...
func (ws *WSConn) fightEnd(arg *pb.CFightingEnd) {
	arg.Userid = ws.user.ID
	NodePid.Request(arg, ws.pid)
}
