package git

import "testing"

func TestBrowserRegexp(t *testing.T) {
	_, err := parseBrowserURL("https://github.com/dyweb/Ayi")
	if err != nil {
		t.Error(err)
	}
	_, err = parseBrowserURL("file:///D:/tmp/mapreduce.pdf")
	errMsg := "not a browser url"
	if !(err.Error() == errMsg) {
		t.Errorf("Error message should be '%s', got '%s' instead", errMsg, err.Error())
	}
}
