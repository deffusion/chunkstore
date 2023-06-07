package main

import (
	"github.com/deffusion/chunkstore/cmd/commands"
	"github.com/deffusion/chunkstore/store"
	"github.com/deffusion/chunkstore/store/kv/level_kv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: commands.Root,
		Metadata: map[string]interface{}{},
	}
	db, err := level_kv.New(store.KVRoot)
	if err != nil {
		log.Fatal("main.app: open leveldb: ", err)
	}
	cs := store.New(db, store.ChunkRoot)
	defer cs.Close()
	app.Metadata["chunkstore"] = cs

	app.Run(os.Args)
}
