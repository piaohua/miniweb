/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-27 21:26:59
 * Filename      : gen.go
 * Description   : 生成工具
 * *******************************************************/

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	Init()
	//Gen()
}

//TODO 内部协议通信
//TODO 通过json配置

var (
	protoPacket = make(map[string]uint32) //响应协议
	protoUnpack = make(map[string]uint32) //请求协议
	protoSid    = make(map[uint32]uint32) //映射关系
	//
	packetPath  = "../pb/packet.go"  //打包协议文件路径
	unpackPath  = "../pb/unpack.go"  //解包协议文件路径
	rPacketPath = "../pb/rpacket.go" //机器人打包协议
	rUnpackPath = "../pb/runpack.go" //机器人解包协议
	//
	packetFunc  = "Packet"  //打包协议函数名字
	unpackFunc  = "Unpack"  //解包协议函数名字
	rPacketFunc = "Rpacket" //机器人打包协议函数名字
	rUnpackFunc = "Runpack" //机器人解包协议函数名字
	//
	luaPath  = "MsgID.lua"  //lua文件
	jsPath   = "MsgID.js"   //js文件
	jsonPath = "MsgID.json" //json文件
)

type proto struct {
	name string
	code uint32
}

var protosUnpack = map[string][]proto{
	//game
	"game": {
		//login
		{code: 1001, name: "CWxLogin"},
		{code: 1002, name: "CLogin"},
		//user
		{code: 1004, name: "CUserData"},
		{code: 1005, name: "CGateData"},
		{code: 1006, name: "CPing"},
		{code: 1007, name: "CPropData"},
		{code: 1008, name: "CGetCurrency"},
		//shop
		{code: 1011, name: "CShop"},
		{code: 1012, name: "CBuy"},
		//play
		{code: 1013, name: "COverData"},
		{code: 1014, name: "CCard"},
		{code: 1015, name: "CLoginPrize"},
		{code: 1016, name: "CUseProp"},
		{code: 1017, name: "CStart"},
	},
}

var protosPacket = map[string][]proto{
	//game
	"game": {
		//login
		{code: 1001, name: "SWxLogin"},
		{code: 1002, name: "SLogin"},
		{code: 1003, name: "SLoginOut"},
		//user
		{code: 1004, name: "SUserData"},
		{code: 1005, name: "SGateData"},
		{code: 1006, name: "SPing"},
		{code: 1007, name: "SPropData"},
		{code: 1008, name: "SGetCurrency"},
		{code: 1009, name: "SPushCurrency"},
		{code: 1010, name: "SPushProp"},
		//shop
		{code: 1011, name: "SShop"},
		{code: 1012, name: "SBuy"},
		//play
		{code: 1013, name: "SOverData"},
		{code: 1014, name: "SCard"},
		{code: 1015, name: "SLoginPrize"},
		{code: 1016, name: "SUseProp"},
		{code: 1017, name: "SStart"},
	},
}

var sids = map[string]uint32{
	"game": 0,
}

//Init 初始化
func Init() {
	var packetStr string
	var unpackStr string
	var rpacketStr string
	var runpackStr string
	//request
	for _, k := range []string{"game"} {
		m := protosUnpack[k] //有序
		//for k, m := range protosUnpack {
		//初始化
		protoPacket = make(map[string]uint32) //响应协议
		protoUnpack = make(map[string]uint32) //请求协议
		protoSid = make(map[uint32]uint32)    //映射关系
		//组装
		for _, v := range m {
			registUnpack(v.name, v.code)
		}
		for _, v := range protosPacket[k] {
			registPacket(v.name, v.code)
			protoSid[v.code] = sids[k]
		}
		//最后生成文件
		prefix := k + "-" //文件前缀
		//genMsgID(prefix)
		//genjsMsgID(prefix)
		//genjsonMsgID(prefix)
		////打包组装
		//packetStr += bodyPacket()
		//unpackStr += bodyUnpack()
		//rpacketStr += bodyClientPacket()
		//runpackStr += bodyClientUnpack()
		//有序
		genMsgID2(prefix, protosPacket[k], m)
		genjsMsgID2(prefix, protosPacket[k], m)
		genjsonMsgID2(prefix, protosPacket[k], m)
		//打包组装
		packetStr += bodyPacket2(protosPacket[k])
		unpackStr += bodyUnpack2(m)
		rpacketStr += bodyClientPacket2(m)
		runpackStr += bodyClientUnpack2(protosPacket[k])
	}
	//server
	genPacket(packetStr)
	genUnpack(unpackStr)
	//client
	genClientPacket(rpacketStr)
	genClientUnpack(runpackStr)
}

