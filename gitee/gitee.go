package gitee

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "gitee",
		Usage: "gitee API",
		Subcommands: []*cli.Command{
			GetV5UserRepos(),
		},
	}
}
