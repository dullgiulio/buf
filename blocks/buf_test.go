package blocks_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/dullgiulio/buf/blocks"
)

func TestNosql(t *testing.T) {
	b := blocks.New()
	b.Write([]byte("01234567890abcdefghilmnopqrst///"))
	b.Write([]byte("---------------------------\000\000--\n"))
	bb := make([]byte, 512)
	if n, err := b.Read(bb); err != nil && err != io.EOF {
		t.Error(err)
	} else {
		bb = bb[:n]
	}
	expected := []byte("01234567890abcdefghilmnopqrst///---------------------------\000\000--\n")
	if bytes.Compare(bb, expected) != 0 {
		t.Error("Unexpected value read: ", bb)
	}
}
