# API

## Usage

```
host http://127.0.0.1:8080

get session by js_code
http://127.0.0.1:8080/ws?js_code=xxx

get session by js_code (post request)
http://127.0.0.1:8080/code
curl -d "js_code=xxx" -H "token:your-token" host/code

websocket connection by session
http://127.0.0.1:8080/ws/login?3rd_session=xxx

get request

/show/prize
/show/gate
/show/shop
/show/prop

post request

/set/prize
json = {
    "id": "5b8a252bc3666ed1e4225d99",
    "day": 0,
    "prize": [
      {
        "type": 2,
        "number": 10000
      },
      {
        "type": 10,
        "number": 2
      },
      {
        "type": 6,
        "number": 1
      },
      {
        "type": 7,
        "number": 1
      }
    ],
    "del": 0,
    "ctime": "2018-09-01T13:35:39.88+08:00"
  }
curl -d "json" -H "token:your-token" host/set/prize

/set/gate
json = {
    "id": "ObjectIdHex(\"5b8a252bc3666ed1e4225d99\")",
    "gateid": 1,
    "type": 1,
    "star": 3,
    "data": [],
    "temp_shop": ["17", "18", "19"],
    "prize": [],
    "incr": true,
    "del": 0,
    "ctime": "2018-09-01T13:35:39.88+08:00"
  }
curl -d "json" host/set/gate

/set/shop
json = {
    "id": "1",
    "status": 0,
    "ptype": 1,
    "payway": 0,
    "number": 120,
    "price": 12,
    "name": "钻石",
    "info": "12元兑换120钻石",
    "prize": [
      {
        "type": 1,
        "number": 4
      }
    ],
    "del": 0,
    "etime": "2018-12-10T13:35:02.521391891+08:00",
    "ctime": "2018-09-01T13:35:02.521+08:00"
  }
curl -d "json" -H "token:your-token" host/set/shop

/set/prop
json = {
    "id": "4",
    "type": 4,
    "name": "提示卡",
    "attr": 0,
    "ctime": "2018-09-01T13:35:02.521+08:00"
  }
curl -d "json" -H "token:your-token" host/set/prop

/set/coin
curl -d "userid=xxx&num=xxx" -H "token:your-token" host/set/coin

/set/diamond
curl -d "userid=xxx&num=xxx" -H "token:your-token" host/set/diamond

/set/close
curl -d "" -H "token:your-token" host/set/close

```
