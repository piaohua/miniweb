/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-25 12:26:39
 * Filename      : genid.go
 * Description   : 生成id
 * *******************************************************/

package models

import (
	"errors"
	"miniweb/libs"
	"time"

	"github.com/astaxie/beego"
)

var (
	// Channel for new code.
	gencodeCh chan string
)

//GenCodeChInit gencodeCh init
func GenCodeChInit() {
	n, err := beego.AppConfig.Int("genid.cache")
	if err != nil {
		beego.Error("gen.cache config err: ", err)
	}
	if n <= 0 {
		n = 10
	}
	gencodeCh = make(chan string, n)
	go genCode()
}

//gen code
func genCode() {
	for {
		code := libs.RandStr(6)
		if len(code) != 6 {
			beego.Error("genCode len err")
			break
		}
		if HasCode(code) {
			continue
		}
		gencodeCh <- code
	}
}

//get new code
func getNewCode() (string, error) {
	select {
	case s, ok := <-gencodeCh:
		if !ok {
			return "", errors.New("gencodeCh closed")
		}
		if HasCode(s) {
			return getNewCode()
		}
		return s, nil
	case <-time.After(time.Millisecond * 200):
	}
	return "", errors.New("get code timeout")
}
