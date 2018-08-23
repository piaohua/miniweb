# 小程序

[中文文档](./README_ZH.md)

This sample is about using http and WebSocket to build a wechat miniprogram based on beego.

## Installation

```
cd $GOPATH/src/miniweb
go get github.com/gorilla/websocket
go get github.com/beego/i18n
bee run
```

## Usage

```
get session by js_code
http://127.0.0.1:8080/ws?js_code=xxx

websocket connection by session
http://127.0.0.1:8080/ws/login?3rd_session=xxx

```
