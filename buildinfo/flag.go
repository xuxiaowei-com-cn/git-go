package buildinfo

import "github.com/urfave/cli/v2"

func ExactMatchFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  ExactMatch,
		Value: false,
	}
}

func AbbrevFlag() cli.Flag {
	return &cli.StringFlag{
		Name: Abbrev,
	}
}

func HeadFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  Head,
		Value: false,
	}
}
