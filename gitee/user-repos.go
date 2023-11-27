package gitee

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/git-go/flag"
)

// GetV5UserRepos 列出授权用户的所有仓库 https://gitee.com/api/v5/swagger#/getV5UserRepos
func GetV5UserRepos() *cli.Command {
	return &cli.Command{
		Name:  "get-v5-user-repos",
		Usage: "GetV5UserRepos 列出授权用户的所有仓库 https://gitee.com/api/v5/swagger#/getV5UserRepos",
		Flags: []cli.Flag{flag.VisibilityFlag(), flag.AffiliationFlag()},
		Action: func(context *cli.Context) error {

			return nil
		},
	}
}
