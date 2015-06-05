package nosql_test

import (
	"io"
	"testing"

	"github.com/dullgiulio/buf/nosql"
)

func TestNosql(t *testing.T) {
	b := nosql.New()
	b.Write([]byte("Some bytes"))
	b.Write([]byte("Some bytes"))
	bb := make([]byte, 10)
	if n, err := b.Read(bb); err != nil && err != io.EOF {
		t.Error(err)
	} else {
		if n != 10 {
			t.Error("Expected 10 bytes, got ", n)
		}
	}
}
