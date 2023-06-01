package chunker

import "github.com/whyrusleeping/chunker"

type BasicChunk struct {
	chunk *chunker.Chunk
}

func (c *BasicChunk) Data() []byte {
	return c.chunk.Data
}

func (c *BasicChunk) Digest() []byte {
	return c.chunk.Digest
}
