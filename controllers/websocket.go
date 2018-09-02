// Copyright 2013 Beego Miniweb authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"miniweb/models"
	"miniweb/pb"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	jscode := this.GetString("js_code")
	if len(jscode) == 0 {
		this.Redirect("/", 302)
		return
	}
	beego.Info("websocket get jscode", jscode)

	if this.isPost() {
		this.Redirect("/", 302)
		return
	}

	ip := this.getClientIp()
	session, err := models.GetSession(jscode, ip)
	wsaddr := beego.AppConfig.String("ws.addr")

	jsonData := &models.SessionResult{
		Session: session,
		WsAddr:  wsaddr + session,
	}
	if err != nil {
		jsonData.WxErr = models.WxErr{
			ErrCode: int(pb.WSGetFailed),
			ErrMsg:  err.Error(),
		}
	}
	this.jsonResult(jsonData)
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Login() {
	session := this.GetString("3rd_session")
	if len(session) == 0 {
		this.Redirect("/", 302)
		return
	}
	beego.Info("websocket join 3rd_session", session)

	if this.isPost() {
		this.Redirect("/", 302)
		return
	}

	if Handler.close {
		jsonData := &models.WxErr{
			ErrCode: int(pb.WSLoginFailed),
			ErrMsg:  "closed",
		}
		this.jsonResult(jsonData)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := Handler.upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if val, ok := err.(websocket.HandshakeError); ok {
		beego.Error("Not a websocket handshake: ", val.Error())
		//http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		jsonData := &models.WxErr{
			ErrCode: int(pb.WSLoginFailed),
			ErrMsg:  err.Error(),
		}
		this.jsonResult(jsonData)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		jsonData := &models.WxErr{
			ErrCode: int(pb.WSLoginFailed),
			ErrMsg:  err.Error(),
		}
		this.jsonResult(jsonData)
		return
	}

	Handler.Add(ws)
	ws.SetReadLimit(int64(Handler.maxMsgLen))

	//handler
	wsConn := newWSConn(ws, Handler.pendingWriteNum, Handler.maxMsgLen)
	wsConn.session = session
	wsConn.pid = wsConn.initWs()
	//start pid
	wsConn.pid.Tell(new(pb.ServeStart))
	//defer Leave handler
	defer wsConn.Leave()
	//Send Message
	go wsConn.writePump()
	go wsConn.pingPump()
	// Message receive loop.
	wsConn.readPump()
}
