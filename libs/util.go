package libs

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

// ExecCmd 执行shell命令
func ExecCmd(c string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	return out, err
}

//Load load info by json
func Load(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

//TodayTime today time
func TodayTime() time.Time {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return today
}

//RandStr ...
func RandStr(i int) (s string) {
	for i > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		s += strconv.FormatInt(r.Int63n(10), 10)
		i--
	}
	return
}
