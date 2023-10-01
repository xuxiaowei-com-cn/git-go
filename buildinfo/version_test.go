package buildinfo

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestInfo(t *testing.T) {
	versionInfo := Version{
		Name: "Your App Name",
		BuildVersion: BuildVersion{
			BuildDate:         "2023-09-20",
			Compiler:          "your_compiler_version",
			GitCommitSha:      "your_git_commit_sha_hash",
			GitCommitShortSha: "your_git_commit_short_sha_hash",
			GitCommitTag:      "your_git_commit_tag",
			GitCommitBranch:   "your_git_commit_branch",
			GitTreeState:      "your_git_tree_state",
			GitVersion:        "your_git_version",
			GoVersion:         "your_go_version",
			Major:             "1",
			Minor:             "2",
			Revision:          "4",
			Platform:          "your_platform",
			CiPipelineId:      "your_ci_pipeline_id",
			CiJobId:           "your_ci_job_id",
		},
		Organization: Organization{
			Name:  "Your Organization Name",
			Url:   "Your Organization URL",
			Email: "your_organization_email",
		},
	}

	yamlData, err := yaml.Marshal(versionInfo)
	if err != nil {
		fmt.Println("版本信息无法转换为 YAML 格式:", err)
		panic(err)
	}
	fmt.Println(string(yamlData))

	jsonData, err := json.MarshalIndent(versionInfo, "", "    ")
	if err != nil {
		fmt.Println("版本信息无法转换为 JSON 格式:", err)
		return
	}
	fmt.Println(string(jsonData))

	// Example test case
	if versionInfo.Name != "Your App Name" {
		t.Errorf("Expected name to be 'Your App Name', but got '%s'", versionInfo.Name)
	}
}
