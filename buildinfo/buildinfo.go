package buildinfo

import (
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
