package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	_ "miniweb/routers"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestWS is a sample to run an endpoint test
func TestWS(t *testing.T) {
	ws := "ws://localhost:8080/ws/login?3rd_session=813c37947b55ca2fe2648cf6e91912df"
	beego.Trace("ws %s\n", ws)
	c, _, err := websocket.DefaultDialer.Dial(ws, nil)
	beego.Trace("err -> %+v\n", err)
	if err != nil {
		beego.Trace("err -> %+v\n", err)
	}
	if c != nil {
		<-time.After(time.Duration(1) * time.Second)
		buff := make([]byte, 0)
		beego.Trace("send msg: %d\n", len(buff))
		c.WriteMessage(websocket.BinaryMessage, buff)
		<-time.After(time.Duration(1) * time.Second)
		beego.Trace("close conn\n")
		c.Close()
	}
}
