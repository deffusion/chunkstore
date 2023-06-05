package memory_kv

import (
	"fmt"
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

func TestRW(t *testing.T) {
	memoKV, err := New()
	defer memoKV.Close()
	if err != nil {
		log.Fatal(err)
	}
	get(memoKV, "memoKV")
	put(memoKV, "memoKV", fmt.Sprint("wrote at: ", time.Now()))
	get(memoKV, "memoKV")
}
func TestClose(t *testing.T) {
	memoKV, err := New()
	defer memoKV.Close()
	if err != nil {
		log.Fatal(err)
	}
	put(memoKV, "memoKV", fmt.Sprint("wrote at: ", time.Now()))
	memoKV.Close()
	get(memoKV, "memoKV")
}
