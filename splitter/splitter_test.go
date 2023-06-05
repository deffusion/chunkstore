package splitter

import (
	"fmt"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"github.com/deffusion/chunkstore/merkle"
	"github.com/deffusion/chunkstore/store"
	"log"
	"os"
	"os/user"
	"testing"
)

func TestSplitter(t *testing.T) {
	file, err := os.Open("test.pdf")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	digests, err := SplitIntoFiles(store.ChunkRoot, file, digest_hash.SHA256)
	if err != nil {
		log.Fatal("split err:", err)
	}
	fmt.Println(digests)
	root, err := merkle.Root(digests)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("root:", root)
}

func TestUserDir(t *testing.T) {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentUser.HomeDir)
}
