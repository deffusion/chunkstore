package store

import (
	"fmt"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/store/kv/memory_kv"
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

func add(cs *ChunkStore, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	root := cs.Add(file)
	fmt.Println("added: ", root)
}

func get(cs *ChunkStore, d digest.Digest) {
	ds := cs.Get(d)
	if ds != nil {
		fmt.Println(ds)
	}
}

func TestChunkStore_Add(t *testing.T) {
	db := memory_kv.New()
	cs := New(db, ChunkRoot)
	add(cs, "../splitter/test.pdf")
}

func TestChunkStore_Get(t *testing.T) {
	db := memory_kv.New()
	cs := New(db, ChunkRoot)
	d, err := digest.New("s8fa18167acc80d63070ccaf4001865f5b37027980ac711911684cecd51b79360e")
	if err != nil {
		log.Fatal(err)
	}
	get(cs, d)
	add(cs, "../splitter/test.pdf")
	get(cs, d)
}
