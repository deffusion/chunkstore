package chunker

import (
	"errors"
	"fmt"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"github.com/jotfs/fastcdc-go"
)

type BasicChunk struct {
	chunk  *fastcdc.Chunk
	digest digest.Digest
}

func (c *BasicChunk) Data() []byte {
	return c.chunk.Data
}

// ChunkDigest returns digest of the chunker used
//func (c *BasicChunk) ChunkDigest() []byte {
//	return c.chunk.Digest
//}

// Digest returns digest hashed by given hash function
func (c *BasicChunk) Digest(dh digest_hash.Hash) (digest.Digest, error) {
	errPrefix := "BasicChunk.Digest:"
	// compute if not computed before
	if c.digest == digest.Null {
		h, err := digest_hash.New(dh)
		if err != nil {
			return digest.Null, errors.New(fmt.Sprint(errPrefix, err))
		}
		h.Write(c.chunk.Data)
		c.digest, err = digest.New(h.Sum([]byte(dh)))
		if err != nil {
			return digest.Null, errors.New(fmt.Sprint(errPrefix, err))
		}
	}
	return c.digest, nil
}
