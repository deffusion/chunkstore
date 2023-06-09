package store

import (
	"fmt"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/store/kv/memory_kv"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
)

func TestPathExist(t *testing.T) {
	_, err := os.Stat(ChunkRoot)
	fmt.Println("check dir:", ChunkRoot)
	if os.IsNotExist(err) {
		fmt.Printf("directory %s is not exist\n", ChunkRoot)
		err := os.MkdirAll(ChunkRoot, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("directory %s was created\n", ChunkRoot)
	}
}

func add(cs *ChunkStore, path string) digest.Digest {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return cs.Add(file)
}

func get(cs *ChunkStore, d digest.Digest) {
	ds := cs.Get(d)
	if ds != nil {
		fmt.Println(ds)
	}
}

func TestChunkStore_Add(t *testing.T) {
	db := memory_kv.New()
	defer db.Close()
	logger, _ := zap.NewProduction(nil)
	cs := New(db, ChunkRoot, logger)
	add(cs, "../splitter/test.pdf")
}

func TestChunkStore_Get(t *testing.T) {
	db := memory_kv.New()
	defer db.Close()
	logger, _ := zap.NewProduction(nil)
	cs := New(db, ChunkRoot, logger)
	root := add(cs, "D:/book/gopl2.pdf")
	fmt.Println("added: ", root)
	//get(cs, root)
}
