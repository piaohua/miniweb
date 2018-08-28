package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	//Cache userinfo Cache
	Cache cache.Cache
)

//TODO redis cache
func init() {
	var err error
	Cache, err = cache.NewCache("memory", `{"interval":0}`)
	//Cache, err = cache.NewCache("redis", `{"key":"collectionName","conn":"127.0.0.1:6039","dbNum":"0","password":"thePassWord"}`)
	if err != nil {
		beego.Error("init cache err : ", err)
		panic(err)
	}
}
