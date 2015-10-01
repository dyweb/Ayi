package util

import (
	"bufio"
	"os"
)

// lineToRemove starts from 1
func RemoveLine(fileName string, lineToRemove int) (error) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(file)
	text := ""
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber != lineToRemove {
			text += scanner.Text() + "\n"
		}
	}
	writer.WriteString(text)
	return writer.Flush()
}