package buildinfo

import (
	"fmt"
	"testing"
)

func Test_CommitAuthor(t *testing.T) {
	fmt.Println(CommitAuthor())
}

func Test_commitAuthor(t *testing.T) {
	fmt.Println(commitAuthor)
}

func Test_CommitBranch(t *testing.T) {
	fmt.Println(CommitBranch())
}

func Test_commitBranch(t *testing.T) {
	fmt.Println(commitBranch)
}

func Test_CommitSha(t *testing.T) {
	fmt.Println(CommitSha())
}

func Test_commitSha(t *testing.T) {
	fmt.Println(commitSha)
}

func Test_CommitShortSha(t *testing.T) {
	fmt.Println(CommitShortSha())
}

func Test_commitShortSha(t *testing.T) {
	fmt.Println(commitShortSha)
}

func Test_CommitTag(t *testing.T) {
	fmt.Println(GitDescribeTags(false, "", true))
}

func Test_commitTag(t *testing.T) {
	fmt.Println(commitTag)
}

func Test_CommitTimestamp(t *testing.T) {
	fmt.Println(CommitTimestamp())
}

func Test_commitTimestamp(t *testing.T) {
	fmt.Println(commitTimestamp)
}

func Test_GoVersion(t *testing.T) {
	fmt.Println(GoVersion())
}

func Test_goVersion(t *testing.T) {
	fmt.Println(goVersion)
}

func Test_GoShortVersion(t *testing.T) {
	fmt.Println(GoShortVersion())
}

func Test_goShortVersion(t *testing.T) {
	fmt.Println(goShortVersion)
}

func Test_GoPlatform(t *testing.T) {
	fmt.Println(GoPlatform())
}

func Test_goPlatform(t *testing.T) {
	fmt.Println(goPlatform)
}
