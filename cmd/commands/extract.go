package commands

import (
	"github.com/deffusion/chunkstore/cmd/flags"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/store"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var Extract = &cli.Command{
	Name:  "extract",
	Usage: "chunkstore extract --path ./output.txt s8a1d36...15f",
	Flags: []cli.Flag{
		flags.Path,
	},
	Action: func(cCtx *cli.Context) error {
		d, err := digest.New(cCtx.Args().First())
		if err != nil {
			return errors.WithMessage(err, "Extract Command")
		}
		path := cCtx.String("path")
		cs := cCtx.App.Metadata["chunkstore"].(*store.ChunkStore)
		err = cs.Extract(d, path)
		if err != nil {
			return errors.WithMessage(err, "Extract Command")
		}
		return nil
	},
}
