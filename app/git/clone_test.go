package git

import (
	"regexp"
	"testing"
)

func TestRegexpGitHub(t *testing.T) {
	gh := "https://github.com/dyweb/Ayi"
	r, _ := regexp.Compile(browserRegexp)
	if r.MatchString(gh) == false {
		t.Fail()
		t.Log("cant match github url")
	}
	a := r.FindStringSubmatch(gh)
	if a[1] != "https" {
		t.Fail()
		t.Log("cant match protocol")
	}
	if a[2] != "github.com" {
		t.Fail()
		t.Log("cant match host")
	}
	if a[3] != "dyweb" {
		t.Fail()
		t.Log("cant match user name/org name")
	}
	if a[4] != "Ayi" {
		t.Fail()
		t.Log("cant match repo name")
	}
}

func TestRegexpGitLab(t *testing.T) {
	gh := "http://git.dongyue.io/dy/support"
	r, _ := regexp.Compile(browserRegexp)
	if r.MatchString(gh) == false {
		t.Fail()
		t.Log("cant match github url")
	}
	a := r.FindStringSubmatch(gh)
	if a[1] != "http" {
		t.Fail()
		t.Log("cant match protocol")
	}
	if a[2] != "git.dongyue.io" {
		t.Fail()
		t.Log("cant match host")
	}
	if a[3] != "dy" {
		t.Fail()
		t.Log("cant match user name/org name")
	}
	if a[4] != "support" {
		t.Fail()
		t.Log("cant match repo name")
	}
}