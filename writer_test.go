package head

import (
	"reflect"
	"testing"
)

func Test_NewWriter(t *testing.T) {
	var deepEqual = reflect.DeepEqual

	w := NewWriter(1)
	w.Write([]byte{'1', '2'})
	if !deepEqual([]byte{'1'}, w.Head()) {
		t.Fatalf("%s", w.Head())
	}

	w.Reset()
	w.Write([]byte{'3', '4', '5'})
	if !deepEqual([]byte{'3'}, w.Head()) {
		t.Fatalf("%s", w.Head())
	}

	w.Reset()
	w.Write([]byte{'6'})
	w.Write([]byte{'7'})
	w.Write([]byte{'8'})
	if !deepEqual([]byte{'6'}, w.Head()) {
		t.Fatalf("%s", w.Head())
	}
}
