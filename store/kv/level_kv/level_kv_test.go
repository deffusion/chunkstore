package level_kv

import (
	"fmt"
	"github.com/deffusion/chunkstore/store"
	"github.com/deffusion/chunkstore/store/kv"
	"log"
	"os"
	"testing"
	"time"
)

func get(db kv.KV, k string) {
	v, err := db.Get([]byte(k))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("%s\n", v)
	}
}

func put(db kv.KV, k, v string) {
	err := db.Put([]byte(k), []byte(v))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
func TestR(t *testing.T) {
	levelKV, err := New(store.KVRoot)
	if err != nil {
		log.Fatal(err)
	}
	defer levelKV.Close()
	get(levelKV, "levelKV")
}

func TestRW(t *testing.T) {
	levelKV, err := New(store.KVRoot)
	defer levelKV.Close()
	if err != nil {
		log.Fatal(err)
	}
	get(levelKV, "levelKV")
	put(levelKV, "levelKV", fmt.Sprint("wrote at: ", time.Now()))
	get(levelKV, "levelKV")
}
func TestClose(t *testing.T) {
	levelKV, err := New(store.KVRoot)
	defer levelKV.Close()
	if err != nil {
		log.Fatal(err)
	}
	put(levelKV, "levelKV", fmt.Sprint("wrote at: ", time.Now()))
	levelKV.Close()
	get(levelKV, "levelKV")
}
