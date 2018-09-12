package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"miniweb/models"
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
	//ws := "wss://www.xxx.com/ws/login?3rd_session=e5e1d338add7094dab230d2fb8a42c82"
	beego.Trace("ws %s\n", ws)
	c, _, err := websocket.DefaultDialer.Dial(ws, nil)
	beego.Trace("err -> %+v\n", err)
	t.Logf("err -> %+v\n", err)
	if err != nil {
		beego.Trace("err -> %+v\n", err)
		t.Fatalf("err -> %+v\n", err)
	}
	if c != nil {
		<-time.After(time.Duration(1) * time.Second)
		buff := make([]byte, 0)
		beego.Trace("send msg: %d\n", len(buff))
		t.Logf("send msg: %d\n", len(buff))
		c.WriteMessage(websocket.BinaryMessage, buff)
		<-time.After(time.Duration(1) * time.Second)
		beego.Trace("close conn\n")
		t.Logf("close conn\n")
		c.Close()
	}
}

// TestSetGate set gate
func TestSetGate(t *testing.T) {
	gate := models.Gate{
		ID:       "ObjectIdHex(\"5b8a252bc3666ed1e4225d99\")",
		Gateid:   1,
		Type:     1,
		Star:     3,
		Data:     []byte{},
		TempShop: []string{"17", "18", "19"},
		//Prize: [],
		Incr:  true,
		Del:   0,
		Ctime: time.Now(),
	}
	body, err := json.Marshal(&gate)
	t.Logf("body %v, err %v\n", body, err)
	gate2 := new(models.Gate)
	err = json.Unmarshal(body, gate2)
	t.Logf("gate2 %#v, err %v\n", gate2, err)
}

// TestSetGate set gate
func TestSetClose(t *testing.T) {
	url := "http://127.0.0.1:8080/set/close"
	b, err := HTTPPost(url, []byte{})
	t.Logf("b %v, err %v\n", b, err)
	//curl -d "" -H "token:your-token" http://127.0.0.1:8080/set/close
}

// HTTPPost post request
func HTTPPost(url string, body []byte) (b []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", "your-token")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)

	return b, err
}
