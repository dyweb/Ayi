package util

import (
	"bufio"
	"os"

	"github.com/go-errors/errors"
)

// RemoveLine, line number starts from 1
func RemoveLine(fileName string, lineToRemove int) error {
	// read the file as string, skip the line to remove
	r, err := os.Open(fileName)
	defer r.Close()
	if err != nil {
		return errors.Wrap(err, 1)
	}
	scanner := bufio.NewScanner(r)
	text := ""
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber != lineToRemove {
			// TODO: what if the line ending is CR-LF
			text += scanner.Text() + "\n"
		}
	}

	// write the file
	w, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	defer w.Close()
	w.WriteString(text)
	err = w.Sync()
	if err != nil {
		return errors.Wrap(err, 1)
	}
	return nil
}
