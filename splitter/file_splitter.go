package splitter

import (
	"fmt"
	"github.com/deffusion/chunkstore/chunker"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"io"
	"log"
	"os"
)

const (
	K         = 1024
	M         = K * K
	ChunkSize = 256 * K
)

func SplitIntoFiles(rootPath string, r io.Reader, h digest_hash.Hash) ([]digest.Digest, error) {
	rc := chunker.NewRabin(r, ChunkSize)
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
		cFile, err := os.Create(fmt.Sprintf("%s%s", rootPath, d))
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
