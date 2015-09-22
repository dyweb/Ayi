package configs

import (
	"testing"
	"regexp"
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

func TestRegexp(t *testing.T) {
	r, _ := regexp.Compile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})\\s*(\\S*)$")
	if r.MatchString("127.0.0.1 localhost") == false {
		t.Fail()
		t.Log("can't match host config")
	}
	a := r.FindStringSubmatch("127.0.0.1 localhost")
	if len(a) != 3 || a[0] != "127.0.0.1 localhost" || a[1] != "127.0.0.1" || a[2] != "localhost" {
		t.Fail()
		t.Log("can't mactch host ip and name from config")
	}

}

func TestParseHost(t *testing.T) {
	host := parseHost("127.0.0.1	localhost")
	if host.name != "localhost" || host.ip != "127.0.0.1" {
		t.Fail()
		t.Log("Can't parse single line of host")
	}
}