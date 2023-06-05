package store

import (
	"fmt"
	"log"
	"os/user"
)

var storeRoot string
var ChunkRoot string
var KVRoot string

func init() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	storeRoot = fmt.Sprint(currentUser.HomeDir, "/chunkstore/")
	ChunkRoot = fmt.Sprint(storeRoot, "chunks/")
	KVRoot = fmt.Sprint(storeRoot, "kv/")
}
