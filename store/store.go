package store

import (
	"fmt"
	"log"
	"os"
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

func init() {
	makeDirIfNotExist(ChunkRoot)
	makeDirIfNotExist(KVRoot)
}

func makeDirIfNotExist(path string) {
	_, err := os.Stat(path)
	fmt.Println("check dir:", path)
	if os.IsNotExist(err) {
		fmt.Printf("directory %s is not exist\n", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("directory %s was created\n", path)
	}
}
