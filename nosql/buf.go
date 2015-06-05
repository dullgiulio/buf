package nosql

import (
	"bufio"
	"io"
)

type Buffer struct {
	size   int
	writes int
	ch     chan []byte
}

func New() *Buffer {
	return &Buffer{ch: make(chan []byte)}
}

// Read receives all data from the pending writes. (UNORDERED!)
func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.size == 0 {
		return n, io.EOF
	}
	for c := range b.ch {
		if len(c) > cap(p) {
			return n, bufio.ErrBufferFull
		}
		m := copy(p[n:], c)
		n = m + n
		b.writes = b.writes - 1
		if b.writes <= 0 {
			err = io.EOF
			break
		}
	}
	return n, err
}

// Write queues the data to be read later.
func (b *Buffer) Write(p []byte) (n int, err error) {
	go func() {
		b.ch <- p
	}()
	b.size = b.size + len(p)
	b.writes = b.writes + 1
	return len(p), nil
}

// Len gives the size of the data it will eventually read.
func (b *Buffer) Len() int { return b.size }

// Cap promises more than it can deliver.
func (b *Buffer) Cap() int { return b.size * 2 }
