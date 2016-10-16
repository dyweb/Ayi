package runner

import (
	"io"
)

// MultiWriter is different from io.MultiWriter in order to solve the isatty problme
// http://stackoverflow.com/questions/33452726/golang-difference-between-os-stdout-and-multiwriter/33768428
type MultiWriter struct {
	writers []io.Writer
}

// NewMultiWriter create a multiwriter that write in the order you created them
func NewMultiWriter(writers ...io.Writer) *MultiWriter {
	m := new(MultiWriter)
	m.writers = []io.Writer{}
	// for _, w := range writers {
	m.writers = append(m.writers, writers...)
	// }
	return m
}

func (m *MultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range m.writers {
		n, err = w.Write(p)
	}
	return n, err
}
