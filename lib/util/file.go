package util

import (
	"bufio"
	"os"
)

// lineToRemove starts from 1
func RemoveLine(fileName string, lineToRemove int) (error) {
	r, err := os.Open(fileName)
	defer r.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(r)
	text := ""
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber != lineToRemove {
			text += scanner.Text() + "\n"
		}
	}
	w, err := os.Create(fileName)
	defer w.Close()
	w.WriteString(text)
	return w.Sync()
}