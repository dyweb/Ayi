package runner

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMultiWriter(t *testing.T) {
	assert := assert.New(t)
	mw := NewMultiWriter(os.Stdout, os.Stderr)
	assert.Equal(2, len(mw.writers))
}

func TestMultiWriterWrite(t *testing.T) {
	assert := assert.New(t)
	mw := NewMultiWriter(os.Stdout, os.Stderr)
	n, err := mw.Write([]byte("hi"))
	assert.Equal(2, n)
	assert.Nil(err)
}
