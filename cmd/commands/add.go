package commands

import (
	"fmt"
	"github.com/deffusion/chunkstore/store"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"os"
)

var Add = &cli.Command{
	Name:  "add",
	Usage: "chunkstore add example.txt",
	Action: func(cCtx *cli.Context) error {
		file, err := os.Open(cCtx.Args().First())
		defer file.Close()
		if err != nil {
			return errors.WithMessage(err, "Add Command")
		}
		cs := cCtx.App.Metadata["chunkstore"].(*store.ChunkStore)
		root, err := cs.Add(file)
		if err != nil {
			return errors.WithMessage(err, "Add Command")
		}
		fmt.Println("file hash:", root)
		return nil
	},
}
