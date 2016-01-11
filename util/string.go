package util

// ByteIndex is copied from https://golang.org/src/net/parser.go
func ByteIndex(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}
