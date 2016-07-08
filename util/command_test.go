package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := Command("ls")
	assert.Equal(1, len(cmd.Args))
	cmd = Command("go test -v -cover $(glide novendor)")
	assert.Equal("test", cmd.Args[1])
}
