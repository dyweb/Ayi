package embed

import (
	"testing"

	"github.com/GeertJohan/go.rice"
	"github.com/stretchr/testify/assert"
)

func TestFindBox(t *testing.T) {
	assert := assert.New(t)
	box, err := rice.FindBox("fixtures")
	assert.Nil(err)
	var content string
	content, err = box.String("a.md")
	assert.Equal("# A", content)
}
