package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	assert := assert.New(t)
	cmd, _ := Command("ls")
	assert.Equal(1, len(cmd.Args))
	// FIXME: $(glide novendor) will not be interpreted
	// TODO: try sh -c "go test -v -cover $(glide novendor)"
	cmd, _ = Command("go test -v -cover $(glide novendor)")
	assert.Equal("test", cmd.Args[1])
	cmd, _ = Command("sh -c \"go test -v -cover $(glide novendor)\"")
	assert.Equal("go test -v -cover $(glide novendor)", cmd.Args[2])
}

func TestRunCommand(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(RunCommand("sh -c \"echo Hi\""))
}
