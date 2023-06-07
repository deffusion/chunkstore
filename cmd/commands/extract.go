package commands

import (
	"fmt"
	"github.com/deffusion/chunkstore/cmd/flags"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/store"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
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
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("create file:", path)
		}
		cs := cCtx.App.Metadata["chunkstore"].(*store.ChunkStore)
		digests := cs.Get(d)
		for _, di := range digests {
			chunkFile, err := os.Open(fmt.Sprint(store.ChunkRoot, di.String()))
			n, err := io.Copy(file, chunkFile)
			chunkFile.Close()
			if err != nil && err != io.EOF {
				fmt.Println(err)
				log.Fatal(n, "bytes were wrote")
			}
		}
		return nil
	},
}
