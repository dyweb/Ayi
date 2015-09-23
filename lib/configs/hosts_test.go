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

func TestRemoveComment(t *testing.T) {
	noComment := "Dia da bu di da bu"
	if noComment != removeComment(noComment){
		t.Fail()
		t.Log("should return original string if no comment")
	}
	prefixComment := "# I have a dream"
	if "" != removeComment(prefixComment){
		t.Fail()
		t.Log("should remove the whole line if found # at start")
	}
	suffixComment := "roast mie! # la da di da di "
	if "roast mie! " != removeComment(suffixComment){
		t.Fail()
		t.Log("should only remove the commented part")
	}
}

func TestParseHost(t *testing.T) {
	host, _ := parseHost("127.0.0.1	localhost")
	if host.name != "localhost" || host.ip != "127.0.0.1" {
		t.Fail()
		t.Log("Can't parse single line of host")
	}
}

func TestParseHostsFile(t *testing.T) {
	hosts := parseHostsFile("/etc/hosts")
	if len(hosts) == 0 {
		t.Fail()
		t.Log("Can't parse hosts file")
	}
}