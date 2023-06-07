package flags

import "github.com/urfave/cli/v2"

var Path = &cli.StringFlag{
	Name:  "path",
	Usage: `target directory or file`,
}
