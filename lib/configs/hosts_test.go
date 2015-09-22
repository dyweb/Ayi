package configs_test

import (
	"testing"
	"github.com/dyweb/Ayi/lib/configs"
)

func TestHostRead(t *testing.T) {
	file, err := configs.ReadHostsFile("../../fixture/hosts")
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	if len(file) > 0 {
		t.Skipped()
	}else {
		t.Fail()
		t.Log("hosts file should not be empty")
	}
}