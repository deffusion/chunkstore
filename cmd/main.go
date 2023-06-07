package main

import (
	"fmt"
	"github.com/deffusion/chunkstore/cmd/commands"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	fmt.Println(app().Run(os.Args))
}

func app() *cli.App {
	return &cli.App{
		Commands: commands.Root,
	}
}
