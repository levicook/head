package head

import (
	"bytes"
	"io"
)

type Header interface {
	Head() []byte
	Reset()
}

type header struct {
	bytes.Buffer
	limit int
}

func (h header) Head() []byte {
	return h.Bytes()
}

func (h *header) stream(p []byte) (n int, err error) {
	n = len(p)

	l := h.limit - h.Len()
	if h.limit > 0 && l > 0 {
		_, err = io.CopyN(h, bytes.NewReader(p), int64(l))
	}

	return
}
