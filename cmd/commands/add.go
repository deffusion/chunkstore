package commands

import (
	"fmt"
	"github.com/deffusion/chunkstore/store"
	"github.com/deffusion/chunkstore/store/kv/level_kv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var Add = &cli.Command{
	Name:  "add",
	Usage: "chunkstore add example.txt",
	Action: func(cCtx *cli.Context) error {
		file, err := os.Open(cCtx.Args().First())
		defer file.Close()
		if err != nil {
			log.Fatal("cmd.Add: ", err)
		}
		db, err := level_kv.New(store.KVRoot)
		defer db.Close()
		if err != nil {
			log.Fatal("cmd.Add: ", err)
		}
		cs := store.New(db, store.ChunkRoot)
		root := cs.Add(file)
		fmt.Println("file hash:", root)
		return nil
	},
}
