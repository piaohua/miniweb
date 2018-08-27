/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-24 11:18:20
 * Filename      : ws_msg.go
 * Description   : ws msg
 * *******************************************************/

package controllers

import (
	"encoding/binary"

	"miniweb/pb"

	"github.com/astaxie/beego"
)

//Router 路由
func (ws *WSConn) Router(id uint32, body []byte) {
	//if aesStatus {
	//	body = aesDe(body) //解密
	//}
	msg, err := pb.Unpack(id, body)
	if err != nil {
		beego.Error("protocol unpack err: ", id, err)
		return
	}
	ws.pid.Tell(msg)
}

//Send 发送消息
func (ws *WSConn) Send(msg interface{}) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		glog.Errorf("msg %#v, err %#v", msg, err)
	//		glog.Error(string(debug.Stack()))
	//	}
	//}()
	if ws.msgCh == nil {
		beego.Error("WSConn msg channel closed: ", msg)
		return
	}
	if len(ws.msgCh) == cap(ws.msgCh) {
		beego.Error("send msg channel full:", len(ws.msgCh))
		return
	}
	select {
	case <-ws.stopCh:
		return
	default:
	}
	select {
	case <-ws.stopCh:
		return
	case ws.msgCh <- msg:
		//beego.Error("message %#v", msg)
	}
}

//封包
func pack(code, sid uint32, msg []byte) []byte {
	//if aesStatus {
	//	msg = aesEn(msg) //加密
	//}
	return append(append(encodeUint32(code), byte(sid)), msg...)
}

//little Endian encode
func encodeUint32(i uint32) (b []byte) {
	b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return
}

//little Endian decode
func decodeUint32(b []byte) (i uint32) {
	i = binary.LittleEndian.Uint32(b)
	return
}
