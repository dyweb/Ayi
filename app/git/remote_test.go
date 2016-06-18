package git

import "testing"

func TestBrowserRegexp(t *testing.T) {
	_, err := parseBrowserURL("https://github.com/dyweb/Ayi")
	if err != nil {
		t.Fail()
		t.Log("cannot parse valid GitHub repo url")
	}
}
