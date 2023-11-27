package flag

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/git-go/constant"
)

func VisibilityFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Visibility,
		Usage: "公开(public)、私有(private)或者所有(all)",
		Value: "all",
	}
}

func AffiliationFlag() cli.Flag {
	return &cli.StringFlag{
		Name: constant.Affiliation,
		Usage: "owner(授权用户拥有的仓库)、\n" +
			"\tcollaborator(授权用户为仓库成员)、\n" +
			"\torganization_member(授权用户为仓库所在组织并有访问仓库权限)、\n" +
			"\tenterprise_member(授权用户所在企业并有访问仓库权限)、\n" +
			"\tadmin(所有有权限的，包括所管理的组织中所有仓库、所管理的企业的所有仓库)。 \n" +
			"\t可以用逗号分隔符组合。\n" +
			"\t如: owner, organization_member 或 owner, collaborator, organization_member",
	}
}
