package buildinfo

type Version struct {
	Name         string       `yaml:"name" json:"name"`                 // 名称
	Description  string       `yaml:"description" json:"description"`   // 描述
	URL          string       `yaml:"url" json:"url"`                   // 仓库 URL
	BugReportUrl string       `yaml:"bugReportUrl" json:"bugReportUrl"` // 错误报告网址
	BuildVersion BuildVersion `yaml:"buildVersion" json:"buildVersion"` // 构建版本
	Organization Organization `yaml:"organization" json:"organization"` // 组织
}

type BuildVersion struct {
	BuildDate          string `yaml:"buildDate" json:"buildDate"`                   // 构建日期
	Compiler           string `yaml:"compiler" json:"compiler"`                     // 编译器
	GitCommitSha       string `yaml:"gitCommitSha" json:"gitCommitSha"`             // 项目为其构建的提交修订
	GitCommitShortSha  string `yaml:"gitCommitShortSha" json:"gitCommitShortSha"`   // 项目为其构建的提交修订的前八个字符
	GitCommitTag       string `yaml:"gitCommitTag" json:"gitCommitTag"`             // 提交标签名称
	GitCommitTimestamp string `yaml:"gitCommitTimestamp" json:"gitCommitTimestamp"` // ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00
	GitCommitBranch    string `yaml:"gitCommitBranch" json:"gitCommitBranch"`       // 提交分支名称
	GitTreeState       string `yaml:"gitTreeState" json:"gitTreeState"`             // Git 树状态
	GitVersion         string `yaml:"gitVersion" json:"gitVersion"`                 // Git 版本
	GoVersion          string `yaml:"goVersion" json:"goVersion"`                   // Go 版本
	Major              string `yaml:"major" json:"major"`                           // 主版本号
	Minor              string `yaml:"minor" json:"minor"`                           // 次版本号
	Revision           string `yaml:"revision" json:"revision"`                     // 修订版本号
	Platform           string `yaml:"platform" json:"platform"`                     // 平台
	InstanceUrl        string `yaml:"instanceUrl" json:"instanceUrl"`               // 实例地址
	CiPipelineId       string `yaml:"ciPipelineId" json:"ciPipelineId"`             // CI 流水线 ID
	CiJobId            string `yaml:"ciJobId" json:"ciJobId"`                       // CI 任务 ID
}

type Organization struct {
	Name        string `yaml:"name" json:"name"`               // 组织名称
	Description string `yaml:"description" json:"description"` // 描述
	Url         string `yaml:"url" json:"url"`                 // 链接
	Email       string `yaml:"email" json:"email"`             // 邮箱
	Address     string `yaml:"address" json:"address"`         // 地址
}
