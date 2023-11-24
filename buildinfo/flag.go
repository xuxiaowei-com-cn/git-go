package buildinfo

import "github.com/urfave/cli/v2"

func DefaultTagFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  DefaultTag,
		Usage: "默认标签",
	}
}