func registUnpack(key string, code uint32) {
	if _, ok := protoUnpack[key]; ok {
		panic(fmt.Sprintf("%s registered: %d", key, code))
	}
	protoUnpack[key] = code
}

func registPacket(key string, code uint32) {
	if _, ok := protoPacket[key]; ok {
		panic(fmt.Sprintf("%s registered: %d", key, code))
	}
	protoPacket[key] = code
}

//Gen 生成文件
func Gen() {
	////server
	//genPacket()
	//genUnpack()
	////client
	//genClientPacket()
	//genClientUnpack()
}

//生成打包文件
func genPacket(body string) {
	var str string
	str += headPacket()
	//str += bodyPacket()
	str += body
	str += endPacket2()
	err := ioutil.WriteFile(packetPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成解包文件
func genUnpack(body string) {
	var str string
	str += headUnpack()
	//str += bodyUnpack()
	str += body
	str += endUnpack()
	err := ioutil.WriteFile(unpackPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

func bodyUnpack() string {
	var str string
	for k, v := range protoUnpack {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v, k, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v, k, bodyUnpackCode(v), resultUnpack())
	}
	return str
}

func bodyPacket() string {
	var str string
	for k, v := range protoPacket {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, k, resultPacket2(v))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, bodyPacketCode(v, k), k, resultPacket2(v))
	}
	return str
}

//有序
func bodyUnpack2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v.code, v.name, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v.code, v.name, bodyUnpackCode(v.name), resultUnpack())
	}
	return str
}

func bodyPacket2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, v.name, resultPacket2(v.code))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, bodyPacketCode(v.code, v.name), v.name, resultPacket2(v.code))
	}
	return str
}

func bodyUnpackCode(code uint32) (str string) {
	str = fmt.Sprintf("//msg.Code = %d", code)
	return
}

func bodyPacketCode(code uint32, name string) (str string) {
	str = fmt.Sprintf("//msg.(*%s).Code = %d", name, code)
	return
}

func headPacket() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Packet 打包消息
func Packet(msg interface{}) (uint32, uint32, []byte, error) {
	switch msg.(type) {
	`)
}

func headUnpack() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Unpack 解包消息
func Unpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	`)
}

func headRpacket() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Rpacket 打包消息
func Rpacket(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	`)
}

func headRunpack() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Runpack 解包消息
func Runpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	`)
}

func resultPacket(code uint32) string {
	return fmt.Sprintf("return %d, b, err", code)
}

func resultPacket2(code uint32) string {
	return fmt.Sprintf("return %d, %d, b, err", code, protoSid[code])
}

func resultUnpack() string {
	return fmt.Sprintf(`err := msg.Unmarshal(b)
		return msg, err`)
}

func endPacket() string {
	return fmt.Sprintf(`default:
		return 0, []byte{}, errors.New("unknown message")
	}
}`)
}

func endPacket2() string {
	return fmt.Sprintf(`default:
		return 0, 0, []byte{}, errors.New("unknown message")
	}
}`)
}

func endUnpack() string {
	return fmt.Sprintf(`default:
		return nil, errors.New("unknown message")
	}
}`)
}

