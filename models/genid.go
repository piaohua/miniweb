package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

var (
	// Channel for new id.
	genidCh chan string
)

//GenIDChInit genidCh init
func GenIDChInit() {
	n, err := beego.AppConfig.Int("genid.cache")
	if err != nil {
		beego.Error("gen.cache config err: ", err)
	}
	if n <= 0 {
		n = 10
	}
	genidCh = make(chan string, n)
	go genID()
}

//gen id
func genID() {
	for {
		id := GenUniqueID()
		if len(id) == 0 {
			beego.Error("genID len err")
			break
		}
		if Cache.IsExist(id) {
			continue
		}
		//if HasID(id) {
		//	continue
		//}
		genidCh <- id
	}
}

//get new id
func getNewID() (string, error) {
	select {
	case s, ok := <-genidCh:
		if !ok {
			return "", errors.New("genidCh closed")
		}
		return s, nil
	case <-time.After(time.Millisecond * 200):
	}
	return "", errors.New("get id timeout")
}

//GenUniqueID 生成id
func GenUniqueID() string {
	return bson.NewObjectId().Hex()
}
