package splitter

import (
	"fmt"
	"github.com/deffusion/chunkstore/chunker"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"github.com/deffusion/chunkstore/store"
	"hash/fnv"
	"io"
	"log"
	"os"
)

const (
	K         = 1024
	M         = K * K
	ChunkSize = 1 * M
)

func SplitIntoFiles(r *os.File, h digest_hash.Hash) ([]digest.Digest, error) {
	rc := chunker.NewRabin(r, fnv.New32(), ChunkSize)
	var digests []digest.Digest
	for {
		chunk, err := rc.NextChunk()
		if err == io.EOF {
			break
		}
		if err != nil {
			return digests, err
		}
		d, err := chunk.Digest(h)
		if err != nil {
			log.Fatal(err)
		}
		cFile, err := os.Create(fmt.Sprintf("%s/%s", store.ChunkRoot, d))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := cFile.Write(chunk.Data()); err != nil {
			log.Fatal(err)
		}
		digests = append(digests, d)
		cFile.Close()
	}
	return digests, nil
}
