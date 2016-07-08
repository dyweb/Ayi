package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := Command("ls")
	assert.Equal(1, len(cmd.Args))
	// FIXME: $(glide novendor) will not be interpreted
	// TODO: try sh -c "go test -v -cover $(glide novendor)"
	cmd = Command("go test -v -cover $(glide novendor)")
	assert.Equal("test", cmd.Args[1])
	// FIXME: this is also broken
	// assert.Nil(RunCommand("sh -c \"echo Hi\""))
}
