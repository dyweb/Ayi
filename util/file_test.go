package util

import (
	"testing"
)

func TestRemoveLine(t *testing.T) {
	fixtureFile := "../fixture/text.txt"
	err := RemoveLine(fixtureFile, 1)
	// TODO: real test logic
	t.Log(err)
}
