package main

import (
	"github.com/deffusion/chunkstore/cmd/commands"
	"github.com/deffusion/chunkstore/store"
	"github.com/deffusion/chunkstore/store/kv/level_kv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	cs := store.New(db, store.ChunkRoot, logger.Named("main"))
	defer cs.Close()
	app.Metadata["chunkstore"] = cs

	app.Run(os.Args)
}
