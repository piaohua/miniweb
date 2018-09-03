package controllers

import (
	"net/http"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

//Handler global handler
var Handler *WSHandler

func init() {
	Handler = &WSHandler{
		maxConnNum:      30000,
		pendingWriteNum: 100,
		maxMsgLen:       4096,
	}
	upgrader := websocket.Upgrader{
		ReadBufferSize:   1024, //default 4096
		WriteBufferSize:  1024, //default 4096
		HandshakeTimeout: 10 * time.Second,
		CheckOrigin:      func(_ *http.Request) bool { return true },
	}
	Handler.upgrader = upgrader
	Handler.conns = make(WebsocketConnSet)
}

//WSHandler ws handler
type WSHandler struct {
	close           bool               //是否关闭监听
	maxConnNum      int                //最大连接数
	pendingWriteNum int                //等待写入消息长度
	maxMsgLen       uint32             //最大消息长
	upgrader        websocket.Upgrader //升级http连接
	conns           WebsocketConnSet   //连接集合
	mutexConns      sync.Mutex         //互斥锁
	wg              sync.WaitGroup     //同步机制
}

//Add add conn
func (handler *WSHandler) Add(conn *websocket.Conn) {
	handler.wg.Add(1)
	defer handler.wg.Done()

	handler.mutexConns.Lock()
	if handler.conns == nil {
		handler.mutexConns.Unlock()
		conn.Close()
		return
	}
	if len(handler.conns) >= handler.maxConnNum {
		handler.mutexConns.Unlock()
		conn.Close()
		beego.Error("too many connections: ", len(handler.conns))
		return
	}
	handler.conns[conn] = struct{}{}
	handler.mutexConns.Unlock()
}

//Close close all conn
func (handler *WSHandler) Close() {

	handler.mutexConns.Lock()
	handler.close = true
	for conn := range handler.conns {
		conn.Close()
	}
	handler.conns = nil
	handler.mutexConns.Unlock()

	handler.wg.Wait()
}
