package commands

import (
	"fmt"
	"github.com/deffusion/chunkstore/store"
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
		cs := cCtx.App.Metadata["chunkstore"].(*store.ChunkStore)
		root := cs.Add(file)
		fmt.Println("file hash:", root)
		return nil
	},
}
