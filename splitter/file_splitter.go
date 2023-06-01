package splitter

import (
	"crypto/sha1"
	"fmt"
	"github.com/deffusion/chunkstore/chunker"
	"io"
	"log"
	"os"
	"os/user"
)

const (
	K = 1024
	M = K * K
)

var storeRoot string

func init() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	storeRoot = fmt.Sprintf("%s/store/", currentUser.HomeDir)
}

func SplitIntoFiles(r io.Reader) error {
	h := sha1.New()
	rc := chunker.NewRabin(r, h, 1*M)

	for {
		chunk, err := rc.NextChunk()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		cFile, err := os.Create(fmt.Sprintf("%s/%x", storeRoot, chunk.Digest()))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := cFile.Write(chunk.Data()); err != nil {
			log.Fatal(err)
		}
		cFile.Close()
	}
	return nil
}
