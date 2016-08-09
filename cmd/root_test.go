package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestBindFlag(t *testing.T) {
	assert := assert.New(t)
	RootCmd.SetArgs([]string{""})
	_ = RootCmd.Execute()
	assert.Equal(false, viper.GetBool("Verbose"))
	assert.Equal(false, viper.GetBool("DryRun"))
	RootCmd.SetArgs([]string{"-v", "-n"})
	_ = RootCmd.Execute()
	assert.Equal(true, viper.GetBool("Verbose"))
	assert.Equal(true, viper.GetBool("DryRun"))
}
