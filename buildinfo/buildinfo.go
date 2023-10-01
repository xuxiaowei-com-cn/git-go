package buildinfo

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strings"
)

var commitAuthor string    // Name <email> 格式的提交作者
var commitBranch string    // 提交分支名称
var commitSha string       // 项目为其构建的提交修订
var commitShortSha string  // 项目为其构建的提交修订的前八个字符
var commitTag string       // 提交标签名称
var commitTimestamp string // ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00

func init() {
	commitAuthor = CommitAuthor()
	commitBranch = CommitBranch()
	commitSha = CommitSha()
	commitShortSha = CommitShortSha()
	commitTag = CommitTag()
	commitTimestamp = CommitTimestamp()
}

func Main() {
	app := &cli.App{
		Name:  "git-go",
		Usage: "基于 Go 语言的 Git 命令行工具",
		Commands: []*cli.Command{
			{
				Name:    "commitAuthor",
				Aliases: []string{"ca"},
				Usage:   "Name <email> 格式的提交作者",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(CommitAuthor())
					return nil
				},
			},
			{
				Name:    "commitBranch",
				Aliases: []string{"cb"},
				Usage:   "提交分支名称",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(CommitBranch())
					return nil
				},
			},
			{
				Name:    "commitSha",
				Aliases: []string{"cs"},
				Usage:   "项目为其构建的提交修订",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(CommitSha())
					return nil
				},
			},
			{
				Name:    "commitShortSha",
				Aliases: []string{"css"},
				Usage:   "项目为其构建的提交修订的前八个字符",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(CommitShortSha())
					return nil
				},
			},
			{
				Name:    "commitTag",
				Aliases: []string{"ct"},
				Usage:   "提交标签名称",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(CommitTag())
					return nil
				},
			},
			{
				Name:    "commitTimestamp",
				Aliases: []string{"cts"},
				Usage:   "ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(CommitTimestamp())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func CommitAuthor() string {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:\"%an <%ae>\"")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}

func CommitBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}

func CommitSha() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}

func CommitShortSha() string {
	cmd := exec.Command("git", "rev-parse", "--short=8", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}

func CommitTag() string {
	cmd := exec.Command("git", "describe", "--tags", "--exact-match", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}

func CommitTimestamp() string {
	cmd := exec.Command("git", "log", "-1", "--format=\"%cd\"", "--date=iso-strict")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}
