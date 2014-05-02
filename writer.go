package head

import "io"

type HeaderWriter interface {
	Header
	io.Writer
}

func NewWriter(limit int) HeaderWriter {
	w := new(writer)
	w.limit = limit
	return w
}

type writer struct{ header }

func (w *writer) Write(p []byte) (n int, err error) {
	return w.stream(p)
}
