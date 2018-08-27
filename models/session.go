/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-25 18:36:20
 * Filename      : session.go
 * Description   : 生成session
 * *******************************************************/

package models

import (
	"errors"
	"time"

	"miniweb/libs"

	"github.com/astaxie/beego"
)

var (
	// Channel for new session.
	sessionCh chan string
)

//SessionChInit sessionCh init
func SessionChInit() {
	n, err := beego.AppConfig.Int("session.cache")
	if err != nil {
		beego.Error("session.cache config err: ", err)
	}
	if n <= 0 {
		n = 10
	}
	sessionCh = make(chan string, n)
	go genSession()
}

//gen session
func genSession() {
	for {
		session, err := GenSession()
		if err != nil {
			beego.Error("genSession err: ", err)
			break
		}
		if len(session) == 0 {
			beego.Error("genSession len err")
			break
		}
		if Cache.IsExist(session) {
			continue
		}
		//if HasSession(session) {
		//	continue
		//}
		sessionCh <- session
	}
}

//get new session
func getNewSession() (string, error) {
	select {
	case s, ok := <-sessionCh:
		if !ok {
			return "", errors.New("sessionCh closed")
		}
		return s, nil
	case <-time.After(time.Millisecond * 200):
	}
	return "", errors.New("get session timeout")
}

//GenSession 生成session
func GenSession() (string, error) {
	//c := "cat /dev/urandom | od -x | tr -d ' ' | head -n 1"
	c := "head /dev/urandom | od -x | tr -d ' ' | cut -c8- | head -c 32"
	out, err := libs.ExecCmd(c)
	if err != nil {
		beego.Error("GenSession err: ", err)
	}
	return string(out), err
}
