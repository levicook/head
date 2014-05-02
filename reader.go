package head

import "io"

type HeaderReader interface {
	Header
	io.Reader
}

func NewReader(limit int) HeaderReader {
	r := new(reader)
	r.limit = limit
	return r
}

type reader struct{ header }

func (r *reader) Read(p []byte) (n int, err error) {
	return r.stream(p)
}
