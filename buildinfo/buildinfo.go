package buildinfo

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var commitAuthor string    // Name <email> 格式的提交作者
var commitBranch string    // 提交分支名称
var commitSha string       // 项目为其构建的提交修订
var commitShortSha string  // 项目为其构建的提交修订的前八个字符
var commitTag string       // 提交标签名称
var commitTimestamp string // ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00
var goVersion string       // go version 命令返回值，如：go version go1.21.1 windows/amd64
var goShortVersion string  // go version 命令返回值截取版本号，如：go1.21.1
var goPlatform string      // go version 命令返回值截取平台信息，如：windows/amd64
var Author = "徐晓伟"
var Email = "xuxiaowei@xuxiaowei.com.cn"
var appVersion = "dev"
var Copyright = "徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>"

func init() {
	commitAuthor = CommitAuthor()
	commitBranch = CommitBranch()
	commitSha = CommitSha()
	commitShortSha = CommitShortSha()
	commitTag = GitDescribeTags(false, "", true)
	commitTimestamp = CommitTimestamp()
	goVersion = GoVersion()
	goShortVersion = GoShortVersion()
	goPlatform = GoPlatform()
}

func Command() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "commit-author",
			Aliases: []string{"commitAuthor", "ca"},
			Usage:   "Name <email> 格式的提交作者",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(CommitAuthor())
				return nil
			},
		},
		{
			Name:    "commit-branch",
			Aliases: []string{"commitBranch", "cb"},
			Usage:   "提交分支名称",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(CommitBranch())
				return nil
			},
		},
		{
			Name:    "commit-sha",
			Aliases: []string{"commitSha", "cs"},
			Usage:   "项目为其构建的提交修订",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(CommitSha())
				return nil
			},
		},
		{
			Name:    "commit-short-sha",
			Aliases: []string{"commitShortSha", "css"},
			Usage:   "项目为其构建的提交修订的前八个字符",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(CommitShortSha())
				return nil
			},
		},
		{
			Name:    "commit-tag",
			Aliases: []string{"commitTag", "ct"},
			Usage:   "提交标签名称",
			Flags:   []cli.Flag{ExactMatchFlag(), AbbrevFlag(), HeadFlag()},
			Action: func(cCtx *cli.Context) error {
				var exactMatch = cCtx.Bool(ExactMatch)
				var abbrev = cCtx.String(Abbrev)
				var head = cCtx.Bool(Head)
				fmt.Println(GitDescribeTags(exactMatch, abbrev, head))
				return nil
			},
		},
		{
			Name:    "commit-timestamp",
			Aliases: []string{"commitTimestamp", "cts"},
			Usage:   "ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(CommitTimestamp())
				return nil
			},
		},
		{
			Name:  "now",
			Usage: "ISO 8601 格式的当前时间戳，如：2023-10-02T02:35:20+08:00",
			Action: func(cCtx *cli.Context) error {
				now := time.Now()
				str := now.Format("2006-01-02T15:04:05+08:00")
				fmt.Println(str)
				return nil
			},
		},
		{
			Name:    "go-version",
			Aliases: []string{"goVersion", "gv"},
			Usage:   "go version 命令返回值，如：go version go1.21.1 windows/amd64",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(GoVersion())
				return nil
			},
		},
		{
			Name:    "go-short-version",
			Aliases: []string{"goShortVersion", "gsv"},
			Usage:   "go version 命令返回值截取版本号，如：go1.21.1",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(GoShortVersion())
				return nil
			},
		},
		{
			Name:    "go-platform",
			Aliases: []string{"goPlatform", "gp"},
			Usage:   "go version 命令返回值截取平台信息，如：windows/amd64",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(GoPlatform())
				return nil
			},
		},
		{
			Name:    "git-tree-state",
			Aliases: []string{"gitTreeState", "gts"},
			Usage:   "git 树状态：clean、dirty",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(GitTreeState())
				return nil
			},
		},
	}
}

func Main() {
	app := &cli.App{
		Name:      "git-go",
		Version:   appVersion,
		Authors:   []*cli.Author{{Name: Author, Email: Email}},
		Usage:     "基于 Go 语言开发的 Git 命令行工具",
		Copyright: Copyright,
		Commands:  Command(),
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

func GitDescribeTags(exactMatch bool, abbrev string, head bool) string {
	cmd := exec.Command("git", "describe", "--tags")

	if exactMatch {
		cmd.Args = append(cmd.Args, "--exact-match")
	}

	if abbrev != "" {
		cmd.Args = append(cmd.Args, "--abbrev="+abbrev)
	}

	if head {
		cmd.Args = append(cmd.Args, "HEAD")
	}

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
	str := strings.ReplaceAll(strings.TrimSpace(string(output)), "\"", "")
	return str
}

func GoVersion() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	return str
}

func GoShortVersion() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	parts := strings.Split(str, " ")
	return parts[2]
}

func GoPlatform() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	parts := strings.Split(str, " ")
	return parts[3]
}

func GitTreeState() string {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	str := strings.TrimSpace(string(output))
	if str == "" {
		return "clean"
	}
	return "dirty"
}
