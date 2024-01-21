package script_test

import (
	"testing"

	script "github.com/238Studio/child-nodes-script-caller"
)

func TestCaller(t *testing.T) {
	re1, err := script.Caller("./test.sh", "1", "2")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(re1))

	s := script.RegisterScriptCaller("", "", "", "./test.sh")
	re, err := s.Caller("1", "2")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(re))
}
