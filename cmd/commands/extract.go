package commands

import (
	"github.com/deffusion/chunkstore/cmd/flags"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/store"
	"github.com/urfave/cli/v2"
	"log"
)

var Extract = &cli.Command{
	Name:  "extract",
	Usage: "chunkstore extract s8a1d36...15f --path ./output.txt",
	Flags: []cli.Flag{
		flags.Path,
	},
	Action: func(cCtx *cli.Context) error {
		d, err := digest.New(cCtx.Args().First())
		if err != nil {
			log.Fatal("cmd.Get: ", err)
		}
		path := cCtx.String("path")
		cs := cCtx.App.Metadata["chunkstore"].(*store.ChunkStore)
		cs.Extract(d, path)
		return nil
	},
}
