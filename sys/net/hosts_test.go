package net

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	r, _ := regexp.Compile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})\\s*(\\S*)\\s*$")
	if r.MatchString("127.0.0.1 localhost") == false {
		t.Fail()
		t.Log("can't match host config")
	}
	a := r.FindStringSubmatch("127.0.0.1 localhost")
	if len(a) != 3 || a[0] != "127.0.0.1 localhost" || a[1] != "127.0.0.1" || a[2] != "localhost" {
		t.Fail()
		t.Log("can't mactch host ip and name from config")
	}
	configWithTab := "127.0.0.1 \t\tlocalhost"
	a = r.FindStringSubmatch(configWithTab)
	if len(a) != 3 || a[2] != "localhost" {
		t.Fail()
		//		t.Log("host is " + a[2])
		t.Log("can't match with tab")
	}
	configWithTrailingSpace := "127.0.0.1 localhost "
	a = r.FindStringSubmatch(configWithTrailingSpace)
	if len(a) != 3 || a[2] != "localhost" {
		t.Fail()
		//		t.Log("host is " + a[2])
		t.Log("can't match with extra space")
	}

}

func TestRemoveComment(t *testing.T) {
	noComment := "Dia da bu di da bu"
	if noComment != removeComment(noComment) {
		t.Fail()
		t.Log("should return original string if no comment")
	}
	prefixComment := "# I have a dream"
	if "" != removeComment(prefixComment) {
		t.Fail()
		t.Log("should remove the whole line if found # at start")
	}
	suffixComment := "roast mie! # la da di da di "
	if "roast mie! " != removeComment(suffixComment) {
		t.Fail()
		t.Log("should only remove the commented part")
	}
}

func TestParseHost(t *testing.T) {
	host, _ := parseHost("127.0.0.1	localhost", 1)
	if host.name != "localhost" || host.ip != "127.0.0.1" || host.line != 1 {
		t.Fail()
		t.Log("Can't parse single line of host")
	}
}

func TestParseHostsFile(t *testing.T) {
	fixtureHostsFile := "../../fixture/hosts"
	hosts := parseHostsFile(fixtureHostsFile)
	if len(hosts) != 3 {
		t.Fail()
		t.Log("Didn't parse all hosts file")
	}
}

func TestAddHostToFile(t *testing.T) {
	fixtureHostsFile := "../../fixture/hosts"
	_, err := addHostToFile(fixtureHostsFile, "127.0.0.1", "localhost")
	if err == nil {
		t.Fail()
		t.Log("Should not add if old one already exists")
	}
	_, err = addHostToFile(fixtureHostsFile, "127.0.0.1", "doubi.lk")
	t.Log(err)
	if err != nil {
		t.Log(err)
		t.Log("Can't add new host")
		t.Fail()
	}
}

func TestRemoveHostFromFile(t *testing.T) {
	fixtureHostsFile := "../../fixture/hosts"
	removed, err := removeHostFromFile(fixtureHostsFile, "doubi.lk")
	if !removed {
		t.Fail()
		t.Log(err)
	}
}
