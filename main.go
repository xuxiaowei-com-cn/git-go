package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/git-go/buildinfo"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "git-go",
		Usage: "基于 Go 语言的 Git 命令行工具",
		Commands: []*cli.Command{
			{
				Name:    "commitAuthor",
				Aliases: []string{"ca"},
				Usage:   "Name <email> 格式的提交作者",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(buildinfo.CommitAuthor())
					return nil
				},
			},
			{
				Name:    "commitBranch",
				Aliases: []string{"cb"},
				Usage:   "提交分支名称",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(buildinfo.CommitBranch())
					return nil
				},
			},
			{
				Name:    "commitSha",
				Aliases: []string{"cs"},
				Usage:   "项目为其构建的提交修订",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(buildinfo.CommitSha())
					return nil
				},
			},
			{
				Name:    "commitShortSha",
				Aliases: []string{"css"},
				Usage:   "项目为其构建的提交修订的前八个字符",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(buildinfo.CommitShortSha())
					return nil
				},
			},
			{
				Name:    "commitTag",
				Aliases: []string{"ct"},
				Usage:   "提交标签名称",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(buildinfo.CommitTag())
					return nil
				},
			},
			{
				Name:    "commitTimestamp",
				Aliases: []string{"cts"},
				Usage:   "ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(buildinfo.CommitTimestamp())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