//生成lua文件
func genMsgID(prefix string) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for k, v := range protoUnpack {
		str += fmt.Sprintf("\n\t%s = %d,", k, v)
	}
	str += fmt.Sprintf("\n")
	for k, v := range protoPacket {
		str += fmt.Sprintf("\n\t%s = %d,", k, v)
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+luaPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成js文件
func genjsMsgID(prefix string) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for k, v := range protoUnpack {
		str += fmt.Sprintf("\n\t%s : %d,", k, v)
	}
	str += fmt.Sprintf("\n")
	length := len(protoPacket)
	var i int
	for k, v := range protoPacket {
		i++
		if i == length {
			str += fmt.Sprintf("\n\t%s : %d", k, v)
		} else {
			str += fmt.Sprintf("\n\t%s : %d,", k, v)
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//{
//	3028:{type:"room",        sendType:"protocol.CChat",            revType:"protocol.SChat",           },
//}
func genjsonMsgID(prefix string) {
	var str string
	str += fmt.Sprintf("{")
	m := make(map[uint32]bool) //协议id不唯一时去重
	//每条协议id唯一
	for k, v := range protoUnpack { //响应
		rsp := ""
		for k2, v2 := range protoPacket { //请求
			if v == v2 {
				rsp = k2
				break
			}
		}
		if len(rsp) == 0 {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"%s\",\t\t},", v, k, rsp)
		} else {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v, k, rsp)
		}
		m[v] = true
	}
	//
	length := len(protoPacket)
	var i int
	for k, v := range protoPacket { //响应
		if _, ok := m[v]; ok {
			continue
		}
		rsp := ""
		for k2, v2 := range protoUnpack { //请求
			if v == v2 {
				rsp = k2
				break
			}
		}
		i++
		if i == length {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t}", v, rsp, k)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t}", v, rsp, k)
			}
		} else {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t},", v, rsp, k)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v, rsp, k)
			}
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsonPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//有序
//生成lua文件
func genMsgID2(prefix string, ps, unps []proto) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for _, v := range unps {
		str += fmt.Sprintf("\n\t%s = %d,", v.name, v.code)
	}
	str += fmt.Sprintf("\n")
	for _, v := range ps {
		str += fmt.Sprintf("\n\t%s = %d,", v.name, v.code)
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+luaPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成js文件
func genjsMsgID2(prefix string, ps, unps []proto) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for _, v := range unps {
		str += fmt.Sprintf("\n\t%s : %d,", v.name, v.code)
	}
	str += fmt.Sprintf("\n")
	length := len(ps)
	var i int
	for _, v := range ps {
		i++
		if i == length {
			str += fmt.Sprintf("\n\t%s : %d", v.name, v.code)
		} else {
			str += fmt.Sprintf("\n\t%s : %d,", v.name, v.code)
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//{
//	3028:{type:"room",        sendType:"protocol.CChat",            revType:"protocol.SChat",           },
//}
func genjsonMsgID2(prefix string, ps, unps []proto) {
	var str string
	str += fmt.Sprintf("{")
	m := make(map[uint32]bool) //协议id不唯一时去重
	//每条协议id唯一
	for _, v := range unps { //响应
		rsp := ""
		for _, v2 := range ps { //请求
			if v.code == v2.code {
				rsp = v2.name
				break
			}
		}
		if len(rsp) == 0 {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"%s\",\t\t},", v.code, v.name, rsp)
		} else {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v.code, v.name, rsp)
		}
		m[v.code] = true
	}
	//
	length := len(ps)
	var i int
	for _, v := range ps { //响应
		if _, ok := m[v.code]; ok {
			continue
		}
		rsp := ""
		for _, v2 := range unps { //请求
			if v.code == v2.code {
				rsp = v2.name
				break
			}
		}
		i++
		if i == length {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t}", v.code, rsp, v.name)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t}", v.code, rsp, v.name)
			}
		} else {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t},", v.code, rsp, v.name)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v.code, rsp, v.name)
			}
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsonPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成机器人打包文件
func genClientPacket(body string) {
	var str string
	str += headRpacket()
	//str += bodyClientPacket()
	str += body
	str += endPacket()
	err := ioutil.WriteFile(rPacketPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成机器人解包文件
func genClientUnpack(body string) {
	var str string
	str += headRunpack()
	//str += bodyClientUnpack()
	str += body
	str += endUnpack()
	err := ioutil.WriteFile(rUnpackPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

func bodyClientPacket() string {
	var str string
	for k, v := range protoUnpack {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, k, resultPacket(v))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, bodyClientPacketCode(v, k), k, resultPacket(v))
	}
	return str
}

func bodyClientUnpack() string {
	var str string
	for k, v := range protoPacket {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v, k, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v, k, bodyClientUnpackCode(v), resultUnpack())
	}
	return str
}

//有序
func bodyClientPacket2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, v.name, resultPacket(v.code))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, bodyClientPacketCode(v.code, v.name), v.name, resultPacket(v.code))
	}
	return str
}

func bodyClientUnpack2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v.code, v.name, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v.code, v.name, bodyClientUnpackCode(v.code), resultUnpack())
	}
	return str
}

func bodyClientUnpackCode(code uint32) (str string) {
	str = fmt.Sprintf("//msg.Code = %d", code)
	return
}

func bodyClientPacketCode(code uint32, name string) (str string) {
	str = fmt.Sprintf("//msg.(*%s).Code = %d", name, code)
	return
}
