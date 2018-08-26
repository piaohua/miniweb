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
	if err != nil {
		beego.Error("init cache err : ", err)
		panic(err)
	}
}
