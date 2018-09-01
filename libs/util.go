package libs

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
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
