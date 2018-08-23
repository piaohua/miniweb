package libs

import (
	"testing"
)

func TestCmd(t *testing.T) {
	c := "cat /dev/urandom | od -x | tr -d ' ' | head -n 1"
	out, err := ExecCmd(c)
	t.Logf("out %s, err %v\n", string(out), err)
}
