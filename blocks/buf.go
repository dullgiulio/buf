package blocks

import "io"

const blocksize = 256

type Buffer struct {
	blocks [][blocksize]byte
	size   int
	read   int
}

func New() *Buffer {
	return &Buffer{blocks: make([][blocksize]byte, 0)}
}

func (b *Buffer) Cap() int {
	return blocksize * len(b.blocks)
}

func (b *Buffer) Len() int {
	return b.size
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	offset := b.size % blocksize
	lastblock := len(b.blocks) - 1

	for {
		if offset == 0 || offset >= blocksize {
			var block [blocksize]byte
			b.blocks = append(b.blocks, block)
			offset = 0
			lastblock = lastblock + 1
		}
		nwrite := len(p) - n
		if nwrite > blocksize {
			nwrite = blocksize
		}
		for ; n < nwrite; n++ {
			b.blocks[lastblock][offset] = p[n]
			offset = offset + 1
		}
		if n >= len(p) {
			break
		}
	}
	b.size = b.size + len(p)
	return len(p), nil
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	for ; n < cap(p); n++ {
		j := n + b.read
		if j >= b.size {
			return n, io.EOF
		}
		p[n] = b.blocks[j/blocksize][j%blocksize]
	}
	b.read = b.read + n
	return n, nil
}
