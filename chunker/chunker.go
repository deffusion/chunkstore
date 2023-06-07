package chunker

import (
	"github.com/jotfs/fastcdc-go"
	"io"
)

type Rabin struct {
	chunker *fastcdc.Chunker
}

func NewRabin(r io.Reader, avgBlkSize int) *Rabin {
	opts := fastcdc.Options{
		MinSize:     avgBlkSize / 4,
		AverageSize: avgBlkSize,
		MaxSize:     avgBlkSize * 4,
	}
	c, _ := fastcdc.NewChunker(r, opts)
	return &Rabin{
		c,
	}
}

func (r *Rabin) NextChunk() (*BasicChunk, error) {
	chunk, err := r.chunker.Next()
	if err != nil {
		return nil, err
	}
	return &BasicChunk{chunk: &chunk}, nil
}
