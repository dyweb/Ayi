package configs

import (
	"testing"
)

func TestHostRead(t *testing.T) {
	_, err := readHostsFile("../../fixture/hosts")
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	//	if len(file) > 0 {
	//		t.Skipped()
	//	}else {
	//		t.Fail()
	//		t.Log("hosts file should not be empty")
	//	}
}

func TestParseHost(t *testing.T) {
	host := parseHost("127.0.0.1	localhost")
	if host.name != "localhost" || host.ip != "127.0.0.1" {
		t.Fail()
		t.Log("Can't parse single line of host")
	}
}