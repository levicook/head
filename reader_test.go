package head

import (
	"reflect"
	"testing"
)

func Test_NewReader(t *testing.T) {
	var deepEqual = reflect.DeepEqual

	r := NewReader(1)
	r.Read([]byte{'1', '2'})
	if !deepEqual([]byte{'1'}, r.Head()) {
		t.Fatalf("%s", r.Head())
	}

	r.Reset()
	r.Read([]byte{'3', '4', '5'})
	if !deepEqual([]byte{'3'}, r.Head()) {
		t.Fatalf("%s", r.Head())
	}

	r.Reset()
	r.Read([]byte{'6'})
	r.Read([]byte{'7'})
	r.Read([]byte{'8'})
	if !deepEqual([]byte{'6'}, r.Head()) {
		t.Fatalf("%s", r.Head())
	}
}
