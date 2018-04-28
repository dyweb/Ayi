package github

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestReadLabelConfig(t *testing.T) {
	assert := asst.New(t)

	labels, err := ReadLabelConfig("testdata/at15.label.yml")
	assert.Nil(err)
	t.Log(labels)
}

func TestFlattenLabelConfigs(t *testing.T) {
	assert := asst.New(t)

	configs, err := ReadLabelConfig("testdata/at15.label.yml")
	assert.Nil(err)
	labels := FlattenLabelConfigs(configs)
	t.Log(labels)
}
