package util

import (
	"os"
)

// FileExists return if a file exists and can be read by current program
// http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
