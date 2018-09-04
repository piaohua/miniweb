# 小程序

[English](./README.md)

本示例通过使用http和WebSocket 基于 beego 构建一个小程序后台程序。

## Usage

[API](./api.md)


```
c2s
[4byte(自定义id) + msgdata(protobuf)]

s2c
[4byte(自定义id) + 1byte(暂时保留) + msgdata(protobuf)]

```
