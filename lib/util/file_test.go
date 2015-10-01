package util

import (
	"testing"
)

func TestRemoveLine(t *testing.T){
	fixtureFile := "../../fixture/text.txt"
	err := RemoveLine(fixtureFile,1)
	t.Log(err)
}