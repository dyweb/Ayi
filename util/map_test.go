package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasKey(t *testing.T) {
	assert := assert.New(t)
	// FIXME: cannot cannot use map[string]int as map[string]interface{}
	// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	// var ages map[string]int
	// ages["jack"] = 12
	// assert.True(HasKey(ages, "jack"))
	ages := make(map[string]interface{})
	ages["jack"] = 12
	assert.True(HasKey(ages, "jack"))
	assert.False(HasKey(ages, "jimmy"))
}

func TestGetWithDeault(t *testing.T) {
	assert := assert.New(t)
	ages := make(map[string]interface{})
	ages["jack"] = 12
	assert.Equal(12, GetWithDefault(ages, "jack", 10086))
	assert.Equal(10086, GetWithDefault(ages, "jimmy", 10086))
}
