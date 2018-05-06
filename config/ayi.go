// Package config contains config of all apps. Config strcut are not defined in their own package to avoid cycle dependencies
package config

// AyiConfig contains config struct of all apps
type AyiConfig struct {
	GitHub GitHubConfig `yaml:"github"`
}
